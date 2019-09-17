package grpc

import (
	"context"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/risk"
	v1 "github.com/dexter/grpcRiskStandalone/pkg/api/v1"
	"github.com/dexter/grpcRiskStandalone/pkg/logger"
	"github.com/dexter/grpcRiskStandalone/pkg/protocol/grpc/middleware"
)

//RunRiskServerServer : Run risk-service runs
func RunRiskServerServer(ctx context.Context, riskAPI risk.RiskServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// register service
	server := grpc.NewServer()
	risk.RegisterRiskServiceServer(server, riskAPI)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			//log.Println("shutting down gRPC server...")
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	//log.Println("starting gRPC server...")
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterToDoServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			//log.Println("shutting down gRPC server...")
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	//log.Println("starting gRPC server...")
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
