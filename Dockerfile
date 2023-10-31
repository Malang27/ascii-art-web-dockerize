FROM golang:1.19

RUN mkdir /app

LABEL Description="malang coulibaly"

LABEL version="1.0"

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN apt-get update && apt-get install -y bash

RUN go build -o main .

CMD  ["/app/main"]
