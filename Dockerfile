FROM golang:1.25 AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum* ./

RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server

# Use minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080

CMD ["./server"]
