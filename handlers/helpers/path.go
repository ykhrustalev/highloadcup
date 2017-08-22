package helpers

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
)

func PathToIdPrefix(req *fasthttp.Request, prefix string) (int, error) {
	return PathToId(req, prefix, "")
}

func PathToId(req *fasthttp.Request, prefix string, suffix string) (int, error) {
	path := string(req.URI().Path())

	idStr := strings.Replace(path, prefix, "", 1)
	idStr = strings.TrimRight(idStr, "/")
	if suffix != "" {
		idStr = strings.Replace(idStr, suffix, "", 1)
	}

	return ToInt(idStr)
}

func ToInt(val string) (int, error) {
	return strconv.Atoi(val)
}
