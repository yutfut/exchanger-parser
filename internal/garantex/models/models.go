package models

type GreenUsdtRub struct {
	Timestamp int   `json:"timestamp"`
	Asks      []Ask `json:"asks"`
	Bids      []Bid `json:"bids"`
}

type Ask struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}

type Bid struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}
