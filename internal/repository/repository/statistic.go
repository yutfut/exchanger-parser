package repository

import (
	"context"
	"exchanger-parser/internal/models"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"time"
)

type StatisticRepository interface {
	GetStatistic(ctx context.Context) ([]models.Chanel, error)
	GetStatisticByExchanger(ctx context.Context, exchanger uint8, exchangerID uint16, limit int) ([]models.Chanel, error)
	GetStatisticByExchangerByLastHourAVG(ctx context.Context, exchanger uint8, exchangerID uint16, time time.Time) (float64, error)
	GetStatisticByExchangerByLastHourMedian(ctx context.Context, exchanger uint8, exchangerID uint16, time time.Time) (float64, error)
	GetStatisticByAllExchangerAVG(ctx context.Context, exchanger uint8, exchangerID uint16, time time.Time) (float64, error)
	GetStatisticByAllExchangerMedian(ctx context.Context, exchanger uint8, exchangerID uint16, time time.Time) (float64, error)
}

type statisticRepository struct {
	driver driver.Conn
	logger *log.Logger
}

func NewStatisticRepository(
	driver driver.Conn,
	logger *log.Logger,
) StatisticRepository {
	return &statisticRepository{
		driver: driver,
		logger: logger,
	}
}

func (r *statisticRepository) GetStatistic(
	ctx context.Context,
) ([]models.Chanel, error) {
	response := make([]models.Chanel, 0, 35)

	err := r.driver.Select(
		ctx,
		&response,
		GetStatistic,
	)
	if err != nil {
		r.logger.Println(err)
		return nil, err
	}

	return response, nil
}

func (r *statisticRepository) GetStatisticByExchanger(
	ctx context.Context,
	exchanger uint8,
	exchangerID uint16,
	limit int,
) ([]models.Chanel, error) {
	response := make([]models.Chanel, 0, limit)

	err := r.driver.Select(
		ctx,
		&response,
		GetStatistic,
		exchanger,
		exchangerID,
		limit,
	)
	if err != nil {
		r.logger.Println(err)
		return nil, err
	}

	return response, nil
}

type Course struct {
	Course float64 `db:"course"`
}

func (r *statisticRepository) GetStatisticByExchangerByLastHourAVG(
	ctx context.Context,
	exchanger uint8,
	exchangerID uint16,
	time time.Time,
) (float64, error) {
	var response float64

	err := r.driver.QueryRow(
		ctx,
		GetStatisticByExchangerByLastHourAVG,
		exchanger,
		exchangerID,
		time,
	).Scan(&response)
	if err != nil {
		r.logger.Println(err)
		return response, err
	}

	return response, nil
}

func (r *statisticRepository) GetStatisticByExchangerByLastHourMedian(
	ctx context.Context,
	exchanger uint8,
	exchangerID uint16,
	time time.Time,
) (float64, error) {
	var response float64

	err := r.driver.QueryRow(
		ctx,
		GetStatisticByExchangerByLastHourMedian,
		exchanger,
		exchangerID,
		time,
	).Scan(&response)
	if err != nil {
		r.logger.Println(err)
		return response, err
	}

	return response, nil
}

func (r *statisticRepository) GetStatisticByAllExchangerAVG(
	ctx context.Context,
	exchanger uint8,
	exchangerID uint16,
	time time.Time,
) (float64, error) {
	var response float64

	err := r.driver.QueryRow(
		ctx,
		GetStatisticByAllExchangerAVG,
		exchanger,
		exchangerID,
		time,
	).Scan(&response)
	if err != nil {
		r.logger.Println(err)
		return response, err
	}

	return response, nil
}

func (r *statisticRepository) GetStatisticByAllExchangerMedian(
	ctx context.Context,
	exchanger uint8,
	exchangerID uint16,
	time time.Time,
) (float64, error) {
	var response float64

	err := r.driver.QueryRow(
		ctx,
		GetStatisticByAllExchangerMedian,
		exchanger,
		exchangerID,
		time,
	).Scan(&response)
	if err != nil {
		r.logger.Println(err)
		return response, err
	}

	return response, nil
}
