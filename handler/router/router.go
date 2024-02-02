package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
)

type middleWare func(http.Handler) http.Handler

type mwStack struct {
	middleWares []middleWare
}

func newMws(mws ...middleWare) mwStack {
	return mwStack{append([]middleWare{}, mws...)}
}

func (m mwStack) Then(h http.Handler) http.Handler {
	for i := range m.middleWares {
		h = m.middleWares[len(m.middleWares)-1-i](h)
	}
	return h
}

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// create middleware stack
	mws := newMws(middleware.Recovery,middleware.GetUserAgent,middleware.Logger)
	mwsWithBasicAuth := newMws(middleware.Recovery,middleware.GetUserAgent,middleware.Logger,middleware.BasicAuth)
	// register routes
	mux := http.NewServeMux()
	todoService := service.NewTODOService(todoDB)
	mux.Handle("/healthz", mws.Then(handler.NewHealthzHandler()))
	mux.Handle("/todos",mwsWithBasicAuth.Then(handler.NewTODOHandler(todoService)))
	mux.Handle("/do-panic",mws.Then(handler.NewpanicHandler()))
	mux.Handle("/delay",mws.Then(handler.NewDelayHandler()))
	return mux
}
