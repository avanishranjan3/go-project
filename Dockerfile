# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .
COPY go.mod .
COPY go.sum .
COPY .env /app
# Build the Go application inside the container
RUN go build -o myapp

# Expose a port if your Go application listens on a specific port
EXPOSE 9090

# Command to run when the container starts
CMD ["./myapp"]
