package useCase

import (
	"context"
	"exchanger-parser/internal/binance/api"
	"exchanger-parser/internal/binance/models"
	m "exchanger-parser/internal/models"
	"github.com/shopspring/decimal"
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

func (u *useCase) calcUZS(
	ctx context.Context,
	request models.P2PRequest,
) {
	if response, err := u.api.Parse(
		ctx,
		request,
	); err == nil {
		if len(response.Data) < 10 {
			u.logger.Println(err)
			return
		}

		var i uint16 = 0
		for ; i < 10; i++ {
			if rate, parseErr := strconv.ParseFloat(
				response.Data[i].Adv.Price,
				64,
			); parseErr == nil {
				u.chanel <- m.Chanel{
					Exchanger:             3,
					ExchangersConditionID: i + 1,
					Course:                rate,
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

func (u *useCase) calcTRY(
	ctx context.Context,
	request models.P2PRequest,
) {
	if response, err := u.api.Parse(
		ctx,
		request,
	); err == nil {
		if len(response.Data) < 2 {
			u.logger.Println(err)
		} else {
			if rate, parseErr := strconv.ParseFloat(
				response.Data[1].Adv.Price,
				64,
			); parseErr == nil {
				u.chanel <- m.Chanel{
					Exchanger:             3,
					ExchangersConditionID: 12,
					Course:                rate,
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

func (u *useCase) calcAZN(
	ctx context.Context,
	request models.P2PRequest,
) {
	if response, err := u.api.Parse(
		ctx,
		request,
	); err == nil {
		if len(response.Data) < 10 {
			u.logger.Println(err)
			return
		}

		total := decimal.Zero

		for i := 0; i < 10; i++ {
			if parseRate, parseErr := decimal.NewFromString(
				response.Data[i].Adv.Price,
			); parseErr == nil {
				total = total.Add(parseRate)
			} else {
				u.logger.Println(parseErr)
				return
			}
		}

		rate, _ := total.Div(decimal.NewFromInt(10)).Round(2).Float64()

		u.chanel <- m.Chanel{
			Exchanger:             3,
			ExchangersConditionID: 11,
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
				go u.calcUZS(
					context.Background(),
					models.P2PRequest{
						Fiat:                      "UZS",
						Page:                      1,
						Rows:                      10,
						TransAmount:               1000000,
						TradeType:                 "BUY",
						Asset:                     "USDT",
						FilterType:                "all",
						AdditionalKycVerifyFilter: 0,
						PayTypes:                  []string{"Humo"},
						Classifies:                []string{"mass", "profession"},
					},
				)

				go u.calcAZN(
					context.Background(),
					models.P2PRequest{
						Fiat:                      "AZN",
						Page:                      1,
						Rows:                      10,
						TradeType:                 "BUY",
						Asset:                     "USDT",
						FilterType:                "all",
						AdditionalKycVerifyFilter: 1,
						PayTypes:                  []string{},
						Countries:                 []string{"AZ"},
					},
				)

				go u.calcTRY(
					context.Background(),
					models.P2PRequest{
						Fiat:                      "TRY",
						Page:                      1,
						Rows:                      10,
						TradeType:                 "BUY",
						TransAmount:               5000,
						Asset:                     "USDT",
						FilterType:                "all",
						AdditionalKycVerifyFilter: 0,
						PayTypes:                  []string{"Papara"},
						Countries:                 []string{"TR"},
					},
				)
			}
		}
	}
}
