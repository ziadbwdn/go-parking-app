package app

import (
	"go-parking-app/internal/domain/services"
	"go-parking-app/internal/infrastructure/commandline"
	"go-parking-app/internal/infrastructure/persistence/memory"
	"go-parking-app/internal/usecases"
)

type Application struct {
	handler *commandline.CommandHandler
}

func NewApplication() *Application {
	pricing := services.NewSimplePricingService()
	repo := memory.NewInMemoryParkingRepository()
	parking := services.NewDefaultParkingService(repo, pricing)

	interactor := usecases.NewParkingUseCases(parking)
	handler := commandline.NewCommandHandler(interactor)

	return &Application{handler: handler}
}

// app.go
func (a *Application) Run(filePath string) error {
	return a.handler.Process(filePath)
}
