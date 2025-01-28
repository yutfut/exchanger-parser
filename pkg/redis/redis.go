package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func Connect(
	ctx context.Context,
	config Redis,
) (*redis.Client, error) {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%s",
				config.Host,
				config.Port,
			),
			Password: config.Password,
			DB:       0,
		},
	)

	return rdb, rdb.Ping(ctx).Err()
}
