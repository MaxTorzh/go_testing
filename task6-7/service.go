package task6_7

import (
	"context"
	"fmt"
	"time"
)

type PaymentGateway interface {
	ProcessPayment(ctx context.Context, amount float64) (string, error)
}

type TimeProvider interface {
	Now() time.Time
	Sleep(d time.Duration)
}

type RealPaymentGateway struct {}

func(r *RealPaymentGateway) ProcessPayment(ctx context.Context, amount float64) (string, error) {
	time.Sleep(100 * time.Millisecond)

	if amount <= 0 {
		return "", fmt.Errorf("amount must be positive")
	}

	return fmt.Sprintf("TXN-%d", time.Now().Unix()), nil
}

type RealTime struct {}

func (r RealTime) Now() time.Time {
	return time.Now()
}

func (r RealTime) Sleep(d time.Duration) {
	time.Sleep(d)
}

type PaymentService struct {
	gateway PaymentGateway
	time    TimeProvider
}

func NewPaymentService(gateway PaymentGateway, timeProvider TimeProvider) *PaymentService {
	return &PaymentService{
		gateway: gateway,
		time:    timeProvider,
	}
}

func (s *PaymentService) ProcessPaymentWithRetry(ctx context.Context, amount float64, maxRetries int) (string, error) {
	var lastErr error
	
	for i := 0; i <= maxRetries; i++ {
		txID, err := s.gateway.ProcessPayment(ctx, amount)
		if err == nil {
			return txID, nil
		}
		
		lastErr = err

		waitTime := time.Duration(i*i) * 100 * time.Millisecond
		s.time.Sleep(waitTime)
		
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
	}
	
	return "", fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

func (s *PaymentService) IsBusinessHours() bool {
	now := s.time.Now()
	
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}
	
	hour := now.Hour()
	return hour >= 9 && hour < 18
}

func (s *PaymentService) CalculateDueDate(days int) time.Time {
	now := s.time.Now()
	
	for i := 0; i < days; i++ {
		now = now.Add(24 * time.Hour)
		
		for now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
			now = now.Add(24 * time.Hour)
		}
	}
	
	return now
}

