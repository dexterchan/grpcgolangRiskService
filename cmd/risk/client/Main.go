package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/riskservice"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := risk.RiskRequest{
		SystemDate:   "20180101",
		TradeId:      "B0001",
		TradeMessage: "<trade></trade>",
	}
	res1, err := c.CalculateRisk(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Output is %v:", res1)
}
