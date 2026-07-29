FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o go-maro-bin ./cmd/main.go

FROM alpine:3.24

WORKDIR /app
COPY --from=builder /app/go-maro-bin .
