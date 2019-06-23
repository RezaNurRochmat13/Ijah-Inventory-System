FROM frolvlad/alpine-glibc

COPY log log

COPY environment environment

COPY ijah ijah

RUN apk add --no-cache bash

ADD main /

EXPOSE 8080

CMD ["/main"]
