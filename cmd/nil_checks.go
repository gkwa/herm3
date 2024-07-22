package cmd

import (
	"github.com/gkwa/herm3/nil_checks"
	"github.com/spf13/cobra"
)

var nilChecksCmd = &cobra.Command{
	Use:   "nil_checks",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running nil_checks command")
		nil_checks.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(nilChecksCmd)
}
