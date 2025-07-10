package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manages workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Managing workspaces...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
}
