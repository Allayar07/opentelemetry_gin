FROM golang:1.20-alpine
COPY .. ./
RUN go mod download
RUN go build -o opt ./cmd/main.go
CMD ["opt"]