package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Username                string
	Password                string
	NodeURI                 string
	Mnemonic                string
	ContractAddress         string
	MaxGoldBadgeTotalQty    int64
	MaxGoldBadgeTransferQty int64
	MaxPointTotalQty        int64
	MaxPointTransferQty     int64
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	maxGoldBadgeTotalQty, err := strconv.ParseInt(os.Getenv("MAX_GOLD_BADGE_TOTAL_QUANTITY"), 10, 64)
	if err != nil {
		return nil, err
	}
	maxGoldBadgeTransferQty, err := strconv.ParseInt(os.Getenv("MAX_GOLD_BADGE_TRANSFER_QUANTITY"), 10, 64)
	if err != nil {
		return nil, err
	}
	maxPointTotalQty, err := strconv.ParseInt(os.Getenv("MAX_POINT_TOTAL_QUANTITY"), 10, 64)
	if err != nil {
		return nil, err
	}
	maxPointTransferQty, err := strconv.ParseInt(os.Getenv("MAX_POINT_TRANSFER_QUANTITY"), 10, 64)
	if err != nil {
		return nil, err
	}
	return &Config{
		Username:                os.Getenv("USERNAME"),
		Password:                os.Getenv("PASSWORD"),
		NodeURI:                 os.Getenv("NODE_URI"),
		Mnemonic:                os.Getenv("MNEMONIC"),
		ContractAddress:         os.Getenv("CONTRACT_ADDRESS"),
		MaxGoldBadgeTotalQty:    maxGoldBadgeTotalQty,
		MaxGoldBadgeTransferQty: maxGoldBadgeTransferQty,
		MaxPointTotalQty:        maxPointTotalQty,
		MaxPointTransferQty:     maxPointTransferQty,
	}, nil
}
