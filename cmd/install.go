package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := LoadManifest("orus.toml")
		if err != nil {
			return err
		}
		if len(m.Dependencies) == 0 {
			fmt.Println("No dependencies to install")
			return nil
		}
		if err := os.MkdirAll(".orus", 0755); err != nil {
			return err
		}
		for name, url := range m.Dependencies {
			dest := filepath.Join(".orus", name)
			if _, err := os.Stat(dest); err == nil {
				fmt.Printf("Dependency %s already installed\n", name)
				continue
			}
			fmt.Printf("Cloning %s into %s...\n", url, dest)
			c := exec.Command("git", "clone", "--depth", "1", url, dest)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				return err
			}
		}
		fmt.Println("Dependencies installed")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
