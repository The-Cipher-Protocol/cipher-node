package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	ibft "github.com/nnlgsakib/neth/command/NLG-ibft"
	"github.com/nnlgsakib/neth/command/backup"

	//"github.com/nnlgsakib/neth/command/bridge"
	"github.com/nnlgsakib/neth/command/genesis"
	"github.com/nnlgsakib/neth/command/helper"
	"github.com/nnlgsakib/neth/command/license"
	"github.com/nnlgsakib/neth/command/monitor"
	"github.com/nnlgsakib/neth/command/peers"

	//"github.com/nnlgsakib/neth/command/polybft"
	//"github.com/nnlgsakib/neth/command/polybftsecrets"
	//"github.com/nnlgsakib/neth/command/regenesis"
	//"github.com/nnlgsakib/neth/command/rootchain"
	"github.com/nnlgsakib/neth/command/secrets"
	"github.com/nnlgsakib/neth/command/server"
	"github.com/nnlgsakib/neth/command/status"
	"github.com/nnlgsakib/neth/command/txpool"
	"github.com/nnlgsakib/neth/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "The go implementation of neth core (developed by https://github.com/nnlgsakib)",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		//	rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
	//	polybftsecrets.GetCommand(),
	//	polybft.GetCommand(),
	//	bridge.GetCommand(),
	//regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
