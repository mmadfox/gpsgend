FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s" ./deployments/bootstrap.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bootstrap bootstrap 
ENTRYPOINT [ "/app/bootstrap"]