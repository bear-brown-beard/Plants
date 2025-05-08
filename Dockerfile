FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o plants .

FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/plants .

# Copy environment file
COPY .env .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./plants"]
