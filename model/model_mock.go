package model

import "time"

func GatewayResponseMock() GatewayResponse {
	now := time.Now().UTC()
	return GatewayResponse{
		Data: []Data{
			{
				Name:   "Bitcoin",
				Symbol: "BTC",
				Quote: Quote{
					Usd: Usd{
						MarketCap:   1.1,
						Price:       2.2,
						LastUpdated: now,
					},
				},
			},
			{
				Name:   "Ethereum",
				Symbol: "ETH",
				Quote: Quote{
					Usd: Usd{
						MarketCap:   3.3,
						Price:       4.4,
						LastUpdated: now,
					},
				},
			},
		},
	}
}
