FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gochat .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/gochat .
COPY --from=builder /app/static ./static
COPY --from=builder /app/config ./config
EXPOSE 8081
CMD ["./gochat", "server"]
