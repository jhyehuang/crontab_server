SHELL=/usr/bin/env bash

all: build
.PHONY: all

unexport GOFLAGS

GOCC?=go

GOVERSION:=$(shell $(GOCC) version | tr ' ' '\n' | grep go1 | sed 's/^go//' | awk -F. '{printf "%d%03d%03d", $$1, $$2, $$3}')
#GOVERSIONMIN:=$(shell cat GO_VERSION_MIN | awk -F. '{printf "%d%03d%03d", $$1, $$2, $$3}')

#ifeq ($(shell expr $(GOVERSION) \< $(GOVERSIONMIN)), 1)
#$(warning Your Golang version is go$(shell expr $(GOVERSION) / 1000000).$(shell expr $(GOVERSION) % 1000000 / 1000).$(shell expr $(GOVERSION) % 1000))
#$(error Update Golang to version to at least $(shell cat GO_VERSION_MIN))
#endif

# git modules that need to be loaded
MODULES:=

CLEAN:=
BINS:=

ldflags=-X=github.com/filecoin-project/lotus/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="$(ldflags)"

build: $(BINS)

filscan-chain-stat:
	rm -f filscan-chain-stat
	$(GOCC) build $(GOFLAGS) -o filscan-chain-stat ./cmd/filscan-chain-stat
.PHONY: filscan-chain-stat
BINS+=filscan-chain-stat

letsfil-job:
	rm -f letsfil-job
	$(GOCC) build $(GOFLAGS) -o letsfil-job main.go
.PHONY: letsfil-job
BINS+=letsfil-job