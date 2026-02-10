# Задание 6-7: Mocking и тестирование времени в Go

## Описание
Комбинированный пример, показывающий использование моков для тестирования внешних зависимостей и моков времени для тестирования time-dependent логики.

## Файлы
- `service.go` - бизнес-логика с зависимостями от времени и внешних сервисов
- `service_test.go` - тесты с моками

## Ключевые концепции

### 1. Интерфейсы для мокинга
```go
type PaymentGateway interface {
    ProcessPayment(ctx context.Context, amount float64) (string, error)
}

type TimeProvider interface {
    Now() time.Time
    Sleep(d time.Duration)
}
```
### 2. Простые моки
```go
type MockPaymentGateway struct {
    ProcessPaymentFunc func(ctx context.Context, amount float64) (string, error)
    Calls []struct{ Amount float64 }
}

func (m *MockPaymentGateway) ProcessPayment(ctx context.Context, amount float64) (string, error) {
    m.Calls = append(m.Calls, struct{ Amount float64 }{Amount: amount})
    if m.ProcessPaymentFunc != nil {
        return m.ProcessPaymentFunc(ctx, amount)
    }
    return "", errors.New("not implemented")
}
```
### 3. Мок времени
```go
type MockTime struct {
    currentTime time.Time
    sleepCalls  []time.Duration
}

func (m *MockTime) Now() time.Time {
    return m.currentTime
}

func (m *MockTime) Sleep(d time.Duration) {
    m.sleepCalls = append(m.sleepCalls, d)
    m.currentTime = m.currentTime.Add(d)
}
```
### 4. Инъекция зависимостей
```go
func NewPaymentService(gateway PaymentGateway, timeProvider TimeProvider) *PaymentService {
    return &PaymentService{
        gateway: gateway,
        time:    timeProvider,
    }
}
```
Запуск тестов
bash
# Все тесты
go test -v

# Только unit-тесты
go test -v -short

# С покрытием кода
go test -cover