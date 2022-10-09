package client

import (
	"context"
	"fmt"

	"github.cbhq.net/engineering/sff-workshop/internal/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEVMClient returns an EVM client that uses a CoinbaseCloud Node
func NewEVMClient(ctx context.Context, cfg *config.Config) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(
		ctx,
		fmt.Sprintf("https://%s:%s@%s", cfg.Username, cfg.Password, cfg.NodeURI),
	)

	return client, err
}
