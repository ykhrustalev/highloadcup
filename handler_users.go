package highloadcup

import (
	"net/http"
	"errors"
)

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

func (h *UsersHandler) NewItem() interface{} {
	return &User{}
}

func (h *UsersHandler) PathToId(req *http.Request) (int, error) {
	return pathToId(req, h.path)
}

func (h *UsersHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}

func (h *UsersHandler) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*User)
	source := theSource.(*User)

	if source.Email != nil {
		target.Email = source.Email
	}
	if source.FirstName != nil {
		target.FirstName = source.FirstName
	}
	if source.LastName != nil {
		target.LastName = source.LastName
	}
	if source.Gender != nil {
		target.Gender = source.Gender
	}
	if source.BirthDate != nil {
		target.BirthDate = source.BirthDate
	}

	return nil
}

func (h *UsersHandler) Add(theTarget interface{}) error {
	target := theTarget.(*User)

	// TODO: static errors

	if target.Email == nil {
		return errors.New("empty email")
	}
	if target.FirstName == nil {
		return errors.New("empty first_name")
	}
	if target.LastName == nil {
		return errors.New("empty last_name")
	}
	if target.Gender == nil {
		return errors.New("empty gender")
	}
	if target.BirthDate == nil {
		return errors.New("empty birth_date")
	}

	h.repo.Save(target)

	return nil
}
