FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/gohttpserver
ADD cert.pem /go/bin/cert.pem
ADD key.pem /go/bin/key.pem

RUN go get \
 github.com/gorilla/context \
 github.com/jessevdk/go-flags \
 github.com/julienschmidt/httprouter \
 github.com/justinas/alice \
 gopkg.in/mgo.v2

RUN go install github.com/shvetsovau/gohttpserver

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/gohttpserver"]

EXPOSE 8083