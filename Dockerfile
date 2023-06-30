FROM alpine:3.18 AS builder

WORKDIR /build/vcloud/

ARG CPI_BUILD_DIR
ADD ${CPI_BUILD_DIR}/cloud-provider-for-cloud-director .

RUN chmod +x /build/vcloud/cloud-provider-for-cloud-director
########################################################
FROM scratch

WORKDIR /opt/vcloud/bin

# copy multiple files at a time to create a single layer
COPY LICENSE.txt NOTICE.txt open_source_license.txt /opt/vcloud/bin/
COPY --from=builder /build/vcloud/cloud-provider-for-cloud-director .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER 65534
ENTRYPOINT ["/opt/vcloud/bin/cloud-provider-for-cloud-director"]
