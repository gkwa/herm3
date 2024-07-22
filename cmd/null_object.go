package cmd

import (
	"github.com/gkwa/herm3/null_object"
	"github.com/spf13/cobra"
)

var nullObjectCmd = &cobra.Command{
	Use:   "null_object",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running null_object command")
		null_object.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(nullObjectCmd)
}
