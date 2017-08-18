package highloadcup

import (
	"strings"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

type UsersHandler struct {
	repo UserRepo
	Path string
}

func NewUsersHandler(repo UserRepo) *UsersHandler {
	return &UsersHandler{
		repo: repo,
		Path: "/users/",
	}
}

func (h *UsersHandler) Handle(w http.ResponseWriter, req *http.Request) {
	id, err := pathToId(req, h.Path)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.repo.Get(id)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	contents := toJson(user)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(contents)))
	w.Write([]byte(contents))

	return
}

/// helpers

func toJson(obj interface{}) []byte {
	enc, _ := json.Marshal(obj)
	return enc
}

func writeResponse(w http.ResponseWriter, code int, contents string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(contents)))
	w.Write([]byte(contents))

}

func pathToId(req *http.Request, prefix string) (int, error) {
	idStr := strings.Replace(req.URL.Path, prefix, "", 1)
	return toInt(idStr)
}

func toInt(val string) (int, error) {
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, ErrorMalformed
	}
	return v, nil
}
