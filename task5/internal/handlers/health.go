package handlers

import (
	"errors"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteError(w, http.StatusMethodNotAllowed,
		errors.New("method not allowed"))
		return
	}

	response := map[string]string{
		"status": "ok",
		"service": "user-api",
	}

	WriteJson(w, http.StatusOK, response)
}