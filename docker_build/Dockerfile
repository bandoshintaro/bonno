FROM ubuntu
MAINTAINER bando

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get upgrade -y
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y software-properties-common python-software-properties build-essential sysv-rc-conf curl git
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y sqlite3 libsqlite3-dev
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y mercurial make binutils bison gcc

ADD gvminstall.sh /usr/local/bin/gvminstall.sh
RUN chmod u+x /usr/local/bin/gvminstall.sh
RUN /usr/local/bin/gvminstall.sh
ENV GOROOT=/root/.gvm/gos/go1.5.3
ENV GOPATH=/root/.gvm/pkgsets/go1.5.3/global
ENV PATH=$PATH:$GOPATH/bin:$GOROOT/bin
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get gopkg.in/gorp.v1
RUN go get github.com/mattn/go-sqlite3

CMD /bin/bash

EXPOSE 9000
