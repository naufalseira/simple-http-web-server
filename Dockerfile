# Stage 1: Build the Go application
FROM golang:1.23.2-bookworm AS builder

# Set the working directory for the build stage
WORKDIR /build

# Copy all files to download dependencies
COPY . .
RUN go mod download

# Build the application
RUN go build -o ./api ./app/main.go

# Stage 2: Create a lightweight container to run the application
FROM gcr.io/distroless/base-debian12 AS runtime

# Set the working directory for the runtime stage
WORKDIR /run

# Copy the built binary from the builder stage
COPY --from=builder /build/api ./api

# Copy assets folder to the runtime container
COPY --from=builder /build/app/assets ./assets

# Run the application
CMD ["/run/api"]
