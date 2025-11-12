# Backend Dockerfile
FROM golang:1.25.4-alpine AS build

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the Go binary
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/main .
COPY .env .

EXPOSE 8080
CMD ["./main"]
