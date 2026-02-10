package task10

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	server := NewServer()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	server.healthHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want %q", ct, "application/json")
	}

	var body map[string]string
	json.NewDecoder(resp.Body).Decode(&body)

	if status, ok := body["status"]; !ok || status != "ok" {
		t.Errorf("body = %v, want status=ok", body)
	}
}

func TestUserHandler(t *testing.T) {
	server := NewServer()

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantName   string
	}{
		{
			name:       "существующий пользователь GET",
			method:     http.MethodGet,
			path:       "/users/1",
			wantStatus: http.StatusOK,
			wantName:   "Alice",
		},
		{
			name:       "несуществующий пользователь GET",
			method:     http.MethodGet,
			path:       "/users/999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "невалидный ID GET",
			method:     http.MethodGet,
			path:       "/users/abc",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "пустой ID GET",
			method:     http.MethodGet,
			path:       "/users/",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "неправильный метод POST",
			method:     http.MethodPost,
			path:       "/users/1",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			server.userHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}

			if tt.wantName != "" {
				var user User
				json.NewDecoder(resp.Body).Decode(&user)
				if user.Name != tt.wantName {
					t.Errorf("user.Name = %q, want %q", user.Name, tt.wantName)
				}
			}
		})
	}
}

func TestCreateUserHandler(t *testing.T) {
	server := NewServer()

	tests := []struct {
		name       string
		method     string
		body       string
		wantStatus int
	}{
		{
			name:       "успешное создание POST",
			method:     http.MethodPost,
			body:       `{"name": "Charlie"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "пустое имя POST",
			method:     http.MethodPost,
			body:       `{"name": ""}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "невалидный JSON POST",
			method:     http.MethodPost,
			body:       `{invalid json}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "неправильный метод GET",
			method:     http.MethodGet,
			body:       "",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/users", bytes.NewReader([]byte(tt.body)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			server.createUserHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestFullServer(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server.Handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("health request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("health status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	createBody := `{"name": "David"}`
	resp, err = http.Post(ts.URL+"/users", "application/json", bytes.NewReader([]byte(createBody)))
	if err != nil {
		t.Fatalf("create user failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("create status = %d, want %d", resp.StatusCode, http.StatusCreated)
	}

	var createdUser User
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &createdUser)

	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(createdUser.ID))
	if err != nil {
		t.Fatalf("get user failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("get status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	var fetchedUser User
	json.NewDecoder(resp.Body).Decode(&fetchedUser)

	if fetchedUser.Name != "David" {
		t.Errorf("fetched user name = %q, want %q", fetchedUser.Name, "David")
	}
}

func TestConcurrentRequests(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server.Handler())
	defer ts.Close()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			resp, err := http.Get(ts.URL + "/health")
			if err == nil {
				resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					done <- true
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestErrorResponseFormat(t *testing.T) {
	server := NewServer()

	req := httptest.NewRequest("GET", "/users/999", nil)
	w := httptest.NewRecorder()

	server.userHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	var errorResp map[string]string
	json.NewDecoder(resp.Body).Decode(&errorResp)

	if _, ok := errorResp["error"]; !ok {
		t.Errorf("error response missing 'error' field: %v", errorResp)
	}
}