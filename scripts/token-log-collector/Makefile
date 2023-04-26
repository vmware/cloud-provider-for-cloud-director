GITROOT := $(shell git rev-parse --show-toplevel)
TOKEN_CONTAINER := tokens_logger
TOKEN_BINARY_DIR := $(GITROOT)/scripts/token-log-collector/bin
TOKEN_BINARY := $(TOKEN_BINARY_DIR)/token-log-collector
all: bin-dir generate-binary
bin-dir:
	mkdir -p $(TOKEN_BINARY_DIR)
generate-binary: bin-dir
	docker build . -t token-loggger
	docker create -ti --name $(TOKEN_CONTAINER) token-loggger:latest bash
	docker cp $(TOKEN_CONTAINER):/build/vcloud/token-log-collector $(TOKEN_BINARY)
	docker rm $(TOKEN_CONTAINER)