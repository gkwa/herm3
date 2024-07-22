package cmd

import (
	"github.com/gkwa/herm3/custom_getters"
	"github.com/spf13/cobra"
)

var customGettersCmd = &cobra.Command{
	Use:   "custom_getters",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running custom_getters command")
		custom_getters.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(customGettersCmd)
}
