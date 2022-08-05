FROM golang:1.16-alpine as builder

RUN go version
RUN apk update

WORKDIR /inventory

RUN mkdir libs/
RUN mkdir types/

COPY libs/ ./libs
COPY types/ ./types
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./

RUN ls

RUN pwd
RUN go build -o main.go .

EXPOSE 9678

CMD ./main.go
