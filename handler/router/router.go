package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	todoService := service.NewTODOService(todoDB)
	mux.Handle("/healthz", handler.NewHealthzHandler())
	mux.Handle("/todos", handler.NewTODOHandler(todoService))
	mux.Handle("/do-panic", middleware.Recovery(handler.NewpanicHandler()))
	return mux
}
