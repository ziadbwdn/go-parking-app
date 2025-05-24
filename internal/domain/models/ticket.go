package models

import "time"

type Ticket struct {
	SlotNumber int
	Car        Car
	EntryTime  time.Time
}

// NewTicket creates a new ticket with current timestamp.
func NewTicket(slotNumber int, car Car) Ticket {
	return Ticket{
		SlotNumber: slotNumber,
		Car:        car,
		EntryTime:  time.Now(),
	}
}
