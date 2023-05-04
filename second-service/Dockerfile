FROM golang:1.20-alpine

ENV GOPATH=/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY ./ ./

RUN go build -o opt ./cmd/main.go

CMD ["./opt"]