package models

type P2PRequest struct {
	Fiat                      string   `json:"fiat"`
	Page                      int      `json:"page"`
	Rows                      int      `json:"rows"`
	TransAmount               uint64   `json:"transAmount"`
	TradeType                 string   `json:"tradeType"`
	Asset                     string   `json:"asset"`
	Countries                 []string `json:"countries"`
	ProMerchantAds            bool     `json:"proMerchantAds"`
	ShieldMerchantAds         bool     `json:"shieldMerchantAds"`
	FilterType                string   `json:"filterType"`
	Periods                   []string `json:"periods"`
	AdditionalKycVerifyFilter int      `json:"additionalKycVerifyFilter"`
	PublisherType             *string  `json:"publisherType"`
	PayTypes                  []string `json:"payTypes"`
	Classifies                []string `json:"classifies"`
}

type GetBinanceMarketRateResponse struct {
	BinanceRates string `json:"binance_rate"`
}

type GetBinanceMarketRateResponseNew struct {
	BinanceRates []string `json:"binance_rate"`
}

type Adv struct {
	Price string `json:"price"`
}

type Data struct {
	Adv  Adv        `json:"adv"`
	User Advertiser `json:"advertiser"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

type Advertiser struct {
	NickName string `json:"nickName"`
}
