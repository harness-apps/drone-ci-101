# syntax=docker/dockerfile:1.4

FROM gcr.io/distroless/base

LABEL org.opencontainers.image.source https://github.com/harness-apps/drone-ci-101

ARG TARGETARCH

COPY server_linux_${TARGETARCH}/server /bin/server

EXPOSE 8080

ENTRYPOINT ["server"]