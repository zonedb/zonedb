# Public Zone Database

[![build status](https://img.shields.io/circleci/project/zonedb/zonedb/master.svg)](https://circleci.com/gh/zonedb/zonedb)
[![godoc](http://img.shields.io/badge/docs-GoDoc-blue.svg)](https://godoc.org/github.com/zonedb/zonedb)

The Public Zone Database (zonedb) is a [free, open-source database](http://opendatacommons.org/licenses/odbl/1.0/) containing a list and associated metadata of public [DNS zones](http://en.wikipedia.org/wiki/DNS_zone) (domain name extensions). It attempts to be exhaustive, including current, retired, and withdrawn top-level domains and subdomains.

The intended use case is programs that interface with the [public domain name system](http://en.wikipedia.org/wiki/Domain_Name_System), including registry and registrar services. The data overlaps with and shares test cases with the [Public Suffix List](http://publicsuffix.org/) maintained by the [Mozilla Foundation](https://mozilla.org/). The source data was originally derived from the internal zone database in use at [Domainr](https://domainr.com/).

## Structure

The database consists of a list of zones (`zones.txt`) and associated metadata in JSON format.

### zones.txt

The `zones.txt` file is a UTF-8 encoded text file containing a list of IDN & lower-case domain names, one per line, followed by a newline (`\n`) character. The domain names in zones.txt are sorted according to the following rules:

1. Number of labels (top-level domains, followed by second- and third-level domains)
2. Parent domain (e.g. `uk`)
3. Subdomain (e.g. `co.uk`)

### Metadata

Each domain with associated metadata will have a JSON file in the `metadata` directory.

### Updates

A [CircleCI workflow](https://circleci.com/docs/2.0/workflows/) updates ZoneDB each night.

#### zones.txt

If a new SLD or third-level domain needs to be added to `zones.txt`, follow these steps locally:

1. add the new zone to the bottom of the `zones.txt` file
1. run `make normalize` to normalize the data
1. run `make update` to run ZoneDB's update process
1. create a pull request for the changes, and confirm the tests are passing

#### json files

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
$ go run cmd/zonedb/main.go -list-wildcards
```

List all zones tagged `geo`:

```shell
$ go run cmd/zonedb/main.go -tags geo
```

List a given zone's tags:

```shell
$ go run cmd/zonedb/main.go -zones capetown -list-tags
```

Add a tag to multiple zones (and write the output):

```shell
$ go run cmd/zonedb/main.go -zones capetown,durban,joburg -add-tags city -w
$ make normalize
```

Remove a tag from a zone (and write the output):

```shell
$ go run cmd/zonedb/main.go -zones la -remove-tags generic -w
```

Add a location to a zone (and write the output):

```shell
$ go run cmd/zonedb/main.go -zones alsace -add-locations fr-a -w
```

## License

This database is licensed under the [Open Database License (ODBl) version 1.0](http://opendatacommons.org/licenses/odbl/1.0/). See [LICENSE.md](https://github.com/domainr/zonedb/blob/master/LICENSE.md) for more information.

Copyright © 2008–2019 the Public Zone Database authors.
