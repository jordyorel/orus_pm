package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Upgrading all packages...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
