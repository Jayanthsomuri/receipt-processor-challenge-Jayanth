# Use the official Golang image to create a build artifact.
FROM golang:1.20-alpine as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go module files.
COPY go.mod .
COPY go.sum .

# Download the dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -o receipt-processor .

# Use a minimal Alpine image for the final container.
FROM alpine:latest

# Set the working directory inside the container.
WORKDIR /root/

# Copy the binary from the builder stage.
COPY --from=builder /app/receipt-processor .

# Expose the application port.
EXPOSE 8080

# Run the application.
CMD ["./receipt-processor"]