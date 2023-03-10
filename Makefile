GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
GITROOT := $(shell git rev-parse --show-toplevel)
GO_CODE := $(shell ls go.mod go.sum **/*.go)
version := $(shell cat ${GITROOT}/release/version)

REGISTRY?="harbor-repo.vmware.com/vcloud"

.PHONY: build-within-docker vendor

crs-artifacts-prod:
	sed -e "s/\.__GIT_COMMIT__//g" -e "s/__VERSION__/$(version)/g" artifacts/default-cloud-director-ccm-crs-airgap.yaml.template > artifacts/cloud-director-ccm-crs-airgap.yaml.template
	sed -e "s/\.__GIT_COMMIT__//g" -e "s/__VERSION__/$(version)/g" -e "s~__REGISTRY__~$(REGISTRY)~g" artifacts/dependencies.txt.template > artifacts/dependencies.txt
	sed -e "s/\.__GIT_COMMIT__//g" -e "s/__VERSION__/$(version)/g" -e "s~__REGISTRY__~$(REGISTRY)~g" artifacts/bom.json.template > artifacts/bom.json
	docker build -f ./artifacts/Dockerfile . -t cpi-crs-airgapped:$(version)
	docker tag cpi-crs-airgapped:$(version) $(REGISTRY)/cpi-crs-airgapped:$(version)
	docker push $(REGISTRY)/cpi-crs-airgapped:$(version)

crs-artifacts-dev:
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" -e "s/__VERSION__/$(version)/g" artifacts/default-cloud-director-ccm-crs-airgap.yaml.template > artifacts/cloud-director-ccm-crs-airgap.yaml.template
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" -e "s/__VERSION__/$(version)/g" -e "s~__REGISTRY__~$(REGISTRY)~g" artifacts/dependencies.txt.template > artifacts/dependencies.txt
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" -e "s/__VERSION__/$(version)/g" -e "s~__REGISTRY__~$(REGISTRY)~g" artifacts/bom.json.template > artifacts/bom.json
	docker build -f ./artifacts/Dockerfile . -t cpi-crs-airgapped:$(GITCOMMIT)
	docker tag cpi-crs-airgapped:$(GITCOMMIT) $(REGISTRY)/cpi-crs-airgapped:$(GITCOMMIT)
	docker push $(REGISTRY)/cpi-crs-airgapped:$(GITCOMMIT)

build-within-docker: vendor
	mkdir -p /build/cloud-provider-for-cloud-director
	go build -ldflags "-X github.com/vmware/cloud-provider-for-cloud-director/version.Version=$(version)" -o /build/vcloud/cloud-provider-for-cloud-director cmd/ccm/main.go

ccm: $(GO_CODE)
	docker build -f Dockerfile . -t cloud-provider-for-cloud-director:$(version)
	docker tag cloud-provider-for-cloud-director:$(version) $(REGISTRY)/cloud-provider-for-cloud-director:$(version)
	docker tag cloud-provider-for-cloud-director:$(version) $(REGISTRY)/cloud-provider-for-cloud-director:$(version).$(GITCOMMIT)
	docker push $(REGISTRY)/cloud-provider-for-cloud-director:$(version)
	touch out/$@

prod: ccm prod-manifest crs-artifacts-prod

dev: ccm dev-manifest crs-artifacts-dev
	docker push $(REGISTRY)/cloud-provider-for-cloud-director:$(version).$(GITCOMMIT)

vendor:
	go mod edit -go=1.17
	go mod tidy -compat=1.17
	go mod vendor

test:
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdclient -cover -count=1
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/config -cover -count=1

integration-test: test
	go test -tags="testing integration" -v github.com/vmware/cloud-provider-for-cloud-director/vcdclient -cover -count=1

dev-manifest:
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" -e "s/__VERSION__/$(version)/g" manifests/cloud-director-ccm.yaml.template > manifests/cloud-director-ccm.yaml
	sed -e "s/__GIT_COMMIT__/$(GITCOMMIT)/g" -e "s/__VERSION__/$(version)/g" manifests/cloud-director-ccm-crs.yaml.template > manifests/cloud-director-ccm-crs.yaml

prod-manifest:
	sed -e "s/\.__GIT_COMMIT__//g" -e "s/__VERSION__/$(version)/g" manifests/cloud-director-ccm.yaml.template > manifests/cloud-director-ccm.yaml
	sed -e "s/\.__GIT_COMMIT__//g" -e "s/__VERSION__/$(version)/g" manifests/cloud-director-ccm-crs.yaml.template > manifests/cloud-director-ccm-crs.yaml
