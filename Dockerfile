# Start from the official Golang image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code (including the 'cmd' folder)
COPY . .

# Set the working directory to the 'cmd' folder, where the main.go is
WORKDIR /app/cmd

# Build the Go application (specify the main.go file location)
RUN go build -o /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
