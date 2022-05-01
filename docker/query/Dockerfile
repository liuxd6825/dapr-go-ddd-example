# current directory must be ./dist

FROM gcr.io/distroless/static:nonroot
ARG PKG_FILES

COPY /config /app/config
COPY /*  /app/

WORKDIR /app

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

ENTRYPOINT [ "./query-service","-envType","test"]
