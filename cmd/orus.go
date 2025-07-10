package cmd

import (
	"os"
	"os/exec"
)

// findOrusExecutable finds the path to the `orus` executable.
// It first checks the ORUS_HOME environment variable, then searches the system's PATH.
func findOrusExecutable() (string, error) {
	orusHome := os.Getenv("ORUS_HOME")
	if orusHome != "" {
		exePath := orusHome + "/orus"
		if _, err := os.Stat(exePath); err == nil {
			return exePath, nil
		}
	}

	// If ORUS_HOME is not set or the executable is not found there, search the PATH.
	return exec.LookPath("orus")
}
