package highloadcup

import (
	"net/http"
)

type UsersHandler struct {
	repo UsersRepo
	Path string
	VisitsPath string
}

func NewUsersHandler(repo UsersRepo) *UsersHandler {
	return &UsersHandler{
		repo: repo,
		Path: "/users/",
		VisitsPath: "/visits",
	}
}

func (h *UsersHandler) New() interface{} {
	return &User{}
}

func (h *UsersHandler) NewPartial() interface{} {
	return &UserPartialRaw{}
}

func (h *UsersHandler) PathToId(req *http.Request) (int, error) {
	return pathToIdPrefix(req, h.Path)
}

func (h *UsersHandler) PathToIdVisits(req *http.Request) (int, error) {
	return pathToId(req, h.Path, h.VisitsPath)
}

func (h *UsersHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *UsersHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*User)
	source := theSource.(*UserPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *UsersHandler) Add(theTarget interface{}) error {
	target := theTarget.(*User)

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
