package task6_7

import (
	"context"
	"errors"
	"testing"
	"time"
)

type MockPaymentGateway struct {
	ProcessPaymentFunc func(ctx context.Context, amount float64) (string, error)
	Calls              []struct {
		Amount float64
	}
}

func (m *MockPaymentGateway) ProcessPayment(ctx context.Context, amount float64) (string, error) {
	m.Calls = append(m.Calls, struct{ Amount float64 }{Amount: amount})
	if m.ProcessPaymentFunc != nil {
		return m.ProcessPaymentFunc(ctx, amount)
	}
	return "", errors.New("not implemented")
}

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

func TestProcessPaymentWithRetry_Success(t *testing.T) {
	mockGateway := &MockPaymentGateway{
		ProcessPaymentFunc: func(ctx context.Context, amount float64) (string, error) {
			return "TXN-123", nil
		},
	}
	
	mockTime := &MockTime{
		currentTime: time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
	}
	
	service := NewPaymentService(mockGateway, mockTime)
	ctx := context.Background()
	
	txID, err := service.ProcessPaymentWithRetry(ctx, 100.0, 3)
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if txID != "TXN-123" {
		t.Errorf("txID = %q, want %q", txID, "TXN-123")
	}
	
	if len(mockGateway.Calls) != 1 {
		t.Errorf("ProcessPayment called %d times, want 1", len(mockGateway.Calls))
	}
}

func TestProcessPaymentWithRetry_WithRetries(t *testing.T) {
	callCount := 0
	mockGateway := &MockPaymentGateway{
		ProcessPaymentFunc: func(ctx context.Context, amount float64) (string, error) {
			callCount++
			if callCount < 3 {
				return "", errors.New("temporary error")
			}
			return "TXN-456", nil
		},
	}
	
	mockTime := &MockTime{
		currentTime: time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
	}
	
	service := NewPaymentService(mockGateway, mockTime)
	ctx := context.Background()
	
	txID, err := service.ProcessPaymentWithRetry(ctx, 200.0, 5)
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if txID != "TXN-456" {
		t.Errorf("txID = %q, want %q", txID, "TXN-456")
	}
	
	if callCount != 3 {
		t.Errorf("ProcessPayment called %d times, want 3", callCount)
	}
	
	expectedSleeps := []time.Duration{
		0,          
		100000000,
	}
	
	if len(mockTime.sleepCalls) != len(expectedSleeps) {
		t.Errorf("Sleep called %d times, want %d", len(mockTime.sleepCalls), len(expectedSleeps))
	}
}

func TestIsBusinessHours(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "рабочий день 10:00",
			time:     time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "рабочий день 17:59",
			time:     time.Date(2026, 1, 1, 17, 59, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "после работы 18:00",
			time:     time.Date(2026, 1, 1, 18, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "до работы 8:59",
			time:     time.Date(2026, 1, 1, 8, 59, 0, 0, time.UTC),
			expected: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTime := &MockTime{currentTime: tt.time}
			service := NewPaymentService(nil, mockTime)
			
			result := service.IsBusinessHours()
			
			if result != tt.expected {
				t.Errorf("IsBusinessHours() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCalculateDueDate(t *testing.T) {
	mockTime := &MockTime{
		currentTime: time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
	}
	
	service := NewPaymentService(nil, mockTime)
	
	tests := []struct {
		days     int
		expected string
	}{
		{1, "2026-01-02"},
		{5, "2026-01-08"},
	}
	
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			dueDate := service.CalculateDueDate(tt.days)
			formatted := dueDate.Format("2006-01-02")
			
			if formatted != tt.expected {
				t.Errorf("CalculateDueDate(%d) = %s, want %s", tt.days, formatted, tt.expected)
			}
		})
	}
}


