FROM --platform=${BUILDPLATFORM} golang:1.15.7-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/agora .

FROM scratch AS bin-unix
COPY --from=build /out/agora /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/agora /agora.exe

FROM bin-${TARGETOS} AS bin