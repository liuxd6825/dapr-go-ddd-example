# current directory must be ./dist

FROM gcr.io/distroless/static:nonroot
ARG PKG_FILES
WORKDIR /
COPY /config     /
COPY /$PKG_FILES /

EXPOSE 9020
EXPOSE 9021
EXPOSE 9022

ENTRYPOINT [ "query-service" ]
