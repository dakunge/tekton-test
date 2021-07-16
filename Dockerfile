FROM alpine:3.11

WORKDIR /tkdir
COPY tk-homework /tkdir/tk-homework
ENTRYPOINT ["/tk-homework"]
