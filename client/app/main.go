package main

import (
	"franciscolarrocca/client/app/circuit_breaker"
	"franciscolarrocca/client/app/custom_errors"
	"log"
	"net/http"
	"time"
)

func main() {
	failureThreshold := 2
	timeout := 5 * time.Second
	circuitBreaker := circuit_breaker.New(failureThreshold, timeout)

	for i := 0; i < 15; i++ {
		excludedErrors := []error{&custom_errors.HttpClientError{}}
		if err := circuitBreaker.CallWithExcludedErrors(doRequest, excludedErrors); err != nil {
			log.Printf("result: %s \n", err.Error())
		} else {
			log.Printf("result: successful \n")
		}
		time.Sleep(1 * time.Second)
	}
}

func doRequest() error {
	resp, err := http.Get("http://server:8080/data")
	if err != nil {
		return err
	}

	if custom_errors.Is5xxError(resp.StatusCode) {
		return &custom_errors.HttpServerError{
			Status:  resp.StatusCode,
			Message: resp.Status,
		}
	}

	if custom_errors.Is4xxError(resp.StatusCode) {
		return &custom_errors.HttpClientError{
			Status:  resp.StatusCode,
			Message: resp.Status,
		}
	}

	return nil
}
