package test

import (
	"go-parking-app/internal/domain/services"
	"go-parking-app/internal/infrastructure/persistence/memory"
	"testing"
)

func TestParkingServiceIntegration(t *testing.T) {
	repo := memory.NewInMemoryParkingRepository()
	pricing := services.NewSimplePricingService()
	service := services.NewDefaultParkingService(repo, pricing)

	err := service.CreateParkingLot(2)
	if err != nil {
		t.Fatalf("unexpected error when creating parking lot: %v", err)
	}

	// Park car 1
	slot1, err := service.ParkCar("KA-01-HH-1234")
	if err != nil {
		t.Fatalf("failed to park car 1: %v", err)
	}
	if slot1 != 1 {
		t.Errorf("expected slot 1, got %d", slot1)
	}

	// Park car 2
	slot2, err := service.ParkCar("KA-01-HH-9999")
	if err != nil {
		t.Fatalf("failed to park car 2: %v", err)
	}
	if slot2 != 2 {
		t.Errorf("expected slot 2, got %d", slot2)
	}

	// Try parking when full
	_, err = service.ParkCar("KA-01-BB-0001")
	if err == nil {
		t.Error("expected error when parking in full lot, got nil")
	}

	// Leave car 1
	leftSlot, charge, err := service.LeaveParkingLot("KA-01-HH-1234", 4)
	if err != nil {
		t.Fatalf("error during leave: %v", err)
	}
	if leftSlot != 1 {
		t.Errorf("expected slot 1 to be freed, got %d", leftSlot)
	}
	if charge != 30 {
		t.Errorf("expected charge 30, got %d", charge)
	}

	// Park again
	slot3, err := service.ParkCar("KA-01-BB-0001")
	if err != nil {
		t.Fatalf("failed to park after leave: %v", err)
	}
	if slot3 != 1 {
		t.Errorf("expected slot 1 to be reused, got %d", slot3)
	}

	// Status check
	status := service.GetStatus()
	if len(status) != 2 {
		t.Errorf("expected 2 cars in status, got %d", len(status))
	}
}
