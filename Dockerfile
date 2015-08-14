FROM golang

MAINTAINER Toan Nguyen <ngdinhtoan@gmail.com>

RUN go get github.com/dancannon/gorethink

VOLUME /go/src/bm-rethinkdb

ENTRYPOINT go test -v -bench . bm-rethinkdb -host ${RETHINKDB_ALIAS_PORT_28015_TCP_ADDR}
