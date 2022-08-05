package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadSecrets(t *testing.T) {
	t.Parallel()
	secrets, err := LoadSecrets("data_test/data_test.yaml")
	assert.NoError(t, err)
	assert.Equal(t, "mocked_api_key", secrets.CoinMarketCapApiKey)
}

func Test_LoadSecrets_Error(t *testing.T) {
	t.Parallel()
	secrets, err := LoadSecrets("data_test/file_not_found.yaml")
	assert.Error(t, err)
	assert.Nil(t, secrets)
}
