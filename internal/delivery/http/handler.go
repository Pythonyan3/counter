package http

import (
	"net/http"

	v1 "github.com/Pythonyan3/Counter/internal/delivery/http/v1"
)

// verify interface compliance for specific structs.
var _ HttpRouteInitor = (*v1.HttpCounterHandler)(nil)

type HttpRouteInitor interface {
	InitRoutes(mux *http.ServeMux) error
}
