package vault

import (
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

const (
	envToken = `VAULT_TOKEN` // vault auth token
	envVault = `VAULT_ADDR`  // vault address
)

// GetSecrets under the key secrets from the store
func GetSecrets(store, secret string) (map[string]string, error) {
	client, err := api.NewClient(&api.Config{Address: os.Getenv(envVault), HttpClient: http.DefaultClient})
	if err != nil {
		return nil, errors.Wrap(err, "getting vault client")
	}

	client.SetToken(os.Getenv(envToken))

	data, err := client.Logical().Read(store + `/data/` + secret)
	if err != nil {
		return nil, errors.Wrapf(err, "reading vault secret: %s/data/%s", store, secret)
	}
	// @todo: check data exists first or we panic
	dataMap, ok := data.Data["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("unexpected datatype")
	}
	dataString := make(map[string]string)
	for k, _ := range dataMap {
		if val, ok := dataMap[k].(string); ok {
			dataString[k] = val
		}
	}
	return dataString, nil
}
