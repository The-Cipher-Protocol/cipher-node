package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/The-Cipher-Protocol/cipher-node/cmd/backup"
	ibft "github.com/The-Cipher-Protocol/cipher-node/cmd/chp"

	//"github.com/The-Cipher-Protocol/cipher-node/cmd/bridge"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/genesis"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/helper"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/license"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/monitor"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/peers"

	//"github.com/The-Cipher-Protocol/cipher-node/cmd/polybft"
	//"github.com/The-Cipher-Protocol/cipher-node/cmd/polybftsecrets"
	//"github.com/The-Cipher-Protocol/cipher-node/cmd/regenesis"
	//"github.com/The-Cipher-Protocol/cipher-node/cmd/rootchain"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/secrets"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/server"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/status"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/txpool"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "The go implementation of cipher core (developed by https://github.com/The-Cipher-Protocol)",
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
