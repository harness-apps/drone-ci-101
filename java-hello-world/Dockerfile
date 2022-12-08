FROM gcr.io/distroless/java11-debian11

LABEL org.opencontainers.image.source https://github.com/harness-apps/drone-ci-101

WORKDIR /app

COPY target/hello-world.jar /app

CMD [ "hello-world.jar" ]