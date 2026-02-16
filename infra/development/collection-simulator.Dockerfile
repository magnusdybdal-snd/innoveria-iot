FROM golang:1.25.6

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY services/collection-simulator/go.mod ./

CMD ["air"]
