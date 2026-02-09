package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/MaxTorzh/go-practice/task5/internal/models"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, models.ErrorResponse{Error: err.Error()})
}