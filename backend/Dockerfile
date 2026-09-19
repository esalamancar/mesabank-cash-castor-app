# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/cashcastor-api .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/cashcastor-api /usr/local/bin/cashcastor-api

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/cashcastor-api"]
