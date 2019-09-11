package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/risk"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := risk.NewRiskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	healthCheck := risk.HealthCheckRequest{
		Service: "Testing",
	}
	res, _ := c.Check(ctx, &healthCheck)
	if res.GetStatus() != risk.HealthCheckResponse_SERVING {
		log.Panic("Health check failed")
	}

	tt := time.Now().In(time.UTC)
	timeStampProto, _ := ptypes.TimestampProto(tt)

	req := risk.ValueRequest{
		SystemDate:      timeStampProto,
		TradeId:         "B0001",
		TradeMessage:    "<trade></trade>",
		OutputType:      risk.ValueRequest_ALL,
		RunType:         risk.ValueRequest_FO,
		WantedRiskSense: make([]string, 0),
	}
	res1, err := c.Calculate(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Output is %v:", res1)
}
