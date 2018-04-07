FROM plugins/base:multiarch

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

ADD server /bin/

EXPOSE 8080

HEALTHCHECK --start-period=2s --interval=10s --timeout=5s \
  CMD ["/bin/server", "-ping"]

ENTRYPOINT ["/bin/server"]
