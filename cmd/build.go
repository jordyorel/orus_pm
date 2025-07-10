package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the Orus project",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := findOrusExecutable()
		if err != nil {
			return fmt.Errorf("orus executable not found. Please ensure Orus is installed and in your PATH, or set the ORUS_HOME environment variable")
		}
		fmt.Println("Building project...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
