package highloadcup

import (
	"errors"
	"net/http"
)

// TODO: static errors

type UsersHandler struct {
	repo UsersRepo
	path string
}

func NewUsersHandler(repo UsersRepo) *UsersHandler {
	return &UsersHandler{
		repo: repo,
		path: "/users/",
	}
}

func (h *UsersHandler) New() interface{} {
	return &User{}
}

func (h *UsersHandler) NewPartial() interface{} {
	return &UserPartialRaw{}
}

func (h *UsersHandler) PathToId(req *http.Request) (int, error) {
	return pathToId(req, h.path)
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
		return errors.New("user with id exists")
	}

	h.repo.Save(target)

	return nil
}
