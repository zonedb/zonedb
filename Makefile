.PHONY: test update metadata/*.json

GO_VERSION := 1.7.4

install-go:
	curl -o /app/go.tgz https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	tar xvf /app/go.tgz -C /app
	mkdir /app/go2
	export GOROOT=/app/go GOPATH=/app/go2 PATH=/app/go/bin:$PATH
	go get -t -d -v ./...

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
