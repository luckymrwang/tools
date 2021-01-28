# Image URL to use all building/pushing image targets;
# Use your own docker registry and image name for dev/test by overridding the
# IMAGE_REPO, IMAGE_NAME and IMAGE_TAG environment variable.
IMAGE_REPO ?= node:14.15-slim
IMAGE_NAME ?= node
IMAGE_TAG ?= v1

all: fmt lint test build image

############################################################
# format section
############################################################

fmt:
	@echo "Run go fmt..."

############################################################
# lint section
############################################################

lint:
	@echo "Runing the golangci-lint..."

############################################################
# test section
############################################################

test:
	@echo "Running the tests for $(IMAGE_NAME)..."

############################################################
# build section
############################################################

build:
	@echo "Building the $(IMAGE_NAME) binary..."
	@CGO_ENABLED=0 go build

############################################################
# image section
############################################################

image: build-node-image build-skyapm-image

build-node-image: build
	@echo "Building the node docker image: $(IMAGE_NAME):$(IMAGE_TAG)..."
	@docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f build/node/Dockerfile .

build-skyapm-image: node
	@echo "Building the skyapm docker image: $(IMAGE_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)..."
	@docker build -t skyapm:v2.0.1 -f build/node/skyapm/Dockerfile .

node:
	@echo "npm install skyapm..."
	@npm config set registry https://registry.npm.taobao.org/
	@npm install skyapm-nodejs --target=14.15.4 --target_platform=linux --target_arch=x64 --target_libc=glibc --save
	@mv -f node_modules build/node/skyapm/

test-deploy:
	@echo "creating deployment banana..."
	@kubectl apply -f build/node/banana/banana.yaml

############################################################
# clean section
############################################################
clean:
	@rm -rf build/node/skyapm/node_modules
	@rm -rf build/node/skyapm/package-lock.json

.PHONY: all fmt lint check test build image clean

