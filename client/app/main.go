package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func main() {
	failureThreshold := 2
	timeout := 5 * time.Second
	circuitBreaker := NewCircuitBreaker(failureThreshold, timeout)

	for i := 0; i < 15; i++ {
		if err := circuitBreaker.Call(doRequest); err != nil {
			log.Printf("result: %s \n", err)
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
	if resp.StatusCode != http.StatusOK {
		return errors.New("server error")
	}
	return nil
}
