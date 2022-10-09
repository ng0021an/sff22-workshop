package main

import (
	"log"
	"net/http"

	"github.cbhq.net/engineering/sff-workshop/internal/server"
	"github.com/rs/cors"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}
	mux := http.NewServeMux()
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.HandleFunc("/gettoken", server.GetToken)
	handler := corsOpts.Handler(mux)
	log.Println(http.ListenAndServe(":8081", handler))
}
