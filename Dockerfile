FROM bando/revel
MAINTAINER bando

ADD ./ $GOPATH/src/bonno

CMD revel run bonno prod && bash

EXPOSE 9000
