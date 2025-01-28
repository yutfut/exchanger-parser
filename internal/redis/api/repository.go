package api

import (
	"context"
	"exchanger-parser/internal/api/models"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type Repository interface {
	Get(ctx context.Context, request models.Request) ([]byte, error)
}

type repository struct {
	redis  *redis.Client
	logger *log.Logger
}

func NewRepository(
	redis *redis.Client,
	logger *log.Logger,
) Repository {
	return &repository{
		redis:  redis,
		logger: logger,
	}
}

func (r *repository) Get(
	ctx context.Context,
	request models.Request,
) (
	[]byte,
	error,
) {
	response, err := r.redis.Get(
		ctx,
		fmt.Sprintf(
			"%d:%d",
			request.Exchanger,
			request.ExchangersConditionID,
		),
	).Bytes()
	if err != nil {
		r.logger.Println(err)
		return nil, err
	}

	return response, nil
}
