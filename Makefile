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
SAMPLES=sampleTinyBasic sampleTinyBasic2 sampleTinyBasic3 sampleTinyBasic4 sampleTinyBasic5

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
	$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

buildSamples: build
	@for sample in $(SAMPLES); do \
  		echo "BUILDING SAMPLE: $$sample"; \
		$(BUILD_DIR)/$(BINARY_NAME) $(SAMPLES_DIR)/$$sample.bas $(RESULTS_DIR)/$$sample.c && \
		$(GCC) -o $(RESULTS_DIR)/$$sample $(RESULTS_DIR)/$$sample.c && \
		echo "DONE BUILDING SAMPLE: $$sample\n\n"; \
	done

buildAndRunSamples: buildSamples
	@for sample in $(SAMPLES); do \
		./$(RESULTS_DIR)/$$sample; \
	done
