package handlers

import (
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"strings"
)

type ListVisitsHandler struct {
	repo   *repos.Repo
	prefix string
	suffix string
}

func NewListVisitsHandler(repo *repos.Repo) *ListVisitsHandler {
	return &ListVisitsHandler{
		repo:   repo,
		prefix: "/users/",
		suffix: "/visits",
	}
}

func (h *ListVisitsHandler) PathToIdVisits(req *http.Request) (int, error) {
	return helpers.PathToId(req, h.prefix, h.suffix)
}

type VisitsResponse struct {
	Visits []*models.VisitForUser `json:"visits"`
}

func (h *ListVisitsHandler) Handle(w http.ResponseWriter, req *http.Request) bool {
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

func (h *ListVisitsHandler) doHandle(w http.ResponseWriter, req *http.Request) {
	id, err := h.PathToIdVisits(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user, err := h.repo.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	values := req.URL.Query()
	filters, err := models.VisitsFilterFromValues(&values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	visits := h.repo.FilterVisits(user.Id, filters)

	visitsForUser := make([]*models.VisitForUser, 0)
	for _, visit := range visits {
		visitsForUser = append(visitsForUser, visit.ToVisitForUser())
	}
	response := &VisitsResponse{
		Visits: visitsForUser,
	}

	helpers.WriteResponse(w, helpers.ToJson(response))
}
