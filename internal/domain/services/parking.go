package services

import (
	"container/heap"
	"errors"
	"go-parking-app/internal/domain/models"
	"go-parking-app/internal/infrastructure/persistence/memory"
	"time"
)

type ParkingStatus struct {
	SlotNumber         int
	RegistrationNumber string
}

// ParkingService defines business rules interface.
type ParkingService interface {
	CreateParkingLot(capacity int) error
	ParkCar(registrationNumber string) (int, error)
	LeaveParkingLot(registrationNumber string, hours int) (int, int, error)
	GetStatus() []ParkingStatus
}

// slotHeap used for nearest-slot-first allocation.
type slotHeap []int

func (h slotHeap) Len() int            { return len(h) }
func (h slotHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h slotHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *slotHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *slotHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// defaultParkingService is the concrete implementation.
type defaultParkingService struct {
	slots        []models.Slot
	tickets      map[string]models.Ticket
	repo         memory.ParkingRepository
	pricing      PricingService
	availableMin *slotHeap
}

func NewDefaultParkingService(repo memory.ParkingRepository, pricing PricingService) ParkingService {
	h := &slotHeap{}
	heap.Init(h)
	return &defaultParkingService{
		slots:        []models.Slot{},
		tickets:      make(map[string]models.Ticket),
		repo:         repo,
		pricing:      pricing,
		availableMin: h,
	}
}

func (s *defaultParkingService) CreateParkingLot(capacity int) error {
	if capacity <= 0 {
		return errors.New("invalid capacity")
	}
	s.slots = make([]models.Slot, capacity)
	for i := 0; i < capacity; i++ {
		s.slots[i] = models.Slot{Number: i + 1}
		heap.Push(s.availableMin, i+1)
	}
	return nil
}

func (s *defaultParkingService) ParkCar(regNum string) (int, error) {
	if s.availableMin.Len() == 0 {
		return 0, errors.New("parking lot is full")
	}
	slotNum := heap.Pop(s.availableMin).(int)
	car := &models.Car{RegistrationNumber: regNum}
	s.slots[slotNum-1].IsOccupied = true
	s.slots[slotNum-1].Car = car
	ticket := models.Ticket{
		SlotNumber: slotNum,
		Car:        *car,
		EntryTime:  time.Now(),
	}
	s.repo.SaveTicket(ticket)
	s.tickets[regNum] = ticket
	return slotNum, nil
}

func (s *defaultParkingService) LeaveParkingLot(regNum string, hours int) (int, int, error) {
	ticket, err := s.repo.FindTicketByRegistration(regNum)
	if err != nil {
		return 0, 0, errors.New("ticket not found")
	}
	slot := ticket.SlotNumber
	s.slots[slot-1].IsOccupied = false
	s.slots[slot-1].Car = nil
	heap.Push(s.availableMin, slot)

	charge := s.pricing.CalculateCharge(hours)
	s.repo.DeleteTicket(regNum)
	delete(s.tickets, regNum)

	return slot, charge, nil
}

func (s *defaultParkingService) GetStatus() []ParkingStatus {
	status := []ParkingStatus{}
	for _, slot := range s.slots {
		if slot.IsOccupied && slot.Car != nil {
			status = append(status, ParkingStatus{
				SlotNumber:         slot.Number,
				RegistrationNumber: slot.Car.RegistrationNumber,
			})
		}
	}
	return status
}