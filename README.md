# Scraper
This scraper reads the CoinMarketCap API every 5 seconds and writes a snapshot to an auto-generated local SQLite database in the repository folder. This project uses gorm, you can change the driver to use a different database.

## Setup
1) Generate your API key in https://coinmarketcap.com/api
2) Update config/secrets.yaml with your new API key.
3) Install dependencies: `make install`
4) Run tests: `make test`
5) Execute the scraper: `make run`

## Mockgen
Here is an example to generate the mocks in case you modify the interfaces:
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go install github.com/golang/mock/mockgen@v1.6.0
mockgen -version 

mockgen -destination=mocks/gateway.go -package=mocks github.com/sosamatias/crypto-scraper/gateway Gateway 

mockgen -destination=mocks/repository.go -package=mocks github.com/sosamatias/crypto-scraper/repository Repository 
```
