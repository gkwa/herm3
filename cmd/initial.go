package cmd

import (
	"github.com/gkwa/herm3/initial"
	"github.com/spf13/cobra"
)

// initialCmd represents the initial command
var initialCmd = &cobra.Command{
	Use:   "initial",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		initial.Main(logger)
	},
}

func init() {
	rootCmd.AddCommand(initialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
