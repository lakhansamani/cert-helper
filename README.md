# cert-helper

Cert Helper is a CLI tool that helps generate public & private certificates/keys for HMAC/RSA/ECDSA algorithms

## Installation

### Using golang

If you have golang installed and want to install `cert-helper` as binary that would run with `go/bin` then you can install using

```sh
go install github.com/lakhansamani/cert-helper
```

## Usage

### Flags with `cert-helper generate` cli command

- `-a, --algorithm` string Algorithm for which certificates will be generated. Valid values are RS256, RS384, RS512, ES256, ES384, ES512, HS256, HS384, HS512 (default "RS256")
- `-h, --help` help for cert-helper
- `-k, --key` string Key ID using which certificates will be generated. Default will be random UUID. (default "RANDOM KEY")

### Example

```sh
cert-helper generate
```
