FROM golang:1.8

WORKDIR /go/src/github.com/lucasmagnum/thanks/
ADD . .

RUN go get ./...

EXPOSE 8080
