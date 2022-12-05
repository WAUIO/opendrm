.PHONY:

NPROC=$(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS+=-j$(NPROCS)

TAG=$(shell git describe --tags)
BUILD=$(shell date +'%Y%m%d%H%M%S')
ARCH=$(shell go env GOARCH)
OS=$(shell go env GOOS)

fmt:
	@gofmt -l -w ./..

build: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -mod=readonly -ldflags "-X github.com/wauio/elacity-drm/cmd.Version=$(TAG) -X github.com/wauio/elacity-drm/cmd.Build=$(BUILD)" -o bin/elacity-drm

docker-push:
	docker push gcr.io/github.com/wauio/elacity-drm

image:
	docker build -t gcr.io/github.com/wauio/elacity-drm .
