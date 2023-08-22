package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command
	RootCmd = &cobra.Command{
		Use:   "cert-helper",
		Short: "Certificate Helper",
		Run:   runRootCommand,
	}
	// arguments used by the generate command
	generateCommandArgs struct {
		// Algorithm is the algorithm for which certificates will be generated
		algorithm string
		// Key is the key ID using which certificates will be generated
		key string
	}
)

// Set version of the cli tool
func SetVersion(version string) {
	RootCmd.Version = version
}

func init() {
	// Setup flags
	f := GenerateCommand.Flags()
	f.StringVarP(&generateCommandArgs.algorithm, "algorithm", "a", "RS256", "Algorithm for which certificates will be generated. Valid values are RS256, RS384, RS512, ES256, ES384, ES512, HS256, HS384, HS512")
	f.StringVarP(&generateCommandArgs.key, "key", "k", uuid.NewString(), "Key ID using which certificates will be generated. Default will be random UUID.")
	RootCmd.AddCommand(GenerateCommand)
}

func runRootCommand(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
