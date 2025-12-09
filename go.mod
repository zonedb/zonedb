module github.com/zonedb/zonedb

// Minimum supported Go version
// (current version -2, except when security fixes are only backported to the current version -1)
// Note: this should use the latest dot release in CI, once GitHub has added it to "actions/go-versions"
//
// EDIT[2025.12.09]: golang.org/x/net requires go1.24.0 - so we have to avoid go1.23.x
go 1.24.10

require (
	github.com/PuerkitoBio/goquery v1.11.0
	github.com/miekg/dns v1.1.68
	github.com/wsxiaoys/terminal v0.0.0-20160513160801-0940f3fc43a0
	golang.org/x/net v0.47.0
	golang.org/x/text v0.31.0
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
)
