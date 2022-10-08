package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, value := range env {
		if value.NeedRemove {
			os.Unsetenv(key)
			continue
		}

		os.Setenv(key, value.Value)
	}

	mainCmd := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	mainCmd.Stdout = os.Stdout
	mainCmd.Stderr = os.Stderr
	mainCmd.Stdin = os.Stdin

	if err := mainCmd.Run(); err != nil {
		var execErr *exec.ExitError
		if errors.As(err, &execErr) {
			fmt.Printf("faild to execute command: %s\n", err.Error())
			return execErr.ExitCode()
		}
		return 128
	}

	return 0
}
