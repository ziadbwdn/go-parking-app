package main

import (
	"fmt"
	"os"
	"go-parking-app/internal/app"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: parking <input_file_path>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	application := app.NewApplication()
	if err := application.Run(filePath); err != nil {
		fmt.Printf("Execution error: %v\n", err)
	}
}
