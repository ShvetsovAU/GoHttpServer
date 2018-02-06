FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/GoHttpServer

RUN go install github.com/shvetsovau/GoHttpServer

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/GoHttpServer"]

EXPOSE 8083