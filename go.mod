module github.com/zonedb/zonedb

// Minimum supported Go version
// (current version -2, except when security fixes are only backported to the current version -1)
// Note: this should use the latest dot release in CI, once GitHub has added it to "actions/go-versions"
go 1.21
toolchain go1.23.4

require (
	github.com/PuerkitoBio/goquery v1.10.1
	github.com/miekg/dns v1.1.62
	github.com/wsxiaoys/terminal v0.0.0-20160513160801-0940f3fc43a0
	golang.org/x/net v0.33.0
	golang.org/x/text v0.21.0
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
)
