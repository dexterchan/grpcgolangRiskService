package rest

import (
	"context"
	//"log"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	v1 "github.com/dexter/grpcRiskStandalone/pkg/api/v1"
	"github.com/dexter/grpcRiskStandalone/pkg/logger"
	"github.com/dexter/grpcRiskStandalone/pkg/protocol/rest/middleware"
)

// RunRestServer runs HTTP/REST gateway
func RunRestServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		//log.Fatalf("failed to start HTTP gateway: %v", err)
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}

	srv := &http.Server{
		Addr: ":" + httpPort,
		//Handler: mux,
		Handler: middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	//log.Println("starting HTTP/REST gateway...")
	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
