FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY main.go .
RUN go build -o app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 80
CMD ["./app"]
