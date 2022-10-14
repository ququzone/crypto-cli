package wallet

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/sha3"
)

type SignType struct {
}

func NewSignType() *SignType {
	return &SignType{}
}

func (s *SignType) Command() *cli.Command {
	return &cli.Command{
		Name:    "sign-type",
		Aliases: []string{"st"},
		Usage:   "cast wallet sign-type <DOMAIN_SEPARATOR> <HASH>",
		Action: func(ctx *cli.Context) error {
			separator, err := hexutil.Decode(ctx.Args().Get(0))
			if err != nil {
				return errors.New("decode separator error")
			}
			hash, err := hexutil.Decode(ctx.Args().Get(1))
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
			data := append([]byte{0x19, 0x01}, append(separator, hash...)...)
			sha := sha3.NewLegacyKeccak256()
			sha.Write(data)
			signHash := sha.Sum(nil)

			sig, err := crypto.Sign(signHash, key)
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
