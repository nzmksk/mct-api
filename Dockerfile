# Build stage as builder
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory to /app inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of source code into the container
COPY . .

# Build the application
# CGO_ENABLED=0 disables C bindings so it compiles as a fully static binary
# GOOS=linux sets the target OS to Linux
# -a forces rebuild of packages (clean build)
# -installsuffix cgo prevents cache reuse of C-related packages
# -o main outputs the binary as "main"
# cmd/server/main.go is the entry point of the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create app directory
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy migration files if they exist
COPY --from=builder /app/migrations ./migrations/

# Container listens on port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
