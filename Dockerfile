FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN apt add bash

RUN go build -o main .

CMD  ["/app/main"]
