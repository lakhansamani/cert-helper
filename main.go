package main

import (
	"github.com/lakhansamani/cert-generator/cmd"
)

var (
	// Version is the version of this service
	version = "0.0.1"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic("Failed to run root command")
	}
}
