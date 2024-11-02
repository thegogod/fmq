# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o ./bin/main ./server

# Expose the port your application listens on
EXPOSE 1883
EXPOSE 8080

ENV LOG_PREFIX=fmq/server
ENV PORT=8080

# Command to run when the container starts
CMD ["./bin/main"]
