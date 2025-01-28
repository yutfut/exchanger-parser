package statistic

import (
	"context"
	"encoding/json"
	"exchanger-parser/internal/models"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type Repository interface {
	Set(ctx context.Context, request models.Chanel) error
	Get(ctx context.Context, request models.Chanel) (models.Chanel, error)
	Del(ctx context.Context, request models.Chanel) error
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

func (r *repository) Set(
	ctx context.Context,
	request models.Chanel,
) error {
	data, marshalError := json.Marshal(request)
	if marshalError != nil {
		r.logger.Println(marshalError)
	}

	if err := r.redis.Set(
		ctx,
		fmt.Sprintf(
			"%d:%d",
			request.Exchanger,
			request.ExchangersConditionID,
		),
		data,
		0,
	).Err(); err != nil {
		r.logger.Println(err)
		return err
	}

	return nil
}

func (r *repository) Get(
	ctx context.Context,
	request models.Chanel,
) (
	models.Chanel,
	error,
) {
	result, err := r.redis.Get(
		ctx,
		fmt.Sprintf(
			"%d:%d",
			request.Exchanger,
			request.ExchangersConditionID,
		),
	).Bytes()
	if err != nil {
		r.logger.Println(err)
		return models.Chanel{}, err
	}

	response := models.Chanel{}
	if err = json.Unmarshal(
		result,
		&response,
	); err != nil {
		r.logger.Println(err)
		return models.Chanel{}, err
	}

	return response, nil
}

func (r *repository) Del(
	ctx context.Context,
	request models.Chanel,
) error {
	if err := r.redis.Del(
		ctx,
		fmt.Sprintf(
			"%d:%d",
			request.Exchanger,
			request.ExchangersConditionID,
		),
	).Err(); err != nil {
		r.logger.Println(err)
		return err
	}

	return nil
}
