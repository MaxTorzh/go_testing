# Задание 4: Property-based тестирование с testing/quick

## Описание
Пример property-based тестирования с использованием пакета `testing/quick`. Показывает, как тестировать математические свойства и инварианты.

## Файлы
- `quick_example.go` - функции со свойствами для тестирования
- `quick_example_test.go` - тесты с quick.Check

## Ключевые концепции

### 1. Property-based тестирование
Вместо проверки конкретных примеров проверяем свойства, которые должны выполняться для всех входных данных.

### 2. quick.Check
```go
import "testing/quick"

property := func(x int) bool {
    return IsEven(x) == IsEven(x+2)
}

if err := quick.Check(property, nil); err != nil {
    t.Errorf("IsEven property failed: %v", err)
}
```
### 3. Конфигурация тестов
```go
config := &quick.Config{
    MaxCount: 1000,    
    MaxCountScale: 1.0, 
}
```
### 4. Примеры свойств
```go
func CommutativeAdd(a, b int) bool {
    return (a + b) == (b + a)
}

func ReverseTwice(s string) bool {
    return Reverse(Reverse(s)) == s
}
```
Запуск тестов
bash
# Все тесты
go test -v

# С большим количеством тестов
go test -v -count=100

# С детектором гонок
go test -race
