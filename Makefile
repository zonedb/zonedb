.PHONY: test update metadata/*.json

install:
	go install ./cmd/zonedb

test:
	go run ./cmd/zonedb
	go test ./...

zones.go: zones.txt metadata/*.json internal/* internal/*/*
	go generate

update:
	go run ./cmd/zonedb -update -w -c 10 $(ZONEDB_ARGS)
	$(MAKE) zones.go

normalize:
	go run ./cmd/zonedb -w
	$(MAKE) zones.go

git_revision=$(shell git describe --no-tags --always --dirty --abbrev=0)
number_of_commits=$(shell git rev-list HEAD --count)
major_version=$(shell cat VERSION)
tag_version=v$(major_version).$(number_of_commits)

tag-version: .git/refs/heads/main
	git tag $(tag_version) $(git_revision)
	git push --tags
