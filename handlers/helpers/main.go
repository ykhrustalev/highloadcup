package helpers

import (
	"fmt"
	"net/http"
	"github.com/pquerna/ffjson/ffjson"
)

func WriteResponseJson(w http.ResponseWriter, obj interface{}) {
	buf, _ := ffjson.Marshal(obj)
	WriteResponse(w, buf)
	ffjson.Pool(buf)
}

func WriteResponse(w http.ResponseWriter, contents []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(contents)))
	w.WriteHeader(http.StatusOK)
	w.Write(contents)
}
