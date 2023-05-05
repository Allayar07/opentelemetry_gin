FROM golang:1.20-alpine

ENV GOPATH=/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY ./ ./

RUN go build -o service_1 ./cmd/main.go

CMD ["./service_1"]