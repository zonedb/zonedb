.PHONY: test update metadata/*.json

install:
	go install ./build/cmd/zonedb

test:
	go test ./...

zones.go: zones.txt metadata/*.json build/*.go build/*/*/*.go
	go generate

update:
	go run build/cmd/zonedb/main.go -update -w
	make zones.go
	make test
