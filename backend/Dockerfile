FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tms-server .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/tms-server .
COPY .env .env
EXPOSE 8080
CMD ["./tms-server"]
