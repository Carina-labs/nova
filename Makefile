.PHONY: geth android ios evm all test clean

GOBIN = ./build/bin
GO ?= latest
GOBUILD = env GO111MODULE=on go build
GORUN = env GO111MODULE=on go run

all:
	$(GOBUILD) -v -o $(GOBIN)/ ./...
	@echo "Done building."
	@echo "Run \"$(GOBIN)/novachaind\" to launch supernova blockchain."