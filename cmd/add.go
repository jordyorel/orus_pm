package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <package>=<version>",
	Short: "Add dependency to orus.toml",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires package argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		arg := args[0]
		name, version, found := strings.Cut(arg, "=")
		if !found {
			return fmt.Errorf("invalid format, want name=version")
		}
		m, err := LoadManifest("orus.toml")
		if err != nil {
			return err
		}
		if m.Dependencies == nil {
			m.Dependencies = make(map[string]string)
		}
		m.Dependencies[name] = version
		if err := SaveManifest("orus.toml", m); err != nil {
			return err
		}
		fmt.Printf("Added %s %s\n", name, version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
