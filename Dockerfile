FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o firewall-proxy main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/firewall-proxy .
EXPOSE 7860
CMD ["./firewall-proxy"]