package commandline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"go-parking-app/pkg/utils"
	"go-parking-app/internal/usecases"
)

type CommandHandler struct {
	interactor usecases.ParkingInteractor
}

func NewCommandHandler(interactor usecases.ParkingInteractor) *CommandHandler {
	return &CommandHandler{
		interactor: interactor,
	}
}

func (h *CommandHandler) Process(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		h.handleCommand(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}

func (h *CommandHandler) handleCommand(commandLine string) {
	parts := strings.Fields(commandLine)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create_parking_lot":
		if len(parts) < 2 {
			fmt.Println("Invalid create_parking_lot command")
			return
		}
		capacity, _ := strconv.Atoi(parts[1])
		err := h.interactor.CreateParkingLot(capacity)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Created parking lot with %d slots\n", capacity)

	case "park":
		if len(parts) < 2 {
			fmt.Println("Invalid park command")
			return
		}
		slot, err := h.interactor.ParkCar(parts[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Allocated slot number: %d\n", slot)

	case "leave":
		if len(parts) < 3 {
			fmt.Println("Invalid leave command")
			return
		}
		hours, _ := strconv.Atoi(parts[2])
		slot, charge, err := h.interactor.LeaveParkingLot(parts[1], hours)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Registration number %s with Slot Number %d is free with Charge %d\n", parts[1], slot, charge)

	case "status":
		status := h.interactor.GetStatus()
		fmt.Println("Slot No.\tRegistration No.")
		for _, s := range status {
			fmt.Printf("%d\t%s\n", s.SlotNumber, s.RegistrationNumber)
		}

	default:
		fmt.Println("Unknown command:", parts[0])
	}
}