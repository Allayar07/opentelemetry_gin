FROM golang:1.19.5-alpine3.17

WORKDIR /app
COPY ./ ./

RUN go mod download

RUN go build -o service_2 ./cmd/main.go

CMD ["./service_2"]