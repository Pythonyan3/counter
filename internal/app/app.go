package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	delivery_http "github.com/Pythonyan3/Counter/internal/delivery/http"
	v1 "github.com/Pythonyan3/Counter/internal/delivery/http/v1"
	"github.com/Pythonyan3/Counter/internal/repository/postgresql"
	"github.com/Pythonyan3/Counter/internal/server"
	"github.com/Pythonyan3/Counter/internal/service"
)

const name string = "Vitalii Manoilo"

// Run application entrypoint.
func Run() error {
	var (
		err               error
		httpServer        *server.Server
		mux               *http.ServeMux
		counterService    service.CounterServiceInterface
		counterRepository *postgresql.PostgresRepository
		httpHandler       delivery_http.HttpRouteInitor
	)

	// create in memory repository
	counterRepository, err = postgresql.NewPostgresRepository()
	if err != nil {
		log.Fatalf("Error postgresql.NewPostgresRepository: %v", err)
	}

	// create counter service
	counterService = service.NewCounterService(counterRepository)

	// create http mux router
	mux = http.NewServeMux()

	// create http counter handler
	httpHandler = v1.NewHttpCounterHandler(name, counterService)

	// init routes
	httpHandler.InitRoutes(mux)

	// create custom http server
	httpServer = server.NewServer("8080", mux)

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

	log.Println("Closing postgresql connection...")
	err = counterRepository.Close()
	if err != nil {
		return fmt.Errorf("counterRepository.Close: %w", err)
	}

	log.Println("Server is shutted down.")

	return nil
}
