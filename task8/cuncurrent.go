package task8

import (
	"sync"
	"time"
)

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

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Allow(user string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	validRequests := []time.Time{}
	for _, t := range rl.requests[user] {
		if t.After(windowStart) {
			validRequests = append(validRequests, t)
		}
	}

	if len(validRequests) >= rl.limit {
		return false
	}

	validRequests = append(validRequests, now)
	rl.requests[user] = validRequests
	return true
}

func ParallelSum(numbers []int, workers int) int {
	if len(numbers) == 0 {
		return 0
	}

	chunkSize := (len(numbers) + workers - 1) / workers
	results := make(chan int, workers)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		if start >= len(numbers) {
			break
		}

		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			sum := 0
			for _, n := range nums {
				sum += n
			}
			results <- sum
		}(numbers[start:end])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for sum := range results {
		total += sum
	}

	return total
}

func BatchProcess(tasks []func() error, batchSize int) []error {
	var wg sync.WaitGroup
	errors := make([]error, len(tasks))
	var mu sync.Mutex

	for i := 0; i < len(tasks); i += batchSize {
		end := i + batchSize
		if end > len(tasks) {
			end = len(tasks)
		}

		batch := tasks[i:end]
		for j, task := range batch {
			wg.Add(1)
			go func(idx int, t func() error) {
				defer wg.Done()
				err := t()
				if err != nil {
					mu.Lock()
					errors[idx] = err
					mu.Unlock()
				}
			}(i+j, task)
		}
		wg.Wait()
	}

	return errors
}