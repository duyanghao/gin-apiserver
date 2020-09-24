## server version
SERVER_VERSION = v0.1.0
## Folder content generated files
BUILD_FOLDER = ./build
PROJECT_URL  = github.com/duyanghao/gin-apiserver
## command
GO           = go
GO_VENDOR    = go mod
MKDIR_P      = mkdir -p

## Random Alphanumeric String
SECRET_KEY   = $(shell cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

## UNAME
UNAME := $(shell uname)

################################################

.PHONY: all
all: build

.PHONY: pre-build
pre-build:
	$(MAKE) download

.PHONY: build
build: pre-build build.docs
	$(MAKE) src.build

.PHONY: clean
clean:
	$(RM) -rf $(BUILD_FOLDER)

## vendor/ #####################################

.PHONY: download
download:
	$(GO_VENDOR) vendor

## src/ ########################################

.PHONY: src.build
src.build:
	$(MKDIR_P) $(BUILD_FOLDER)/pkg/cmd/gin-apiserver/
	GO111MODULE=on $(GO) build -mod=vendor -v -o $(BUILD_FOLDER)/pkg/cmd/gin-apiserver/gin-apiserver \
	./cmd/...

## dockerfiles/ ########################################

.PHONY: dockerfiles.build
dockerfiles.build:
	docker build --no-cache --rm --tag duyanghao/gin-apiserver:$(SERVER_VERSION) -f ./docker/Dockerfile .

## git tag version ########################################

.PHONY: push.tag
push.tag:
	@echo "Current git tag version:"$(SERVER_VERSION)
	git tag $(SERVER_VERSION)
	git push --tags

.PHONY: build.docs
build.docs:
	@echo "Build swagger docs"
	swag init -g pkg/route/routes.go
