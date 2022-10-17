package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/sha3"
	"golang.org/x/term"
)

type SignType struct {
	key *ecdsa.PrivateKey
}

func NewSignType() *SignType {
	return &SignType{}
}

func (s *SignType) loadDefaultKey() error {
	keyStr := os.Getenv("PRIVATE_KEY")
	if keyStr == "" {
		return errors.New("need PRIVATE_KEY")
	}
	if strings.EqualFold("0x", keyStr[:2]) {
		keyStr = keyStr[2:]
	}
	key, err := crypto.HexToECDSA(keyStr)
	if err != nil {
		return errors.New("load private key error")
	}
	s.key = key
	return nil
}

func (s *SignType) Command() *cli.Command {
	return &cli.Command{
		Name:    "sign-type",
		Aliases: []string{"st"},
		Usage:   "cast wallet sign-type <DOMAIN_SEPARATOR> <HASH>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "keystore",
				Usage: "--keystore <PATH>",
				Action: func(ctx *cli.Context, path string) error {
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
					s.key = key.PrivateKey
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if s.key == nil {
				if err := s.loadDefaultKey(); err != nil {
					return err
				}
			}

			separator, err := hexutil.Decode(ctx.Args().Get(0))
			if err != nil {
				return errors.New("decode separator error")
			}
			hash, err := hexutil.Decode(ctx.Args().Get(1))
			if err != nil {
				return errors.New("decode hash error")
			}

			data := append([]byte{0x19, 0x01}, append(separator, hash...)...)
			sha := sha3.NewLegacyKeccak256()
			sha.Write(data)
			signHash := sha.Sum(nil)

			sig, err := crypto.Sign(signHash, s.key)
			if err != nil {
				return errors.New("sign message error")
			}
			v := new(big.Int).SetBytes([]byte{sig[64] + 27})
			fmt.Printf(
				"Successfully sign type message.\nr: %s\ns: %s\nv: %d\n",
				hexutil.Encode(sig[:32]),
				hexutil.Encode(sig[32:64]),
				v.Uint64(),
			)

			return nil
		},
	}
}
