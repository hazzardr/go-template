package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-template",
	Short: "go-template CLI.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
