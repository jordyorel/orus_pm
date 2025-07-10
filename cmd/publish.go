package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes the package to the Orus public registry",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Publishing package...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
