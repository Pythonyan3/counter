package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Pythonyan3/Counter/internal/server"
)

// Run application entrypoint.
func Run() error {
	var (
		err         error
		httpServer  *server.Server
		httpHandler http.Handler
	)

	// create custom http handler
	httpHandler = http.NewServeMux()

	// create custom http server
	httpServer = server.NewServer("8080", httpHandler)

	// run http server
	go func() {
		if err := httpServer.Serve(); err != nil {
			log.Fatalf("Error occured while running http server: %v", err.Error())
		}
	}()

	// waiting for Ctrl + C to exit application
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// closing http server
	log.Println("Closing http serving...")
	err = httpServer.Close()
	if err != nil {
		return fmt.Errorf("httpServer.Close: %w", err)
	}

	log.Println("Server is shutted down.")

	return nil
}
