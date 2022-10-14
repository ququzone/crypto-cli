package wallet

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "wallet",
		Usage: "Wallet management utilities.",
		Subcommands: []*cli.Command{
			NewNew().Command(),
			NewSign().Command(),
			NewSignType().Command(),
		},
	}
}
