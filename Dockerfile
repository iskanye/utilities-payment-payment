# Сборка
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/payment ./cmd/payment/main.go

# Запуск
FROM alpine
USER root
WORKDIR /home/app
COPY --from=builder /bin/payment ./
COPY --from=builder /app/config ./config
ENTRYPOINT ["./payment"]
CMD ["-config", "./config/dev.yaml"]