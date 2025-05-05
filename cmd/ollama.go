package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ollamaCmd = &cobra.Command{
	Use:   "ollama",
	Short: "Interact with local Ollama instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ollama command invoked")
	},
}

func init() {
	rootCmd.AddCommand(ollamaCmd)
}
