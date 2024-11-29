FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum into the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire app directory into the container
COPY app/ ./

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
