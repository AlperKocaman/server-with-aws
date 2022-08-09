# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint compared to Ubuntu
FROM golang:1.18-alpine AS build

# Create a directory inside image
WORKDIR /app

# Copy all the contents of the go directory(includes go.mod and go.sum)
COPY go/ ./

# Install dependencies
RUN go mod download

# Compile the app and create a binary called server-with-aws
RUN GOOS=linux go build -o /server-with-aws

EXPOSE 8080

# Use image to start a container
CMD [ "/server-with-aws" ]