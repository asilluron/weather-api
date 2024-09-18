GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod
BINARY_NAME=weather-api

MAIN_PACKAGE=./cmd/server/main.go

.PHONY: all build clean test run deps update-go

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE)
	./$(BINARY_NAME)

deps:
	$(GOMOD) tidy
