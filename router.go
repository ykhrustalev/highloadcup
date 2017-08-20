package highloadcup

import (
	"github.com/ykhrustalev/highloadcup/handlers"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"net/http"
)

type Router struct {
	crudHandler       *crud.Handler
	listVisitsHandler *handlers.ListVisitsHandler
}

func NewRouter(crudHandler *crud.Handler, listVisitsHandler *handlers.ListVisitsHandler) *Router {
	return &Router{
		crudHandler:       crudHandler,
		listVisitsHandler: listVisitsHandler,
	}
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
	ok := r.listVisitsHandler.Handle(w, req)
	if ok {
		return
	}

	ok = r.crudHandler.Handle(w, req)
	if ok {
		return
	}

	http.Error(w, "method not supported", http.StatusNotFound)
}
