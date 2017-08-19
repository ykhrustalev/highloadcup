package highloadcup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	users *UsersHandler
}

func NewRouter(users *UsersHandler) *Router {
	return &Router{
		users: users,
	}
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	fmt.Println(path, method)

	if method == "POST" {
		if strings.HasPrefix(path, "/items/") {

		} else if strings.HasPrefix(path, "/locations/") {

		} else if strings.HasPrefix(path, "/visits/") {

		}
	} else if method == "GET" {
		if strings.HasPrefix(path, "/items/") {

			if strings.HasSuffix(path, "/visited") {

			} else {
				r.Get(r.users, w, req)
			}

		} else if strings.HasPrefix(path, "/locations/") {

		} else if strings.HasPrefix(path, "/visits/") {

		}
	}

	http.Error(w, "method not supported", http.StatusBadRequest)
}
func (r *Router) Get(handler Handler, w http.ResponseWriter, req *http.Request) {
	id, err := handler.PathToId(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	contents := r.toJson(user)
	r.writeResponse(w, contents)
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
