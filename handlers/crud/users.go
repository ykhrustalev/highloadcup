package crud

import (
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
)

type Users struct {
	repo   *repos.Repo
	prefix string
}

func NewUsers(repo *repos.Repo) *Users {
	return &Users{
		repo:   repo,
		prefix: "/users/",
	}
}

func (h *Users) Prefix() string {
	return h.prefix
}

func (h *Users) New() interface{} {
	return &models.User{}
}

func (h *Users) NewPartial() interface{} {
	return &models.UserPartialRaw{}
}

func (h *Users) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Users) Get(id int) (interface{}, error) {
	return h.repo.GetUser(id)
}

func (h *Users) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.User)
	source := theSource.(*models.UserPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *Users) Add(theTarget interface{}) error {
	target := theTarget.(*models.User)

	_, err := h.repo.GetUser(target.Id)
	if err == nil {
		return ErrorObjectExists
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveUser(target)

	return nil
}
