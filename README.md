# vault-retriever
Utility to interact with HashiCorp Vault to retrieve Secrets and write to the filesystem

## Configuration

|        ENV VAR       |                              Description                             |
|:--------------------:|:--------------------------------------------------------------------:|
| ROLE_ID              | The Role ID for retrieving the secret value                          |
| SECRET_ID            | The Secret ID associated with the AppRole auth                       |
| SECRET_KEY           | The Key the secret is stored at                                      |
| DESTINATION_FILEPATH | Where to store the secret value on the file system                   |
| SECRET_PATH          | The API path where the secret can be found                           |
| VAULT_ADDR           | The address of the Vault cluster (https://127.0.0.1:8200 is default) |
