FROM golang:1.25.6

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY services/context-service/go.mod ./

EXPOSE 8080

CMD ["air"]
