.SILENT :

export GO111MODULE=on

# Base package
BASE_PACKAGE=github.com/ncarlier

# App name
APPNAME=feedpushr

# Go configuration
GOOS?=$(shell go env GOHOSTOS)
GOARCH?=$(shell go env GOHOSTARCH)

# Add exe extension if windows target
is_windows:=$(filter windows,$(GOOS))
EXT:=$(if $(is_windows),".exe","")
LDLAGS_LAUNCHER:=$(if $(is_windows),-ldflags "-H=windowsgui",)

# Archive name
ARCHIVE=$(APPNAME)-$(GOOS)-$(GOARCH).tgz

# Main executable name
MAIN_EXE=$(APPNAME)$(EXT)

# CLI executable name
CLI_EXE=$(APPNAME)-ctl$(EXT)

# Launcher filename
LAUNCHER_EXE=$(APPNAME)-launcher$(EXT)

# Plugin name
PLUGIN?=twitter

# Plugin filename
PLUGIN_SO=$(APPNAME)-$(PLUGIN).so

# Extract version infos
PKG_VERSION:=github.com/ncarlier/$(APPNAME)/v3/pkg/version
VERSION:=`git describe --tags`
GIT_COMMIT:=`git rev-list -1 HEAD --abbrev-commit`
BUILT:=`date`
define LDFLAGS
-X '$(PKG_VERSION).Version=$(VERSION)' \
-X '$(PKG_VERSION).GitCommit=$(GIT_COMMIT)' \
-X '$(PKG_VERSION).Built=$(BUILT)'
endef

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
	echo ">>> Generating code ..."
	goagen bootstrap -o autogen -d $(BASE_PACKAGE)/$(APPNAME)/v3/design
	echo ">>> Moving Swagger files to assets ..."
	cp -f $(root_dir)/autogen/swagger/** $(root_dir)/var/assets/

## Build web UI
ui:
	-rm -rf pkg/assets/statik.go var/assets/ui
	make pkg/assets/statik.go
.PHONY: ui

## Start web UI dev server
ui-dev-server:
	cd ui && REACT_APP_API_ROOT="http://localhost:8080" npm start
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

# Build SYSO Windows file
contrib/launcher/main.syso:
	# go get github.com/akavel/rsrc
	rsrc -arch="amd64" -ico var/assets/ui/favicon.ico -o contrib/launcher/main.syso

## Build executable
build: autogen pkg/assets/statik.go
	-mkdir -p release
	echo ">>> Building: $(MAIN_EXE) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o release/$(MAIN_EXE)
.PHONY: build

## Build CLI executable
cli:
	-mkdir -p release
	echo ">>> Building: $(CLI_EXE) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	cd ./autogen/tool/$(APPNAME)-cli && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o ../../../release/$(CLI_EXE)
.PHONY: cli

## Build launcher executable
launcher: contrib/launcher/main.syso
	echo ">>> Building: $(LAUNCHER_EXE) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	cd contrib/launcher && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDLAGS_LAUNCHER) -o ../../release/$(LAUNCHER_EXE)
.PHONY: launcher

release/$(MAIN_EXE): build

## Run tests
test:
	-golint pkg/...
	go test `go list ./... | grep -v autogen`
.PHONY: test

## Install executable
install: release/$(MAIN_EXE)
	echo ">>> Installing $(MAIN_EXE) to ${HOME}/.local/bin/$(MAIN_EXE) ..."
	cp release/$(MAIN_EXE) ${HOME}/.local/bin/$(MAIN_EXE)
.PHONY: install

## Create Docker image
image:
	echo ">>> Building Docker image ..."
	docker build --rm -t ncarlier/$(APPNAME) .
.PHONY: image

## Generate changelog
changelog:
	standard-changelog --first-release
.PHONY: changelog

## Create archive
archive:
	echo ">>> Creating release/$(ARCHIVE) archive..."
	tar czf release/$(ARCHIVE) \
		--exclude=*.tgz \
	 	README.md \
		LICENSE \
		CHANGELOG.md \
		-C release/ $(subst release/,,$(wildcard release/*))
	find release/ -type f -not -name '*.tgz' -delete
.PHONY: archive

## Create distribution binaries
distribution:
	GOOS=linux GOARCH=amd64 make build cli launcher plugins archive
	GOOS=linux GOARCH=arm64 make build cli archive
	GOOS=linux GOARCH=arm make build cli archive
	GOOS=windows make build cli launcher archive
	GOOS=darwin make build cli archive
.PHONY: distribution

## Bulid plugin (defined by PLUGIN variable)
plugin:
	-mkdir -p release
	echo ">>> Building: $(PLUGIN_SO) $(VERSION) for $(GOOS)-$(GOARCH) ..."
	cd contrib/$(PLUGIN) && GOOS=$(GOOS) GOARCH=$(GOARCH) go build -buildmode=plugin -o ../../release/$(PLUGIN_SO)
.PHONY: plugin

## Build all plugins
plugins:
	GOARCH=amd64 PLUGIN=twitter-selenium make plugin	
	GOARCH=amd64 PLUGIN=twitter make plugin
	GOARCH=amd64 PLUGIN=mastodon make plugin
	GOARCH=amd64 PLUGIN=kafka make plugin
	GOARCH=amd64 PLUGIN=rdbms make plugin
	GOARCH=amd64 PLUGIN=rake make plugin
	GOARCH=amd64 PLUGIN=prose make plugin
.PHONY: plugins
