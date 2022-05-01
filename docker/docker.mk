#
# Copyright 2021 The Dapr Authors
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


# Docker image build and push setting
DOCKER:=docker
DOCKERFILE_DIR?=./docker

APP_SYSTEM_IMAGE_NAME=$(RELEASE_NAME)
APP_RUNTIME_IMAGE_NAME=example
APP_CMD_IMAGE_NAME=cmd-service
APP_QUERY_IMAGE_NAME=query-service

# build docker image for linux
BIN_PATH=$(OUT_DIR)/$(TARGET_OS)_$(TARGET_ARCH)

ifeq ($(TARGET_OS), windows)
  DOCKERFILE:=Dockerfile-windows
  BIN_PATH := $(BIN_PATH)/release
else ifeq ($(origin DEBUG), undefined)
  DOCKERFILE:=Dockerfile
  BIN_PATH := $(BIN_PATH)/release
else ifeq ($(DEBUG),0)
  DOCKERFILE:=Dockerfile
  BIN_PATH := $(BIN_PATH)/release
else
  DOCKERFILE:=Dockerfile-debug
  BIN_PATH := $(BIN_PATH)/debug
endif

ifeq ($(TARGET_ARCH),arm)
  DOCKER_IMAGE_PLATFORM:=$(TARGET_OS)/arm/v7
else ifeq ($(TARGET_ARCH),arm64)
  DOCKER_IMAGE_PLATFORM:=$(TARGET_OS)/arm64/v8
else
  DOCKER_IMAGE_PLATFORM:=$(TARGET_OS)/amd64
endif

# Supported docker image architecture
DOCKERMUTI_ARCH=linux-amd64 linux-arm linux-arm64 windows-amd64

################################################################################
# Target: docker-build, docker-push                                            #
################################################################################

LINUX_BINS_OUT_DIR=$(OUT_DIR)/linux_$(GOARCH)
DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_SYSTEM_IMAGE_NAME):$(APP_TAG)
APP_CMD_DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_CMD_IMAGE_NAME):$(APP_TAG)
APP_QUERY_DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_QUERY_IMAGE_NAME):$(APP_TAG)

ifeq ($(LATEST_RELEASE),true)
APP_CMD_DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_CMD_IMAGE_NAME):$(LATEST_TAG)
APP_QUERY_DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_QUERY_IMAGE_NAME):$(LATEST_TAG)
endif


# To use buildx: https://github.com/docker/buildx#docker-ce
export DOCKER_CLI_EXPERIMENTAL=enabled

# check the required environment variables
check-docker-env:
ifeq ($(APP_REGISTRY),)
	$(error APP_REGISTRY environment variable must be set)
endif
ifeq ($(APP_TAG),)
	$(error APP_TAG environment variable must be set)
endif

check-arch:
ifeq ($(TARGET_OS),)
	$(error TARGET_OS environment variable must be set)
endif
ifeq ($(TARGET_ARCH),)
	$(error TARGET_ARCH environment variable must be set)
endif


docker-build: check-docker-env check-arch
	$(info Building $(DOCKER_IMAGE_TAG) docker image ...)
ifeq ($(TARGET_ARCH),amd64)
	$(DOCKER) build PKG_FILES=cmd-service -f $(DOCKERFILE_DIR)/cmd/$(DOCKERFILE) $(BIN_PATH) -t $(APP_CMD_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)
	$(DOCKER) build PKG_FILES=query-service -f $(DOCKERFILE_DIR)/query/$(DOCKERFILE) $(BIN_PATH) -t $(APP_QUERY_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)
else
	$(DOCKER) buildx build --load --build-arg PKG_FILES=cmd-service --platform $(DOCKER_IMAGE_PLATFORM) -f $(DOCKERFILE_DIR)/cmd/$(DOCKERFILE) $(BIN_PATH) -t $(APP_CMD_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)
	$(DOCKER) buildx build --load --build-arg PKG_FILES=query-service --platform $(DOCKER_IMAGE_PLATFORM) -f $(DOCKERFILE_DIR)/query/$(DOCKERFILE) $(BIN_PATH) -t $(APP_QUERY_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)
endif


docker-push:
	docker push $(APP_CMD_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)
	docker push $(APP_QUERY_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)


docker-rmi:
	docker images|grep none|awk '{print $3}'|xargs docker rmi
	docker rmi -f cmd-service
	docker rmi -f query-service


