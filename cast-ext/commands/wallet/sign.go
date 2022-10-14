package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

type Sign struct {
	raw bool
}

func NewSign() *Sign {
	return &Sign{}
}

func (s *Sign) Command() *cli.Command {
	return &cli.Command{
		Name:    "sign",
		Aliases: []string{"s"},
		Usage:   "cast wallet sign [OPTIONS] <MESSAGE>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "raw",
				Aliases: []string{"r"},
				Usage:   "sign with raw message",
				Action: func(ctx *cli.Context, b bool) error {
					s.raw = b
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			hashStr := ctx.Args().Get(0)
			hash, err := hex.DecodeString(hashStr[2:])
			if err != nil {
				return errors.New("decode hash error")
			}
			keyStr := os.Getenv("PRIVATE_KEY")
			if keyStr == "" {
				return errors.New("need PRIVATE_KEY")
			}
			if strings.EqualFold("0x", keyStr[:2]) {
				keyStr = keyStr[2:]
			}
			key, err := crypto.HexToECDSA(keyStr)
			if err != nil {
				fmt.Print(err)
				return errors.New("load private key error")
			}
			if s.raw {
				sig, err := crypto.Sign(hash, key)
				if err != nil {
					return errors.New("sign message error")
				}
				v := new(big.Int).SetBytes([]byte{sig[64] + 27})
				fmt.Printf(
					"Successfully sign message.\nr: %s\ns: %s\nv: %d\n",
					hexutil.Encode(sig[:32]),
					hexutil.Encode(sig[32:64]),
					v.Uint64(),
				)
			}
			return nil
		},
	}
}
