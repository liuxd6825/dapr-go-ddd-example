

################################################################################
# Variables                                                                    #
################################################################################

export GO111MODULE ?= on
export GOPROXY ?= https://proxy.golang.org
export GOSUMDB ?= sum.golang.org

GIT_COMMIT  = $(shell git rev-list -1 HEAD)
GIT_VERSION = $(shell git describe --always --abbrev=7 --dirty)
# By default, disable CGO_ENABLED. See the details on https://golang.org/cmd/cgo
CGO         ?= 0
BINARIES    ?= cmd-service query-service
HA_MODE     ?= false
# Force in-memory log for placement
FORCE_INMEM ?= true

# Add latest tag if LATEST_RELEASE is true
LATEST_RELEASE ?=

CONFIG_PATH=./config

PROTOC ?=protoc
# name of protoc-gen-go when protoc-gen-go --version is run.
PROTOC_GEN_GO_NAME = "protoc-gen-go"
ifdef REL_VERSION
	APP_VERSION := $(REL_VERSION)
else
	APP_VERSION := edge
endif

LOCAL_ARCH := $(shell uname -m)
ifeq ($(LOCAL_ARCH),x86_64)
	TARGET_ARCH_LOCAL=amd64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),armv8)
	TARGET_ARCH_LOCAL=arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 4),armv)
	TARGET_ARCH_LOCAL=arm
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),arm64)
	TARGET_ARCH_LOCAL=arm64
else
	TARGET_ARCH_LOCAL=amd64
endif
export GOARCH ?= $(TARGET_ARCH_LOCAL)

ifeq ($(GOARCH),amd64)
	LATEST_TAG=latest
else
	LATEST_TAG=latest-$(GOARCH)
endif

LOCAL_OS := $(shell uname)
ifeq ($(LOCAL_OS),Linux)
   TARGET_OS_LOCAL = linux
else ifeq ($(LOCAL_OS),Darwin)
   TARGET_OS_LOCAL = darwin
else
   TARGET_OS_LOCAL ?= windows
   PROTOC_GEN_GO_NAME := "protoc-gen-go.exe"
endif
export GOOS ?= $(TARGET_OS_LOCAL)

PROTOC_GEN_GO_NAME+= "v1.26.0"

# Default docker container and e2e test targst.
TARGET_OS ?= linux
TARGET_ARCH ?= amd64
TEST_OUTPUT_FILE_PREFIX ?= ./test_report

ifeq ($(GOOS),windows)
BINARY_EXT_LOCAL:=.exe
GOLANGCI_LINT:=golangci-lint.exe
export ARCHIVE_EXT = .zip
else
BINARY_EXT_LOCAL:=
GOLANGCI_LINT:=golangci-lint
export ARCHIVE_EXT = .tar.gz
endif

export BINARY_EXT ?= $(BINARY_EXT_LOCAL)

OUT_DIR := ./dist

# Helm template and install setting
HELM:=helm
RELEASE_NAME?=go-ddd-example
APP_NAMESPACE?=dapr-system
APP_MTLS_ENABLED?=true
HELM_CHART_ROOT:=./charts
HELM_CHART_DIR:=$(HELM_CHART_ROOT)/dapr
HELM_OUT_DIR:=$(OUT_DIR)/install
HELM_MANIFEST_FILE:=$(HELM_OUT_DIR)/$(RELEASE_NAME).yaml
HELM_REGISTRY?=daprio.azurecr.io


################################################################################
# Go build details                                                             #
################################################################################
BASE_PACKAGE_NAME := github.com/dapr/dapr
LOGGER_PACKAGE_NAME := github.com/dapr/kit/logger

DEFAULT_LDFLAGS:=-X $(BASE_PACKAGE_NAME)/pkg/version.gitcommit=$(GIT_COMMIT) \
  -X $(BASE_PACKAGE_NAME)/pkg/version.gitversion=$(GIT_VERSION) \
  -X $(BASE_PACKAGE_NAME)/pkg/version.version=$(APP_VERSION) \
  -X $(LOGGER_PACKAGE_NAME).DaprVersion=$(APP_VERSION)

