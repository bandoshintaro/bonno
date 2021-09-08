FROM golang:1.6
MAINTAINER bando

RUN apt update && apt-get install -y sqlite3 libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

RUN set -x \
    && go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel \
    && go get gopkg.in/gorp.v1 \
    && go get github.com/mattn/go-sqlite3

ADD ./ $GOPATH/src/bonno

CMD revel run bonno prod
EXPOSE 9000
