PROJECT_ROOT          = github.com/alexey-zayats/netspeed
BUILD_PATH           ?= bin
DOCKERFILE_PATH      ?= $(CURDIR)/deploy/docker

IMAGE_VERSION        ?= $(shell git describe --tags --long)
BRANCH               ?= $(shell git rev-parse --abbrev-ref HEAD)

NETSPEED_BINARY      ?= $(BUILD_PATH)/netspeed
NETSPEED_PATH        ?= $(PROJECT_ROOT)/cmd/netspeed
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


.PHONY all:
all: binary

.PHONY binary:
binary: netspeed-binary ## build binaries

.PHONY docker:
docker: docker-build docker-push ## build docker images

netspeed-binary: | ; $(info $(M) build netspeed binary…) @
	$Q go build -o $(NETSPEED_BINARY) $(NETSPEED_PATH)

docker-build: docker-netspeed

docker-netspeed: | ; $(info $(M) building netspeed docker image…) @
	$Q docker build -f $(NETSPEED_DOCKERFILE) -t $(NETSPEED_IMAGE):$(IMAGE_VERSION) -t $(NETSPEED_IMAGE):latest .
	$Q docker image prune -f --filter label=stage=netspeed-intermediate

docker-push: docker-push-netspeed ## push docker images

docker-push-netspeed:
	$Q docker push $(NETSPEED_IMAGE):$(IMAGE_VERSION)
	$Q docker push $(NETSPEED_IMAGE):latest

test: ## Run tests
	$Q go test $(GO_TEST_FLAGS) ./...

test-bench: ## Run benchmarks
	$Q go test -bench=.

test-cover: ## Run test with coverage
	$Q go test -v -coverprofile $(COVER_PROFILE) ./...
	$Q go tool cover -html=$(COVER_PROFILE) && rm $(COVER_PROFILE)

.PHONY tag:
tag: | ; $(info $(M) bump version…) @ ## Bump version by tags, use TAG_CMD as major, minor, patch, default is patch
	$Q $(eval VERSION := $(shell docker run --rm -v "$PWD":/app treeder/bump --input "$(VERSION)" $(TAG_CMD)))
	$Q git tag -a "v$(VERSION)" -m "version $(VERSION)"
	$Q git push --follow-tags


.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./...); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
		done ; exit $$ret

.PHONY: vendor
vendor: ; $(info $(M) vendoring…) @ ## Vendoring the app
	$Q go mod vendor
	$Q go mod tidy

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-23s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:  ## Show version
	@echo APP: $(VERSION)
	@echo IMAGE: $(IMAGE_VERSION)
