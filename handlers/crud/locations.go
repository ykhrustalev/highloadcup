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
	return &models.LocationPartial{}
}

func (h *Locations) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Locations) Get(id int) (interface{}, bool) {
	return h.repo.GetLocation(id)
}

func (h *Locations) GetRaw(id int) (interface{}, bool) {
	return h.Get(id)
}

func (h *Locations) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Location)
	source := theSource.(*models.LocationPartial)

	return h.repo.UpdateLocation(target, source)
}

func (h *Locations) Add(theTarget interface{}) error {
	target := theTarget.(*models.Location)

	//_, found := h.repo.GetLocation(target.Id)
	//if found {
	//	return ErrorObjectExists
	//}

	err := target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveLocation(target)

	return nil
}
