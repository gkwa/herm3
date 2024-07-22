package cmd

import (
	"github.com/gkwa/herm3/new_pattern"
	"github.com/spf13/cobra"
)

var newPatternCmd = &cobra.Command{
	Use:   "new_pattern",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running new_pattern command")
		new_pattern.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(newPatternCmd)
}