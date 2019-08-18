package riskservice

import (
	"context"
	"fmt"
	"time"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/riskservice"
)

const fakeTenor = 10

//FakeRiskCalcHelper : Fake Risk calculation
type FakeRiskCalcHelper struct {
	numberOfTenor int
}

//NewRiskServiceServer : creates ToDo service
func NewRiskServiceServer(numTenor int) risk.RiskServiceServer {
	return &FakeRiskCalcHelper{numTenor}
}

//Check : for help check
func (fake FakeRiskCalcHelper) Check(c context.Context, req *risk.HealthCheckRequest) (*risk.HealthCheckResponse, error) {
	res := risk.HealthCheckResponse{
		Status: risk.HealthCheckResponse_SERVING,
	}

	return &res, nil
}

//CalculateRisk : for fake risk calculation
func (fake FakeRiskCalcHelper) CalculateRisk(c context.Context, req *risk.RiskRequest) (*risk.RiskResponse, error) {
	timestampStr := fmt.Sprintf("%v", time.Now())
	res := risk.RiskResponse{
		Status:  risk.RiskResponse_SUCCESS,
		Time:    timestampStr,
		TradeId: "xxx",
		Ccy:     "USD",
		Npv:     1000,
		Risk:    make([]*risk.RiskResponse_Tenor, 0),
	}
	for i := 0; i < fake.numberOfTenor; i++ {
		tenor := risk.RiskResponse_Tenor{
			Label: string(i),
			Value: float64(i+1) * 1000.00,
		}
		res.Risk = append(res.Risk, &tenor)
	}

	return &res, nil
}
