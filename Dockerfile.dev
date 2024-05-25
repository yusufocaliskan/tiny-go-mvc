# Dockerfile

FROM golang:1.21.3

# Install fresh and swag
RUN go install github.com/gravityblast/fresh@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY . .

EXPOSE 8080

ENV GIN_MODE=debug
ENV GO111MODULE=on

# Install dependencies
RUN go mod download

# Start the application in hot reload mode using Makefile
CMD ["make", "hot"]
