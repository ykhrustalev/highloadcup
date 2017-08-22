package crud

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"strings"
	"bytes"
)

type Adapter interface {
	Prefix() string
	PathToId(req *fasthttp.Request) (int, error)
	New() interface{}
	NewPartial() interface{}
	Get(id int) (interface{}, bool)
	GetRaw(id int) (interface{}, bool)
	Update(interface{}, interface{}) error
	Add(interface{}) error
}

type Handler struct {
	users     *Users
	locations *Locations
	visits    *Visits
}

func NewHandler(repo *repos.Repo) *Handler {
	return &Handler{
		users:     NewUsers(repo),
		locations: NewLocations(repo),
		visits:    NewVisits(repo),
	}
}

func (r *Handler) Handle(ctx *fasthttp.RequestCtx) bool {
	ok := r.chain(r.users, ctx)
	if ok {
		return ok
	}

	ok = r.chain(r.locations, ctx)
	if ok {
		return ok
	}

	ok = r.chain(r.visits, ctx)
	if ok {
		return ok
	}

	return false
}

func (r *Handler) chain(adapter Adapter, ctx *fasthttp.RequestCtx) bool {
	path := string(ctx.Path())
	method := string(ctx.Method())
	prefix := adapter.Prefix()

	if !strings.HasPrefix(path, prefix) {
		return false
	}

	if method == "POST" {
		if strings.HasSuffix(path, "/new") {
			r.Add(adapter, ctx)
			return true
		} else {
			r.Update(adapter, ctx)
			return true
		}
	}

	if method == "GET" {
		r.Get(adapter, ctx)
		return true
	}

	return false
}

func (r *Handler) Get(adapter Adapter, ctx *fasthttp.RequestCtx) {
	id, err := adapter.PathToId(&ctx.Request)
	if err != nil {
		ctx.Error(err.Error(), http.StatusNotFound)
		return
	}

	item, found := adapter.GetRaw(id)
	if !found {
		ctx.Error(ErrorNotFound.Error(), http.StatusNotFound)
		return
	}

	helpers.WriteResponseJson(ctx, item)
}

func (r *Handler) Update(adapter Adapter, ctx *fasthttp.RequestCtx) {
	id, err := adapter.PathToId(&ctx.Request)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	target, found := adapter.Get(id)
	if !found {
		ctx.Error(ErrorNotFound.Error(), http.StatusNotFound)
		return
	}

	source := adapter.NewPartial()

	decoder := ffjson.NewDecoder()
	err = decoder.DecodeReader(bytes.NewReader(ctx.PostBody()), source)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = adapter.Update(target, source)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	helpers.WriteResponse(ctx, []byte("{}"))
}

func (r *Handler) Add(adapter Adapter, ctx *fasthttp.RequestCtx) {
	target := adapter.New()

	decoder := ffjson.NewDecoder()
	err := decoder.DecodeReader(bytes.NewReader(ctx.PostBody()), target)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = adapter.Add(target)
	if err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	helpers.WriteResponse(ctx, []byte("{}"))
}
