package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// setup apis for exposing the functionalities of shitpostercli

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with external APIs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("API command invoked")
	},
}

// Add subcommands for different APIs

func init() {
	apiCmd.AddCommand(ollamaCmd)
	apiCmd.AddCommand(mcpCmd)
}

// API for sending a response to queries to ollama instance

var ollamaResponseCmd = &cobra.Command{
	Use:   "ollama-response",
	Short: "Send a response to the local Ollama instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ollama response command invoked")
	},
}
