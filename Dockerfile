# Use the official Go image to build the application
FROM golang:1.22.5 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Then copy the rest of the code
COPY . .

# Build the application for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Use a smaller image to run the application
FROM alpine:latest

# Install the necessary dependencies to run the application
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
