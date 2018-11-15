package main

import (
	"os"

	"github.com/tensor-programming/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