ifeq ($(origin DEBUG), undefined)
  BUILDTYPE_DIR:=release
  LDFLAGS:="$(DEFAULT_LDFLAGS) -s -w"
else ifeq ($(DEBUG),0)
  BUILDTYPE_DIR:=release
  LDFLAGS:="$(DEFAULT_LDFLAGS) -s -w"
else
  BUILDTYPE_DIR:=debug
  GCFLAGS:=-gcflags="all=-N -l"
  LDFLAGS:="$(DEFAULT_LDFLAGS)"
  $(info Build with debugger information)
endif

APP_OUT_DIR := $(OUT_DIR)/$(GOOS)_$(GOARCH)/$(BUILDTYPE_DIR)
APP_LINUX_OUT_DIR := $(OUT_DIR)/linux_$(GOARCH)/$(BUILDTYPE_DIR)


################################################################################
# Target: build                                                                #
################################################################################
.PHONY: build
APP_BINS:=$(foreach ITEM,$(BINARIES),$(APP_OUT_DIR)/$(ITEM)$(BINARY_EXT))
build:
	$(APP_BINS)
	@cp -rf $(CONFIG_PATH) $(APP_OUT_DIR)
# Generate builds for dapr binaries for the target
# Params:
# $(1): the binary name for the target
# $(2): the binary main directory
# $(3): the target os
# $(4): the target arch
# $(5): the output directory
define genBinariesForTarget
.PHONY: $(5)/$(1)
$(5)/$(1):
	CGO_ENABLED=$(CGO) GOOS=$(3) GOARCH=$(4) go build $(GCFLAGS) \
	-o $(5)/$(1) $(2)/;
endef

# Generate binary targets
$(foreach ITEM,$(BINARIES),$(eval $(call genBinariesForTarget,$(ITEM)$(BINARY_EXT),./cmd/$(ITEM),$(GOOS),$(GOARCH),$(APP_OUT_DIR))))

################################################################################
# Target: build-linux                                                          #
################################################################################
BUILD_LINUX_BINS:=$(foreach ITEM,$(BINARIES),$(APP_LINUX_OUT_DIR)/$(ITEM))
build-linux: $(BUILD_LINUX_BINS)
	@cp -rf ./config $(APP_LINUX_OUT_DIR)
# Generate linux binaries targets to build linux docker image
ifneq ($(GOOS), linux)
$(foreach ITEM,$(BINARIES),$(eval $(call genBinariesForTarget,$(ITEM),./cmd/$(ITEM),linux,$(GOARCH),$(APP_LINUX_OUT_DIR))))
endif

################################################################################
# Target: archive                                                              #
################################################################################
ARCHIVE_OUT_DIR ?= $(APP_OUT_DIR)
ARCHIVE_FILE_EXTS:=$(foreach ITEM,$(BINARIES),archive-$(ITEM)$(ARCHIVE_EXT))

archive: $(ARCHIVE_FILE_EXTS)

# Generate archive files for each binary
# $(1): the binary name to be archived
# $(2): the archived file output directory
define genArchiveBinary
ifeq ($(GOOS),windows)
archive-$(1).zip:
	7z.exe a -tzip "$(2)\\$(1)_$(GOOS)_$(GOARCH)$(ARCHIVE_EXT)" "$(APP_OUT_DIR)\\$(1)$(BINARY_EXT)"
else
archive-$(1).tar.gz:
	tar czf "$(2)/$(1)_$(GOOS)_$(GOARCH)$(ARCHIVE_EXT)" -C "$(APP_OUT_DIR)" "$(1)$(BINARY_EXT)"
endif
endef

# Generate archive-*.[zip|tar.gz] targets
$(foreach ITEM,$(BINARIES),$(eval $(call genArchiveBinary,$(ITEM),$(ARCHIVE_OUT_DIR))))

uninstall-k8s:
	kubectl delete -f ./k8s/cmd-service.yaml
	kubectl delete -f ./k8s/query-service.yaml

install-k8s:
	kubectl apply -f ./k8s/cmd-service.yaml
	kubectl apply -f ./k8s/query-service.yaml





################################################################################
# Target: docker                                                               #
################################################################################
include docker/docker.mk
