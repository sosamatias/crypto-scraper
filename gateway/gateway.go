package gateway

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sosamatias/crypto-scraper/model"
)

type Gateway interface {
	List(Sort) (*model.GatewayResponse, error)
}

// Sort type API order
type Sort string

// SortMarketCap CoinMarketCap's market cap rank
const SortMarketCap Sort = "market_cap"

// SortDateAdded date cryptocurrency was added to the system
const SortDateAdded Sort = "date_added"

// SortCirculatingSupply approximate number of coins currently in circulation
const SortCirculatingSupply Sort = "circulating_supply"

// SortTotalSupply approximate total amount of coins in existence right now
const SortTotalSupply Sort = "total_supply"

// SortVolume24h rolling 24 hour adjusted trading volume
const SortVolume24h Sort = "volume_24h"

// ErrAPIKeyRequired api key required
var ErrAPIKeyRequired = errors.New("api key required")

const timeout = time.Duration(4) * time.Second
const defaultBaseUrl = "https://pro-api.coinmarketcap.com"

func NewGateway(apiKey, baseUrl string, client *http.Client) (Gateway, error) {
	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}
	if baseUrl == "" {
		baseUrl = defaultBaseUrl
	}
	if client == nil {
		client = &http.Client{}
	}
	client.Timeout = timeout
	g := gateway{
		client:     client,
		baseUrl:    baseUrl,
		apiKey:     apiKey,
		logEnabled: false,
	}
	return g, nil
}

type gateway struct {
	client     *http.Client
	baseUrl    string
	apiKey     string
	logEnabled bool
}

// List cryptocurrencies with latest market data
func (g gateway) List(sort Sort) (*model.GatewayResponse, error) {
	const endpoint = "/v1/cryptocurrency/listings/latest"

	req, err := http.NewRequest(http.MethodGet, g.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("sort", string(sort))

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", g.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := model.GatewayResponse{}
	err = json.Unmarshal(data, &response)
	return &response, err
}
