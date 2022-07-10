package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sosamatias/crypto-scraper/config"
	"github.com/sosamatias/crypto-scraper/gateway"
	"github.com/sosamatias/crypto-scraper/repository"
)

func main() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(5).Seconds().Do(job)

	scheduler.StartAsync()

	scheduler.StartBlocking()
}

func job() {
	fmt.Printf("[%s] executing job.. \n", time.Now().UTC())

	secrets, err := config.LoadSecrets("config/secrets.yaml")
	if err != nil {
		panic(err)
	}

	gateway, err := gateway.NewGateway(secrets.CoinMarketCapApiKey, "", nil)
	if err != nil {
		panic(err)
	}

	repository, err := repository.NewRepository("repository/sqlite.db")
	if err != nil {
		panic(err)
	}

	err = executeJob(gateway, repository)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s] job executed successfully! \n\n", time.Now().UTC())
}

func executeJob(g gateway.Gateway, r repository.Repository) error {
	gatewayResponse, err := g.List(gateway.SortMarketCap)
	if err != nil {
		return err
	}

	cryptoSnapshots := gatewayResponse.ToCryptoSnapshot()

	return r.CreateInBatches(cryptoSnapshots, 50)
}
