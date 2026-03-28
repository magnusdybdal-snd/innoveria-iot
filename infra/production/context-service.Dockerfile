# build stage
FROM golang:1.25.6 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy only necessary files
COPY pkg ./pkg
COPY services/context-service ./services/context-service

WORKDIR /app/services/context-service

RUN go mod download
RUN go build -ldflags='-s -w' -o server ./cmd/main.go

# Run stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/services/context-service/server /app/server

USER nonroot:nonroot

CMD ["/app/server"]
