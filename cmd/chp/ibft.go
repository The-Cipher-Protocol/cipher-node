package ibft

import (
	"github.com/The-Cipher-Protocol/cipher-node/cmd/chp/candidates"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/chp/propose"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/chp/quorum"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/chp/snapshot"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/chp/status"
	_switch "github.com/The-Cipher-Protocol/cipher-node/cmd/chp/switch"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/helper"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "nlgbft",
		Short: "Top level NLG-IBFT command for interacting with the NLG-IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
