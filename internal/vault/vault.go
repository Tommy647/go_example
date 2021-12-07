package vault

import (
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

const (
	// envToken vault token
	envToken = `VAULT_TOKEN`
	// envVault address
	envVault = `VAULT_ADDR`

	// dataKey map key we should find vault data on
	dataKey = `data`
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

	dataMap, ok := data.Data[dataKey]
	if !ok {
		return nil, errors.New("no data")
	}

	d, ok := dataMap.(map[string]interface{})
	if !ok {
		return nil, errors.New("unexpected data type")
	}

	dataString := make(map[string]string)
	for k := range d {
		if val, ok := d[k].(string); ok {
			dataString[k] = val
		}
	}
	return dataString, nil
}
