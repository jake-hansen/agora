# syntax = docker/dockerfile:1-experimental
# special thanks to https://github.com/chris-crone/containerized-go-dev

FROM --platform=${BUILDPLATFORM} golang:1.15.7-alpine AS base
RUN apk --no-cache add curl
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/agora .

FROM base AS unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir /out && go test ./... -v -coverprofile=/out/cover.out ./...

FROM golangci/golangci-lint:v1.31.0-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 10m0s ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM base AS build-image
WORKDIR /app
COPY --from=build /out/agora .
ENTRYPOINT [ "/app/agora" ]

FROM scratch AS bin-unix
COPY --from=build /out/agora /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/agora /agora.exe

FROM bin-${TARGETOS} AS bin
