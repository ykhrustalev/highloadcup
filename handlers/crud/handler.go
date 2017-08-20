package crud

import (
	"encoding/json"
	"github.com/ykhrustalev/highloadcup/handlers/helpers"
	"github.com/ykhrustalev/highloadcup/repos"
	"net/http"
	"strings"
)

type Adapter interface {
	Prefix() string
	PathToId(req *http.Request) (int, error)
	New() interface{}
	NewPartial() interface{}
	Get(id int) (interface{}, error)
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

func (r *Handler) Handle(w http.ResponseWriter, req *http.Request) bool {
	ok := r.chain(r.users, w, req)
	if ok {
		return ok
	}

	ok = r.chain(r.locations, w, req)
	if ok {
		return ok
	}

	ok = r.chain(r.visits, w, req)
	if ok {
		return ok
	}

	return false
}

func (r *Handler) chain(adapter Adapter, w http.ResponseWriter, req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	prefix := adapter.Prefix()

	if !strings.HasPrefix(path, prefix) {
		return false
	}

	if method == "POST" {
		if strings.HasSuffix(path, "/new") {
			r.Add(adapter, w, req)
			return true
		} else {
			r.Update(adapter, w, req)
			return true
		}
	}

	if method == "GET" {
		r.Get(adapter, w, req)
		return true
	}

	return false
}

func (r *Handler) Get(adapter Adapter, w http.ResponseWriter, req *http.Request) {
	id, err := adapter.PathToId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	item, err := adapter.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helpers.WriteResponse(w, helpers.ToJson(item))
}

func (r *Handler) Update(adapter Adapter, w http.ResponseWriter, req *http.Request) {
	id, err := adapter.PathToId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	target, err := adapter.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	source := adapter.NewPartial()

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = adapter.Update(target, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.WriteResponse(w, []byte("{}"))
}

func (r *Handler) Add(adapter Adapter, w http.ResponseWriter, req *http.Request) {
	target := adapter.New()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = adapter.Add(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.WriteResponse(w, []byte("{}"))
}
