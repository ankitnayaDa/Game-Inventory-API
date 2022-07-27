FROM postgres:10.0-alpine

USER postgres

RUN chmod 0700 /var/lib/postgresql/data &&\
    initdb /var/lib/postgresql/data &&\
    echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf &&\
    pg_ctl start &&\
    psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'games'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE games" &&\
    psql -c "ALTER USER postgres WITH ENCRYPTED PASSWORD 'postgres';"

EXPOSE 5432

FROM golang:1.14-alpine as builder

RUN go version
RUN apk update

ENV GOPATH=/root/go

RUN mkdir -p $GOPATH/src
RUN mkdir -p $GOPATH/bin
WORKDIR $GOPATH/src/gameInventoryAPI

COPY src/ $GOPATH/src

RUN pwd
RUN go build -o main.go .

EXPOSE 9678

CMD ./main.go

