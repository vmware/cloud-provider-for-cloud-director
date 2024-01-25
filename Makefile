LINT_VERSION := 1.51.2
GOSEC_VERSION := "v2.16.0"

GOLANGCI_EXIT_CODE ?= 1
# Set PATH to pick up cached tools. The additional 'sed' is required for cross-platform support of quoting the args to 'env'
SHELL := /usr/bin/env PATH=$(shell echo $(GITROOT)/bin:${PATH} | sed 's/ /\\ /g') bash
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
	GOBIN = $(shell go env GOPATH)/bin
else
	GOBIN = $(shell go env GOBIN)
endif

GITCOMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null)
GITROOT ?= $(shell git rev-parse --show-toplevel)

CPI_IMG := cloud-provider-for-cloud-director
ARTIFACT_IMG := cpi-crs-airgapped
VERSION ?= $(shell cat $(GITROOT)/release/version)

PLATFORM ?= linux/amd64
OS ?= linux
ARCH ?= amd64
CGO ?= 0

GOLANGCI_LINT ?= bin/golangci-lint
GOSEC ?= bin/gosec
SHELLCHECK ?= bin/shellcheck


.PHONY: all
all: vendor lint dev

.PHONY: cpi
cpi: vendor lint docker-build-cpi ## Run checks, and build cpi docker image.

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php



help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)



##@ Development

.PHONY: vendor
vendor: ## Update go mod dependencies.
	go mod edit -go=1.20
	go mod tidy -compat=1.20
	go mod vendor

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Run golangci-lint against code.
	$(GOLANGCI_LINT) run --issues-exit-code $(GOLANGCI_EXIT_CODE)

.PHONY: gosec
gosec: $(GOSEC) ## Run gosec against code.
	$(GOSEC) -conf .gosec.json ./...

.PHONY: shellcheck
shellcheck: $(SHELLCHECK) ## Run shellcheck against code.
	find . -name '*.*sh' -not -path '*/vendor/*' | xargs $(SHELLCHECK) --color

.PHONY: lint
lint: lint-deps golangci-lint gosec shellcheck ## Run golangci-lint, gosec, shellcheck.

.PHONY: lint-fix
lint-fix: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fix

.PHONY: test
test:
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdclient -cover -count=1
	go test -tags testing -v github.com/vmware/cloud-provider-for-cloud-director/pkg/config -cover -count=1

.PHONY: integration-test
integration-test: test
	go test -tags="testing integration" -v github.com/vmware/cloud-provider-for-cloud-director/vcdclient -cover -count=1

##@ Gobuild

.PHONY: gobuild
gobuild: vendor manifests release-prep build docker-build docker-archive publish

.PHONY: sandbox
# BUILD_TAG will be set during the build so it is not defined as it's not expected to be used by anything else here.
sandbox: VERSION := $(VERSION)-${BUILD_TAG}-$(GITCOMMIT)
sandbox: gobuild

.PHONY: official
official: gobuild

# docker-archive target saves the artifacts as .tar.gz to build/docker path which gets published as deliverables.
.PHONY: docker-archive
docker-archive: build/docker
	docker save -o build/docker/$(CPI_IMG)_$(VERSION).tar projects-stg.registry.vmware.com/vmware-cloud-director/$(CPI_IMG):$(VERSION)
	docker save -o build/docker/$(ARTIFACT_IMG)_$(VERSION).tar projects-stg.registry.vmware.com/vmware-cloud-director/$(ARTIFACT_IMG):$(VERSION)
	gzip build/docker/$(CPI_IMG)_$(VERSION).tar
	gzip build/docker/$(ARTIFACT_IMG)_$(VERSION).tar

# publish target publishes all the contents inside of build/docker. PUBLISH_DIR will be set during the build so it is not defined.
.PHONY: publish
publish:
	cp -R build/docker ${PUBLISH_DIR}

# build/docker is used as part of the gobuild process. In the end, we publish everything inside this folder. See target publish.
build/docker:
	mkdir -p build/docker

##@ Build

.PHONY: build
build: bin ## Build CSI binary. To be used from within a Dockerfile
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=$(CGO) go build -ldflags "-s -w -X github.com/vmware/cloud-provider-for-cloud-director/release.Version=$(VERSION)" -o bin/cloud-provider-for-cloud-director cmd/ccm/main.go

