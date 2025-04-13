package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	WriteResponse(w, statusCode, map[string]string{"error": err.Error()})
}
