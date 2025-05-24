package usecases

import (
	"go-parking-app/internal/domain/services"
)

// ParkingInteractor defines a contract for use cases that orchestrate parking operations.
type ParkingInteractor interface {
	CreateParkingLot(capacity int) error
	ParkCar(registrationNumber string) (int, error)
	LeaveParkingLot(registrationNumber string, hours int) (int, int, error)
	GetStatus() []services.ParkingStatus
}

// parkingUseCases provides a concrete implementation of the ParkingInteractor interface.
type parkingUseCases struct {
	parkingService services.ParkingService
}

// NewParkingUseCases constructs a ParkingInteractor using the provided ParkingService.
func NewParkingUseCases(parkingService services.ParkingService) ParkingInteractor {
	return &parkingUseCases{parkingService: parkingService}
}

func (u *parkingUseCases) CreateParkingLot(capacity int) error {
	return u.parkingService.CreateParkingLot(capacity)
}

func (u *parkingUseCases) ParkCar(registrationNumber string) (int, error) {
	return u.parkingService.ParkCar(registrationNumber)
}

func (u *parkingUseCases) LeaveParkingLot(registrationNumber string, hours int) (int, int, error) {
	return u.parkingService.LeaveParkingLot(registrationNumber, hours)
}

func (u *parkingUseCases) GetStatus() []services.ParkingStatus {
	return u.parkingService.GetStatus()
}