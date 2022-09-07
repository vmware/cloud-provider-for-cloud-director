GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
GITROOT := $(shell git rev-parse --show-toplevel)
GO_CODE := $(shell ls go.mod go.sum **/*.go)
version := $(shell cat ${GITROOT}/release/version)

REGISTRY?="harbor-repo.vmware.com/vcloud"

.PHONY: build-within-docker vendor

build-within-docker:
	mkdir -p /build/cloud-provider-for-cloud-director
	go build -ldflags "-X github.com/vmware/cloud-provider-for-cloud-director/version.Version=$(version)" -o /build/vcloud/cloud-provider-for-cloud-director cmd/ccm/main.go

ccm: $(GO_CODE)
	docker build -f Dockerfile . -t cloud-provider-for-cloud-director:$(version)
	docker tag cloud-provider-for-cloud-director:$(version) $(REGISTRY)/cloud-provider-for-cloud-director:$(version)
	docker tag cloud-provider-for-cloud-director:$(version) $(REGISTRY)/cloud-provider-for-cloud-director:$(version).$(GITCOMMIT)
	docker push $(REGISTRY)/cloud-provider-for-cloud-director:$(version)
	touch out/$@

prod: ccm
	sed -e "s/\.__GIT_COMMIT__//g" manifests/cloud-director-ccm.yaml.template > manifests/cloud-director-ccm.yaml
	sed -e "s/\.__GIT_COMMIT__//g" manifests/cloud-director-ccm-crs.yaml.template > manifests/cloud-director-ccm-crs.yaml

dev: ccm
	docker push $(REGISTRY)/cloud-provider-for-cloud-director:$(version).$(GITCOMMIT)
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" manifests/cloud-director-ccm.yaml.template > manifests/cloud-director-ccm.yaml
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" manifests/cloud-director-ccm-crs.yaml.template > manifests/cloud-director-ccm-crs.yaml

vendor:
	go mod edit -go=1.17
	go mod tidy
	go mod vendor

test:
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdclient -cover -count=1
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/config -cover -count=1

integration-test: test
	go test -tags="testing integration" -v github.com/vmware/cloud-provider-for-cloud-director/vcdclient -cover -count=1
