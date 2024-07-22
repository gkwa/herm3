package cmd

import (
	"github.com/gkwa/herm3/generics"
	"github.com/spf13/cobra"
)

var genericsCmd = &cobra.Command{
	Use:   "generics",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running generics command")
		generics.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(genericsCmd)
}
