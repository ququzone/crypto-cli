package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

type New struct {
}

func NewNew() *New {
	return &New{}
}

func (n *New) Command() *cli.Command {
	return &cli.Command{
		Name:    "new",
		Aliases: []string{"w"},
		Usage:   "cast-ext wallet new [PATH]",
		Action: func(ctx *cli.Context) error {
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				return errors.New("generate private key error")
			}
			privateKeyBytes := crypto.FromECDSA(privateKey)
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				return errors.New("generate public key error")
			}
			publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
			address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
			fmt.Printf(
				"Successfully created new keypair.:\nAddress: %s\nPrivate Key: %s\nPublic Key: %s\n",
				address,
				hexutil.Encode(privateKeyBytes)[2:],
				hexutil.Encode(publicKeyBytes)[2:],
			)
			return nil
		},
	}
}
