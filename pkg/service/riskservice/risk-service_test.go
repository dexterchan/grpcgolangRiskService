package riskservice

import (
	"testing"

	risk "github.com/dexter/grpcRiskStandalone/pkg/api/riskservice"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRisk(t *testing.T) {
	helper := FakeRiskCalcHelper{
		10,
	}

	tests := []struct {
		name    string
		req     risk.RiskRequest
		wantErr bool
	}{
		{
			name: "basic",
			req: risk.RiskRequest{
				SystemDate:   "20190101",
				TradeId:      "A0001",
				TradeMessage: "<trade></trade>",
			},
			wantErr: false,
		},
		{
			name: "basic",
			req: risk.RiskRequest{
				SystemDate:   "20180101",
				TradeId:      "B0001",
				TradeMessage: "<trade></trade>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := helper.CalculateRisk(tt.req)
			//Check result
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateRisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//Check Risk
			if len(res.Risk) != helper.numberOfTenor {
				t.Errorf("CalculateRisk() Risk tenor not match=%v, wanted %v", len(res.Risk), helper.numberOfTenor)
			}
			//Check each tenor
			for i, tenor := range res.Risk {
				//expected value
				expectedValue := float64(i+1) * 1000.00
				assert.InDelta(t, expectedValue, tenor.GetValue(), 0.0000001)
			}
		})
	}

}