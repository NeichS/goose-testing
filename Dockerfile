FROM golang:1.24.6 AS builder

WORKDIR /app

COPY ./app/ /app

RUN go install github.com/air-verse/air@latest

RUN go install github.com/pressly/goose/v3/cmd/goose@latest


# RUN go build -o main .

CMD ["air", "-c", ".air.toml"]