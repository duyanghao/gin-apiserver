## server version
SERVER_VERSION = v0.1.0
## Folder content generated files
BUILD_FOLDER = ./build
PROJECT_URL  = github.com/duyanghao/GinApiServer
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
all: build test

.PHONY: pre-build
pre-build:
	$(MAKE) download

.PHONY: build
build: pre-build build.docs
	$(MAKE) src.build

.PHONY: test
test: build
	$(MAKE) src.test

.PHONY: clean
clean:
	$(RM) -rf $(BUILD_FOLDER)

## vendor/ #####################################

.PHONY: download
download:
	$(GO_VENDOR) download

## src/ ########################################

.PHONY: src.build
src.build:
	$(MKDIR_P) $(BUILD_FOLDER)/pkg/cmd/GinApiServer/
	GO111MODULE=on $(GO) build -mod=vendor -v -o $(BUILD_FOLDER)/pkg/cmd/GinApiServer/GinApiServer \
	./cmd/...

.PHONY: src.test
src.test:
	$(GO) test -count=1 -v ./src/...

.PHONY: src.install
src.install:
	GO111MODULE=on $(GO) install -v ./src/...

## dockerfiles/ ########################################

.PHONY: dockerfiles.build
dockerfiles.build:
	docker build --no-cache --rm --tag duyanghao/ginapiserver:$(SERVER_VERSION) -f ./docker/Dockerfile .

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
