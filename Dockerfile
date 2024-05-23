# Dockerfile

FROM golang:1.21.3

# Install fresh and swag
RUN go install github.com/gravityblast/fresh@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

# Copy the project files
COPY . .

# Expose the port your app runs on (default 8080, can be overridden by env variable)
EXPOSE 8080

# Set environment variables for fresh
ENV GIN_MODE=debug
ENV GO111MODULE=on

# Install dependencies
RUN go mod download

# Start the application in hot reload mode using Makefile
CMD ["make", "hot"]
