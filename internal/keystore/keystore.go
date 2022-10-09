package keystore

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.cbhq.net/engineering/sff-workshop/internal/config"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type Signer interface {
	Sign(chainId *big.Int, unsignedTx *types.Transaction) (*types.Transaction, error)
	Address() *common.Address
}

type signer struct {
	privateKey *ecdsa.PrivateKey
	address    *common.Address
}

func NewSigner(cfg *config.Config) (Signer, error) {
	seed := bip39.NewSeed(cfg.Mnemonic, "")

	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("error getting master key: %v", err)
	}

	// This gives the path: m/44H
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.Child(hdkeychain.HardenedKeyStart + 60)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.Child(0)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H/0/0
	acc44H60H0H00, err := acc44H60H0H0.Child(0)
	if err != nil {
		return nil, err
	}

	btcecPrivKey, err := acc44H60H0H00.ECPrivKey()
	if err != nil {
		return nil, err
	}

	privateKey := btcecPrivKey.ToECDSA()

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &signer{
		privateKey: privateKey,
		address:    &address,
	}, nil
}

// Sign the transaction (OFFLINE)
func (s *signer) Sign(chainId *big.Int, unsignedTx *types.Transaction) (*types.Transaction, error) {
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(chainId), s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}

	return signedTx, nil
}

func (s *signer) Address() *common.Address {
	return s.address
}
