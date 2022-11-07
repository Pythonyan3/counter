package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Pythonyan3/Counter/internal/service"
)

type HttpCounterHandler struct {
	name    string
	service service.CounterServiceInterface
}

// NewHttpCounterHandler HttpCounterHandler struct constructor function.
func NewHttpCounterHandler(name string, service service.CounterServiceInterface) *HttpCounterHandler {
	return &HttpCounterHandler{name: name, service: service}
}

// InitRoutes initialize http.ServeMux with handler funcs.
func (handler *HttpCounterHandler) InitRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("/", handler.GetCounterValue)
	mux.HandleFunc("/stat", handler.IncrementCounterValue)
	mux.HandleFunc("/about", handler.About)
	return nil
}

// GetCounterValue handle root route to retrieve counter current value.
func (handler *HttpCounterHandler) GetCounterValue(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		counterValue int64
		response     CounterResponse
	)

	counterValue, err = handler.service.Get()
	if err != nil {
		log.Println(fmt.Sprintf("handler.service.Get: %v", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response = CounterResponse{Counter: counterValue}
	json.NewEncoder(w).Encode(response)
}

// IncrementCounterValue handle route to retrieve increment counter current value by one and return new value.
func (handler *HttpCounterHandler) IncrementCounterValue(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		counter  int64
		response CounterResponse
	)

	counter, err = handler.service.Increment(fmt.Sprintf("UserAgent: %v", r.UserAgent()))
	if err != nil {
		log.Println(fmt.Sprintf("handler.service.Increment: %v", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response = CounterResponse{Counter: counter}

	json.NewEncoder(w).Encode(response)
}

// About handle route to retrieve author name as html page.
func (handler *HttpCounterHandler) About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(fmt.Sprintf("<h3> Hello, %v.</h3>", handler.name)))
}
