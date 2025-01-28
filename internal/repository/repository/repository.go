package repository

import (
	"context"
	"log"

	"exchanger-parser/internal/models"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Repository interface {
	Insert(ctx context.Context, request []models.Chanel) error
}

type repository struct {
	driver driver.Conn
	logger *log.Logger
}

func NewRepository(
	driver driver.Conn,
	logger *log.Logger,
) Repository {
	return &repository{
		driver: driver,
		logger: logger,
	}
}

func (r *repository) Insert(
	ctx context.Context,
	request []models.Chanel,
) error {
	batch, err := r.driver.PrepareBatch(
		ctx,
		InsertBatch,
	)
	if err != nil {
		r.logger.Println(err)
		return err
	}

	for _, item := range request {
		if err = batch.AppendStruct(
			&item,
		); err != nil {
			r.logger.Println(err)
			return err
		}
	}

	if err = batch.Send(); err != nil {
		r.logger.Println(err)
		return err
	}

	return nil
}
