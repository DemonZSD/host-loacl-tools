FROM golang:latest
LABEL maintainer="Weshzhu"
WORKDIR /opt/host-local-tools
ENV CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOPATH=/opt/host-local-tools
COPY ./src /opt/host-local-tools/src
COPY ./src/config-writer/resource /opt/host-local-tools/resource

RUN cd /opt/host-local-tools/src/config-writer && \
    go build -o $GOPATH/host-local-tools
# RUN chmod +x 777 $GOPATH/host-local-tools
ENTRYPOINT [ "./host-local-tools"]
