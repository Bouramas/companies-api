FROM golang:1.22-alpine as build

RUN apk update && apk add --no-cache bash make

# Use sh instead of bash
SHELL ["/bin/sh", "-c"]

# Set the working directory to /companies-api
RUN mkdir -p /companies-api
WORKDIR /companies-api

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY ./ /companies-api

# Build the Go app
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the command to run the executable
CMD ["./companies-api"]
