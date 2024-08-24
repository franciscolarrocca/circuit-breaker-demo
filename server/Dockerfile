FROM golang:1.21.8-alpine

WORKDIR /server

COPY . .

RUN go mod tidy

RUN go test ./app

RUN go build -o server ./app

EXPOSE 8080

CMD ["./server"]
