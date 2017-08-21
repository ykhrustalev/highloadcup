package crud

import (
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
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
	return &models.VisitPartial{}
}

func (h *Visits) PathToId(req *http.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Visits) Get(id int) (interface{}, bool) {
	return h.repo.GetVisit(id)
}

func (h *Visits) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Visit)
	source := theSource.(*models.VisitPartial)

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (h *Visits) Add(theTarget interface{}) error {
	target := theTarget.(*models.Visit)

	//_, found := h.repo.GetVisit(target.Id)
	//if found {
	//	return ErrorObjectExists
	//}

	err := target.Validate()
	if err != nil {
		return err
	}

	h.repo.SaveVisit(target)

	return nil
}
