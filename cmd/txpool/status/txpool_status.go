package status

import (
	"context"

	command "github.com/The-Cipher-Protocol/cipher-node/cmd"
	"github.com/The-Cipher-Protocol/cipher-node/cmd/helper"
	"github.com/spf13/cobra"

	txpoolOp "github.com/The-Cipher-Protocol/cipher-node/utils/txpool/proto"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Returns the number of transactions in the transaction pool",
		Run:   runCommand,
	}
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	statusResponse, err := getTxPoolStatus(helper.GetGRPCAddress(cmd))
	if err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(&TxPoolStatusResult{
		Transactions: statusResponse.Length,
	})
}

func getTxPoolStatus(grpcAddress string) (*txpoolOp.TxnPoolStatusResp, error) {
	client, err := helper.GetTxPoolClientConnection(
		grpcAddress,
	)
	if err != nil {
		return nil, err
	}

	return client.Status(context.Background(), &empty.Empty{})
}
