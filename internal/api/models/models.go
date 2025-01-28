package models

type Request struct {
	Exchanger             uint8  `json:"exchanger"`
	ExchangersConditionID uint16 `json:"exchangers_condition_id"`
}
