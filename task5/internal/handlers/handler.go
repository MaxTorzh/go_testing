package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/MaxTorzh/go-practice/task5/internal/storage"
)

func GetUserHandler(s *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			WriteError(w, http.StatusMethodNotAllowed,
			errors.New("method not allowed"))
			return 
		}

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			WriteError(w, http.StatusBadRequest,
			errors.New("id parameter is required"))
			return 
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			WriteError(w, http.StatusBadRequest,
			fmt.Errorf("invalid id: %w", err))
			return 
		}

		user, err := s.GetUser(id)
		if err != nil {
			WriteError(w, http.StatusNotFound, err)
			return 
		}

		WriteJson(w, http.StatusOK, user)
	}
}

func CreateUserHandler(s *storage.UserStorage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if r.Method != http.MethodPost {
			WriteError(w, http.StatusMethodNotAllowed,
			errors.New("method not allowed"))
			return 
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			WriteError(w, http.StatusUnsupportedMediaType,
			errors.New("content-type must be application/json"))
			return 
		}

		var req struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteError(w, http.StatusBadRequest,
			fmt.Errorf("invalid request body: %w", err))
			return 
		}

		user, err := s.CreateUser(req.Name, req.Age) 
		if err != nil {
			WriteError(w, http.StatusBadRequest, err)
			return 
		}
		WriteJson(w, http.StatusCreated, user)
	}
}