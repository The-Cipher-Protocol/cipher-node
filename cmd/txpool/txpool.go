package txpool

import (
	"github.com/The-Cipher-Protocol/cipher-node/cmd/helper"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/txpool/status"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/txpool/subscribe"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	txPoolCmd := &cobra.Command{
		Use:   "txpool",
		Short: "Top level command for interacting with the transaction pool. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(txPoolCmd)

	registerSubcommands(txPoolCmd)

	return txPoolCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// txpool status
		status.GetCommand(),
		// txpool subscribe
		subscribe.GetCommand(),
	)
}
