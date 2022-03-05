PROJECT_ROOT          = github.com/alexey-zayats/netspeed
BUILD_PATH           ?= bin
DOCKERFILE_PATH      ?= $(CURDIR)/deploy/docker

IMAGE_VERSION        ?= $(shell git describe --tags --long)
BRANCH               ?= $(shell git rev-parse --abbrev-ref HEAD)

NETSPEED_BINARY      ?= $(BUILD_PATH)/server
NETSPEED_PATH        ?= $(PROJECT_ROOT)/cmd/server
NETSPEED_DOCKERFILE  ?= $(DOCKERFILE_PATH)/Dockerfile.netspeed
NETSPEED_IMAGE       ?= aazayats/netspeed

GO_TEST_FLAGS        ?= -v -cover
GO_PACKAGES          := $(shell go list ./... | grep -v vendor)
GIT_BRANCH           := $(shell git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')
CWD                  := $(shell pwd)
COVER_PROFILE        := /tmp/go-cover.tmp

TAG_CMD              ?= patch # major, minor, patch

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

VERSION := $(shell git fetch --tags && git tag --sort=-v:refname --list "v[0-9]*" | head -n 1)
ifeq ($(VERSION),)
	VERSION := "0.0.0"
endif

fmt:
	$Q go fmt $(GO_PACKAGES)

test:
	$Q go test $(GO_TEST_FLAGS) ./...

bench:
	$Q go test -bench=.

test-cover:
	$Q go test -v -coverprofile $(COVER_PROFILE) ./...
	$Q go tool cover -html=$(COVER_PROFILE) && rm $(COVER_PROFILE)

.PHONY tag:
tag: | ; $(info $(M) bump version…) @
	$Q $(eval VERSION := $(shell docker run --rm -v "$PWD":/app treeder/bump --input "$(VERSION)" $(TAG_CMD)))
	$Q git tag -a "v$(VERSION)" -m "version $(VERSION)"
	$Q git push --follow-tags

.PHONY all:
all: binary

.PHONY binary:
binary: | ; $(info $(M) build netspeed binary…) @
	$Q go build -o $(NETSPEED_BINARY) $(NETSPEED_PATH)

.PHONY docker:
docker: ## build docker images
	docker-netspeed

docker-netspeed: | ; $(info $(M) building netspeed docker image…) @ ## build server docker image
	$Q docker build -f $(NETSPEED_DOCKERFILE) -t $(NETSPEED_IMAGE):$(IMAGE_VERSION) -t $(NETSPEED_IMAGE):latest .
	$Q docker image prune -f --filter label=stage=netspeed-intermediate
