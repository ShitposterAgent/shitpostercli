package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with external APIs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("API command invoked")
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
