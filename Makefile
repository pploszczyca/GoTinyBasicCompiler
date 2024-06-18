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
CHMOD=chmod +x

# Targets
.PHONY: all build clean test format vet run runSample runSample2 runSampleAndRunCode runSampleAndRunCode2

all: build

build:
	mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	$(CHMOD) $(BUILD_DIR)/$(BINARY_NAME)

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
	$(BUILD_DIR)/$(BINARY_NAME) ./input.txt ./output.txt

runSample: build
	$(BUILD_DIR)/$(BINARY_NAME) ./samples/sampleTinyBasic.bas ./results/sampleTinyBasic.c

runSample2: build
	$(BUILD_DIR)/$(BINARY_NAME) ./samples/sampleTinyBasic2.bas ./results/sampleTinyBasic2.c

runSampleAndRunCode: build
	 $(BUILD_DIR)/$(BINARY_NAME) ./samples/sampleTinyBasic.bas ./results/sampleTinyBasic.c
	 gcc -o ./results/sampleTinyBasic ./results/sampleTinyBasic.c
	 ./results/sampleTinyBasic

 runSampleAndRunCode2: build
	 $(BUILD_DIR)/$(BINARY_NAME) ./samples/sampleTinyBasic2.bas ./results/sampleTinyBasic2.c
	 gcc -o ./results/sampleTinyBasic2 ./results/sampleTinyBasic2.c
	 ./results/sampleTinyBasic2