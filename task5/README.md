# Задание 5: Тестирование HTTP-обработчиков в Go

## Описание
Пример тестирования HTTP-обработчиков с использованием пакета `httptest`. Показывает, как тестировать REST API без запуска реального сервера.

## Файлы
- `handler.go` - HTTP обработчики и модели
- `handler_test.go` - тесты обработчиков

## Ключевые концепции

### 1. httptest.NewRecorder и httptest.NewRequest
```go
import "net/http/httptest"

req := httptest.NewRequest("GET", "/health", nil)
w := httptest.NewRecorder()

handler(w, req)
resp := w.Result()
```
### 2. Проверка статус-кодов
```go
if resp.StatusCode != http.StatusOK {
    t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
}
```
### 3. Проверка заголовков
```go
contentType := resp.Header.Get("Content-Type")
if contentType != "application/json" {
    t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
}
```
### 4. Парсинг JSON ответов
```go
var result map[string]string
json.NewDecoder(resp.Body).Decode(&result)

if status, ok := result["status"]; !ok || status != "ok" {
    t.Errorf("status = %q, want 'ok'", status)
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