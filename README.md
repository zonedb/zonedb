# Public Zone Database

[![build status](https://img.shields.io/github/actions/workflow/status/zonedb/zonedb/go.yaml.svg)](https://github.com/zonedb/zonedb/actions)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-blue.svg?logo=go&logoColor=white)](https://pkg.go.dev/github.com/zonedb/zonedb)


The Public Zone Database (zonedb) is a [free, open-source database](http://opendatacommons.org/licenses/odbl/1.0/) containing a list and associated metadata of public [DNS zones](http://en.wikipedia.org/wiki/DNS_zone) (domain name extensions). It attempts to be exhaustive, including current, retired, and withdrawn top-level domains and subdomains.

The intended use case is programs that interface with the [public domain name system](http://en.wikipedia.org/wiki/Domain_Name_System), including registry and registrar services. The data overlaps with and shares test cases with the [Public Suffix List](http://publicsuffix.org/) maintained by the [Mozilla Foundation](https://mozilla.org/). The source data was originally derived from the internal zone database in use at [Domainr](https://domainr.com/).

## Structure

The database consists of a list of zones (`zones.txt`) and associated metadata in JSON format.

### CLI

The `zonedb` CLI is fully documented. Run `go run cmd/zonedb/main.go -h` to see its arguments.

### zones.txt

The `zones.txt` file is a UTF-8 encoded text file containing a list of IDN & lower-case domain names, one per line, followed by a newline (`\n`) character. The domain names in zones.txt are sorted according to the following rules:

1. Number of labels (top-level domains, followed by second- and third-level domains)
2. Parent domain (e.g. `uk`)
3. Subdomain (e.g. `co.uk`)

### Metadata

Each domain with associated metadata will have a JSON file in the `metadata` directory.

### Updates

A [GitHub Actions workflow](https://github.com/zonedb/zonedb/blob/HEAD/.github/workflows/update.yaml) updates ZoneDB each night.

#### zones.txt

If a new SLD or third-level domain needs to be added to `zones.txt`, follow these steps locally:

1. Add the new zone to the bottom of the `zones.txt` file.
1. Run `make normalize` to normalize the data changes.
1. Run `go run cmd/zonedb/main.go -update -w -zones {new zone}` to update its metadata.
1. Create a pull request for the changes, and confirm the tests are passing.

#### JSON files

If an update to one of the `{zone}.json` files is needed, do this locally:

1. `make normalize` # tidies the json files, and runs `go generate`
1. commit the changes

## Implementations

### Go

`go get github.com/zonedb/zonedb`

This repository contains a reference implementation of the database written in [Go](https://golang.org). Unlike the source data, which is in Unicode, the Go implementation represents domain names in ASCII [IDNA](http://en.wikipedia.org/wiki/Internationalized_domain_name#Internationalizing_Domain_Names_in_Applications) form for interop with [existing](https://godoc.org/net) [libraries](https://godoc.org/golang.org/x/net/idna).

## Contributing

Fork this repository, make changes, and send a pull request. Before submitting a PR, run `make normalize` to normalize any changes. The data is structured to minimize diff size.

### zonedb build tool

This database is generated and validated using the `zonedb` tool in this directory. To install, make sure you have a working [Go](https://golang.org) installation (1.4+) and run this command:

```shell
go get -u github.com/zonedb/zonedb/cmd/zonedb
```

You can also run the `zonedb` tool directly: `go run cmd/zonedb/main.go`

### Example commands

List all zones that have wildcarded DNS:

```shell
zonedb -list-wildcards
```

List all zones tagged `geo`:

```shell
zonedb -tags geo
```

List a given zone's tags:

```shell
zonedb -zones capetown -list-tags
```

Add a tag to multiple zones (and write the output):

```shell
zonedb -zones capetown,durban,joburg -add-tags city -w
```

Remove a tag from a zone (and write the output):

```shell
zonedb -zones la -remove-tags generic -w
```

Add a location to a zone (and write the output):

```shell
zonedb -zones alsace -add-locations fr-a -w
```

Add `zh-Hans-HK` (Hong Kong Simplified Chinese) as a language (and write the output):

```shell
zonedb -zones 香港 -add-languages zh-Hans-HK -w
```

List all zones with only one language:

```shell
zonedb -grep '"languages"\: \[\s+\S+\s+\]'
```

List all zones  with the `geo` tag that support a Chinese language variant:

```shell
zonedb  -tags geo -grep zh-
```

List all IDN zones with the `geo` tag:

```shell
zonedb -grep '[^\x{0000}-\x{007F}]' -tags geo
```

Or, using the `-idn` tag, all IDN TLDs with the `geo` tag:

```shell
zonedb -tlds -idn -tags geo
```

## License

This database is licensed under the [Open Database License (ODBl) version 1.0](http://opendatacommons.org/licenses/odbl/1.0/). See [LICENSE.md](https://github.com/zonedb/zonedb/blob/HEAD/LICENSE.md) for more information.

Copyright © 2008–2020 the Public Zone Database authors.
