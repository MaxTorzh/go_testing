# Задание 1: Базовые тестирование в Go

## Описание
Пример базовых unit-тестов для математических функций. Демонстрирует подход table-driven тестирования в Go.

## Файлы
- `math.go` - функции для тестирования (Add, Subtract, Multiply, Divide)
- `math_test.go` - тесты с использованием table-driven подхода

## Ключевые концепции

### 1. Table-driven тесты
```go
tests := []struct {
    name     string
    a, b     int
    expected int
}{
    {"positive numbers", 2, 3, 5},
    {"negative numbers", -2, -3, -5},
}
```

### 2. Подтесты (t.Run)
```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result := Add(tt.a, tt.b)
    })
}
```

### 3. Проверка ошибок
```go
result, err := Divide(a, b)
if err != nil {
    t.Errorf("Divide(%d, %d) error: %v", a, b, err)
}
```

## Запуск тестов
bash
# Все тесты
go test -v

# Конкретный тест
go test -v -run TestAdd

# С покрытием кода
go test -cover
Пример вывода
text
=== RUN   TestAdd
=== RUN   TestAdd/positive_numbers
=== RUN   TestAdd/negative_numbers
--- PASS: TestAdd (0.00s)
