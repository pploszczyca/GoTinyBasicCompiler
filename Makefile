# Variables
BINARY_NAME=GoTinyBasicCompiler
SRC_DIR=.
BUILD_DIR=build
MAIN_PACKAGE=$(SRC_DIR)/main.go

# Commands
GO_BUILD=go build
GO_CLEAN=go clean
GO_TEST=go test
GO_FMT=go fmt
GO_VET=go vet

# Targets
.PHONY: all build clean test format vet run

all: build

build:
	mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

clean:
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR)

test:
	$(GO_TEST) ./...

format:
	$(GO_FMT) ./...

vet:
	$(GO_VET) ./...

run: build
	$(BUILD_DIR)/$(BINARY_NAME) input.txt output.txt

