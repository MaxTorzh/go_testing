package task10

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewServer() *Server {
	return &Server{
		users: map[int]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
		nextID: 3,
	}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/users/", s.userHandler)
	mux.HandleFunc("/users", s.createUserHandler)
	
	return http.ListenAndServe(addr, mux)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	s.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	if path == "" {
		s.writeError(w, http.StatusBadRequest, "user id is required")
		return
	}
	
	id, err := strconv.Atoi(path)
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	s.mu.RLock()
	user, exists := s.users[id]
	s.mu.RUnlock()

	if !exists {
		s.writeError(w, http.StatusNotFound, "user not found")
		return
	}

	s.writeJSON(w, http.StatusOK, user)
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Name == "" {
		s.writeError(w, http.StatusBadRequest, "name is required")
		return
	}

	s.mu.Lock()
	user := User{ID: s.nextID, Name: req.Name}
	s.users[s.nextID] = user
	s.nextID++
	s.mu.Unlock()

	s.writeJSON(w, http.StatusCreated, user)
}

func (s *Server) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) writeError(w http.ResponseWriter, status int, message string) {
	s.writeJSON(w, status, map[string]string{"error": message})
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/users/", s.userHandler)
	mux.HandleFunc("/users", s.createUserHandler)
	return mux
}