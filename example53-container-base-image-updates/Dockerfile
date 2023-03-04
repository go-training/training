# syntax=docker/dockerfile:1
FROM alpine:3.17

ARG TARGETOS
ARG TARGETARCH

WORKDIR /

COPY release/${TARGETOS}/${TARGETARCH}/example53 /bin/

EXPOSE 8080

ENTRYPOINT ["/bin/example53"]
