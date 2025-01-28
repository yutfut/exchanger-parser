package main

import (
	"context"
	"errors"
	"exchanger-parser/config"
	"exchanger-parser/internal/api/handler/grpc"
	handler "exchanger-parser/internal/api/handler/http"
	apiUseCase "exchanger-parser/internal/api/useCase"
	binanceAPI "exchanger-parser/internal/binance/api"
	binanceUseCase "exchanger-parser/internal/binance/useCase"
	bybitAPI "exchanger-parser/internal/bybit/api"
	bybitUseCase "exchanger-parser/internal/bybit/useCase"
	garantexAPI "exchanger-parser/internal/garantex/api"
	garantexUseCase "exchanger-parser/internal/garantex/useCase"
	"exchanger-parser/internal/models"
	"exchanger-parser/internal/redis/api"
	"exchanger-parser/internal/redis/statistic"
	"exchanger-parser/internal/repository/repository"
	"exchanger-parser/internal/repository/useCase"
	statisticUseCase "exchanger-parser/internal/statistic/useCase"
	"exchanger-parser/pkg/clickhouse"
	"exchanger-parser/pkg/redis"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	logger := log.New(
		os.Stdout,
		"",
		log.Ldate|log.Ltime|log.Llongfile,
	)

	logger.Println("starting exchanger-parser")

	secret, err := config.LoadConfig("./config/config.json")
	if err != nil {
		logger.Fatalln(err)
	}

	insertChan := make(chan models.Chanel, 10000)

	clickhouseConn, err := clickhouse.Connect(
		clickhouse.ClickHouse{
			Host:     secret.ClickHouse.Host,
			Port:     secret.ClickHouse.Port,
			Database: secret.ClickHouse.Database,
			Username: secret.ClickHouse.Username,
			Password: secret.ClickHouse.Password,
			Debug:    secret.ClickHouse.Debug,
		},
	)
	if err != nil {
		logger.Fatal(err)
	}

	redisConn, err := redis.Connect(ctx, secret.Redis)
	if err != nil {
		log.Fatal(err)
	}

	clickhouseRepository := repository.NewRepository(
		clickhouseConn,
		logger,
	)

	repUseCase := useCase.NewUseCase(
		clickhouseRepository,
		logger,
		insertChan,
		1000,
		time.Minute,
	)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go repUseCase.Execute(ctx, wg)

	garantexAPIInterface := garantexAPI.NewAPI(false)
	garantexUseCaseInterface := garantexUseCase.NewUseCase(garantexAPIInterface, insertChan, logger, 5*time.Second)

	wg.Add(1)
	go garantexUseCaseInterface.Parse(ctx, wg)

	bybitAPIInterface := bybitAPI.NewAPI(false)
	bybitUseCaseInterface := bybitUseCase.NewUseCase(bybitAPIInterface, insertChan, logger, 5*time.Second)

	wg.Add(1)
	go bybitUseCaseInterface.Parse(ctx, wg)

	binanceAPIInterface := binanceAPI.NewAPI(false)
	binanceUseCaseInterface := binanceUseCase.NewUseCase(binanceAPIInterface, insertChan, logger, 5*time.Second)

	wg.Add(1)
	go binanceUseCaseInterface.Parse(ctx, wg)

	statisticRepository := statistic.NewRepository(redisConn, logger)

	stat := repository.NewStatisticRepository(clickhouseConn, logger)
	statUC := statisticUseCase.NewUseCase(stat, statisticRepository, logger)

	wg.Add(1)
	go statUC.DoSomething(ctx)

	router := chi.NewRouter()

	apiRepository := api.NewRepository(redisConn, logger)

	apiUseCaseInterface := apiUseCase.NewUseCase(apiRepository, logger)

	handler.NewRouting(router, handler.NewHandler(apiUseCaseInterface, logger))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "0.0.0.0", "11051"),
		Handler: router,
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	grpcServer := grpc.NewGrpcServer(apiUseCaseInterface, logger)
	grpcServer.Run(logger)

	log.Printf("Server is running on %v\n", server.Addr)

	<-ctx.Done()
	stop()

	wg.Wait()

	logger.Println("main done")
}
