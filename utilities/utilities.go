package utilities

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(wrapper)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	type errorMessage struct {
		Message string `json:"message"`
	}
	errMsg := errorMessage{
		Message: message,
	}
	wrapper := make(map[string]interface{})
	wrapper["error"] = errMsg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(wrapper)
}
