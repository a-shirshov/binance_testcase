package models

type MyStruct struct {
	Ticker string `json:"symbol"`
	Price  string `json:"openPrice"`
	Difference string `json:"priceChangePercent"`
}

type Symbols struct {
	Symbols []Pair `json:"symbols"`
}

type Pair struct {
	Symbol string `json:"symbol"`
}