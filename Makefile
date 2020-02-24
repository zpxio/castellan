.DEFAULT_GOAL := build

# Aliases
GOCMD=go
GO_BUILD=$(GOCMD) build
GO_CLEAN=$(GOCMD) clean
GO_TEST=$(GOCMD) test
GO_GET=$(GOCMD) get

BUILD_DIR := out
BUILD_EXE := castellan
BUILD_TARGET := ${BUILD_DIR}/${BUILD_EXE}

build:
	@-mkdir -p ${BUILD_DIR}
	@-echo "BUILD: ${BUILD_TARGET}"
	$(GO_BUILD) -o $(BUILD_TARGET) -v cmd/castellan/main.go

run: build
	./${BUILD_TARGET}

test:
	$(GO_TEST) -v ./...

clean:
	$(GO_CLEAN)
	rm -f $(BUILD_TARGET)
