package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"github.com/MaxTorzh/go-practice/task5/internal/handlers"
    "github.com/MaxTorzh/go-practice/task5/internal/storage"
)

func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    w := httptest.NewRecorder()

    handlers.HealthHandler(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
    }
}

func TestGetUserHandler_Success(t *testing.T) {
    s := storage.NewUserStorage()
    handler := handlers.GetUserHandler(s)
    
    tests := []struct {
        name       string
        userID     int
        wantStatus int
    }{
        {"get user 1", 1, http.StatusOK},
        {"get user 2", 2, http.StatusOK},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            url := "/user?id=" + strconv.Itoa(tt.userID)
            req := httptest.NewRequest(http.MethodGet, url, nil)
            w := httptest.NewRecorder()
            
            handler(w, req)
            
            resp := w.Result()
            defer resp.Body.Close()
            
            if resp.StatusCode != tt.wantStatus {
                t.Errorf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
            }
        })
    }
}

func TestGetUserHandler_Errors(t *testing.T) {
    s := storage.NewUserStorage()
    handler := handlers.GetUserHandler(s)
    
    tests := []struct {
        name       string
        query      string
        method     string
        wantStatus int
    }{
        {
            name:       "missing id",
            query:      "/user",
            method:     http.MethodGet,
            wantStatus: http.StatusBadRequest,
        },
        {
            name:       "wrong method",
            query:      "/user?id=1",
            method:     http.MethodPost,
            wantStatus: http.StatusMethodNotAllowed,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(tt.method, tt.query, nil)
            w := httptest.NewRecorder()
            
            handler(w, req)
            
            resp := w.Result()
            defer resp.Body.Close()
            
            if resp.StatusCode != tt.wantStatus {
                t.Errorf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
            }
        })
    }
}

func TestCreateUserHandler_Success(t *testing.T) {
    s := storage.NewUserStorage()
    handler := handlers.CreateUserHandler(s)
    
    body := `{"name": "Charlie", "age": 35}`
    req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    handler(w, req)
    
    resp := w.Result()
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusCreated {
        t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusCreated)
    }
}

func TestUserFlow(t *testing.T) {
    s := storage.NewUserStorage()
    createHandler := handlers.CreateUserHandler(s)
    getHandler := handlers.GetUserHandler(s)
    
    createBody := `{"name": "David", "age": 40}`
    createReq := httptest.NewRequest(http.MethodPost, "/users", 
        bytes.NewReader([]byte(createBody)))
    createReq.Header.Set("Content-Type", "application/json")
    
    createW := httptest.NewRecorder()
    createHandler(createW, createReq)
    
    createResp := createW.Result()
    defer createResp.Body.Close()
    
    if createResp.StatusCode != http.StatusCreated {
        t.Fatalf("Create failed: status = %d", createResp.StatusCode)
    }
    
    var createdUser struct {
        ID int `json:"id"`
    }
    body, _ := io.ReadAll(createResp.Body)
    json.Unmarshal(body, &createdUser)
    
    getURL := "/user?id=" + strconv.Itoa(createdUser.ID)
    getReq := httptest.NewRequest(http.MethodGet, getURL, nil)
    
    getW := httptest.NewRecorder()
    getHandler(getW, getReq)
    
    getResp := getW.Result()
    defer getResp.Body.Close()
    
    if getResp.StatusCode != http.StatusOK {
        t.Errorf("Get failed: status = %d", getResp.StatusCode)
    }
}