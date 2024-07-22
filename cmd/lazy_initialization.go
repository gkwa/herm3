package cmd

import (
	"github.com/gkwa/herm3/lazy_initialization"
	"github.com/spf13/cobra"
)

var lazyInitializationCmd = &cobra.Command{
	Use:   "lazy_initialization",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running lazy_initialization command")
		lazy_initialization.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(lazyInitializationCmd)
}
