package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ToJson(obj interface{}) []byte {
	enc, _ := json.Marshal(obj)
	return enc
}


func WriteResponse(w http.ResponseWriter, contents []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(contents)))
	w.WriteHeader(http.StatusOK)
	w.Write(contents)
}
