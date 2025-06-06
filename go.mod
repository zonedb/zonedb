module github.com/zonedb/zonedb

// Minimum supported Go version
// (current version -2, except when security fixes are only backported to the current version -1)
// Note: this should use the latest dot release in CI, once GitHub has added it to "actions/go-versions"
go 1.23.8

// Internal packages rely on newer Go versions. GitHub Actions is set up to test versions between the go and toolchain declarations.
toolchain go1.24.3

require (
	github.com/PuerkitoBio/goquery v1.10.3
	github.com/miekg/dns v1.1.66
	github.com/wsxiaoys/terminal v0.0.0-20160513160801-0940f3fc43a0
	golang.org/x/net v0.40.0
	golang.org/x/text v0.25.0
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
)
