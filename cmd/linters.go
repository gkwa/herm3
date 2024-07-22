package cmd

import (
	"github.com/gkwa/herm3/linters"
	"github.com/spf13/cobra"
)

var lintersCmd = &cobra.Command{
	Use:   "linters",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running linters command")
		linters.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(lintersCmd)
}
