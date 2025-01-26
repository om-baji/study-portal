package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Error   interface{} `json:"error,omitempty"`
}

func ToJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
