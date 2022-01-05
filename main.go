package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/rizkybiz/vault-retriever/config"
)

// TODO: Implement TLS
//       Implement Wrapping token?

func main() {

	// setup config from environment
	var c config.Config
	err := config.GetConfig(&c)
	if err != nil {
		log.Fatal(fmt.Sprintf("error getting config: %s", err))
	}

	// authenticate with vault
	config := vault.DefaultConfig()
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("error authenticating with vault: %s", err)
	}
	secretID := &auth.SecretID{FromString: c.SecretID}
	appRoleAuth, err := auth.NewAppRoleAuth(c.RoleID, secretID)
	if err != nil {
		log.Fatalf("error authenticating with vault: %s", err)
	}
	authInfo, err := client.Auth().Login(context.Background(), appRoleAuth)
	if err != nil {
		log.Fatalf("error authenticating with vault: %s", err)
	}
	if authInfo == nil {
		log.Fatal("error authenticating with vault: no auth info was returned after login")
	}

	// fetch the secret
	secret, err := client.Logical().Read(c.SecretPath)
	if err != nil {
		log.Fatalf("error fetching secret: %s", err)
	}

	// parse the return
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"])
	}

	//grab the secret at the key
	value, ok := data[c.SecretKey].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", data[c.SecretKey], data[c.SecretKey])
	}

	// write the secret to file
	absPath, err := filepath.Abs(c.DestFilePath)
	if err != nil {
		log.Fatalf("error parsing destination filepath: %s", err)
	}
	err = os.WriteFile(absPath, []byte(value), 0666)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}
}
