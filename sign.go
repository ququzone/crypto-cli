package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func main() {
	var raw bool

	app := &cli.App{
		Name:  "sign",
		Usage: "sign message",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "raw",
				Aliases: []string{"r"},
				Usage:   "sign with raw",
				Action: func(ctx *cli.Context, b bool) error {
					raw = b
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
			if strings.EqualFold("0x", keyStr[:2]) {
				keyStr = keyStr[2:]
			}
			key, err := crypto.HexToECDSA(keyStr)
			if err != nil {
				fmt.Print(err)
				return errors.New("load private key error")
			}
			if raw {
				sig, err := crypto.Sign(hash, key)
				if err != nil {
					return errors.New("sign message error")
				}
				v := new(big.Int).SetBytes([]byte{sig[64] + 27})
				fmt.Printf(
					"Signature:\n\tr: 0x%s,\n \ts: 0x%s,\n \tv: %d\n",
					hex.EncodeToString(sig[:32]),
					hex.EncodeToString(sig[32:64]),
					v.Uint64(),
				)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
