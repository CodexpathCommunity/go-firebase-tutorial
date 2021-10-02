package utils

import (
	"encoding/json"
	"net/http"
)

func SUCCESS(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		SUCCESS(w, statusCode, map[string]string{"error": err.Error()})
	}
	SUCCESS(w, http.StatusBadRequest, nil)
}
