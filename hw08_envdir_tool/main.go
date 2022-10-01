package main

import (
	"fmt"
	"os"
)

func main() {
	env, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Printf("error during execution program: %s\n", err.Error())
		os.Exit(1)
	}

	exitCode := RunCmd(os.Args[2:], env)
	if exitCode != 0 {
		fmt.Println("error during execute command")
		os.Exit(exitCode)
	}

	os.Exit(0)
}
