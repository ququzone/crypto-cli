package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

type Detail struct {
	showPrivate bool
}

func NewDetail() *Detail {
	return &Detail{}
}

func (d *Detail) Command() *cli.Command {
	return &cli.Command{
		Name:    "detail",
		Aliases: []string{"d"},
		Usage:   "cast-ext wallet detail [OPTIONS] [PATH]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "private",
				Aliases: []string{"p"},
				Usage:   "show private key",
				Action: func(ctx *cli.Context, b bool) error {
					d.showPrivate = b
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			path := ctx.Args().First()
			fmt.Println("Enter password:")
			password, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return errors.New("read password error")
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return errors.New("read keystore error")
			}

			key, err := keystore.DecryptKey(data, string(password))
			if err != nil {
				return errors.New("decrypt keystore error")
			}
			publicKey := key.PrivateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				return errors.New("generate public key error")
			}
			publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

			if d.showPrivate {
				privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
				fmt.Printf(
					"Successfully load %s keypair.:\nAddress: %s\nPrivate Key: %s\nPublic Key: %s\n",
					path,
					key.Address,
					hex.EncodeToString(privateKeyBytes),
					hex.EncodeToString(publicKeyBytes),
				)
			} else {
				fmt.Printf(
					"Successfully load %s keypair.:\nAddress: %s\nPublic Key: %s\n",
					path,
					key.Address,
					hex.EncodeToString(publicKeyBytes),
				)
			}
			return nil
		},
	}
}
