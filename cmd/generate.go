package cmd

import (
	"encoding/json"

	"github.com/lakhansamani/cert-helper/pkg"
	"github.com/spf13/cobra"
)

var (
	// Genearte command is the command to generate certificates
	GenerateCommand = &cobra.Command{
		Use:   "generate",
		Short: "Command to generate certificates",
		Run:   runRenerateCommand,
	}
)

// Output is the output of the generate command
type Output struct {
	EncryptionKey string `json:"encryption_key"`
	Secret        string `json:"secret"`
	PrivateKey    string `json:"private_key"`
	PublicKey     string `json:"public_key"`
}

func runRenerateCommand(cmd *cobra.Command, args []string) {
	algo := generateCommandArgs.algorithm
	key := generateCommandArgs.key
	output := Output{
		EncryptionKey: key,
	}
	if pkg.IsHMACA(algo) {
		secret, _, err := pkg.NewHMACKey(algo, key)
		if err != nil {
			panic(err)
		}
		output.Secret = secret
	} else if pkg.IsECDSA(algo) {
		_, privateKey, publicKey, _, err := pkg.NewECDSAKey(algo, key)
		if err != nil {
			panic(err)
		}
		output.PrivateKey = privateKey
		output.PublicKey = publicKey
	} else if pkg.IsRSA(algo) {
		_, privateKey, publicKey, _, err := pkg.NewRSAKey(algo, key)
		if err != nil {
			panic(err)
		}
		output.PrivateKey = privateKey
		output.PublicKey = publicKey
	} else {
		panic("Invalid algo")
	}

	// Print stringified JSON
	b, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	println("-----BEGIN ENCRYPTION KEY-----")
	println(generateCommandArgs.key)
	println("-----BEGIN ENCRYPTION KEY-----\n")
	println(output.PrivateKey)
	println(output.PublicKey)
	println("-----BEGIN JSON OUTPUT-----")
	println(string(b))
	println("-----END JSON OUTPUT-----\n")
}
