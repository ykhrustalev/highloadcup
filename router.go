package highloadcup

import (
	"github.com/ykhrustalev/highloadcup/handlers"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"net/http"
)

type Router struct {
	crudHandler         *crud.Handler
	listVisitsHandler   *handlers.ListVisitsHandler
	locationsAvgHandler *handlers.LocationsAvgHandler
}

func NewRouter(
	crudHandler *crud.Handler,
	listVisitsHandler *handlers.ListVisitsHandler,
	locationsAvgHandler *handlers.LocationsAvgHandler,
) *Router {
	return &Router{
		crudHandler:         crudHandler,
		listVisitsHandler:   listVisitsHandler,
		locationsAvgHandler: locationsAvgHandler,
	}
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
	ok := r.listVisitsHandler.Handle(w, req)
	if ok {
		return
	}

	ok = r.locationsAvgHandler.Handle(w, req)
	if ok {
		return
	}

	ok = r.crudHandler.Handle(w, req)
	if ok {
		return
	}

	http.Error(w, "method not supported", http.StatusNotFound)
}
