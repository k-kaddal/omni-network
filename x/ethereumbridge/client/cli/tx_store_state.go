package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStoreState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store-state [address] [slot]",
		Short: "Broadcast message store-state",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argSlot := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStoreState(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argSlot,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
