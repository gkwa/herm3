package cmd

import (
	"github.com/gkwa/herm3/zero_values"
	"github.com/spf13/cobra"
)

var zeroValuesCmd = &cobra.Command{
	Use:   "zero_values",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running zero_values command")
		zero_values.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(zeroValuesCmd)
}
