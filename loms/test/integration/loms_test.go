//go:build integration

package integration

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/grpc/interceptor"
	loms_grpc "github.com/CatMacales/route256/loms/internal/grpc/loms"
	"github.com/CatMacales/route256/loms/internal/repository/memory/order"
	"github.com/CatMacales/route256/loms/internal/repository/memory/stock"
	"github.com/CatMacales/route256/loms/internal/service/loms"
	desc "github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const (
	stocksInitPath = "../../stock-data.json"
	bufSize        = 1024 * 1024
)

func TestIntegration_Run(t *testing.T) {
	suite.Run(t, new(LOMSSuite))
}

type LOMSSuite struct {
	suite.Suite
	client    desc.LOMSClient
	server    *grpc.Server
	listener  net.Listener
	stockRepo *stock_repository.Repository
	orderRepo *order_repository.Repository
}
type Config struct {
	Grpc GRPCConfig
}

type GRPCConfig struct {
	Host string
	Port uint32
}

func (ls *LOMSSuite) SetupSuite() {
	orderRepository := order_repository.NewRepository()
	stockRepository := stock_repository.NewRepository()

	lomsService := loms.NewService(orderRepository, stockRepository)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.Logger,
			interceptor.Panic,
			interceptor.Validate,
		),
	)

	loms_grpc.RegisterServer(grpcServer, lomsService)

	l := bufconn.Listen(bufSize)

	go func() {
		if err := grpcServer.Serve(l); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return l.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		ls.T().Fatalf("grpc server connection failed: %v", err)
	}

	ls.client = desc.NewLOMSClient(conn)
	ls.stockRepo = stockRepository
	ls.orderRepo = orderRepository
	ls.listener = l
	ls.server = grpcServer
}

func (ls *LOMSSuite) TearDownSuite() {
	ls.server.Stop()
	err := ls.listener.Close()
	if err != nil {
		log.Printf("error closing listener: %v", err)
	}
}

func (ls *LOMSSuite) TearDownTest() {
	ls.stockRepo.Clear(context.Background())
	ls.orderRepo.Clear(context.Background())
}
