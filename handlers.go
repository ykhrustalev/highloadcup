package highloadcup

import (
	"strings"
	"net/http"
	"strconv"
)

type Handler interface {
	PathToId(req *http.Request) (int, error)
	Get(id int) (interface{}, error)
	Update(int, map[string]string) (map[string]interface{}, error)
}

/// helpers

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
