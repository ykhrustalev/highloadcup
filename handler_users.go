package highloadcup

import (
	"net/http"
)

type UsersHandler struct {
	repo UsersRepo
	path string
}

func NewUsersHandler(repo UsersRepo) *UsersHandler {
	return &UsersHandler{
		repo: repo,
		path: "/users/",
	}
}

func (h *UsersHandler) PathToId(req *http.Request) (int, error) {
	return pathToId(req, h.path)
}

func (h *UsersHandler) Get(id int) (interface{}, error) {
	return h.repo.Get(id)
}
func (h *UsersHandler) Update(id int, values map[string]string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
