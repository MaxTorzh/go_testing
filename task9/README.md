# Задание 9: Тестирование конкретных ошибок в Go

## Описание
Пример тестирования функций, которые возвращают конкретные ошибки. Показывает работу с кастомными ошибками, их проверкой и обертыванием.

## Файлы
- `errors.go` - функции с возвратом конкретных ошибок
- `errors_test.go` - тесты проверки конкретных ошибок

## Ключевые концепции

### 1. Определение кастомных ошибок
```go
var (
    ErrInvalidInput   = errors.New("invalid input")
    ErrNegativeNumber = errors.New("negative number not allowed")
    ErrDivByZero      = errors.New("division by zero")
    ErrTooLarge       = errors.New("number too large")
)
```
### 2. Проверка через errors.Is()
```go
err := ParsePositiveNumber("-10")
if !errors.Is(err, ErrNegativeNumber) {
    t.Errorf("ожидали ErrNegativeNumber, получили: %v", err)
}
```
### 3. Оборачивание ошибок
```go
return fmt.Errorf("%w: empty string", ErrInvalidInput)
```
### 4. Table-driven тесты для ошибок
```go
tests := []struct {
    name     string
    input    string
    wantErr  error
    wantCont string
}{
    {"empty string", "", ErrInvalidInput, "empty string"},
    {"negative number", "-10", ErrNegativeNumber, "negative number"},
}
```
Запуск тестов
bash
# Все тесты
go test -v

# Конкретный тест
go test -v -run TestParsePositiveNumber

# С покрытием кода
go test -cover