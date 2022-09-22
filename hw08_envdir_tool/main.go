package main

import (
	"fmt"
)

func main() {
	env, err := ReadDir("./testdata/env")
	fmt.Println(env)
	fmt.Println(err)
}
