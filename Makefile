#!/usr/bin/make -f

BINDIR ?= $(GOPATH)/bin
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::') # grab everything after the space in "github.com/tendermint/tendermint v0.34.7"
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)


export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(NOVACHAIN_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=nova \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=novad \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(NOVACHAIN_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ifeq (,$(findstring nostrip,$(NOVACHAIN_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(NOVACHAIN_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

#$(info $$BUILD_FLAGS is [$(BUILD_FLAGS)])

PACKAGES_UNIT=$(shell go list ./...)
COVERAGE_TXT=coverage.txt
COVERAGE_HTML=coverage.html

all: build lint test

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/novad

build:
	go build $(BUILD_FLAGS) -o ./build/novad ./cmd/novad

########################################################################################################
#####                                     Test & Linting                                           #####
########################################################################################################

test:
	@go test -v ./x/...

test-cover:
	@go test -mod=readonly -timeout 30m -coverprofile=$(COVERAGE_TXT) -tags='norace' -covermode=atomic $(PACKAGES_UNIT) && go tool cover -html=$(COVERAGE_TXT) -o $(COVERAGE_HTML)

cover-report:
	@go tool cover -html=$(COVERAGE_TXT) -o $(COVERAGE_HTML)

lint:
	@echo "--> Running linter"
	@docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run --out-format=tab --timeout=10m

format:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./... --fix
	@go run mvdan.cc/gofumpt -l -w x/ app/ ante/ tests/

########################################################################################################
#####                                        Protobuf                                              #####
########################################################################################################

protoConVer=v0.3.0
protoImgName=a41ventures/nova-protogen:$(protoConVer)
protogenConName=nova-proto-gen-$(protoConVer)
protogenApiConName=nova-proto-gen-api-$(protoConVer)

protogen-all: protogen protogen-api

protogen:
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${protogenConName}$$"; \
	then docker start -a $(protogenConName); \
	else docker run --name $(protogenConName) -v $(CURDIR):/workspace --workdir /workspace $(protoImgName) \
	bash ./scripts/protogen.sh; fi

protogen-api:
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${protogenApiConName}$$"; \
	then docker start -a $(protogenApiConName); \
	else docker run --name $(protogenApiConName) -v $(CURDIR):/workspace --workdir /workspace $(protoImgName) \
	bash ./scripts/protogen-api.sh; fi

.PHONY: protogen protogen-api protogen-all all install build test-local

test-local:
	bash ./scripts/run_single_node.sh

########################################################################################################
#####                                       Docs (Swagger)                                         #####
########################################################################################################
protoVer=v0.7
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenAny=$(PROJECT_NAME)-proto-gen-any-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)

docs-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protoc-swagger-gen.sh; fi

docs-update:
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m