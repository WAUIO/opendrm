FROM golang:1.17 as builder

ENV ROOTPATH=github.com/wauio/elacity-drm
ENV GO111MODULE=on

ADD . /go${ROOTPATH}
WORKDIR /go${ROOTPATH}
RUN CGO_ENABLED=0 GOOS=linux \
      && go build -ldflags "-X github.com/open-zhy/elacity-drm/cmd.Version=$(git describe --tags) -X github.com/open-zhy/ibcore/cmd.Build=$(date +'%Y%m%d%H%M%S')" -a -o /go/bin/elacity-drm

FROM alpine
RUN apk update && \
    apk add ca-certificates libxml2-dev && \
    rm -rf /var/cache/apk/*
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN touch /.env
COPY --from=builder /go/bin/elacity-drm /usr/bin
ENTRYPOINT ["/usr/bin/elacity-drm"]
