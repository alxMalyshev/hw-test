package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	dirPath := "./testdata/env"

	expextedMap := map[string]EnvValue{
		"HELLO": {
			Value:      "\"hello\"",
			NeedRemove: false,
		},
		"BAR": {
			Value:      "bar",
			NeedRemove: false,
		},
		"EMPTY": {
			Value:      "",
			NeedRemove: false,
		},
		"FOO": {
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		},
		"UNSET": {
			Value:      "--UNSET--",
			NeedRemove: true,
		},
	}

	t.Run("Getting env from Dir", func(t *testing.T) {
		env, _ := ReadDir(dirPath)

		require.Equal(t, expextedMap["HELLO"], env["HELLO"])
		require.Equal(t, expextedMap["BAR"], env["BAR"])
		require.Equal(t, expextedMap["FOO"], env["FOO"])
		require.Equal(t, expextedMap["UNSET"], env["UNSET"])
	})
}
