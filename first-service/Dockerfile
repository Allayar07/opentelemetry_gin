FROM golang:1.19-alpine

ENV GOPATH=/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY ./ ./

RUN go build -o service_1 ./cmd/main.go

EXPOSE 8089

CMD ["./service_1"]