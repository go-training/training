# syntax=docker/dockerfile:1
FROM scratch

WORKDIR /

COPY release/${TARGETOS}/${TARGETARCH}/example53 /bin/

EXPOSE 8080

ENTRYPOINT ["/example53"]
