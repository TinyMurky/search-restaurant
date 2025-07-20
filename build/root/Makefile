.PHONY: all search test

all: search

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
project_root := $(abspath $(dir $(mkfile_path))/../..)
CONFIG ?= configs/config.yaml
ENV ?= configs/.env
GO_RUN = go run .

search:
	$(GO_RUN) -c $(project_root)/$(CONFIG) -e $(project_root)/$(ENV) search

test:
	go test ./... -v
