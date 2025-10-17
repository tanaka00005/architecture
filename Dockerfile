FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache tzdata ca-certificates
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go mod download

CMD ["air"]