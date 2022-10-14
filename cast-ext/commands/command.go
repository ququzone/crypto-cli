package commands

import (
	"github.com/ququzone/crypto-cli/cast-ext/commands/wallet"
	"github.com/urfave/cli/v2"
)

func Commonds() []*cli.Command {
	return []*cli.Command{
		wallet.Command(),
	}
}
