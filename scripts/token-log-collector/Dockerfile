FROM golang:1.18 as builder

ENV GOOS=linux
ENV GOARCH=amd64


RUN apt-get update && \
    apt-get -y install \
        bash \
        git  \
        make

ADD . /go/src/gitlab.eng.vmware.com/shamasundara/token-log-collector
WORKDIR /go/src/gitlab.eng.vmware.com/shamasundara/token-log-collector

ENV GOPATH /go
ENV CONVERSION_GEN_BIN token-log-collector

RUN mkdir -p /build/vcloud
RUN go build  -o /build/vcloud/token-log-collector cmd/main.go

