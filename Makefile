# Variables
BINARY_NAME := hello
PACKAGE_PATH := cmd/main.go
BUILD_PATH := bin

.PHONY: all install check fix clean build run

all: clean install build

install:
	go mod tidy

check:
	go vet ./...

fix:
	go fmt ./...

clean:
	rm -rf $(BUILD_PATH)

build:
	go build -o $(BUILD_PATH)/$(BINARY_NAME) $(PACKAGE_PATH)

run: build
	./$(BUILD_PATH)/$(BINARY_NAME)

help:
	@echo "Available targets:"
	@echo "  all: Clean, install and builds the program"
	@echo "  install: install packages"
	@echo "  check: Report issues in the code"
	@echo "  fix: Format and fix source code"
	@echo "  clean: Cleans up build artifacts"
	@echo "  build: Builds the program"
	@echo "  run: Run the final program"
	@echo "  help: Displays this help message"
