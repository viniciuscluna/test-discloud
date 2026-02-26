# Stage 1: Builder
# Use the official Golang image for building
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker's cache
COPY go.mod .
COPY go.sum .

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application binary
# CGO_ENABLED=0 disables CGo, making the binary static and portable
RUN CGO_ENABLED=0 go build -o /myapp ./cmd/api/main.go

# Stage 2: Runtime
# Use a minimal base image, like scratch (most minimal) or alpine
FROM scratch

# Copy the built binary from the builder stage into the final image
COPY --from=builder /myapp /myapp

# Expose the port your application listens on (optional, if a web app)
EXPOSE 8080

# Command to run the executable when the container starts2
CMD ["/myapp"]
