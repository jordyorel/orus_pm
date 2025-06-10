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
		if _, err := os.Stat("orus.toml"); os.IsNotExist(err) {
			f, err := os.Create("orus.toml")
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Fprint(f, "[package]\nname = \"example\"\nversion = \"0.1.0\"\n\n[dependencies]\n")
		}
		if _, err := os.Stat("orus.lock"); os.IsNotExist(err) {
			f, err := os.Create("orus.lock")
			if err != nil {
				return err
			}
			f.Close()
		}
		mainPath := filepath.Join("src", "main.orus")
		if _, err := os.Stat(mainPath); os.IsNotExist(err) {
			f, err := os.Create(mainPath)
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Fprint(f, "fn main() {\n    print(\"Hello from Orus!\")\n}\n")
		}
		fmt.Println("Initialized new Orus project")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
