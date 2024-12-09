# Start from the official Golang base image.
FROM golang:1.23.3-alpine as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy go module files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the current directory to the container's working directory.
COPY . .

# Build the application for a production environment from the new main.go location.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api-gateway

# Start a new stage from scratch for a smaller image size.
FROM alpine:latest

# Set the current working directory.
WORKDIR /root/

# Copy the binary file from the builder stage.
COPY --from=builder /app/main .

ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY

ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

# Specify the command to run on container start.
CMD ["./main"]