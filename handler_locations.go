package highloadcup

import (
	"net/http"
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
	return &Location{}
}

func (h *LocationsHandler) NewPartial() interface{} {
	return &LocationPartialRaw{}
}

func (h *LocationsHandler) PathToId(req *http.Request) (int, error) {
	return pathToId(req, h.Path)
}

func (h *LocationsHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *LocationsHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*Location)
	source := theSource.(*LocationPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *LocationsHandler) Add(theTarget interface{}) error {
	target := theTarget.(*Location)

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
