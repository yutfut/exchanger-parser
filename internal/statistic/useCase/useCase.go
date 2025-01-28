package useCase

import (
	"context"
	"exchanger-parser/internal/redis/statistic"
	"exchanger-parser/internal/repository/repository"
	"log"
	"time"
)

type UseCase interface {
	DoSomething(ctx context.Context)
}

type useCase struct {
	repository     repository.StatisticRepository
	statisticRedis statistic.Repository
	logger         *log.Logger
	ticker         *time.Ticker
}

func NewUseCase(
	repository repository.StatisticRepository,
	statisticRedis statistic.Repository,
	logger *log.Logger,
) UseCase {
	return &useCase{
		repository:     repository,
		statisticRedis: statisticRedis,
		logger:         logger,
		ticker:         time.NewTicker(5 * time.Second),
	}
}

func (u *useCase) DoSomething(
	ctx context.Context,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-u.ticker.C:
			response, err := u.repository.GetStatistic(ctx)
			if err != nil {
				return
			}

			for _, item := range response {
				checkTime := time.Now().Sub(item.Time)

				if time.Hour > checkTime && checkTime > time.Minute {
					course, errStat := u.repository.GetStatisticByExchangerByLastHourAVG(
						ctx,
						item.Exchanger,
						item.ExchangersConditionID,
						item.Time.Add(-time.Hour),
					)
					if errStat != nil {
						u.logger.Println(err)
					}

					item.Course = course
				} else if time.Hour < checkTime {
					course, errStat := u.repository.GetStatisticByAllExchangerAVG(
						ctx,
						item.Exchanger,
						item.ExchangersConditionID,
						item.Time,
					)
					if errStat != nil {
						u.logger.Println(err)
					}

					item.Course = course[0].Course
				}

				if err = u.statisticRedis.Set(
					ctx,
					item,
				); err != nil {
					u.logger.Println(err)
				}
			}
		}
	}
}
