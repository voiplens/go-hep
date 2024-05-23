BUF_VERSION  := v1.31.0
GOGO_VERSION := v1.4.12
GOLANGCI_VERSION := v1.58.2

default: help
.PHONY: install-linter lint lint-imports generate-proto install-generators test help

##
## tools and help
##
install-linter: ## install go linters.
ifeq (, $(shell which golangci-lint))
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
endif
ifeq (, $(shell which impi))
	@go install github.com/pavius/impi/cmd/impi@latest
endif

lint: ## run go linter.
lint: lint-imports install-linter
	golangci-lint run

lint-imports: ## checks if the go imports are correctly grouped in each file.
	@impi --local go.voiplens.io --scheme stdThirdPartyLocal --ignore-generated=true ./...


# GENERATE

install-generators: ## install protobuf generation tools.
	go install github.com/cosmos/gogoproto/protoc-gen-gogoslick@$(GOGO_VERSION)
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

generate-proto: ## re-generate all protos.
	@cd $(CURDIR)/encoding/protobuf/proto2 && buf generate
	@cd $(CURDIR)/encoding/protobuf/proto3 && buf generate

################################################################################
# Help target
################################################################################
help: ## show this help text
	@awk -vG=$$(tput setaf 2) -vR=$$(tput sgr0) ' \
	  match($$0, "^(([^#:]*[^ :]) *:)?([^#]*)##([^#].+|)$$",a) { \
	    if (a[2] != "") { printf "    make %s%-18s%s %s\n", G, a[2], R, a[4]; next }\
	    if (a[3] == "") { print a[4]; next }\
	    printf "\n%-36s %s\n","",a[4]\
	  }' $(MAKEFILE_LIST)
	@printf "\n" # blank line at the end
