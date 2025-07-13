# Stage 1: Build
FROM golang:alpine AS builder

WORKDIR /usr/local/app

# Copy go mod and sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp .

# Stage 2: Run
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache \
    ca-certificates \
    postgresql-client \
    tzdata

WORKDIR /usr/local/app

# Copy the built binary
COPY --from=builder /usr/local/app/myapp .

# Expose application port
EXPOSE 8000

# Run the binary
CMD ["./myapp"]