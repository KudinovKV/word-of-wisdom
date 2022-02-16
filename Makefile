VERSION?=0.1.0
PROJECT_NAME := $(shell basename "$(PWD)")
GO=go
SUDO=sudo
DOCKER=docker
BUILD_PATH=build
DEPLOYMENT_PATH=deployments
SERVER_NAME=server
CLIENT_NAME=client
POW_SERVER_PORT?=9999

.PHONY: build-bin-server
build-bin-server:
	@echo "> Build server..."
	@$(GO) build -o $(BUILD_PATH)/$(SERVER_NAME) cmd/$(SERVER_NAME)/main.go

.PHONY: build-bin-client
build-bin-client:
	@echo "> Build client..."
	@$(GO) build -o $(BUILD_PATH)/$(CLIENT_NAME) cmd/$(CLIENT_NAME)/main.go

.PHONY: build-bin
build-bin: build-bin-server build-bin-client

.PHONY: clean
clean:
	@echo "> Clean build..."
	@rm -rf $(BUILD_PATH)

.PHONY: build-image-server
build-image-server: build-bin-server
	@$(SUDO) $(DOCKER) build \
		--tag $(PROJECT_NAME)-$(SERVER_NAME):$(VERSION) \
		-f $(DEPLOYMENT_PATH)/$(SERVER_NAME)/Dockerfile \
		.

.PHONY: build-image-client
build-image-client: build-bin-client
	@$(SUDO) $(DOCKER) build \
		--tag $(PROJECT_NAME)-$(CLIENT_NAME):$(VERSION) \
		-f $(DEPLOYMENT_PATH)/$(CLIENT_NAME)/Dockerfile \
		.

.PHONY: build-images
build-images: build-image-server build-image-client

.PHONY: delete-image-server
delete-image-server:
	@$(SUDO) $(DOCKER) rmi --force $(PROJECT_NAME)-$(SERVER_NAME):$(VERSION)

.PHONY: delete-image-client
delete-image-client:
	@$(SUDO) $(DOCKER) rmi --force $(PROJECT_NAME)-$(CLIENT_NAME):$(VERSION)

.PHONY: delete-images
delete-images: delete-image-server delete-image-client

.PHONY: run-server
run-server: build-image-server
	$(SUDO) $(DOCKER) run \
		--name $(PROJECT_NAME)-$(SERVER_NAME) \
		--publish $(POW_SERVER_PORT):$(POW_SERVER_PORT) \
		--env POW_SERVER_PORT=$(POW_SERVER_PORT) \
		$(PROJECT_NAME)-$(SERVER_NAME):$(VERSION)

.PHONY: run-client
run-client: build-image-client
	$(SUDO) $(DOCKER) run \
		--name $(PROJECT_NAME)-$(CLIENT_NAME) \
		--publish $(POW_SERVER_PORT):$(POW_SERVER_PORT) \
		--env POW_SERVER_PORT=$(POW_SERVER_PORT) \
		$(PROJECT_NAME)-$(CLIENT_NAME):$(VERSION)