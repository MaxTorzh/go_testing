# Задание 8: Тестирование конкурентности в Go

## Описание
Пример тестирования конкурентного кода, включая потокобезопасные структуры, ограничители запросов и параллельные алгоритмы.

## Файлы
- `concurrent.go` - конкурентные структуры и алгоритмы
- `concurrent_test.go` - тесты с детектором гонок данных

## Ключевые концепции

### 1. Потокобезопасные структуры
```go
type SafeCounter struct {
    mu    sync.RWMutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}
```
### 2. Ограничитель запросов
```go
type RateLimiter struct {
    mu       sync.Mutex
    requests map[string][]time.Time
    limit    int
    window   time.Duration
}
```
### 3. Параллельные алгоритмы
```go
func ParallelSum(numbers []int, workers int) int {
    results := make(chan int, workers)
    var wg sync.WaitGroup
    
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func(nums []int) {
            defer wg.Done()
            sum := 0
            for _, n := range nums {
                sum += n
            }
            results <- sum
        }(chunk)
    }
}
```
Запуск тестов
bash
# Обычные тесты
go test -v

# С детектором гонок данных (обязательно!)
go test -race -v

# Бенчмарки
go test -bench=.

# Покрытие кода
go test -cover