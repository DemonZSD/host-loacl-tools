FROM golang:latest
LABEL maintainer="Weshzhu"
ENV CGO_ENABLED=0 \
    GOARCH=amd64
WORKDIR $GOPATH/host-local-tools
COPY . $GOPATH/host-local-tools
COPY ./src/config-writer/resource $GOPATH/host-local-tools
RUN go build src/config-writer/host-local-tools.go -o $GOPATH/host-local-tools
ENTRYPOINT [ "host-local-tools"]