.PHONY: test update metadata/*.json

install:
	go install ./cmd/zonedb

test:
	go run cmd/zonedb/main.go
	go test ./...

zones.go: zones.txt metadata/*.json internal/* internal/*/*
	go generate

update:
	go run cmd/zonedb/main.go -update -w -c 100
	make zones.go
	make test

normalize:
	go run cmd/zonedb/main.go -w
	make zones.go

formatted_date = $(shell date -u "+%Y.%m.%d.%H")

tag-version: .git/refs/heads/master
	git tag v1.$(formatted_date) $(shell cat .git/refs/heads/master)
	git push --tags
