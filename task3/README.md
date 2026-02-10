# Задание 3: Тестирование слайсов и мапов в Go

## Описание
Пример тестирования функций, работающих со слайсами и мапами. Показывает различные подходы к сравнению коллекций.

## Файлы
- `collections.go` - функции для работы со слайсами и мапами
- `collections_test.go` - тесты с использованием reflect.DeepEqual и slices.Equal

## Ключевые концепции

### 1. Сравнение слайсов
```go
// Старый способ (работает всегда)
import "reflect"
if !reflect.DeepEqual(got, expected) {
    t.Errorf("слайсы не равны")
}

// Новый способ (Go 1.21+)
import "slices"
if !slices.Equal(got, expected) {
    t.Errorf("слайсы не равны")
}
```
### 2. Сравнение мапов
```go
// Для мапов используем только reflect.DeepEqual
if !reflect.DeepEqual(got, expected) {
    t.Errorf("мапы не равны")
}
```
### 3. Тестирование уникальности
```go
func TestUnique(t *testing.T) {
    got := Unique([]string{"a", "b", "a", "c"})
    expected := []string{"a", "b", "c"}
    
    if !slices.Equal(got, expected) {
        t.Errorf("Unique() = %v, want %v", got, expected)
    }
}
```
### 4. Тестирование агрегаций
```go
func TestCountWords(t *testing.T) {
    got := CountWords([]string{"a", "b", "a"})
    expected := map[string]int{"a": 2, "b": 1}
    
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("CountWords() = %v, want %v", got, expected)
    }
}
```
Запуск тестов
bash
# Все тесты
go test -v

# Конкретный тест
go test -v -run TestUnique

# Бенчмарки
go test -bench=.

# Покрытие кода
go test -cover