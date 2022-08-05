package model

import (
	"time"
)

type GatewayResponse struct {
	Status Status `json:"status"`
	Data   []Data `json:"data"`
}

type Status struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Notice       string      `json:"notice"`
	TotalCount   int         `json:"total_count"`
}

type Usd struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}

type Quote struct {
	Usd Usd `json:"USD"`
}

type Data struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	NumMarketPairs    int       `json:"num_market_pairs"`
	DateAdded         time.Time `json:"date_added"`
	Tags              []string  `json:"tags"`
	MaxSupply         int       `json:"max_supply"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	CmcRank           int       `json:"cmc_rank"`
	LastUpdated       time.Time `json:"last_updated"`
	Quote             Quote     `json:"quote"`
}

type CryptoSnapshot struct {
	ID        int    `gorm:"primaryKey,autoIncrement"`
	Symbol    string `gorm:"index:idx_symbol"`
	Name      string `gorm:"index:idx_name"`
	MarketCap float64
	USDPrice  float64
	UpdatedAt time.Time
}

func (gr GatewayResponse) ToCryptoSnapshot() []CryptoSnapshot {
	coins := make([]CryptoSnapshot, 0, len(gr.Data))
	for _, data := range gr.Data {
		coin := CryptoSnapshot{
			Symbol:    data.Symbol,
			Name:      data.Name,
			MarketCap: data.Quote.Usd.MarketCap,
			USDPrice:  data.Quote.Usd.Price,
			UpdatedAt: data.Quote.Usd.LastUpdated,
		}
		coins = append(coins, coin)
	}
	return coins
}
