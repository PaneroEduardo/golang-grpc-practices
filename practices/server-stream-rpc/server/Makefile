

define help
Usage: make <command>
Commands:
  help:	Show this help information
  tidy:	Execute: go mod tidy
  build: Build the project
  start: Start the service
endef
export help

.PHONY: help
help:
	@echo "$$help"

.PHONY: help
tidy:
	$(info) "Executing go mod tidy in the project"
	go mod tidy

.PHONY: build
build:
	$(info) "Building the project"
	mkdir -p build/bin
	go build -v -ldflags='-w -s' -o build/bin $(get_packages)

.PHONY: start
start: build
	$(info) "Starting the project"
	./build/bin/server


info := @printf '\033[32m01%s\033[0m\n'
get_packages := $$(go list ./... | grep -v test/acceptance)