# Makefile

# Variables
IMAGE_NAME = hub.docker.smartforge.eu/radar-app
CONTAINER_NAME = radar

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run -d --name $(CONTAINER_NAME) -v $(pwd):/app $(IMAGE_NAME) /bin/bash


# Stop the Docker container
stop:
	docker stop $(CONTAINER_NAME)

# Remove the Docker container
remove:
	docker rm $(CONTAINER_NAME)

# Remove the Docker image
clean:
	docker rmi $(IMAGE_NAME)

.PHONY: build run stop remove clean
