package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Interact with MCP servers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MCP command invoked")
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)
}
