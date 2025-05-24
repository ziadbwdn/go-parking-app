package memory

import (
	"errors"
	"go-parking-app/internal/domain/models"
	"sync"
)

// ParkingRepository defines the interface expected by ParkingService.
type ParkingRepository interface {
	SaveTicket(ticket models.Ticket) error
	FindTicketByRegistration(regNum string) (models.Ticket, error)
	DeleteTicket(regNum string) error
}

// inMemoryParkingRepository implements ParkingRepository.
type inMemoryParkingRepository struct {
	tickets map[string]models.Ticket
	mu      sync.RWMutex
}

// NewInMemoryParkingRepository is a constructor for in-memory storage.
func NewInMemoryParkingRepository() ParkingRepository {
	return &inMemoryParkingRepository{
		tickets: make(map[string]models.Ticket),
	}
}

func (r *inMemoryParkingRepository) SaveTicket(ticket models.Ticket) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tickets[ticket.Car.RegistrationNumber] = ticket
	return nil
}

func (r *inMemoryParkingRepository) FindTicketByRegistration(regNum string) (models.Ticket, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ticket, exists := r.tickets[regNum]
	if !exists {
		return models.Ticket{}, errors.New("ticket not found")
	}
	return ticket, nil
}

func (r *inMemoryParkingRepository) DeleteTicket(regNum string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tickets[regNum]; !ok {
		return errors.New("ticket not found")
	}
	delete(r.tickets, regNum)
	return nil
}