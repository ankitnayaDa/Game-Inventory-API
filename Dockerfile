FROM golang:1.14-alpine as builder

RUN go version
RUN apk update

USER root

ENV GOPATH=/root/go

RUN mkdir -p $GOPATH/src
RUN mkdir -p $GOPATH/bin
WORKDIR $GOPATH/src/gameInventoryAPI

COPY src/ $GOPATH/src

RUN pwd
RUN go build -o main.go .

EXPOSE 9678

CMD ./main.go
