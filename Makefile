LAST_COMMIT_SHA = $(shell git rev-parse --short HEAD)
GIT_TREE_STATE = $(shell (git status --porcelain | grep -q .) && echo dirty || echo clean)

all: bin/agora

PLATFORM=local

.PHONY: bin/agora
bin/agora:
	@DOCKER_BUILDKIT=1 docker build . --target bin --output bin/ --platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@DOCKER_BUILDKIT=1 docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@DOCKER_BUILDKIT=1 docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@DOCKER_BUILDKIT=1 docker build . --target lint

.PHONY: build-image
build-image:
	@DOCKER_BUILDKIT=1 docker build . --target build-image --platform linux/amd64 \
	-t agora:latest

	@DOCKER_BUILDKIT=1 docker build database/migrations/ -t agora-migrate:latest

ifeq ($(GIT_TREE_STATE), clean)
	@docker tag agora:latest agora:$(LAST_COMMIT_SHA)
	@docker tag agora-migrate:latest agora-migrate:$(LAST_COMMIT_SHA)
endif