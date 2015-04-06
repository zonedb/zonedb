test:
	go test ./...

zones.go: zones.txt metadata/*.json
	go generate

update:
	go run build/cmd/zonedb/main.go -update -w
	make zones.go
	make test
