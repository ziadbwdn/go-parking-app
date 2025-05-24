package models

type Slot struct {
	Number     int
	IsOccupied bool
	Car        *Car
}

func (s *Slot) ParkCar(car *Car) {
	s.Car = car
	s.IsOccupied = true
}

func (s *Slot) LeaveSlot() {
	s.Car = nil
	s.IsOccupied = false
}