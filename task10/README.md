# Задание 10: Тестирование HTTP-сервера в Go

## Описание
Пример полного тестирования HTTP-сервера, включая отдельные обработчики и полную интеграцию. Показывает паттерны маршрутизации Go 1.22+.

## Файлы
- `server.go` - HTTP сервер с обработчиками
- `server_test.go` - тесты сервера и обработчиков

## Ключевые концепции

### 1. httptest для тестирования обработчиков
```go
req := httptest.NewRequest("GET", "/health", nil)
w := httptest.NewRecorder()

server.healthHandler(w, req)
resp := w.Result()
```
### 2. httptest.NewServer для интеграционных тестов
```go
ts := httptest.NewServer(server.Handler())
defer ts.Close()

resp, err := http.Get(ts.URL + "/health")
```
### 3. Паттерны маршрутизации Go 1.22+
```go
mux.HandleFunc("GET /users/{id}", s.getUserHandler)
```
### 4. Проверка JSON ответов
```go
var user User
json.NewDecoder(resp.Body).Decode(&user)

if user.Name != "Alice" {
    t.Errorf("user.Name = %q, want 'Alice'", user.Name)
}
```
Запуск тестов
bash
# Все тесты
go test -v

# Конкретный тест
go test -v -run TestHealthHandler

# С покрытием кода
go test -cover

# С детектором гонок (для конкурентных тестов)
go test -race