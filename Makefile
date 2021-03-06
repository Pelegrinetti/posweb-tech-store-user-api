TARGET ?= development
PORT ?= 3001
PWD = $(shell pwd)
API_NAME = user-api
CONTAINER_NAME = tech-store-$(API_NAME)

build-docker-image:
	@echo "Build docker image..."
	@docker build --tag $(CONTAINER_NAME):$(TARGET) --target $(TARGET) .
	@echo "Built!"

check-if-docker-image-exists:
ifeq ($(shell docker images -q $(CONTAINER_NAME):$(TARGET) 2> /dev/null | wc -l), 0)
	@make build-docker-image
else
	@echo "Docker image already exists!"
endif

copy-dotenv-sample:
	@echo "Creating .env file..."
	@cp .env.sample .env
	@echo "Done! .Env created!"

check-dotenv:
ifeq ($(shell ls -la | grep .env 2> /dev/null | wc -l), 1)
	@make copy-dotenv-sample
endif

install-vendor: 
	@echo "Installing modules..."
	@go mod vendor
	@echo "Done! All modules installed."

check-vendor:
ifeq ($(shell ls -la | grep vendor 2> /dev/null | wc -l), 0)
	@make install-vendor
endif

build:
	@go build -o bin/server cmd/main.go

start: check-dotenv check-if-docker-image-exists
	@echo "Running $(API_NAME) at $(PORT) port."
	@docker-compose --env-file .env up

stop:
	@docker-compose down