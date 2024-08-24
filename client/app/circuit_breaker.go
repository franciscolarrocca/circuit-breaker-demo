package main

import (
	"fmt"
	"time"
)

const (
	STATE_CLOSED    = "closed"
	STATE_OPEN      = "open"
	STATE_HALF_OPEN = "half_open"
)

type CircuitBreaker struct {
	failureCount     int
	failureThreshold int
	state            string
	timeout          time.Duration
	lastAttemptTime  time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		state:            STATE_CLOSED,
		timeout:          timeout,
	}
}

func (cb *CircuitBreaker) Call(operation func() error) error {
	if cb.state == STATE_OPEN {
		if time.Since(cb.lastAttemptTime) < cb.timeout {
			return fmt.Errorf("circuit breaker is OPEN waiting %dsec", int(cb.timeout.Seconds()))
		}
		cb.state = STATE_HALF_OPEN
	}

	if err := operation(); err != nil {
		cb.failureCount++
		cb.lastAttemptTime = time.Now()

		if cb.failureCount >= cb.failureThreshold {
			cb.state = STATE_OPEN
		}

		return err
	}

	cb.failureCount = 0
	cb.state = STATE_CLOSED

	return nil
}