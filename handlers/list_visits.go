package handlers

import (
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"strings"
	"github.com/valyala/fasthttp"
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

func (h *ListVisitsHandler) PathToIdVisits(req *fasthttp.Request) (int, error) {
	return helpers.PathToId(req, h.prefix, h.suffix)
}

func (h *ListVisitsHandler) Handle(ctx *fasthttp.RequestCtx) bool {
	path := string(ctx.Path())
	method := string(ctx.Method())

	if method != "GET" {
		return false
	}

	if !(strings.HasPrefix(path, h.prefix) && strings.HasSuffix(path, h.suffix)) {
		return false
	}

	h.doHandle(ctx)

	return true
}

func (h *ListVisitsHandler) doHandle(ctx *fasthttp.RequestCtx) {
	id, err := h.PathToIdVisits(&ctx.Request)
	if err != nil {
		ctx.Error(err.Error(), http.StatusNotFound)
		return
	}

	user, found := h.repo.GetUser(id)
	if !found {
		ctx.Error(crud.ErrorNotFound.Error(), http.StatusNotFound)
		return
	}

	values := ctx.QueryArgs()
	filters, err := models.VisitsFilterFromValues(values)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.VisitsResponse{
		Visits: h.repo.FilterVisitsForUser(user.Id, filters),
	}

	helpers.WriteResponseJson(ctx, response)
}
