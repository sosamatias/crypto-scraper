package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToCryptoSnapshot(t *testing.T) {
	t.Parallel()
	mock := GatewayResponseMock()
	coins := mock.ToCryptoSnapshot()
	assert.Equal(t, 2, len(coins))
	assert.Equal(t, mock.Data[0].Name, coins[0].Name)
	assert.Equal(t, mock.Data[0].Symbol, coins[0].Symbol)
	assert.Equal(t, mock.Data[0].Quote.Usd.MarketCap, coins[0].MarketCap)
	assert.Equal(t, mock.Data[0].Quote.Usd.Price, coins[0].USDPrice)
	assert.Equal(t, mock.Data[0].Quote.Usd.LastUpdated, coins[0].UpdatedAt)
	assert.Equal(t, mock.Data[1].Name, coins[1].Name)
	assert.Equal(t, mock.Data[1].Symbol, coins[1].Symbol)
	assert.Equal(t, mock.Data[1].Quote.Usd.MarketCap, coins[1].MarketCap)
	assert.Equal(t, mock.Data[1].Quote.Usd.Price, coins[1].USDPrice)
	assert.Equal(t, mock.Data[1].Quote.Usd.LastUpdated, coins[1].UpdatedAt)
}
