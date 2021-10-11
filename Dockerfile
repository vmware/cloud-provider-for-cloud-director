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

FROM photonos-docker-local.artifactory.eng.vmware.com/photon4:4.0-GA

WORKDIR /opt/vcloud/bin

COPY --from=builder /go/src/github.com/vmware/cloud-provider-for-cloud-director/LICENSE.txt .
COPY --from=builder /go/src/github.com/vmware/cloud-provider-for-cloud-director/NOTICE.txt .
COPY --from=builder /go/src/github.com/vmware/cloud-provider-for-cloud-director/open_source_license.txt .
COPY --from=builder /build/vcloud/cloud-provider-for-cloud-director .

RUN chmod +x /opt/vcloud/bin/cloud-provider-for-cloud-director

# USER nobody
ENTRYPOINT ["/bin/bash", "-l", "-c"]
