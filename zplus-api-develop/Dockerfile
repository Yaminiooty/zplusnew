# Use an official Golang runtime as a parent image
FROM golang:1.21.1


ENV GONOSUMDB=*
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app
COPY . /app

RUN go mod download

# Install "air" for hot-reloading
#RUN go install github.com/cosmtrek/air@latest

# Copy the local package files to the container's workspace
RUN go build -o main

# Update Go packages with go get
# Expose the port the application will run on
EXPOSE 3000

# Start the Go application with hot-reloading using "air"
CMD ["./main"]