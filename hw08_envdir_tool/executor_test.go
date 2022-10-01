package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	cmdArgs := []string{"./testdata/env", "/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"}

	t.Run("success exit code", func(t *testing.T) {
		env, _ := ReadDir(cmdArgs[0])
		exitCode := RunCmd(cmdArgs[1:], env)

		require.Equal(t, 0, exitCode)
	})

	t.Run("unsuccess exit code", func(t *testing.T) {
		env, _ := ReadDir(cmdArgs[0])
		cmdArgs[2] = "./testdata/echo.sh1"
		exitCode := RunCmd(cmdArgs[1:], env)

		require.Equal(t, 127, exitCode)
	})

	t.Run("unsuccess main command exit code", func(t *testing.T) {
		env, _ := ReadDir(cmdArgs[0])
		cmdArgs[1] = "/bin/bash1"
		exitCode := RunCmd(cmdArgs[1:], env)

		require.Equal(t, 128, exitCode)
	})

	t.Run("check system env", func(t *testing.T) {
		env, _ := ReadDir(cmdArgs[0])
		_ = RunCmd(cmdArgs[1:], env)

		helloEnv, exits := os.LookupEnv("HELLO")
		require.Equal(t, true, exits)
		require.Equal(t, "\"hello\"", helloEnv)

		barEnv, exits := os.LookupEnv("BAR")
		require.Equal(t, true, exits)
		require.Equal(t, "bar", barEnv)

		_, ok := os.LookupEnv("UNSET")
		require.Equal(t, false, ok)

		emptyEnv, exits := os.LookupEnv("EMPTY")
		require.Equal(t, true, exits)
		require.Equal(t, "", emptyEnv)
	})
}
