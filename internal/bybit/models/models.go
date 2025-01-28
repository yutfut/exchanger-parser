package models

type P2PCourseParams struct {
	UserId     string   `json:"userId"`
	TokenId    string   `json:"tokenId"`
	CurrencyId string   `json:"currencyId"`
	Payment    []string `json:"payment"`
	Side       string   `json:"side"`
	Size       string   `json:"size"`
	Page       string   `json:"page"`
	Amount     string   `json:"amount"`
	AuthMaker  bool     `json:"authMaker"`
	CanTrade   bool     `json:"canTrade"`
}

type P2PCourseResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	Result  struct {
		Count int `json:"count"`
		Items []struct {
			Price     string `json:"price"`
			MinAmount string `json:"minAmount"`
			MaxAmount string `json:"maxAmount"`
		} `json:"items"`
	} `json:"result"`
}

type GetBybitMarketRateResponseNew struct {
	Rate float64 `json:"rate"`
}
