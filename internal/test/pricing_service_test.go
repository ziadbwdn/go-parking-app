// test/pricing_service_test.go
package test

import (
	"testing"

	"go-parking-app/internal/domain/services"
)

func TestCalculateCharge(t *testing.T) {
	service := services.NewSimplePricingService()
	tests := []struct {
		hours    int
		expected int
	}{
		{1, 10},
		{2, 10},
		{3, 20},
		{4, 30},
		{5, 40},
	}

	for _, tt := range tests {
		charge := service.CalculateCharge(tt.hours)
		if charge != tt.expected {
			t.Errorf("expected %d for %d hours, got %d", tt.expected, tt.hours, charge)
		}
	}
}
