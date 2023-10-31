# Makefile for building a Go application without modules and with multi-arch support

# Name of the generated binary
BINARY_NAME=tofuinit

# Go commands
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_FOLDER := bin

# Architectures and OS to build for
ARCHS=amd64
OSES=linux windows darwin

# By default, just run "all" which will test and then build
all: format test build

# # Multi-arch build
# multiarch:
# 	for os in $(OSES); do \
# 		for arch in $(ARCHS); do \
# 			GO111MODULE=off GOOS=$$os GOARCH=$$arch $(GOBUILD) -o $(BINARY_NAME)-$$os-$$arch; \
# 		done \
# 	done

multiarch:
	mkdir -p $(BINARY_FOLDER)
	for os in $(OSES); do \
		for arch in $(ARCHS); do \
			GO111MODULE=on GOOS=$$os GOARCH=$$arch $(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME)-$$os-$$arch; \
		done \
	done

# Format the code using gofmt
format:
	find . -name "*.go" -type f | xargs gofmt -s -w


# Build the binary
build:
	GO111MODULE=on $(GOBUILD) -o $(BINARY_NAME) -v .


# Test the application
test: format
	GO111MODULE=on $(GOTEST) -v ./cmd 

# Clean up the binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)*
	rm -rf $(BINARY_FOLDER)
# Run the application
run:
	GO111MODULE=on $(GOBUILD) -o $(BINARY_NAME) -v .
	./$(BINARY_NAME)

# If you want a help message
help:
	@echo "Use: make [target] where target is one of"
	@echo "    build     to build the application for the local architecture"
	@echo "    multiarch to build the application for multiple architectures"
	@echo "    test      to run tests"
	@echo "    clean     to remove the binary files and any temporary files"
	@echo "    run       to build and run the application for the local architecture"

.PHONY: all multiarch build test clean run help
