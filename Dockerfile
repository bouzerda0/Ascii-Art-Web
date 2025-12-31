# Use the official lightweight Golang image based on Alpine Linux
FROM golang:1.23-alpine

# Set the working directory inside the container for organization
WORKDIR /app

# Copy all local files to the container's working directory
COPY . .

# Compile the Go application into a binary named 'app'
RUN go build -o app .

# Expose port 8000 as defined in the Go server logic
EXPOSE 8000

# Run the compiled binary when the container starts
CMD ["./app"]