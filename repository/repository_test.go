package repository

import (
	"testing"

	"github.com/sosamatias/crypto-scraper/model"
	"github.com/stretchr/testify/assert"
)

func Test_CreateInBatches(t *testing.T) {
	repo, err := NewRepository("file::memory:?cache=shared")
	assert.NoError(t, err)

	gatewayResponse := model.GatewayResponseMock()
	cryptoSnapshots := gatewayResponse.ToCryptoSnapshot()

	err = repo.CreateInBatches(cryptoSnapshots, 2)
	assert.NoError(t, err)

	count, err := repo.Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}

func Test_Open_Error(t *testing.T) {
	_, err := NewRepository("user:pass@tcp(127.0.0.1:3306)/dbname")
	assert.Error(t, err)
}
