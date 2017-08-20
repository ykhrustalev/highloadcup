package highloadcup

import (
	"net/http"
	"github.com/ykhrustalev/highloadcup/models"
)

type LocationsHandler struct {
	repo LocationsRepo
	Path string
}

func NewLocationsHandler(repo LocationsRepo) *LocationsHandler {
	return &LocationsHandler{
		repo: repo,
		Path: "/locations/",
	}
}

func (h *LocationsHandler) New() interface{} {
	return &models.Location{}
}

func (h *LocationsHandler) NewPartial() interface{} {
	return &models.LocationPartialRaw{}
}

func (h *LocationsHandler) PathToId(req *http.Request) (int, error) {
	return pathToIdPrefix(req, h.Path)
}

func (h *LocationsHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *LocationsHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Location)
	source := theSource.(*models.LocationPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *LocationsHandler) Add(theTarget interface{}) error {
	target := theTarget.(*models.Location)

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
