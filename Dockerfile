# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Copy go.mod and go.sum to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o email-service ./cmd/main.go

# Stage 2: Create a lightweight final image
FROM alpine:latest

WORKDIR /app

# Install necessary packages (like ca-certificates for secure connections)
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/email-service .

# Copy config files (optional)
COPY config/config.yaml ./config/config.yaml

# Use a non-root user for security
RUN addgroup -S appgroup
