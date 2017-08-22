package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"strings"
)

type LocationsAvgHandler struct {
	repo   *repos.Repo
	prefix string
	suffix string
}

func NewLocationsAvgHandler(repo *repos.Repo) *LocationsAvgHandler {
	return &LocationsAvgHandler{
		repo:   repo,
		prefix: "/locations/",
		suffix: "/avg",
	}
}

func (h *LocationsAvgHandler) PathToLocationId(req *fasthttp.Request) (int, error) {
	return helpers.PathToId(req, h.prefix, h.suffix)
}

func (h *LocationsAvgHandler) Handle(ctx *fasthttp.RequestCtx) bool {
	if string(ctx.Method()) != "GET" {
		return false
	}

	path := string(ctx.Path())

	if !(strings.HasPrefix(path, h.prefix) && strings.HasSuffix(path, h.suffix)) {
		return false
	}

	h.doHandle(ctx)

	return true
}

func (h *LocationsAvgHandler) doHandle(ctx *fasthttp.RequestCtx) {
	id, err := h.PathToLocationId(&ctx.Request)
	if err != nil {
		ctx.Error(err.Error(), http.StatusNotFound)
		return
	}

	location, found := h.repo.GetLocation(id)
	if !found {
		ctx.Error(crud.ErrorNotFound.Error(), http.StatusNotFound)
		return
	}

	values := ctx.QueryArgs()
	filters, err := models.LocationsAvgFilterFromValues(values)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	value := h.repo.AverageLocationMark(location.Id, filters)
	contents := fmt.Sprintf("{\"avg\": %.5f}", value)

	helpers.WriteResponse(ctx, []byte(contents))
}

type LocationsAvgResponse struct {
	Avg float32 `json:"avg"`
}
