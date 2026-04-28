# build stage
FROM golang:1.25.6 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# copy only required modules
COPY pkg ./pkg
COPY services/erp-agent-service ./services/erp-agent-service

WORKDIR /app/services/erp-agent-service

RUN go mod download
RUN go build -ldflags='-s -w' -o server ./cmd/main.go

# runtime stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/services/erp-agent-service/server /app/server

USER nonroot:nonroot

CMD ["/app/server"]
