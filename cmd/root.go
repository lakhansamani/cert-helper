package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command
	RootCmd = &cobra.Command{
		Use:   "cert-helper",
		Short: "Certificate Helper",
		Run:   runRootCommand,
	}
)

// Set version of the cli tool
func SetVersion(version string) {
	RootCmd.Version = version
}

func init() {
	RootCmd.AddCommand(GenerateCommand)
}

func runRootCommand(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
