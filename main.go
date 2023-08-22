package main

import (
	"github.com/lakhansamani/cert-helper/cmd"
)

var (
	version = "0.0.0"
)

func main() {
	cmd.SetVersion(version)
	if err := cmd.RootCmd.Execute(); err != nil {
		panic("Failed to run root command")
	}
}
