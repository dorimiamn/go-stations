package handler

import (
	"io"
	"net/http"
	"time"
)

type DelayHandler struct{}

// NewDelayHandler returns DelayHandler based http.Handler.
func NewDelayHandler() *DelayHandler {
	return &DelayHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *DelayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	io.WriteString(w, "Hello, world!\n")
}
