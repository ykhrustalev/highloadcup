package highloadcup

import (
	"github.com/ykhrustalev/highloadcup/handlers"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"github.com/valyala/fasthttp"
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

func (r *Router) Handle(ctx *fasthttp.RequestCtx) {

	ok := r.listVisitsHandler.Handle(ctx)
	if ok {
		return
	}

	ok = r.locationsAvgHandler.Handle(ctx)
	if ok {
		return
	}

	ok = r.crudHandler.Handle(ctx)
	if ok {
		return
	}

	ctx.Error("not found", fasthttp.StatusNotFound)
}
