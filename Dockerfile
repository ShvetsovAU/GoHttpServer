FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/GoHttpServer

RUN go get \
 github.com/andrushk/calculations \
 github.com/andrushk/excgor \
 github.com/kikinteractive/go-gcm \
 github.com/gorilla/context \
 github.com/jessevdk/go-flags \
 github.com/jpillora/backoff \
 github.com/julienschmidt/httprouter \
 github.com/justinas/alice \
 github.com/mattn/go-xmpp \
 github.com/pborman/uuid \
 github.com/robfig/cron \
 gopkg.in/gomail.v2 \
 gopkg.in/mgo.v2 \
 github.com/streadway/amqp \
 github.com/gorilla/schema \
 github.com/wuman/firebase-server-sdk-go

RUN go install github.com/shvetsovau/GoHttpServer

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/GoHttpServer"]

EXPOSE 8083