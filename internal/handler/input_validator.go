package handler

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.cbhq.net/engineering/sff-workshop/contract"
	"github.cbhq.net/engineering/sff-workshop/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type limitSetting struct {
	transfer  int64
	ownership int64
}

type InputValidator struct {
	contractInstance *contract.Contract
	limits           map[int64]*limitSetting
}

func NewInputValidator(
	ctx context.Context,
	client *ethclient.Client,
	cfg *config.Config,
) (*InputValidator, error) {
	contractAddr := common.HexToAddress(cfg.ContractAddress)
	contractInstance, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	goldBadgeId, err := contractInstance.GoldBadge(callOpts)
	if err != nil {
		return nil, fmt.Errorf("error getting gold badge id: %v", err)
	}
	pointId, err := contractInstance.Points(callOpts)
	if err != nil {
		return nil, fmt.Errorf("error getting point id: %v", err)
	}

	limits := make(map[int64]*limitSetting)
	limits[goldBadgeId.Int64()] = &limitSetting{
		transfer:  cfg.MaxGoldBadgeTransferQty,
		ownership: cfg.MaxGoldBadgeTotalQty,
	}
	limits[pointId.Int64()] = &limitSetting{
		transfer:  cfg.MaxPointTransferQty,
		ownership: cfg.MaxPointTotalQty,
	}

	return &InputValidator{
		contractInstance: contractInstance,
		limits:           limits,
	}, nil
}

// Check if we exceed the ownership or transfer limit
// and return error in such case
func (v *InputValidator) CanTransfer(
	ctx context.Context,
	to string,
	id int64,
	quantity int64,
) error {
	limitSetting, ok := v.limits[id]
	if !ok {
		return fmt.Errorf("unrecognized token id")
	}
	if quantity > limitSetting.transfer {
		return fmt.Errorf("transfer limit exceeded")
	}
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	toAddr := common.HexToAddress(to)
	balance, err := v.contractInstance.BalanceOf(
		callOpts,
		toAddr,
		big.NewInt(id),
	)
	if err != nil {
		return fmt.Errorf("error calling BalanceOf: %v", err)
	}

	log.Printf("Balance: %v", balance)
	newBal := &big.Int{}
	newBal.Add(balance, big.NewInt(quantity))
	log.Printf("New balance: %v", newBal)
	if newBal.Cmp(big.NewInt(limitSetting.ownership)) > 0 {
		return fmt.Errorf("ownership limit exceeded")
	}
	return nil
}
