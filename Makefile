.DEFAULT_GOAL := default

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: build
build: 
	go build -o build/dcim

.PHONY: default
default:
	@echo "No Target Selected"; exit 1
