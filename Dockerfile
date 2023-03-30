# Use the official Golang image as the parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /app

# Copy the current directory contents into the container at /go/src/app
COPY . /app

# Install any needed dependencies
RUN go get -d -v ./...

# Compile the Go program
RUN go build cmd/main.go

# # Expose port 8080 for the container
EXPOSE 8080

# Run the Go program when the container starts
CMD ["./main"]
