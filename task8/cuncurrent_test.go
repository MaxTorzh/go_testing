package task8

import (
	"sync"
	"testing"
	"time"
)

func TestSafeCounter(t *testing.T) {
	counter := &SafeCounter{}
	const goroutines = 1000
	const increments = 100

	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	expected := goroutines * increments
	if counter.Value() != expected {
		t.Errorf("counter.Value() = %d, want %d", counter.Value(), expected)
	}
}

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(5, time.Second)

	for i := 0; i < 5; i++ {
		if !rl.Allow("user1") {
			t.Errorf("request %d should be allowed", i+1)
		}
	}

	if rl.Allow("user1") {
		t.Error("6th request should be denied")
	}

	if !rl.Allow("user2") {
		t.Error("user2 should be allowed")
	}
}

func TestRateLimiter_Concurrent(t *testing.T) {
	rl := NewRateLimiter(100, time.Second)
	const users = 10
	const requests = 20

	var wg sync.WaitGroup
	allowed := make([]int, users)

	for i := 0; i < users; i++ {
		wg.Add(1)
		userID := i
		go func() {
			defer wg.Done()
			for j := 0; j < requests; j++ {
				if rl.Allow(string(rune('A' + userID))) {
					allowed[userID]++
				}
			}
		}()
	}

	wg.Wait()

	for i, count := range allowed {
		if count > 100 {
			t.Errorf("user %d got %d requests, max 100", i, count)
		}
	}
}

func TestParallelSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		workers int
		want    int
	}{
		{
			name:    "empty slice",
			numbers: []int{},
			workers: 4,
			want:    0,
		},
		{
			name:    "single worker",
			numbers: []int{1, 2, 3, 4, 5},
			workers: 1,
			want:    15,
		},
		{
			name:    "multiple workers",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			workers: 3,
			want:    55,
		},
		{
			name:    "more workers than numbers",
			numbers: []int{1, 2, 3},
			workers: 10,
			want:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParallelSum(tt.numbers, tt.workers)
			if got != tt.want {
				t.Errorf("ParallelSum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParallelSum_Concurrent(t *testing.T) {
	const size = 10000
	numbers := make([]int, size)
	expected := 0

	for i := 0; i < size; i++ {
		numbers[i] = i
		expected += i
	}

	workers := []int{1, 2, 4, 8, 16}
	for _, w := range workers {
		t.Run(string(rune('0'+w)), func(t *testing.T) {
			got := ParallelSum(numbers, w)
			if got != expected {
				t.Errorf("ParallelSum with %d workers = %d, want %d", w, got, expected)
			}
		})
	}
}

func TestBatchProcess(t *testing.T) {
	tasks := make([]func() error, 10)
	for i := 0; i < 10; i++ {
		taskNum := i
		tasks[i] = func() error {
			if taskNum%3 == 0 {
				return &testError{msg: "task failed"}
			}
			return nil
		}
	}

	errors := BatchProcess(tasks, 3)

	errorCount := 0
	for i, err := range errors {
		if i%3 == 0 {
			if err == nil {
				t.Errorf("task %d should have error", i)
			} else {
				errorCount++
			}
		} else if err != nil {
			t.Errorf("task %d should not have error, got %v", i, err)
		}
	}

	if errorCount != 4 {
		t.Errorf("expected 4 errors, got %d", errorCount)
	}
}

func TestConcurrentReadWrite(t *testing.T) {
	counter := &SafeCounter{}

	var wg sync.WaitGroup
	readers := 50
	writers := 50

	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}()
	}

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				val := counter.Value()
				if val < 0 {
					t.Error("counter value should not be negative")
				}
			}
		}()
	}

	wg.Wait()
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}