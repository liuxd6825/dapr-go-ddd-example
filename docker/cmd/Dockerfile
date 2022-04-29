# current directory must be ./dist

FROM gcr.io/distroless/static:nonroot
ARG PKG_FILES
WORKDIR /
COPY /config     /
COPY /$PKG_FILES /

EXPOSE 9010
EXPOSE 9011
EXPOSE 9012

ENTRYPOINT [ "cmd-service -envType=test" ]
