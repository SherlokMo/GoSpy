FROM golang:1.16

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

ADD . /usr/src/app/

RUN go mod download

VOLUME ["/usr/src/app" ]

EXPOSE 8080

CMD ["go run main.go"]