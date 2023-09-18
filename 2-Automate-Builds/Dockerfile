# Start from golang base image for build stage
FROM golang:alpine AS build

# Enable go modules
ENV GO111MODULE=on

# Install git (alpine image does not it, and go uses it to download dependencies)
RUN apk update && apk add --no-cache git

# Set current working directory - /app is the convention, but not mandatory
WORKDIR /app

# Copy go mod and sum files - these have the dependencies our app needs
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello ./main.go

# Start from a minimal base image for runtime stage
FROM alpine

# Set current working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app .

# Expose the necessary port - not super important, but convention
EXPOSE 8080

# The Command that will get executed when we run the container
CMD ["./hello"]%