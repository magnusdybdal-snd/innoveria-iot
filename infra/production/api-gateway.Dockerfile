# bulid stage
FROM golang:1.25.6 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy only neccesary files
COPY pkg ./pkg
COPY services/api-gateway ./services/api-gateway

WORKDIR /app/services/api-gateway

RUN go mod download
RUN go build -ldflags='-s -w' -o server ./cmd/main.go

# Runstage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/services/api-gateway/server /app/server

USER nonroot:nonroot

EXPOSE 8080

CMD ["/app/server"]
