package useCase

import (
	"context"
	"exchanger-parser/internal/api/models"
	"exchanger-parser/internal/redis/api"
	"log"
)

type UseCase interface {
	Get(ctx context.Context, request models.Request) ([]byte, error)
}

type useCase struct {
	redisAPI api.Repository
	logger   *log.Logger
}

func NewUseCase(
	redisAPI api.Repository,
	logger *log.Logger,
) UseCase {
	return &useCase{
		redisAPI: redisAPI,
		logger:   logger,
	}
}

func (u *useCase) Get(
	ctx context.Context,
	request models.Request,
) (
	[]byte,
	error,
) {
	response, err := u.redisAPI.Get(
		ctx,
		request,
	)
	if err != nil {
		u.logger.Println(err)
	}
	return response, err
}
