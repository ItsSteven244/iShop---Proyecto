# Etapa 1: builder (tiene el toolchain de Go)
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0: binario estático. Como usamos glebarez/sqlite (driver Go puro),
# no necesitamos gcc ni libs de C ni siquiera para la compilación local.
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

# Etapa 2: runner (imagen mínima, sin el toolchain de Go)
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /bin/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]