FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s" ./cmd/gpsgend

FROM scratch
WORKDIR /app
COPY --from=builder /app/gpsgend gpsgend
COPY --from=builder /app/deployments/gpsgend.yaml gpsgend.yaml 
ENTRYPOINT [ "/app/gpsgend", "-config", "/app/gpsgend.yaml" ]