# Задание 2: Тестирование ошибок в Go

## Описание
Пример тестирования функций, которые возвращают ошибки. Показывает работу с кастомными ошибками и их проверкой.

## Файлы
- `validator.go` - функции с возвратом ошибок
- `validator_test.go` - тесты проверки ошибок

## Ключевые концепции

### 1. Создание кастомных ошибок
```go
var (
    ErrNegativeNumber = errors.New("negative number not allowed")
    ErrZeroNumber     = errors.New("zero not allowed")
)
```
### 2. Проверка ошибок через errors.Is()
```go
err := ValidateNumber(-5)
if !errors.Is(err, ErrNegativeNumber) {
    t.Errorf("ожидали ErrNegativeNumber, получили: %v", err)
}
```
### 3. Оборачивание ошибок
```go
return fmt.Errorf("process failed: %w", err)
```
### 4. Table-driven тесты для ошибок
```go
tests := []struct {
    name    string
    input   int
    wantErr error
}{
    {"positive number", 42, nil},
    {"negative number", -5, ErrNegativeNumber},
}
```
Запуск тестов
bash
# Все тесты
go test -v

# С детектором гонок данных
go test -race

# Покрытие кода
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

## Пример проверки ошибки
```go
func TestValidateNumber(t *testing.T) {
    err := ValidateNumber(-10)
    
    if err == nil {
        t.Fatal("ожидали ошибку, получили nil")
    }
    
    if !errors.Is(err, ErrNegativeNumber) {
        t.Errorf("неправильный тип ошибки: %v", err)
    }
    
    if !strings.Contains(err.Error(), "negative") {
        t.Errorf("текст ошибки не содержит 'negative'")
    }
}
```