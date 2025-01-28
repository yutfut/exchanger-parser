package grpc

import (
	"context"
	"encoding/json"
	"exchanger-parser/internal/api/models"
	"exchanger-parser/internal/api/useCase"
	m "exchanger-parser/internal/models"
	pb "exchanger-parser/pkg/pb/pb"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type Handler struct {
	useCase useCase.UseCase
	logger  *log.Logger
	pb.UnimplementedPriceParserServiceServer
}

func (h *Handler) GetRate(ctx context.Context, in *pb.GetRateRequest) (*pb.GetRateResponse, error) {
	var startRequest time.Time = time.Now()

	response, err := h.useCase.Get(
		ctx,
		models.Request{
			Exchanger:             uint8(in.Exchange),
			ExchangersConditionID: uint16(in.ExchangersConditional),
		},
	)
	if err != nil {
		h.logger.Printf("error getting rate: %v\n", err)
		return nil, err
	}
	h.logger.Println(fmt.Sprintf("grpc - GetRate [%s] processing time - [%d ms]", string(response), time.Since(startRequest).Milliseconds()))

	resp := m.Chanel{}
	if err = json.Unmarshal(response, &resp); err != nil {
		h.logger.Println("error unmarshalling response", zap.Error(err))
	}

	return &pb.GetRateResponse{Rate: fmt.Sprintf("%f", resp.Course)}, nil
}

type GRPCServer interface {
	Run(logger *log.Logger)
}

type grpcServer struct {
	useCase useCase.UseCase
	logger  *log.Logger
}

func NewGrpcServer(
	useCase useCase.UseCase,
	logger *log.Logger,
) GRPCServer {
	return &grpcServer{
		useCase: useCase,
		logger:  logger,
	}
}

func (a *grpcServer) Run(logger *log.Logger) {
	var err error

	listen, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", "9001"))
	if err != nil {
		logger.Println(fmt.Sprintf("failed to listen on port %s, error: %s", "9000", err.Error()))
	}

	opt := []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	}

	server := grpc.NewServer(opt...)
	reflection.Register(server)
	pb.RegisterPriceParserServiceServer(server, &Handler{useCase: a.useCase, logger: logger})

	if err = server.Serve(listen); err != nil {
		logger.Println(fmt.Sprintf("failed to serve grpc on port %s", "900"))
	}
}
