.SILENT :

export GO111MODULE=on

# Base package
BASE_PACKAGE=github.com/ncarlier

# App name
APPNAME=feedpushr

# Go app path
APPBASE=${GOPATH}/src/$(BASE_PACKAGE)

# Go configuration
GOOS?=linux
GOARCH?=amd64

# Add exe extension if windows target
is_windows:=$(filter windows,$(GOOS))
EXT:=$(if $(is_windows),".exe","")

# Archive name
ARCHIVE=$(APPNAME)-$(GOOS)-$(GOARCH).tgz

# Executable name
EXECUTABLE=$(APPNAME)$(EXT)

# CTL executable name
CTL_EXECUTABLE=$(APPNAME)-ctl$(EXT)

# Extract version infos
VERSION:=`git describe --tags`
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

all: build

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/makefiles
include $(makefiles)/help.Makefile

## Clean built files
clean:
	echo ">>> Removing generated files ..."
	-rm -rf release autogen var/assets/ui pkg/assets/statik.go
.PHONY: clean

## Run code generation
autogen:
	-rm $(APPBASE)/$(APPNAME)
	echo ">>> Creating GO src link: $(APPBASE)/$(APPNAME) ..."
	mkdir -p $(APPBASE)
	ln -s $(root_dir) $(APPBASE)/$(APPNAME)
	echo ">>> Generating code ..."
	cd $(APPBASE)/$(APPNAME) && goagen bootstrap -o autogen -d $(BASE_PACKAGE)/$(APPNAME)/design
	echo ">>> Moving Swagger files to assets ..."
	cp -f autogen/swagger/** var/assets/
	echo ">>> Removing GO src link: $(APPBASE)/$(APPNAME) ..."
	rm $(APPBASE)/$(APPNAME)

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
build: autogen pkg/assets/statik.go
	-mkdir -p release
	echo ">>> Building: $(EXECUTABLE) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o release/$(EXECUTABLE)
	echo ">>> Building: $(CTL_EXECUTABLE) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	cd ./autogen/tool/$(APPNAME)-cli && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o ../../../release/$(CTL_EXECUTABLE)
.PHONY: build

release/$(EXECUTABLE): build

## Run tests
test:
	-golint pkg/...
	go test `go list ./... | grep -v autogen`
.PHONY: test

## Install executable
install: release/$(EXECUTABLE)
	echo ">>> Installing $(EXECUTABLE) to ${HOME}/.local/bin/$(EXECUTABLE) ..."
	cp release/$(EXECUTABLE) ${HOME}/.local/bin/$(EXECUTABLE)
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

## Create archive
archive:
	echo ">>> Creating release/$(ARCHIVE) archive..."
	tar czf release/$(ARCHIVE) README.md LICENSE CHANGELOG.md -C release/ $(EXECUTABLE) $(CTL_EXECUTABLE)
	rm release/$(EXECUTABLE) release/$(CTL_EXECUTABLE)
.PHONY: archive

## Create distribution binaries
distribution:
	GOARCH=amd64 make build archive
	GOARCH=arm64 make build archive
	GOARCH=arm make build archive
	GOOS=darwin make build archive
.PHONY: distribution

