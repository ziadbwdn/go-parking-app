package test

import (
	"go-parking-app/internal/domain/models"
	"go-parking-app/internal/infrastructure/persistence/memory"
	"testing"
	"time"
)

func TestInMemoryParkingRepository(t *testing.T) {
	repo := memory.NewInMemoryParkingRepository()

	ticket := models.Ticket{
		SlotNumber: 1,
		Car:        models.Car{RegistrationNumber: "KA-01-HH-1234"},
		EntryTime:  time.Now(),
	}

	// Save
	err := repo.SaveTicket(ticket)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Find
	found, err := repo.FindTicketByRegistration("KA-01-HH-1234")
	if err != nil {
		t.Fatalf("expected ticket to be found, got error: %v", err)
	}
	if found.Car.RegistrationNumber != "KA-01-HH-1234" {
		t.Errorf("expected registration number to match, got %s", found.Car.RegistrationNumber)
	}

	// Delete
	err = repo.DeleteTicket("KA-01-HH-1234")
	if err != nil {
		t.Fatalf("expected delete to succeed, got error: %v", err)
	}

	// Ensure deletion
	_, err = repo.FindTicketByRegistration("KA-01-HH-1234")
	if err == nil {
		t.Errorf("expected error after deletion, got none")
	}
}
