package risk

import (
	"context"
	"fmt"
	"math"
	"time"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/risk"
)

const fakeTenor = 10
const numOfAsset = 2
const numOfRiskSense = 2

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
func (fake FakeRiskCalcHelper) Calculate(c context.Context, req *risk.ValueRequest) (*risk.ValueResponse, error) {
	timestampStr := fmt.Sprintf("%v", time.Now())
	res := risk.ValueResponse{
		Status:              risk.ValueResponse_SUCCESS,
		Time:                timestampStr,
		TradeId:             "xxx",
		AssetSensitivityLst: make([]*risk.ValueResponse_AssetSensivity, 0),
	}

	//Fill in asset
	for asset := 0; asset < numOfAsset; asset++ {
		assetBuilder := risk.ValueResponse_AssetSensivity{
			AssetId:  "asset" + string(asset),
			Ccy:      "USD",
			Npv:      math.Pow(-1.0, float64(asset)) * 1000.0,
			SenseLst: make([]*risk.ValueResponse_Sensitivity, 0),
		}
		for i := 0; i < numOfRiskSense; i++ {
			sens := risk.ValueResponse_Sensitivity{
				RiskLabel: "risk" + string(0),
				Ccy:       "USD",
				Tenors:    make([]*risk.ValueResponse_Sensitivity_Tenor, 0),
			}
			for j := 0; j < fakeTenor; j++ {
				tenor := risk.ValueResponse_Sensitivity_Tenor{
					Label: "tenor" + string(j),
					Value: math.Pow(-1.0, float64(asset)) * 1000.0 * float64(j+1),
				}
				sens.Tenors = append(sens.Tenors, &tenor)
			}
			assetBuilder.SenseLst = append(assetBuilder.SenseLst, &sens)
		}
		res.AssetSensitivityLst = append(res.AssetSensitivityLst, &assetBuilder)
	}
	return &res, nil
}
