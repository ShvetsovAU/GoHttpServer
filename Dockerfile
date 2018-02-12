FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/GoHttpServer

RUN go get \
 github.com/gorilla/context \
 github.com/jessevdk/go-flags \
 github.com/julienschmidt/httprouter \

RUN go install github.com/shvetsovau/GoHttpServer

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/GoHttpServer"]

EXPOSE 8083