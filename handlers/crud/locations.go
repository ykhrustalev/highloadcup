package crud

import (
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
)

type Locations struct {
	repo   *repos.Repo
	prefix string
}

func NewLocations(repo *repos.Repo) *Locations {
	return &Locations{
		repo:   repo,
		prefix: "/locations/",
	}
}

func (h *Locations) Prefix() string {
	return h.prefix
}

func (h *Locations) New() interface{} {
	return &models.Location{}
}

func (h *Locations) NewPartial() interface{} {
	return &models.LocationPartialRaw{}
}

func (h *Locations) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Locations) Get(id int) interface{} {
	return h.repo.GetLocation(id)
}

func (h *Locations) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Location)
	source := theSource.(*models.LocationPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *Locations) Add(theTarget interface{}) error {
	target := theTarget.(*models.Location)

	_, err := h.repo.GetLocation(target.Id)
	if err == nil {
		return ErrorObjectExists
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveLocation(target)

	return nil
}
