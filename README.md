# vault-retriever
Utility to interact with HashiCorp Vault to retrieve Secrets and write to the filesystem

## Configuration

|        CLI Flag       |        ENV_VAR        |                                          Description                                         |
|:---------------------:|:---------------------:|:--------------------------------------------------------------------------------------------:|
| -role-id              | ROLE_ID               | The Role ID for retrieving the secret value                                                  |
| -secret-id            | SECRET_ID             | The Secret ID associated with the AppRole auth                                               |
| -secret-key           | SECRET_KEY            | The Key the secret is stored at                                                              |
| -destination-filepath | DESTINATION_FILEPATH  | Where to store the secret value on the file system                                           |
| -secret-path          | SECRET_PATH           | The API path where the secret can be found                                                   |
| -addr                 | VAULT_ADDR            | The address of the Vault cluster (https://127.0.0.1:8200 is default)                         |
| -ca-cert              | VAULT_CA_CERT         | Path to a PEM-encoded CA cert file to use to verify the vault server SSL certificate.        |
| -ca-path              | VAULT_CA_PATH         | Path to a directory of PEM-encoded CA cert files to verify the vault server SSL certificate. |
| -client-cert          | VAULT_CLIENT_CERT     | Path to the certificate for Vault communication.                                             |
| -client-key           | VAULT_CLIENT_KEY      | Path to the private key for vault communication.                                             |
| -tls-server-name      | VAULT_TLS_SERVER_NAME | If set, this is used to set the SNI host when connecting via TLS.                            |
| -insecure             | VAULT_INSECURE        | Enables or disables SSL verification.                                                        |