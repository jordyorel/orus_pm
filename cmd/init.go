package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Orus project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll("src", 0755); err != nil {
			return err
		}
		if _, err := os.Stat("ka.toml"); os.IsNotExist(err) {
			f, err := os.Create("ka.toml")
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Fprint(f, "[package]\nname = \"example\"\nversion = \"0.1.0\"\n\n[dependencies]\n")
		}
		if _, err := os.Stat("ka.lock"); os.IsNotExist(err) {
			f, err := os.Create("ka.lock")
			if err != nil {
				return err
			}
			f.Close()
		}
		if _, err := os.Stat("README.md"); os.IsNotExist(err) {
			f, err := os.Create("README.md")
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Fprint(f, "# New Orus Project\n")
		}
		mainPath := filepath.Join("src", "main.orus")
		if _, err := os.Stat(mainPath); os.IsNotExist(err) {
			f, err := os.Create(mainPath)
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Fprint(f, "print(\"Hello from Orus!\")\n")
		}
		fmt.Println("Initialized new Orus project in the current directory.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
