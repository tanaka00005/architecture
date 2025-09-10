FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache tzdata ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

CMD ["go", "run", "main.go"]