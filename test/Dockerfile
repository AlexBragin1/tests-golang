FROM ubuntu:latest
LABEL authors="alexbragin"

FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o banner-click-counter cmd/main.go

CMD ["./banner-click-counter"]
