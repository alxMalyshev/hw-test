package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func NewEnvironment() Environment {
	return make(map[string]EnvValue)
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := NewEnvironment()

	dirWithFiles, err := os.Open(dir)
	if err != nil {
		fmt.Printf("error to open dir %s\n", err.Error())
		return nil, err
	}

	defer dirWithFiles.Close()

	listFiles, err := dirWithFiles.ReadDir(-1)
	if err != nil {
		fmt.Printf("error get list of files %s\n", err.Error())
		return nil, err
	}

	for _, file := range listFiles {
		readFile, err := os.Open(filepath.Join(dir, file.Name()))
		if err != nil {
			fmt.Println("faild to", err.Error())
			return nil, err
		}

		fileStat, err := readFile.Stat()
		if err != nil {
			fmt.Println("faild to get stat file ", err.Error())
			return nil, err
		}

		if fileStat.Size() > 0 {
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)

			envValue := getEnv(fileScanner)

			env[file.Name()] = EnvValue{
				Value:      envValue,
				NeedRemove: false,
			}
		} else {
			env[file.Name()] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
		}

		readFile.Close()
	}

	return env, nil
}

func getEnv(fileScanner *bufio.Scanner) string {
	var envValue string
	fileScanner.Scan()
	nullByte := bytes.Replace([]byte(fileScanner.Text()), []byte("\x00"), []byte("\n"), 1)
	envValue = strings.ReplaceAll(strings.TrimRight(string(nullByte), "\t "), "=", "")

	return envValue
}
