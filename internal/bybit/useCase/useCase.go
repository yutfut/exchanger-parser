package useCase

import (
	"context"
	"exchanger-parser/internal/bybit/api"
	"exchanger-parser/internal/bybit/models"
	m "exchanger-parser/internal/models"
	"github.com/shopspring/decimal"
	"log"
	"sync"
	"time"
)

type UseCase interface {
	Parse(ctx context.Context, wg *sync.WaitGroup)
}

type useCase struct {
	api           api.API
	chanel        chan m.Chanel
	logger        *log.Logger
	requestTicker *time.Ticker
}

func NewUseCase(
	api api.API,
	chanel chan m.Chanel,
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

func (u *useCase) calc(
	ctx context.Context,
	request models.P2PCourseParams,
	exchangersConditionID uint16,
) {
	if response, err := u.api.Parse(
		ctx,
		request,
	); err == nil {
		if len(response.Result.Items) < 10 {
			u.logger.Println(err)
		}

		total := decimal.Zero

		for i, item := range response.Result.Items {
			if i >= 6 {
				break
			}
			if parseRate, parseErr := decimal.NewFromString(item.Price); parseErr == nil {
				total = total.Add(parseRate)
			} else {
				u.logger.Println(err)
			}

		}

		rate, _ := total.Div(decimal.NewFromInt(6)).Round(2).Float64()

		u.chanel <- m.Chanel{
			Exchanger:             2,
			ExchangersConditionID: exchangersConditionID,
			Course:                rate,
			Time:                  time.Now(),
		}
	} else {
		u.logger.Println(err)
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
				go u.calc(
					context.Background(),
					models.P2PCourseParams{
						TokenId:    "USDT",
						CurrencyId: "RUB",
						Payment:    []string{"582"},
						Side:       "1",
						Size:       "50",
						Page:       "1",
						AuthMaker:  false,
						CanTrade:   false,
						Amount:     "100000",
					},
					1,
				)

				go u.calc(
					context.Background(),
					models.P2PCourseParams{
						TokenId:    "USDT",
						CurrencyId: "RUB",
						Payment:    []string{"581"},
						Side:       "1",
						Size:       "50",
						Page:       "1",
						AuthMaker:  false,
						CanTrade:   false,
						Amount:     "100000",
					},
					2,
				)

				go u.calc(
					context.Background(),
					models.P2PCourseParams{
						TokenId:    "USDT",
						CurrencyId: "RUB",
						Payment:    []string{"585"},
						Side:       "1",
						Size:       "50",
						Page:       "1",
						AuthMaker:  false,
						CanTrade:   false,
						Amount:     "100000",
					},
					3,
				)
			}
		}
	}
}
