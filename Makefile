.PHONY: test update metadata/*.json

GO_VERSION := 1.7.4

install-go:
	./install-go.sh

install:
	go install ./build/cmd/zonedb

test:
	go run build/cmd/zonedb/main.go
	go test ./...

zones.go: zones.txt metadata/*.json build/*.go build/*/*/*.go
	go generate

update:
	go run build/cmd/zonedb/main.go -update -w -c 100
	make zones.go
	make test

normalize:
	go run build/cmd/zonedb/main.go -w
	make zones.go
