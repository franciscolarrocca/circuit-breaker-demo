FROM golang:1.21.8-alpine

WORKDIR /client

COPY . .

RUN go mod tidy

RUN go test ./app

RUN go build -o client ./app

CMD ["./client"]
