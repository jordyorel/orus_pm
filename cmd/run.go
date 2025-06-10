package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:           "run [file]",
	Short:         "Run the main project file or specified .orus file",
	Args:          cobra.MaximumNArgs(1),
	SilenceErrors: true, // Prevent Cobra from printing errors
	SilenceUsage:  true, // Prevent Cobra from displaying usage on error
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
		err := c.Run()
		if err != nil {
			// Don't display raw "exit status X" errors - these are already handled by stderr output
			if err.Error() != fmt.Sprintf("exit status %d", c.ProcessState.ExitCode()) {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			os.Exit(c.ProcessState.ExitCode()) // Exit with the same code as the called program
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func main() {
	err := exec.Command("go", "build", "-o", "orus", "main.go").Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building project: %v\n", err)
		os.Exit(1)
	}
}
