LAST_COMMIT_SHA := $(shell git rev-parse --short HEAD)

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
	-t agora:$(LAST_COMMIT_SHA) -t agora:latest