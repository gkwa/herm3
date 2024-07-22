package cmd

import (
	"github.com/gkwa/herm3/constructor_functions"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "constructor_functions",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		constructor_functions.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
