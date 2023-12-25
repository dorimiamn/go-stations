package handler

import (
	"net/http"
	// "fmt"
)

type panicHandler struct{}

// NewpanicHandler returns panicHandler based http.Handler.
func NewpanicHandler() *panicHandler {
	return &panicHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("panic")
	// fmt.Fprint(w, "Hello, world!")
}
