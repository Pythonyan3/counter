package main

import (
	"log"

	"github.com/Pythonyan3/Counter/internal/app"
)

// main service entrypoint.
func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("app.Run: %v", err)
	}
}
