package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Run the main project file or specified .orus file",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := "src/main.orus"
		if len(args) == 1 {
			file = args[0]
		}
		if _, err := os.Stat(file); err != nil {
			return fmt.Errorf("file %s not found", file)
		}
		c := exec.Command("orusc", file)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
