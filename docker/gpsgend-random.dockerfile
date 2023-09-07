FROM golang:1.21 AS builder
LABEL stage=builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s" ./cmd/gpsgend-random

FROM scratch
COPY --from=builder /app/gpsgend-random /app/gpsgend-random
ENTRYPOINT ["/app/gpsgend-random"]