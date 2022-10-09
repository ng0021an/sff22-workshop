package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.cbhq.net/engineering/sff-workshop/internal/client"
	"github.cbhq.net/engineering/sff-workshop/internal/config"
	"github.cbhq.net/engineering/sff-workshop/internal/handler"
	"github.cbhq.net/engineering/sff-workshop/internal/keystore"
)

type getTokenRequest struct {
	ctx        context.Context
	to         string
	id         int64
	quantity   int64
	resChannel chan *getTokenResponse
}

type getTokenResponse struct {
	res string
	err error
}

type Server struct {
	transactionHandler *handler.TransactionHandler
	inputValidator     *handler.InputValidator
	queue              chan *getTokenRequest
}

func NewServer() (*Server, error) {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	evmClient, err := client.NewEVMClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	signer, err := keystore.NewSigner(cfg)
	if err != nil {
		return nil, err
	}

	inputValidator, err := handler.NewInputValidator(ctx, evmClient, cfg)
	if err != nil {
		return nil, err
	}

	transactionHandler, err := handler.NewTransactionHandler(ctx, evmClient, cfg, signer, inputValidator)
	if err != nil {
		return nil, err
	}

	queue := make(chan *getTokenRequest, 500)
	s := &Server{
		transactionHandler: transactionHandler,
		inputValidator:     inputValidator,
		queue:              queue,
	}
	go s.startTransactionProcessor(queue)

	return s, nil
}

func (s *Server) GetToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GetToken request")
	query := r.URL.Query()

	to := query.Get("to")
	id, err := getInt64(&query, "id")
	if err != nil {
		handleError(w, err)
		return
	}
	quantity, err := getInt64(&query, "quantity")
	if err != nil {
		handleError(w, err)
		return
	}

	resChannel := make(chan *getTokenResponse, 1)
	req := &getTokenRequest{
		ctx:        r.Context(),
		to:         to,
		id:         id,
		quantity:   quantity,
		resChannel: resChannel,
	}
	s.queue <- req
	result := <-resChannel
	if result.err != nil {
		handleError(w, result.err)
		return
	}
	res := result.res
	log.Println(res)
	_, writeErr := w.Write([]byte(res))
	if writeErr != nil {
		log.Printf("Error writing response %v", writeErr)
	}
}

func getInt64(query *url.Values, field string) (int64, error) {
	val := query.Get(field)
	return strconv.ParseInt(val, 10, 64)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, writeErr := w.Write([]byte(fmt.Sprintf("%v", err)))
	if writeErr != nil {
		log.Printf("Error writing error response %v", writeErr)
	}
}

func (s *Server) startTransactionProcessor(queue chan *getTokenRequest) {
	for {
		req := <-queue
		log.Printf("Received request from queue %v", req)
		txHash, err := s.transactionHandler.ERC1155Transfer(req.ctx, req.to, req.id, req.quantity)
		req.resChannel <- &getTokenResponse{
			res: txHash,
			err: err,
		}
	}
}
