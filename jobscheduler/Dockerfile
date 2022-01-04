FROM golang:1.16

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go mod tidy

CMD go run main.go