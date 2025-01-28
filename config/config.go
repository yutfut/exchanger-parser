package config

import (
	"encoding/json"
	"exchanger-parser/pkg/clickhouse"
	"exchanger-parser/pkg/redis"
	"os"
)

type Secret struct {
	ClickHouse clickhouse.ClickHouse `json:"clickHouse"`
	Redis      redis.Redis           `json:"redis"`
}

func LoadConfig(path string) (Secret, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Secret{}, err
	}

	response := Secret{}

	if err = json.Unmarshal(data, &response); err != nil {
		return Secret{}, err
	}

	return response, nil
}
