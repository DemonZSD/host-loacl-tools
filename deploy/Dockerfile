FROM alpine:3.10
LABEL maintainer="Weshzhu"
WORKDIR /opt/host-local-tools
RUN mkdir -p /opt/cni/
ADD ./host-local-tools /opt/host-local-tools/
ADD ./resource /opt/host-local-tools/resource
ENTRYPOINT [ "./host-local-tools"]
