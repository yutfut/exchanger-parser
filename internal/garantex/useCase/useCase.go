package useCase

import (
	"context"
	"exchanger-parser/internal/garantex/api"
	"exchanger-parser/internal/models"
	"log"
	"strconv"
	"sync"
	"time"
)

type UseCase interface {
	Parse(ctx context.Context, wg *sync.WaitGroup)
}

type useCase struct {
	api           api.API
	chanel        chan models.Chanel
	logger        *log.Logger
	requestTicker *time.Ticker
}

func NewUseCase(
	api api.API,
	chanel chan models.Chanel,
	logger *log.Logger,
	requestTicker time.Duration,
) UseCase {
	return &useCase{
		api:           api,
		chanel:        chanel,
		logger:        logger,
		requestTicker: time.NewTicker(requestTicker),
	}
}

func (u *useCase) Parse(
	ctx context.Context,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			select {
			case <-u.requestTicker.C:
				if resp, err := u.api.Parse(
					context.Background(),
				); err == nil {
					var i uint16 = 0
					for ; i < 20; i++ {
						if course, parseErr := strconv.ParseFloat(
							resp.Bids[i].Price,
							64,
						); parseErr == nil {
							u.chanel <- models.Chanel{
								Exchanger:             1,
								ExchangersConditionID: i + 1,
								Course:                course,
								Time:                  time.Now(),
							}
						} else {
							u.logger.Println(parseErr)
						}
					}
				} else {
					u.logger.Println(err)
				}
			}
		}
	}
}
