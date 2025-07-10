package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticates with the registry",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Logging in...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
