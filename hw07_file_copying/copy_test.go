package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	var (
		inputFilePath  = "testdata/input.txt"
		resultFilePath = "testdata/out_test.txt"
	)

	t.Run("ErrUnsupportedFile", func(t *testing.T) {
		err := Copy("/dev/urandom", resultFilePath, 0, 0)

		require.Equal(t, err, ErrUnsupportedFile)
	})

	t.Run("ErrOffsetExceedsFileSize", func(t *testing.T) {
		err := Copy(inputFilePath, resultFilePath, 10000, 0)

		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("Offset0Limit0", func(t *testing.T) {
		_ = Copy(inputFilePath, resultFilePath, 0, 0)

		testFile, _ := os.Open("testdata/out_offset0_limit0.txt")
		resultFile, _ := os.Open(resultFilePath)

		testFileBytes, _ := io.ReadAll(testFile)
		resultFileBytes, _ := io.ReadAll(resultFile)

		fileCloser(testFile, resultFile)
		_ = os.Remove(resultFilePath)

		require.Equal(t, testFileBytes, resultFileBytes)
	})

	t.Run("Offset0Limit10000", func(t *testing.T) {
		_ = Copy(inputFilePath, resultFilePath, 0, 10000)

		testFile, _ := os.Open("testdata/out_offset0_limit10000.txt")
		resultFile, _ := os.Open(resultFilePath)

		testFileBytes, _ := io.ReadAll(testFile)
		resultFileBytes, _ := io.ReadAll(resultFile)

		fileCloser(testFile, resultFile)
		_ = os.Remove(resultFilePath)

		require.Equal(t, testFileBytes, resultFileBytes)
	})

	t.Run("Offset6000Limit1000", func(t *testing.T) {
		_ = Copy(inputFilePath, resultFilePath, 6000, 1000)

		testFile, _ := os.Open("testdata/out_offset6000_limit1000.txt")
		resultFile, _ := os.Open(resultFilePath)

		testFileBytes, _ := io.ReadAll(testFile)
		resultFileBytes, _ := io.ReadAll(resultFile)

		fileCloser(testFile, resultFile)
		_ = os.Remove(resultFilePath)

		require.Equal(t, testFileBytes, resultFileBytes)
	})
}
