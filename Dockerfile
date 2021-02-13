FROM golang:1.15.2-alpine

WORKDIR /go/src/calender
RUN apk add --no-cache git
ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=

RUN go get github.com/go-sql-driver/mysql &&\
    go get github.com/beego/bee &&\
    go get github.com/astaxie/beego

CMD bee run