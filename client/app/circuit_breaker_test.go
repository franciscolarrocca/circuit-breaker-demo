package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CircuitBreaker_NewCircuitBreaker(t *testing.T) {
	failureThreshold := 2
	duration := 2 * time.Second

	circuitBreaker := NewCircuitBreaker(failureThreshold, duration)

	assert.NotNil(t, circuitBreaker)
	assert.Equal(t, 2, circuitBreaker.failureThreshold)
	assert.Equal(t, 2*time.Second, circuitBreaker.timeout)
}

func Test_CircuitBreaker_Call(t *testing.T) {
	failureThreshold := 2
	duration := 2 * time.Second
	circuitBreaker := &CircuitBreaker{
		failureThreshold: failureThreshold,
		state:            STATE_CLOSED,
		timeout:          duration,
	}

	err := circuitBreaker.Call(mock_error)
	assert.Error(t, err)
	assert.ErrorContains(t, err, mock_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)

	_ = circuitBreaker.Call(mock_error)
	assert.Equal(t, circuitBreaker.state, STATE_OPEN)

	time.Sleep(3 * time.Second)

	_ = circuitBreaker.Call(mock_success)
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)
}

func mock_error() error {
	return errors.New("an error ocurred")
}

func mock_success() error {
	return nil
}
