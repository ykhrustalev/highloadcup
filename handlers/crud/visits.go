package crud

import (
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"github.com/valyala/fasthttp"
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
	return &models.VisitRaw{}
}

func (h *Visits) NewPartial() interface{} {
	return &models.VisitPartial{}
}

func (h *Visits) PathToId(req *fasthttp.Request) (int, error) {
	return helpers.PathToIdPrefix(req, h.prefix)
}

func (h *Visits) Get(id int) (interface{}, bool) {
	return h.repo.GetVisit(id)
}

func (h *Visits) GetRaw(id int) (interface{}, bool) {
	obj, found := h.repo.GetVisit(id)
	if found {
		return obj.VisitRaw(), found
	}
	return obj, found
}

func (h *Visits) Update(theTarget interface{}, theSource interface{}) error {
	target := theTarget.(*models.Visit)
	source := theSource.(*models.VisitPartial)
	return h.repo.UpdateVisit(target, source)
}

func (h *Visits) Add(theTarget interface{}) error {
	target := theTarget.(*models.VisitRaw).Visit()

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
