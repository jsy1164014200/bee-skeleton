FROM golang

MAINTAINER gng <gng@bingyan.net>
WORKDIR /go/src/kuakuaAi
COPY . .
RUN go build
