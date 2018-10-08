.SILENT :

# Author
AUTHOR=github.com/ncarlier

# App name
APPNAME=feedpushr

# Go configuration
GOOS?=linux
GOARCH?=amd64

# Add exe extension if windows target
is_windows:=$(filter windows,$(GOOS))
EXT:=$(if $(is_windows),".exe","")

# Go app path
APPBASE=${GOPATH}/src/$(AUTHOR)

# Artefact name
ARTEFACT=release/$(APPNAME)-$(GOOS)-$(GOARCH)$(EXT)
ARTEFACT_CTL=release/$(APPNAME)-ctl-$(GOOS)-$(GOARCH)$(EXT)

# Extract version infos
VERSION:=`git describe --tags`
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

all: build

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/makefiles
include $(makefiles)/help.Makefile

$(APPBASE)/$(APPNAME):
	echo ">>> Creating GO src link: $(APPBASE)/$(APPNAME) ..."
	mkdir -p $(APPBASE)
	ln -s $(root_dir) $(APPBASE)/$(APPNAME)

## Clean built files
clean:
	echo ">>> Removing generated files ..."
	-rm -rf release autogen var/assets/ui pkg/assets/statik.go
.PHONY: clean

deps:
	echo ">>> Installing dependencies ..."
	cd $(APPBASE)/$(APPNAME) && dep ensure
.PHONY: deps

Gopkg.lock:
	make deps

## Run code generation
autogen:
	echo ">>> Generating code ..."
	#go get -u github.com/goadesign/goa/...
	-mv vendor vendor_bak
	cd $(APPBASE)/$(APPNAME) && goagen bootstrap -o autogen -d $(AUTHOR)/$(APPNAME)/design
	mv vendor_bak vendor
	cp -f autogen/swagger/** var/assets/

## Build web UI
ui:
	-rm -rf pkg/assets/statik.go var/assets/ui
	make pkg/assets/statik.go
.PHONY: ui

## Start web UI dev server
ui-dev-server:
	cd ui && REACT_APP_API_ROOT="http://localhost:8080/v1" npm start
.PHONY: ui-dev-server

# Build web UI
var/assets/ui:
	echo ">>> Building web UI ..."
	cd ui && npm install && npm run-script build
	mv ui/build var/assets/ui

# Build assets as Go file
pkg/assets/statik.go:
	make var/assets/ui
	echo ">>> Generating \"pkg/assets/statik.go\" ..."
	go get -u github.com/rakyll/statik
	statik -p assets -src var/assets -dest pkg -f

## Build executable
build: Gopkg.lock autogen pkg/assets/statik.go $(APPBASE)/$(APPNAME)
	-mkdir -p release
	echo ">>> Building: $(ARTEFACT) ..."
	cd $(APPBASE)/$(APPNAME) && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(ARTEFACT)
	cd $(APPBASE)/$(APPNAME)/autogen/tool/$(APPNAME)-cli && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(APPBASE)/$(APPNAME)/$(ARTEFACT_CTL)
.PHONY: build

$(ARTEFACT): build

## Run tests
test:
	-golint pkg/...
	cd $(APPBASE)/$(APPNAME) && go test `go list ./... | grep -v vendor`
.PHONY: test

## Install executable
install: $(ARTEFACT)
	echo ">>> Installing $(ARTEFACT) to ${HOME}/.local/bin/$(APPNAME) ..."
	cp $(ARTEFACT) ${HOME}/.local/bin/$(APPNAME)
.PHONY: install

## Create Docker image
image:
	echo ">>> Building Docker inage ..."
	docker build --rm -t ncarlier/$(APPNAME) .
.PHONY: image

## Generate changelog
changelog:
	standard-changelog --first-release
.PHONY: changelog

## GZIP executable
gzip:
	gzip $(ARTEFACT)
	gzip $(ARTEFACT_CTL)
.PHONY: gzip

## Create distribution binaries
distribution:
	GOARCH=amd64 make build gzip
	GOARCH=arm64 make build gzip
	GOARCH=arm make build gzip
	GOOS=darwin make build gzip
.PHONY: distribution
