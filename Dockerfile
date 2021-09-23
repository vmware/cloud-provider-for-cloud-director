FROM golang:1.16 AS builder

RUN apt-get update && \
    apt-get -y install \
        bash \
        git  \
        make

ADD . /go/src/github.com/vmware/cloud-provider-for-cloud-director
WORKDIR /go/src/github.com/vmware/cloud-provider-for-cloud-director

ENV GOPATH /go
RUN ["make", "build-within-docker"]

########################################################

FROM photon:4.0-20210910

WORKDIR /opt/vcloud/bin

COPY --from=builder /build/vcloud/cloud-provider-for-cloud-director .

RUN chmod +x /opt/vcloud/bin/cloud-provider-for-cloud-director

# USER nobody
ENTRYPOINT ["/bin/bash", "-l", "-c"]
