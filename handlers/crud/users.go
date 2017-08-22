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
	return &models.UserRaw{}
}

func (h *Users) NewPartial() interface{} {
	return &models.UserPartial{}
}

func (h *Users) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Users) Get(id int) (interface{}, bool) {
	return h.repo.GetUser(id)
}

func (h *Users) GetRaw(id int) (interface{}, bool) {
	obj, found := h.repo.GetUser(id)
	if found {
		return obj.UserRaw(), found
	}
	return obj, found
}

func (h *Users) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.User)
	source := theSource.(*models.UserPartial)
	return h.repo.UpdateUser(target, source)
}

func (h *Users) Add(theTarget interface{}) error {
	target := theTarget.(*models.UserRaw).User()

	//_, found := h.repo.GetUser(target.Id)
	//if found {
	//	return ErrorObjectExists
	//}

	err := target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveUser(target)

	return nil
}
