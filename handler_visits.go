package highloadcup

import (
	"net/http"
)

type VisitsHandler struct {
	repo VisitsRepo
	path string
}

func NewVisitsHandler(repo VisitsRepo) *VisitsHandler {
	return &VisitsHandler{
		repo: repo,
		path: "/visits/",
	}
}

func (h *VisitsHandler) New() interface{} {
	return &Visit{}
}

func (h *VisitsHandler) NewPartial() interface{} {
	return &VisitPartialRaw{}
}

func (h *VisitsHandler) PathToId(req *http.Request) (int, error) {
	return pathToId(req, h.path)
}

func (h *VisitsHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *VisitsHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*Visit)
	source := theSource.(*VisitPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *VisitsHandler) Add(theTarget interface{}) error {
	target := theTarget.(*Visit)

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
