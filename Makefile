PROJECT_NAME := $(shell basename `pwd`)
PKG := "sbp.gitlab.schubergphilis.com/shoekstra/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all test-dep lint test build-dep build-snapshot build-tag help

all: lint build-snapshot

test-dep: ## Get testing dependencies
	@go get -u github.com/golang/lint/golint

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST} -cover

build-dep: ## Get build dependencies
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

build-snapshot: ## Build from current master
	@goreleaser --snapshot --rm-dist

build-tag: ## Build from latest tag
	@goreleaser

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
