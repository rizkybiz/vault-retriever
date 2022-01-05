package config

import (
	"errors"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/namsral/flag"
)

type Config struct {
	RoleID       string
	SecretID     string
	DestFilePath string
	SecretKey    string
	SecretPath   string
	VaultAddr    string
	TLS          vault.TLSConfig
}

var config Config

func GetConfig() (Config, error) {
	flag.StringVar(&config.RoleID, "role-id", "", "The configured Role ID for accessing the secret.")
	flag.StringVar(&config.SecretID, "secret-id", "", "The Secret ID retrieved from vault.")
	flag.StringVar(&config.DestFilePath, "destination-filepath", "", "Filepath of where to store the secret once retrieved. Should include filename.")
	flag.StringVar(&config.SecretKey, "secret-key", "", "The name of the key where the value is stored in vault.")
	flag.StringVar(&config.SecretPath, "secret-path", "", "The path where the KV secret engine was enabled. Generally looks like: \"{engine path name}/data/{secret name}\".")
	vfs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "VAULT", 0)
	vfs.StringVar(&config.TLS.CACert, "ca-cert", "", "Path to a PEM-encoded CA cert file to use to verify the vault server SSL certificate.")
	vfs.StringVar(&config.TLS.CAPath, "ca-path", "", "Path to a directory of PEM-encoded CA cert files to verify the vault server SSL certificate.")
	vfs.StringVar(&config.TLS.ClientCert, "client-cert", "", "Path to the certificate for Vault communication.")
	vfs.StringVar(&config.TLS.ClientKey, "client-key", "", "Path to the private key for vault communication.")
	vfs.StringVar(&config.TLS.TLSServerName, "tls-server-name", "", "If set, this is used to set the SNI host when connecting via TLS.")
	vfs.StringVar(&config.VaultAddr, "addr", "https://127.0.0.1:8200", "The address of the Vault cluster.")
	vfs.BoolVar(&config.TLS.Insecure, "insecure", true, "Enables or disables SSL verification.")
	vfs.Parse(os.Args[1:])
	flag.Parse()

	if config.RoleID == "" || config.SecretID == "" {
		return config, errors.New("ROLE_ID and SECRET_ID both need to be set to authenticate with Vault")
	}
	if config.DestFilePath == "" {
		return config, errors.New("DESTINATION_FILEPATH needs to be set in order to write the secret to the filesystem")
	}
	if config.SecretKey == "" {
		return config, errors.New("SECRET_KEY needs to be set to identify the Key:Value secret to be retrieved")
	}
	if config.SecretPath == "" {
		return config, errors.New("SECRET_PATH needs to be set in order to request the secret")
	}
	return config, nil
}

func getTLSConfig(c *Config) error {
	return nil
}
