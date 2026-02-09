package server

import (
    "log"
    "net/http"
    "github.com/MaxTorzh/go-practice/task5/internal/handlers"
    "github.com/MaxTorzh/go-practice/task5/internal/storage"
)

func main() {
    storage := storage.NewUserStorage()
    
    mux := http.NewServeMux()
    mux.HandleFunc("/health", handlers.HealthHandler)
    mux.HandleFunc("/user", handlers.GetUserHandler(storage))
    mux.HandleFunc("/users", handlers.CreateUserHandler(storage))
    
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}