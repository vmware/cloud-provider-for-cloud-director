FROM golang:1.19 AS builder

RUN apt-get update && \
    apt-get -y install \
        bash \
        git  \
        make

ADD . /go/src/github.com/vmware/cloud-provider-for-cloud-director
WORKDIR /go/src/github.com/vmware/cloud-provider-for-cloud-director

ENV GOPATH /go
RUN make build-within-docker && \
    chmod +x /build/vcloud/cloud-provider-for-cloud-director

########################################################

FROM scratch

WORKDIR /opt/vcloud/bin

# copy multiple files at a time to create a single layer
COPY --from=builder /go/src/github.com/vmware/cloud-provider-for-cloud-director/LICENSE.txt /go/src/github.com/vmware/cloud-provider-for-cloud-director/NOTICE.txt /go/src/github.com/vmware/cloud-provider-for-cloud-director/open_source_license.txt .
COPY --from=builder /build/vcloud/cloud-provider-for-cloud-director .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER 65534
ENTRYPOINT ["/opt/vcloud/bin/cloud-provider-for-cloud-director"]
