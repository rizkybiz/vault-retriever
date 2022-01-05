package retriever

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/rizkybiz/vault-retriever/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Retriever struct {
	config      config.Config
	vaultClient *vault.Client
	log         zerolog.Logger
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// New creates and returns a new Retriever
func New() (*Retriever, error) {

	// initialize the Retriever
	var r Retriever
	// set the Retriever's config
	err := r.getConfig()
	if err != nil {
		return nil, err
	}
	// set the Retriever's vault client
	err = r.setVaultClient()
	if err != nil {
		return nil, err
	}
	// set the Retriever's logger
	//r.log = zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &r, nil
}

func (r *Retriever) Run() error {

	// use the vault client to retrieve secret
	secret, err := r.getSecret()
	if err != nil {
		return err
	}
	// write secret to file
	err = r.writeSecret(secret)
	if err != nil {
		return err
	}
	return nil
}

func (r *Retriever) getConfig() error {

	// create a new config and get the values
	c, err := config.GetConfig()
	if err != nil {
		log.Err(err).Msg("error getting the config")
		return err
	}
	r.config = c
	return nil
}

func (r *Retriever) setVaultClient() error {

	// authenticate with vault using AppRole auth
	config := vault.DefaultConfig()
	err := config.ConfigureTLS(&r.config.TLS)
	if err != nil {
		log.Err(err).Msg("error configuring TLS within vault client")
		return err
	}
	config.Address = r.config.VaultAddr
	client, err := vault.NewClient(config)
	if err != nil {
		log.Err(err).Msg("error creating vault client")
		return err
	}
	secretID := &auth.SecretID{FromString: r.config.SecretID}
	appRoleAuth, err := auth.NewAppRoleAuth(r.config.RoleID, secretID)
	if err != nil {
		log.Err(err).Msg("error creating vault client")
		return err
	}
	authInfo, err := client.Auth().Login(context.Background(), appRoleAuth)
	if err != nil {
		log.Err(err).Msg("error creating vault client")
		return err
	}
	if authInfo == nil {
		err = errors.New("no authentication information was returned")
		log.Err(err).Msg("error creating vault client")
		return err
	}
	r.vaultClient = client
	return nil
}

func (r *Retriever) getSecret() (string, error) {

	// fetch the secret
	secret, err := r.vaultClient.Logical().Read(r.config.SecretPath)
	if err != nil {
		log.Err(err).Msg("error fetching secret")
		return "", err
	}
	// parse the return
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		err = errors.New(fmt.Sprintf("data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"]))
		log.Err(err).Msg("error fetching secret")
		return "", err
	}
	//grab the secret at the key
	value, ok := data[r.config.SecretKey].(string)
	if !ok {
		err = errors.New(fmt.Sprintf("value type assertion failed: %T %#v", data[r.config.SecretKey], data[r.config.SecretKey]))
		log.Err(err).Msg("error fetching secret")
		return "", err
	}
	return value, nil
}

func (r *Retriever) writeSecret(secret string) error {

	// write the secret to file
	absPath, err := filepath.Abs(r.config.DestFilePath)
	if err != nil {
		log.Err(err).Msg("error parsing destination filepath")
		return err
	}
	err = os.WriteFile(absPath, []byte(secret), 0666)
	if err != nil {
		log.Err(err).Msg("error writing file")
		return err
	}
	return nil
}
