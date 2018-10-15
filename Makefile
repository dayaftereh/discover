# Version and Relase 
VERSION	= 0.0.1
RELEASE = $(shell date '+%d-%m-%Y %H:%M:%S')

# Main package
MAIN_PKG =  github.com/dayaftereh/discover/server/main

# Executabels
NAME = discover

GOBIN = $(GOPATH)/bin

# GO commands
GO			= go
GOGET		= $(GO) get
GORUN 		= $(GO) run
GOBUILD		= $(GO) build
GOPKG 		= $(GOBIN)/dep
_GOOS		= linux
_GOARCH		= amd64

# GO FLAGS
GO_LD_FLAGS  = -X "main.VERSION=$(VERSION)"  -X 'main.RELEASE=$(RELEASE)'

# destination directory
DIST_DIR = $(CURDIR)/dist

OUTPUT = 

# Default Build variables
EXE_ENDOING =

.PHONY: clean

all: build-windows build-linux

# ------- release -------

release: GO_LD_FLAGS += -s -w
release: build-windows build-linux

# ------- build -------

build-windows: _GOOS=windows
build-windows: OUTPUT=$(DIST_DIR)/$(NAME)-$(_GOOS)-$(_GOARCH)-$(VERSION).exe
build-windows: dependencies
	$(MAKE) GO_LD_FLAGS="$(GO_LD_FLAGS)" EXE_ENDOING=$(EXE_ENDOING) _GOOS=$(_GOOS) OUTPUT=$(OUTPUT) _GOARCH=$(_GOARCH) build

build-linux: OUTPUT=$(DIST_DIR)/$(NAME)-$(_GOOS)-$(_GOARCH)-$(VERSION)
build-linux: dependencies
	$(MAKE) GO_LD_FLAGS="$(GO_LD_FLAGS)" EXE_ENDOING=$(EXE_ENDOING) _GOOS=$(_GOOS) OUTPUT=$(OUTPUT) _GOARCH=$(_GOARCH) build

build: 	
	@GOOS=$(_GOOS) GOARCH=$(_GOARCH) $(GOBUILD) -buildmode=exe -ldflags "$(GO_LD_FLAGS)" -o $(OUTPUT) $(MAIN_PKG)

# ------- run -------

run-server: dependencies
	$(GORUN) $(MAIN_PKG)

# ------- dependencies -------

dependencies: gopkg
	$(GOPKG) ensure

gopkg:
	$(GOGET) github.com/golang/dep/cmd/dep

# ------- clean -------

clean:
	@-rm -r $(DIST_DIR)