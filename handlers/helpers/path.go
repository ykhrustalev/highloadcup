package helpers

import (
	"net/http"
	"strconv"
	"strings"
)

func PathToIdPrefix(req *http.Request, prefix string) (int, error) {
	return PathToId(req, prefix, "")
}

func PathToId(req *http.Request, prefix string, suffix string) (int, error) {
	idStr := strings.Replace(req.URL.Path, prefix, "", 1)
	idStr = strings.TrimRight(idStr, "/")
	if suffix != "" {
		idStr = strings.Replace(idStr, suffix, "", 1)
	}

	return ToInt(idStr)
}

func ToInt(val string) (int, error) {
	return strconv.Atoi(val)
}
