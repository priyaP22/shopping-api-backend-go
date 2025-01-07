# Start from the official Golang image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
