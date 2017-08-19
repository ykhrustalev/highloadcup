package highloadcup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	users     *UsersHandler
	locations *LocationsHandler
	visits    *VisitsHandler
}

func NewRouter(users *UsersHandler, locations *LocationsHandler, visits *VisitsHandler) *Router {
	return &Router{
		users:     users,
		locations: locations,
		visits:    visits,
	}
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	if method == "POST" {
		if strings.HasSuffix(path, "/new") {
			if strings.HasPrefix(path, r.users.Path) {
				r.Add(r.users, w, req)
				return
			} else if strings.HasPrefix(path, r.locations.Path) {
				r.Add(r.locations, w, req)
				return
			} else if strings.HasPrefix(path, r.visits.Path) {
				r.Add(r.visits, w, req)
				return
			}
		} else {
			if strings.HasPrefix(path, r.users.Path) {
				r.Update(r.users, w, req)
				return
			} else if strings.HasPrefix(path, r.locations.Path) {
				r.Update(r.locations, w, req)
				return
			} else if strings.HasPrefix(path, r.visits.Path) {
				r.Update(r.visits, w, req)
				return
			}
		}
	} else if method == "GET" {
		if strings.HasPrefix(path, r.users.Path) {
			if strings.HasSuffix(path, "/visited") {

			} else {
				r.Get(r.users, w, req)
				return
			}
		} else if strings.HasPrefix(path, r.locations.Path) {
			r.Get(r.locations, w, req)
			return
		} else if strings.HasPrefix(path, r.visits.Path) {
			r.Get(r.visits, w, req)
			return
		}
	}

	http.Error(w, "method not supported", http.StatusNotFound)
}

func (r *Router) Get(handler Handler, w http.ResponseWriter, req *http.Request) {
	id, err := handler.PathToId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := handler.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	contents := r.toJson(item)
	r.writeResponse(w, contents)
}

func (r *Router) Update(handler Handler, w http.ResponseWriter, req *http.Request) {
	id, err := handler.PathToId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	target, err := handler.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	source := handler.NewPartial()

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = handler.Update(target, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r.writeResponse(w, []byte("{}"))
}

func (r *Router) Add(handler Handler, w http.ResponseWriter, req *http.Request) {
	target := handler.New()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//defer req.Body.Close() // seems not needed

	err = handler.Add(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r.writeResponse(w, []byte("{}"))
}

// helper

func (r *Router) toJson(obj interface{}) []byte {
	enc, _ := json.Marshal(obj)
	return enc
}

func (r *Router) writeResponse(w http.ResponseWriter, contents []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(contents)))
	w.WriteHeader(http.StatusOK)
	w.Write(contents)
}
