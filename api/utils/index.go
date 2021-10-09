package utils

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, e error, header int) {
	w.WriteHeader(header)
	message := make(map[string]string)
	message["error"] = e.Error()
	json.NewEncoder(w).Encode(message)
}
