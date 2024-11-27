package keeper

import (
	"github.com/spf13/cobra"

	"github.com/goplugin/pluginv3.0/core/scripts/chaincli/config"
	"github.com/goplugin/pluginv3.0/core/scripts/chaincli/handler"
)

// jobCmd represents the command to run the service
var jobCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Add job to keeper nodes",
	Long:  `This command creates a job on keepers.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		hdlr := handler.NewKeeper(cfg)
		hdlr.CreateJob(cmd.Context())
	},
}
