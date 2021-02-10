LAST_COMMIT_SHA = $(shell git rev-parse --short HEAD)
GIT_TREE_STATE = $(shell (git status --porcelain | grep -q .) && echo dirty || echo clean)

all: bin/agora

PLATFORM=local

.PHONY: bin/agora
bin/agora:
	@docker build . --target bin --output bin/ --platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: build-image
build-image:
	@docker build . --target build-image --platform linux/amd64 \
	-t agora:latest

ifeq ($(GIT_TREE_STATE), clean)
	@docker tag agora:latest agora:$(LAST_COMMIT_SHA)
endif