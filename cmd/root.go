package cmd

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/lakhansamani/cert-generator/pkg"
)

var (
	// RootCmd is the root (and only) command of this service
	RootCmd = &cobra.Command{
		Use:   "cert-generator",
		Short: "Certificate Generator",
		Run:   runRootCmd,
	}
	rootArgs struct {
		// Algorithm is the algorithm for which certificates will be generated
		algorithm string
		// Key is the key ID using which certificates will be generated
		key string
	}
)

func init() {
	// Setup flags
	f := RootCmd.Flags()
	f.StringVarP(&rootArgs.algorithm, "algorithm", "a", "RS256", "Algorithm for which certificates will be generated. Valid values are RS256, RS384, RS512, ES256, ES384, ES512, HS256, HS384, HS512")
	f.StringVarP(&rootArgs.key, "key", "k", uuid.NewString(), "Key ID using which certificates will be generated. Default will be random UUID.")
}

type Output struct {
	Secret     string `json:"secret"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func runRootCmd(cmd *cobra.Command, args []string) {
	algo := rootArgs.algorithm
	key := rootArgs.key
	output := Output{}
	if pkg.IsHMACA(algo) {
		secret, _, err := pkg.NewHMACKey(algo, key)
		if err != nil {
			panic(err)
		}
		println(secret)
		output.Secret = secret
	} else if pkg.IsECDSA(algo) {
		_, privateKey, publicKey, _, err := pkg.NewECDSAKey(algo, key)
		if err != nil {
			panic(err)
		}
		println(privateKey)
		println(publicKey)
		output.PrivateKey = privateKey
		output.PublicKey = publicKey
	} else if pkg.IsRSA(algo) {
		_, privateKey, publicKey, _, err := pkg.NewRSAKey(algo, key)
		if err != nil {
			panic(err)
		}
		println(privateKey)
		println(publicKey)
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
	println("Stringified JSON Output")
	println(string(b))
}
