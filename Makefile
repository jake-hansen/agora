all: bin/agora

PLATFORM=local

.PHONY: bin/agora
bin/agora:
	@docker build . --target bin --output bin/ --platform ${PLATFORM}