package circuit_breaker

import (
	"errors"
	"franciscolarrocca/client/app/custom_errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CircuitBreaker_New(t *testing.T) {
	failureThreshold := 2
	duration := 2 * time.Second

	circuitBreaker := New(failureThreshold, duration)

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

func Test_CircuitBreaker_CallWithExcludedErrors_HttpServerError(t *testing.T) {
	failureThreshold := 2
	duration := 2 * time.Second
	circuitBreaker := &CircuitBreaker{
		failureThreshold: failureThreshold,
		state:            STATE_CLOSED,
		timeout:          duration,
	}

	excludedErrors := []error{&custom_errors.HttpClientError{}}

	err := circuitBreaker.CallWithExcludedErrors(mock_server_error, excludedErrors)
	assert.Error(t, err)
	assert.ErrorContains(t, err, mock_server_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)

	err = circuitBreaker.CallWithExcludedErrors(mock_server_error, excludedErrors)
	assert.ErrorContains(t, err, mock_server_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_OPEN)

	time.Sleep(3 * time.Second)

	_ = circuitBreaker.Call(mock_success)
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)
}

func Test_CircuitBreaker_CallWithExcludedErrors_HttpClientError(t *testing.T) {
	failureThreshold := 2
	duration := 2 * time.Second
	circuitBreaker := &CircuitBreaker{
		failureThreshold: failureThreshold,
		state:            STATE_CLOSED,
		timeout:          duration,
	}

	excludedErrors := []error{&custom_errors.HttpClientError{Status: 403, Message: "Bad Request"}}

	err := circuitBreaker.CallWithExcludedErrors(mock_client_error, excludedErrors)
	assert.Error(t, err)
	assert.ErrorContains(t, err, mock_client_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)

	err = circuitBreaker.CallWithExcludedErrors(mock_client_error, excludedErrors)
	assert.ErrorContains(t, err, mock_client_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)

	time.Sleep(3 * time.Second)

	err = circuitBreaker.CallWithExcludedErrors(mock_client_error, excludedErrors)
	assert.ErrorContains(t, err, mock_client_error().Error())
	assert.Equal(t, circuitBreaker.state, STATE_CLOSED)
}

func mock_error() error {
	return errors.New("an error ocurred")
}

func mock_server_error() error {
	return &custom_errors.HttpServerError{Status: 503, Message: "Service Unavailble"}
}

func mock_client_error() error {
	return &custom_errors.HttpClientError{Status: 403, Message: "Bad Request"}
}

func mock_success() error {
	return nil
}
