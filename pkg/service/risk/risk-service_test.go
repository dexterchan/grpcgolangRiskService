package risk

import (
	"context"
	"math"
	"testing"
	"time"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/risk"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	numOfTenor := 10
	ctx1 := context.Background()
	helper := NewRiskServiceServer(numOfTenor)

	tests := []struct {
		name    string
		ctx     context.Context
		req     risk.HealthCheckRequest
		wantErr bool
	}{
		{
			name: "basic",
			ctx:  ctx1,
			req: risk.HealthCheckRequest{
				Service: "Testing",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := helper.Check(tt.ctx, &tt.req)

			//Check result
			if (err != nil) != tt.wantErr {
				t.Errorf("Health Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, res.GetStatus(), risk.HealthCheckResponse_SERVING, "Status should be ok")
		})
	}

}

func TestCalculateRisk(t *testing.T) {
	numOfTenor := 10
	ctx1 := context.Background()
	helper := NewRiskServiceServer(numOfTenor)
	tt := time.Now().In(time.UTC)
	timeStampProto, _ := ptypes.TimestampProto(tt)
	tests := []struct {
		name    string
		ctx     context.Context
		req     risk.ValueRequest
		wantErr bool
	}{
		{
			name: "basic",
			ctx:  ctx1,
			req: risk.ValueRequest{
				SystemDate:      timeStampProto,
				TradeId:         "A0001",
				TradeMessage:    "<trade></trade>",
				OutputType:      risk.ValueRequest_ALL,
				RunType:         risk.ValueRequest_FO,
				WantedRiskSense: make([]string, 0),
			},
			wantErr: false,
		},
		{
			name: "basic",
			ctx:  ctx1,
			req: risk.ValueRequest{
				SystemDate:      timeStampProto,
				TradeId:         "B0001",
				TradeMessage:    "<trade></trade>",
				OutputType:      risk.ValueRequest_ALL,
				RunType:         risk.ValueRequest_FO,
				WantedRiskSense: make([]string, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := helper.Calculate(tt.ctx, &tt.req)
			//Check result
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//Check each tenor
			for asset, assetSens := range res.AssetSensitivityLst {
				for _, riskSens := range assetSens.SenseLst {

					//Check Risk
					if len(riskSens.Tenors) != numOfTenor {
						t.Errorf("CalculateRisk() Risk tenor not match=%v, wanted %v", len(riskSens.Tenors), numOfTenor)
					}
					for i, tenor := range riskSens.Tenors {
						//expected value
						expectedValue := float64(i+1) * 1000.00 * math.Pow(-1.0, float64(asset))
						assert.InDelta(t, expectedValue, tenor.GetValue(), 0.0000001)
					}
				}

			}
		})
	}

}
