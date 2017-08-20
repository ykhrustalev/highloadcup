package highloadcup

import (
	"net/http"
	"strconv"
	"strings"
)

type Handler interface {
	PathToId(req *http.Request) (int, error)
	New() interface{}
	NewPartial() interface{}
	Get(id int) (interface{}, error)
	Update(interface{}, interface{}) error
	Add(interface{}) error
}

/// helpers

func pathToIdPrefix(req *http.Request, prefix string) (int, error) {
	return pathToId(req, prefix, "")
}

func pathToId(req *http.Request, prefix string, suffix string) (int, error) {
	idStr := strings.Replace(req.URL.Path, prefix, "", 1)
	idStr = strings.TrimRight(idStr, "/")
	if suffix != "" {
		idStr = strings.Replace(idStr, suffix, "", 1)
	}

	return toInt(idStr)
}

func toInt(val string) (int, error) {
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, ErrorMalformed
	}
	return v, nil
}
