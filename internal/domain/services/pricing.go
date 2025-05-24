package services

// PricingService defines how to calculate parking charges.
type PricingService interface {
	CalculateCharge(hours int) int
}

// simplePricingService is a concrete implementation of PricingService.
type simplePricingService struct{}

// NewSimplePricingService returns a new instance of PricingService.
func NewSimplePricingService() PricingService {
	return &simplePricingService{}
}

// CalculateCharge calculates the total cost based on parking duration.
// $10 for the first 2 hours, $10 for each additional hour.
func (s *simplePricingService) CalculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}