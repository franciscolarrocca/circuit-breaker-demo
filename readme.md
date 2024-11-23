# circuit-breaker-demo

## Introduction

Project with the purpose to apply and show the usage of "circuit breaker" pattern, implemented with Golang.

## Prerequisites

- Go v1.21.8
- Docker (optional)

## Using Docker

In the root of the project (`/circuit-breaker-demo`) run the following command to create the images and run the containers:

```bash
docker-compose up --build 
```

## Without Docker (On windows)

In the root of the project (`/circuit-breaker-demo`) run the following commands to start the server:
```
cd server
go run .\app\
```

Before starting the client, update the server host URL in `main.go` (line 26) from `http://server:8080/data` to `http://localhost:8080/data`.

Again in the root of the project run (`/circuit-breaker-demo`) run the following commands to start the client:
```
cd client
go run .\app\
```

## How it works?

The logs will show the status of requests from the client to the server. When the number of consecutive failures reaches the defined threshold (2 requests with error in this example), the circuit breaker will be in state OPEN and will stop making requests for a defined duration of time (5 seconds in this example) avoiding server overload. After this timeout, the circuit breaker moves to the HALF_OPEN state. If the next request is successful, the failure counter resets to zero, and the circuit breaker returns to the CLOSED state.

```bash
2024-08-25 21:30:15 server  | server running on port 8080
2024-08-25 21:30:15 client  | 2024/08/26 00:30:15 result: successful 
2024-08-25 21:30:16 client  | 2024/08/26 00:30:16 result: an HTTP server error ocurred: '503: 503 Service Unavailable'
2024-08-25 21:30:17 client  | 2024/08/26 00:30:17 result: an HTTP server error ocurred: '503: 503 Service Unavailable' 
2024-08-25 21:30:18 client  | 2024/08/26 00:30:18 result: circuit breaker is OPEN waiting 5sec 
2024-08-25 21:30:19 client  | 2024/08/26 00:30:19 result: circuit breaker is OPEN waiting 5sec 
2024-08-25 21:30:20 client  | 2024/08/26 00:30:20 result: circuit breaker is OPEN waiting 5sec 
2024-08-25 21:30:21 client  | 2024/08/26 00:30:21 result: circuit breaker is OPEN waiting 5sec
2024-08-25 21:30:22 client  | 2024/08/26 00:30:15 result: successful
```