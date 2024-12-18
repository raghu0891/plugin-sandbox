package command

import (
	"github.com/spf13/cobra"

	"github.com/goplugin/pluginv3.0/core/scripts/chaincli/config"
	"github.com/goplugin/pluginv3.0/core/scripts/chaincli/handler"
)

// RevertReasonCmd takes in a failed tx hash and tries to give you the reason
var RevertReasonCmd = &cobra.Command{
	Use:   "reason",
	Short: "Revert reason for failed TX.",
	Long:  `Given a failed TX tries to find the revert reason. args = tx hex address`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		baseHandler := handler.NewBaseHandler(cfg)
		baseHandler.RevertReason(args[0])
	},
}
