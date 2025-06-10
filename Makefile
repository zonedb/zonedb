# Clear many built-in suffix rules to speed things up
.SUFFIXES:

%.json ::
	@# drop all implicit rules for .json files.

export PATH := $(PATH):$(shell go env GOROOT)/misc/wasm
# wasm tools moved in Go 1.24:
export PATH := $(PATH):$(shell go env GOROOT)/lib/wasm

GO_TEST_ARGS ?= -v ./...

.PHONY: install
install:
	go install ./cmd/zonedb

.PHONY: test
test:
	go run ./cmd/zonedb
	go test $(GO_TEST_ARGS)

.PHONY: test-wasm
test-wasm:
	GOOS=wasip1 GOARCH=wasm go test $(GO_TEST_ARGS)

.PHONY: test-tinygo
test-tinygo:
	tinygo test $(GO_TEST_ARGS)

.PHONY: test-tinygo-wasm
test-tinygo-wasm:
	tinygo test -target wasip1 $(GO_TEST_ARGS)

zones.go: zones.txt metadata/*.json internal/* internal/*/*
	go generate -x

.PHONY: update
update: update-zones zones.go

.PHONY: update-zones
update-zones:
	go run ./cmd/zonedb -update -w -c 100 $(ZONEDB_ARGS)

.PHONY: normalize
normalize:
	go run ./cmd/zonedb -w
	$(MAKE) zones.go

git_revision=$(shell git describe --no-tags --always --dirty --abbrev=0)
number_of_commits=$(shell git rev-list HEAD --count)
major_version=$(shell cat VERSION)
tag_version=v$(major_version).$(number_of_commits)

.PHONY: tag-version
tag-version: .git/refs/heads/main
	git tag $(tag_version) $(git_revision)
	git push --tags

.PHONY: test-makefile-integrity
test-makefile-integrity:
	@echo ">>> Testing 'make install'..."
	$(MAKE) install
	@echo ">>> Testing 'make test'…"
	$(MAKE) test
	@echo ">>> Testing 'make test-wasm'…"
	$(MAKE) test-wasm
	@echo ">>> Testing 'make zones.go'…"
	$(MAKE) zones.go
	@echo ">>> Testing 'make normalize'…"
	# This target modifies files, which is acceptable in an isolated CI test job.
	$(MAKE) normalize
	@echo ">>> Makefile integrity tests completed successfully."
