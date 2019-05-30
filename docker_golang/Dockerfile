FROM golang:alpine

RUN apk update \
    && apk add --no-cache git \
    && git clone --depth=1 https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/globalsign/mgo

WORKDIR /usr/src/app

COPY . .

EXPOSE 8080

