package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New() // Singleton

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload) // Decode json
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v) // Encode json
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
