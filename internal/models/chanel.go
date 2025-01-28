package models

import "time"

type Chanel struct {
	Exchanger             uint8     `ch:"exchanger" json:"exchanger"`
	ExchangersConditionID uint16    `ch:"exchangers_condition_id" json:"exchangers_condition_id"`
	Course                float64   `ch:"course" json:"course"`
	Time                  time.Time `ch:"time" json:"time"`
}
