package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
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
			path := ctx.Args().First()
			var password string

			if path != "" {
				fmt.Println("Enter password:")
				pw, err := term.ReadPassword(int(syscall.Stdin))
				if err != nil {
					return errors.New("read password error")
				}

				fmt.Println("Re-enter password:")
				pw1, err := term.ReadPassword(int(syscall.Stdin))
				if err != nil {
					return errors.New("read password error")
				}

				if string(pw) != string(pw1) {
					return errors.New("password dismatch")
				}
				password = string(pw)
			}

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
			address := crypto.PubkeyToAddress(*publicKeyECDSA)

			if path == "" {
				fmt.Printf(
					"Successfully created new keypair.:\nAddress: %s\nPrivate Key: %s\nPublic Key: %s\n",
					address.Hex(),
					hexutil.Encode(privateKeyBytes)[2:],
					hexutil.Encode(publicKeyBytes)[2:],
				)
			} else {
				key := &keystore.Key{
					Id:         uuid.New(),
					Address:    address,
					PrivateKey: privateKey,
				}

				data, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
				if err != nil {
					return errors.New("encrypt keystore error")
				}
				file, err := os.Create(path + address.Hex())
				if err != nil {
					return errors.New("open keystore file error")
				}
				if _, err = file.Write(data); err != nil {
					return errors.New("write keystore file error")
				}
				fmt.Printf(
					"Successfully created new keypair.:\nAddress: %s\nPublic Key: %s\nPath: %s\n",
					address.Hex(),
					hexutil.Encode(publicKeyBytes)[2:],
					path+address.Hex(),
				)
			}
			return nil
		},
	}
}
