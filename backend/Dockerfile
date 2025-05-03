FROM golang:1.24

WORKDIR ${GOPATH}/image-storage/
COPY .. ${GOPATH}/image-storage/

RUN go build -o /build ./cmd  && \
go clean -cache -modcache

EXPOSE 8000

CMD ["/build"]