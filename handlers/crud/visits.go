package crud

import (
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
)

type Visits struct {
	repo   *repos.Repo
	prefix string
}

func NewVisits(repo *repos.Repo) *Visits {
	return &Visits{
		repo:   repo,
		prefix: "/visits/",
	}
}

func (h *Visits) Prefix() string {
	return h.prefix
}

func (h *Visits) New() interface{} {
	return &models.Visit{}
}

func (h *Visits) NewPartial() interface{} {
	return &models.VisitPartialRaw{}
}

func (h *Visits) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Visits) Get(id int) interface{} {
	return h.repo.GetVisit(id)
}

func (h *Visits) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Visit)
	source := theSource.(*models.VisitPartialRaw)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *Visits) Add(theTarget interface{}) error {
	target := theTarget.(*models.Visit)

	_, err := h.repo.GetVisit(target.Id)
	if err == nil {
		return ErrorObjectExists
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveVisit(target)

	return nil
}
