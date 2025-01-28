package useCase

import (
	"context"
	"exchanger-parser/internal/models"
	"exchanger-parser/internal/repository/repository"
	"log"
	"sync"
	"time"
)

type UseCase interface {
	Execute(ctx context.Context, wg *sync.WaitGroup)
}

type useCase struct {
	repository   repository.Repository
	logger       *log.Logger
	insertChan   chan models.Chanel
	batch        []models.Chanel
	batchSize    int
	insertTicker *time.Ticker
}

func NewUseCase(
	repository repository.Repository,
	logger *log.Logger,
	insertChan chan models.Chanel,
	batchSize int,
	insertTicker time.Duration,
) UseCase {
	return &useCase{
		repository:   repository,
		logger:       logger,
		insertChan:   insertChan,
		batch:        make([]models.Chanel, batchSize),
		insertTicker: time.NewTicker(insertTicker),
	}
}

func (u *useCase) Execute(
	ctx context.Context,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			if err := u.repository.Insert(
				context.Background(),
				u.batch[0:u.batchSize],
			); err != nil {
				u.logger.Println(err)
			}

			batch := make([]models.Chanel, 0, len(u.insertChan))

			for item := range u.insertChan {
				batch = append(batch, item)
			}

			if err := u.repository.Insert(
				context.Background(),
				u.batch[0:u.batchSize],
			); err != nil {
				u.logger.Println(err)
			}

			u.logger.Println("insert gracefully shutdown done")
			return
		default:
			select {
			case <-u.insertTicker.C:
				if err := u.repository.Insert(
					context.Background(), //todo: thinking
					u.batch[0:u.batchSize],
				); err != nil {
					u.logger.Println(err)
				}
				u.batchSize = 0
			case item := <-u.insertChan:
				u.batch[u.batchSize] = item
				u.batchSize++

				if u.batchSize >= cap(u.batch) {
					if err := u.repository.Insert(
						context.Background(), //todo: thinking
						u.batch[0:u.batchSize],
					); err != nil {
						u.logger.Println(err)
					}
					u.batchSize = 0
				}
			}
		}
	}
}
