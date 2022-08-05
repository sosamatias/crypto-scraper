package scraper

import (
	"github.com/sosamatias/crypto-scraper/gateway"
	"github.com/sosamatias/crypto-scraper/repository"
)

func Execute(g gateway.Gateway, r repository.Repository) error {
	gatewayResponse, err := g.List(gateway.SortMarketCap)
	if err != nil {
		return err
	}

	cryptoSnapshots := gatewayResponse.ToCryptoSnapshot()

	return r.CreateInBatches(cryptoSnapshots, 50)
}
