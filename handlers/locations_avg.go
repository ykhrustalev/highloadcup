package handlers

import (
	"fmt"
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

func (h *LocationsAvgHandler) PathToLocationId(req *http.Request) (int, error) {
	return helpers.PathToId(req, h.prefix, h.suffix)
}

func (h *LocationsAvgHandler) Handle(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != "GET" {
		return false
	}

	path := req.URL.Path

	if !(strings.HasPrefix(path, h.prefix) && strings.HasSuffix(path, h.suffix)) {
		return false
	}

	h.doHandle(w, req)

	return true
}

func (h *LocationsAvgHandler) doHandle(w http.ResponseWriter, req *http.Request) {
	id, err := h.PathToLocationId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	location, found := h.repo.GetLocation(id)
	if !found {
		http.Error(w, crud.ErrorNotFound.Error(), http.StatusNotFound)
		return
	}

	values := req.URL.Query()
	filters, err := models.LocationsAvgFilterFromValues(&values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := h.repo.AverageLocationMark(location.Id, filters)
	contents := fmt.Sprintf("{\"avg\": %.5f}", value)

	helpers.WriteResponse(w, []byte(contents))
}

type LocationsAvgResponse struct {
	Avg float32 `json:"avg"`
}
