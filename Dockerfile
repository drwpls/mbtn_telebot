FROM golang:1.22.0-alpine3.18 AS builder
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download
COPY . /app
RUN go build -o /app/main

FROM alpine:3.7
COPY --from=builder /app/main /app/main
CMD ["/app/main", "-config", "/app/config.yaml"]
