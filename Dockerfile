FROM golang:1.23.2-bookworm AS builder

WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o ./api

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/api ./api
CMD ["/app/api"]