# Variables
BINARY_NAME=GoTinyBasicCompiler
SRC_DIR=.
BUILD_DIR=build
MAIN_PACKAGE=$(SRC_DIR)/main.go
SAMPLES_DIR=./samples
RESULTS_DIR=./results

# Commands
GO_BUILD=go build
GO_CLEAN=go clean
GO_TEST=go test
GO_FMT=go fmt
GO_VET=go vet
CHMOD=chmod +x
GCC=gcc

# Sample files
SAMPLES=sampleTinyBasic sampleTinyBasic2

# Targets
.PHONY: all build clean test format vet run buildSamples buildAndRunSamples

all: build

build:
	mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	$(CHMOD) $(BUILD_DIR)/$(BINARY_NAME)

clean:
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR) $(RESULTS_DIR)

test:
	$(GO_TEST) ./...

format:
	$(GO_FMT) ./...

vet:
	$(GO_VET) ./...

run: build
	$(BUILD_DIR)/$(BINARY_NAME) ./input.txt ./output.txt

buildSamples: build
	@for sample in $(SAMPLES); do \
		$(BUILD_DIR)/$(BINARY_NAME) $(SAMPLES_DIR)/$$sample.bas $(RESULTS_DIR)/$$sample.c; \
		$(GCC) -o $(RESULTS_DIR)/$$sample $(RESULTS_DIR)/$$sample.c; \
	done

buildAndRunSamples: buildSamples
	@for sample in $(SAMPLES); do \
		./$(RESULTS_DIR)/$$sample; \
	done
