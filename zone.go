package zonedb

//go:generate go run build/cmd/zonedb/main.go -generate-go

type Zone struct {
	Domain      string
	Parent      *Zone
	Subdomains  []Zone
	CodePoints  []rune
	NameServers []string
	WhoisServer string
	WhoisURL    string
	InfoURL     string
}