.PHONY: docker-build-cpi
docker-build-cpi: manifests build
	docker build  \
		--platform $(PLATFORM) \
		--file Dockerfile \
		--tag $(CPI_IMG):$(VERSION) \
		--build-arg CPI_BUILD_DIR=bin \
		.

.PHONY: docker-build-artifacts
docker-build-artifacts: release-prep
	docker build  \
		--platform $(PLATFORM) \
		--file artifacts/Dockerfile \
		--tag $(ARTIFACT_IMG):$(VERSION) \
		.

.PHONY: docker-build
docker-build: docker-build-cpi docker-build-artifacts ## Build CPI docker image and artifact image.

##@ Publish

.PHONY: dev
dev: VERSION := $(VERSION)-$(GITCOMMIT)
dev: git-check release ## Build development images and push to registry.

.PHONY: release
release: docker-build docker-push ## Build release images and push to registry.

.PHONY: release-prep
release-prep: ## Generate BOM and dependencies files.
	sed -e "s/__VERSION__/$(VERSION)/g" artifacts/default-cloud-director-ccm-crs-airgap.yaml.template > artifacts/cloud-director-ccm-crs-airgap.yaml.template
	sed -e "s/__VERSION__/$(VERSION)/g" artifacts/dependencies.txt.template > artifacts/dependencies.txt
	sed -e "s/__VERSION__/$(VERSION)/g" artifacts/bom.json.template > artifacts/bom.json

.PHONY: manifests
manifests: ## Generate CPI manifests
	sed -e "s/__VERSION__/$(VERSION)/g" manifests/cloud-director-ccm.yaml.template > manifests/cloud-director-ccm.yaml
	sed -e "s/__VERSION__/$(VERSION)/g" manifests/cloud-director-ccm-crs.yaml.template > manifests/cloud-director-ccm-crs.yaml

.PHONY: docker-push-cpi
docker-push-cpi: # Push CPI image to registry.
	docker tag $(CPI_IMG):$(VERSION) projects-stg.registry.vmware.com/vmware-cloud-director/$(CPI_IMG):$(VERSION)
	docker push projects-stg.registry.vmware.com/vmware-cloud-director/$(CPI_IMG):$(VERSION)

.PHONY: docker-push-artifacts
docker-push-artifacts: # Push artifacts image to registry
	docker tag $(ARTIFACT_IMG):$(VERSION) projects-stg.registry.vmware.com/vmware-cloud-director/$(ARTIFACT_IMG):$(VERSION)
	docker push projects-stg.registry.vmware.com/vmware-cloud-director/$(ARTIFACT_IMG):$(VERSION)

.PHONY: docker-push
docker-push: docker-push-cpi docker-push-artifacts ## Push images to container registry.



##@ Dependencies

.PHONY: lint-deps
lint-deps: $(GOLANGCI_LINT) $(GOSEC) $(SHELLCHECK) ## Download lint dependencies locally.





.PHONY: clean
clean:
	rm -rf bin
	rm \
		artifacts/cloud-director-ccm-crs-airgap.yaml.template \
		artifacts/bom.json \
		artifacts/dependencies.txt

bin:
	mkdir -p bin

$(GOLANGCI_LINT): bin
	@set -o pipefail && \
		wget -q -O - https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GITROOT)/bin v$(LINT_VERSION);

$(GOSEC): bin
	@GOBIN=$(GITROOT)/bin go install github.com/securego/gosec/v2/cmd/gosec@${GOSEC_VERSION}

$(SHELLCHECK): bin
	@cd bin && \
		set -o pipefail && \
		wget -q -O - https://github.com/koalaman/shellcheck/releases/download/stable/shellcheck-stable.$$(uname).x86_64.tar.xz | tar -xJv --strip-components=1 shellcheck-stable/shellcheck && \
		chmod +x $(GITROOT)/bin/shellcheck

.PHONY: git-check
git-check:
	@git diff --exit-code --quiet artifacts/ cmd/ manifests/ pkg/ Dockerfile || (echo 'Uncommitted changes found. Please commit your changes before proceeding.'; exit 1)
