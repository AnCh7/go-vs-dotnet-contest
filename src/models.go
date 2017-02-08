package main

type PriceBar struct {
	TickDate       string `json:"date"`
	TickYear       string `json:"year"`
	MarketID       string `json:"marketId"`
	Open           string `json:"open_price"`
	Close          string `json:"close_price"`
	High           string `json:"high_price"`
	Low            string `json:"low_price"`
	OpenTickdate   string `json:"open_date"`
	OpenVersionNo  string `json:"open_version_no"`
	CloseTickDate  string `json:"close_date"`
	CloseVersionNo string `json:"close_version_no"`
	Spike          string `json:"spike"`
	Gap            string `json:"gap"`
	OpenBid        string `json:"open_bid_price"`
	OpenAsk        string `json:"open_ask_price"`
	HighBid        string `json:"high_bid_price"`
	HighAsk        string `json:"high_ask_price"`
	LowBid         string `json:"low_bid_price"`
	LowAsk         string `json:"low_ask_price"`
	CloseBid       string `json:"close_bid_price"`
	CloseAsk       string `json:"close_ask_price"`
}
