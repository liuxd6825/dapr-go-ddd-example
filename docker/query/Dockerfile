# current directory must be ./dist

FROM scratch
ARG PKG_FILES

COPY /config /app/config
COPY /*  /app/

WORKDIR /app

EXPOSE 8080
EXPOSE 3500
EXPOSE 50001

ENTRYPOINT [ "./query-service","-envType","test"]
