package highloadcup

import (
	"net/http"
	"github.com/ykhrustalev/highloadcup/models"
)

type VisitsHandler struct {
	repo VisitsRepo
	Path string
}

func NewVisitsHandler(repo VisitsRepo) *VisitsHandler {
	return &VisitsHandler{
		repo: repo,
		Path: "/visits/",
	}
}

func (h *VisitsHandler) New() interface{} {
	return &models.Visit{}
}

func (h *VisitsHandler) NewPartial() interface{} {
	return &models.VisitPartialRaw{}
}

func (h *VisitsHandler) PathToId(req *http.Request) (int, error) {
	return pathToIdPrefix(req, h.Path)
}

func (h *VisitsHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *VisitsHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Visit)
	source := theSource.(*models.VisitPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *VisitsHandler) Add(theTarget interface{}) error {
	target := theTarget.(*models.Visit)

	_, err := h.repo.Get(target.Id)
	if err == nil {
		return ErrorObjectExists
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	h.repo.Save(target)

	return nil
}

func (h *VisitsHandler) Filter(userId int, filters *models.VisitsFilter) []*models.Visit {
	return h.repo.Filter(userId, filters)
}
