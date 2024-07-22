package cmd

import (
	"github.com/gkwa/herm3/default_values"
	"github.com/spf13/cobra"
)

// defaultValues represents the impementDefaultValues command
var defaultValues = &cobra.Command{
	Use:   "default_values",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running implement_default_values command")
		default_values.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(defaultValues)
}
