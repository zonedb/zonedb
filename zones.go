// Automatically generated

package zonedb

import "os"

func init() {
	if os.Getenv("ZONEDB_SKIP_INIT") != "" {
		return
	}
	initZones() // Separate function to report allocs in initialization
}

// Type w is an alias for []string to generate smaller source code
type w []string

// Constants to generate smaller source code
const (
	t = true
	f = false
	e = ""
)

// Nil variables to generate smaller source code
var (
	n []string // == nil
	x []Zone   // nil (no subdomains)
	r *Zone    // nil (root zone for TLDs)
)

// Tags are stored in a single integer as a bit field.
type Tags uint32

// Tag values corresponding to bit shifts (1 << iota)
const (
	TagAdult          = 1
	TagBrand          = 2
	TagCity           = 4
	TagClosed         = 8
	TagCommunity      = 16
	TagCountry        = 32
	TagGeneric        = 64
	TagGeo            = 128
	TagInfrastructure = 256
	TagPrivate        = 512
	TagRegion         = 1024
	TagReserved       = 2048
	TagRetired        = 4096
	TagSpecial        = 8192
	TagSponsored      = 16384
	TagTest           = 32768
	TagWithdrawn      = 65536
	numTags           = 17
)

// TagStrings maps integer tag values to strings.
var TagStrings = map[Tags]string{
	TagAdult:          "adult",
	TagBrand:          "brand",
	TagCity:           "city",
	TagClosed:         "closed",
	TagCommunity:      "community",
	TagCountry:        "country",
	TagGeneric:        "generic",
	TagGeo:            "geo",
	TagInfrastructure: "infrastructure",
	TagPrivate:        "private",
	TagRegion:         "region",
	TagReserved:       "reserved",
	TagRetired:        "retired",
	TagSpecial:        "special",
	TagSponsored:      "sponsored",
	TagTest:           "test",
	TagWithdrawn:      "withdrawn",
}

// TagValues maps tag names to integer values.
var TagValues = map[string]Tags{
	"adult":          TagAdult,
	"brand":          TagBrand,
	"city":           TagCity,
	"closed":         TagClosed,
	"community":      TagCommunity,
	"country":        TagCountry,
	"generic":        TagGeneric,
	"geo":            TagGeo,
	"infrastructure": TagInfrastructure,
	"private":        TagPrivate,
	"region":         TagRegion,
	"reserved":       TagReserved,
	"retired":        TagRetired,
	"special":        TagSpecial,
	"sponsored":      TagSponsored,
	"test":           TagTest,
	"withdrawn":      TagWithdrawn,
}

// z is a static slice of Zones.
// Other global variables have pointers into this slice.
var z = make([]Zone, 5230)

// Zones is a slice of all Zones in the database.
var Zones = z[:]

// TLDs is a slice of all top-level domain Zones.
var TLDs = z[:1757]

// ZoneMap maps domain names to Zones.
var ZoneMap map[string]*Zone

//go:noinline
func i0() {
	z[0] = Zone{"aaa", r, x, 0x42, "American Automobile Association, Inc.", "http://nic.aaa/", w{"a.nic.aaa", "b.nic.aaa", "c.nic.aaa", "ns1.dns.nic.aaa", "ns2.dns.nic.aaa", "ns3.dns.nic.aaa"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.aaa", e, w{"https://rdap.nic.aaa/"}, t}
	z[1] = Zone{"aarp", r, x, 0x42, "AARP", "https://www.aarp.org/everywhere/aarp-nic/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.aarp", e, w{"https://tld-rdap.verisign.com/aarp/v1/"}, t}
	z[2] = Zone{"abarth", r, x, 0x1042, "Fiat Chrysler Automobiles N.V.", "http://nic.abarth/", n, n, n, n, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/abarth/"}, f}
	z[3] = Zone{"abb", r, x, 0x42, "ABB Ltd", "https://www.icann.org/en/registry-agreements/details/abb", w{"a0.nic.abb", "a2.nic.abb", "b0.nic.abb", "c0.nic.abb"}, n, n, n, "whois.nic.abb", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[4] = Zone{"abbott", r, x, 0x42, "Abbott Laboratories, Inc.", "https://nic.abbott/", w{"a0.nic.abbott", "a2.nic.abbott", "b0.nic.abbott", "c0.nic.abbott"}, n, n, n, "whois.nic.abbott", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[5] = Zone{"abbvie", r, x, 0x42, "AbbVie Inc.", "http://nic.abbvie/", w{"dns1.nic.abbvie", "dns2.nic.abbvie", "dns3.nic.abbvie", "dns4.nic.abbvie", "dnsa.nic.abbvie", "dnsb.nic.abbvie", "dnsc.nic.abbvie", "dnsd.nic.abbvie"}, n, n, n, "whois.nic.abbvie", e, w{"https://rdap.nominet.uk/abbvie/"}, f}
	z[6] = Zone{"abc", r, x, 0x42, "Disney Enterprises, Inc.", "https://www.nic.abc/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, "whois.nic.abc", e, w{"https://tld-rdap.verisign.com/abc/v1/"}, f}
	z[7] = Zone{"able", r, x, 0x42, "Able Inc.", "http://nic.able/", w{"a.nic.able", "b.nic.able", "c.nic.able", "ns1.dns.nic.able", "ns2.dns.nic.able", "ns3.dns.nic.able"}, n, n, w{"ja"}, "whois.nic.able", e, w{"https://rdap.nic.able/"}, t}
}

//go:noinline
func i7() {
	z[8] = Zone{"abogado", r, x, 0x40, "Registry Services, LLC", "http://nic.abogado/", w{"a.nic.abogado", "b.nic.abogado", "c.nic.abogado", "x.nic.abogado", "y.nic.abogado", "z.nic.abogado"}, n, n, w{"es"}, "whois.nic.abogado", e, w{"https://rdap.nic.abogado/"}, t}
	z[9] = Zone{"abudhabi", r, x, 0xc4, "Abu Dhabi Systems and Information Centre", "https://www.tamm.abudhabi/nicabudhabi/index-en.html", w{"gtld.alpha.aridns.net.au", "gtld.beta.aridns.net.au", "gtld.delta.aridns.net.au", "gtld.gamma.aridns.net.au"}, n, w{"Abu Dhabi", "AE-AZ"}, w{"ar"}, "whois.nic.abudhabi", e, w{"https://rdap.nic.abudhabi/"}, t}
	z[10] = Zone{"ac", r, z[1757:1762], 0xa0, e, "https://icb.co.uk/", w{"a0.nic.ac", "a2.nic.ac", "b0.nic.ac", "c0.nic.ac"}, n, n, n, "whois.nic.ac", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[11] = Zone{"academy", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.academy", "v0n1.nic.academy", "v0n2.nic.academy", "v0n3.nic.academy", "v2n0.nic.academy", "v2n1.nic.academy"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.academy", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[12] = Zone{"accenture", r, x, 0x42, "Accenture plc", "https://www.accenture.com/us-en/acn-welcome-registry", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://tld-rdap.verisign.com/accenture/v1/"}, t}
	z[13] = Zone{"accountant", r, x, 0x40, "dot Accountant Limited", "http://nic.accountant/", w{"a.nic.accountant", "b.nic.accountant", "c.nic.accountant", "ns1.dns.nic.accountant", "ns2.dns.nic.accountant", "ns3.dns.nic.accountant"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.accountant", e, w{"https://rdap.nic.accountant/"}, t}
	z[14] = Zone{"accountants", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.accountants", "v0n1.nic.accountants", "v0n2.nic.accountants", "v0n3.nic.accountants", "v2n0.nic.accountants", "v2n1.nic.accountants"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.accountants", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[15] = Zone{"acer", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i15() {
	z[16] = Zone{"aco", r, x, 0x42, "ACO Severin Ahlmann GmbH & Co. KG", "https://nic.aco.com/", w{"a.dns.nic.aco", "m.dns.nic.aco", "n.dns.nic.aco"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.aco", e, w{"https://rdap.nic.aco/"}, t}
	z[17] = Zone{"active", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[18] = Zone{"actor", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.actor", "v0n1.nic.actor", "v0n2.nic.actor", "v0n3.nic.actor", "v2n0.nic.actor", "v2n1.nic.actor"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.actor", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[19] = Zone{"ad", r, z[1762:1763], 0xe0, e, "https://www.iana.org/domains/root/db/ad.html", w{"ad.cctld.authdns.ripe.net", "ad.ns.nic.es", "anycast23.irondns.net", "anycast9.irondns.net", "dnsc.ad", "dnsm.ad", "ns3.nic.fr"}, n, n, n, e, e, n, t}
	z[20] = Zone{"adac", r, x, 0x1042, "Allgemeiner Deutscher Automobil-Club e.V. (ADAC)", "https://www.adac.de/nic-adac/", n, n, n, w{"mul-Cyrl", "mul-Grek", "mul-Latn"}, "whois.nic.adac", e, w{"https://rdap.centralnic.com/adac/"}, t}
	z[21] = Zone{"ads", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[22] = Zone{"adult", r, x, 0x41, "ICM Registry AD LLC", "https://nic.adult/", w{"a.nic.adult", "b.nic.adult", "c.nic.adult", "x.nic.adult", "y.nic.adult", "z.nic.adult"}, n, n, w{"ar", "be", "bg", "bs", "cnr", "da", "de", "es", "fr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.adult", e, w{"https://rdap.nic.adult/"}, t}
	z[23] = Zone{"ae", r, z[1763:1772], 0xa0, e, "https://www.nic.ae/", w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, "whois.aeda.net.ae", e, n, t}
}

//go:noinline
func i23() {
	z[24] = Zone{"aeg", r, x, 0x42, "Aktiebolaget Electrolux", "https://nic.aeg/", w{"v0n0.nic.aeg", "v0n1.nic.aeg", "v0n2.nic.aeg", "v0n3.nic.aeg", "v2n0.nic.aeg", "v2n1.nic.aeg"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.aeg", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[25] = Zone{"aero", r, z[1772:1775], 0x4040, "SITA Information Networking Computing USA", "https://information.aero/", w{"a0.nic.aero", "a2.nic.aero", "b0.nic.aero", "b2.nic.aero", "c0.nic.aero"}, n, n, n, "whois.aero", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[26] = Zone{"aetna", r, x, 0x42, "Aetna Life Insurance Company", "http://nic.aetna/", w{"a.nic.aetna", "b.nic.aetna", "c.nic.aetna", "ns1.dns.nic.aetna", "ns2.dns.nic.aetna", "ns3.dns.nic.aetna"}, n, n, w{"es", "pt"}, "whois.nic.aetna", e, w{"https://rdap.nic.aetna/"}, t}
	z[27] = Zone{"af", r, z[1775:1785], 0xa0, e, "https://nic.af/", w{"ns.anycast.nic.af", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.af", e, n, f}
	z[28] = Zone{"afamilycompany", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.afamilycompany", e, w{"https://tld-rdap.verisign.com/afamilycompany/v1/"}, f}
	z[29] = Zone{"afl", r, x, 0x42, "Australian Football League", "http://nic.afl/", w{"a.nic.afl", "b.nic.afl", "c.nic.afl", "x.nic.afl", "y.nic.afl", "z.nic.afl"}, n, n, n, "whois.nic.afl", e, w{"https://rdap.nic.afl/"}, f}
	z[30] = Zone{"africa", r, x, 0xc0, "ZA Central Registry NPC trading as Registry.Africa", "https://registry.africa/", w{"coza1.dnsnode.net", "nsp.netnod.se"}, n, w{"Africa"}, n, "whois.nic.africa", e, w{"https://rdap.nic.africa/rdap/"}, f}
	z[31] = Zone{"africamagic", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i31() {
	z[32] = Zone{"ag", r, z[1785:1791], 0xa0, e, "https://www.nic.ag/", w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, "whois.nic.ag", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[33] = Zone{"agakhan", r, x, 0x42, "Fondation Aga Khan (Aga Khan Foundation)", "https://www.icann.org/en/registry-agreements/details/agakhan", w{"a0.nic.agakhan", "a2.nic.agakhan", "b0.nic.agakhan", "c0.nic.agakhan"}, n, n, n, "whois.nic.agakhan", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[34] = Zone{"agency", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.agency", "v0n1.nic.agency", "v0n2.nic.agency", "v0n3.nic.agency", "v2n0.nic.agency", "v2n1.nic.agency"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.agency", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[35] = Zone{"ai", r, z[1791:1795], 0xa0, e, "http://nic.com.ai/", w{"a.lactld.org", "anycastdns1-cz.nic.ai", "anycastdns2-cz.nic.ai", "pch.whois.ai"}, n, n, n, "whois.nic.ai", e, n, f}
	z[36] = Zone{"aig", r, x, 0x42, "American International Group, Inc.", "https://www.nic.aig/home", w{"a.nic.aig", "b.nic.aig", "c.nic.aig", "ns1.dns.nic.aig", "ns2.dns.nic.aig", "ns3.dns.nic.aig"}, n, n, w{"es"}, "whois.nic.aig", e, w{"https://rdap.nic.aig/"}, t}
	z[37] = Zone{"aigo", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[38] = Zone{"airbus", r, x, 0x42, "Airbus S.A.S.", "https://nic.airbus/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"mul-Arab", "mul-Hani", "mul-Hano", "mul-Latn", "ru"}, "whois.nic.airbus", e, w{"https://tld-rdap.verisign.com/airbus/v1/"}, t}
	z[39] = Zone{"airforce", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.airforce", "v0n1.nic.airforce", "v0n2.nic.airforce", "v0n3.nic.airforce", "v2n0.nic.airforce", "v2n1.nic.airforce"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.airforce", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i39() {
	z[40] = Zone{"airtel", r, x, 0x42, "Bharti Airtel Limited", "https://www.icann.org/en/registry-agreements/details/airtel", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, "whois.nic.airtel", e, w{"https://tld-rdap.verisign.com/airtel/v1/"}, f}
	z[41] = Zone{"akdn", r, x, 0x42, "Fondation Aga Khan (Aga Khan Foundation)", "https://www.icann.org/en/registry-agreements/details/akdn", w{"a0.nic.akdn", "a2.nic.akdn", "b0.nic.akdn", "c0.nic.akdn"}, n, n, n, "whois.nic.akdn", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[42] = Zone{"al", r, z[1795:1805], 0xa0, e, "https://www.iana.org/domains/root/db/al.html", w{"munnari.oz.au", "ns1.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, "http://www.akep.al/sq/kerkoni-domain", n, f}
	z[43] = Zone{"alcon", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[44] = Zone{"alfaromeo", r, x, 0x1042, "Fiat Chrysler Automobiles N.V.", "http://nic.alfaromeo/", n, n, n, n, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/alfaromeo/"}, f}
	z[45] = Zone{"alibaba", r, x, 0x42, "Alibaba Group Holding Limited", "http://nic.alibaba/", w{"a0.nic.alibaba", "a2.nic.alibaba", "b0.nic.alibaba", "c0.nic.alibaba"}, n, n, n, "whois.nic.alibaba", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[46] = Zone{"alipay", r, x, 0x42, "Alibaba Group Holding Limited", "http://nic.alipay/", w{"a0.nic.alipay", "a2.nic.alipay", "b0.nic.alipay", "c0.nic.alipay"}, n, n, n, "whois.nic.alipay", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[47] = Zone{"allfinanz", r, x, 0x42, "Allfinanz Deutsche Vermögensberatung Aktiengesellschaft", "https://nic.allfinanz/", w{"a.nic.allfinanz", "b.nic.allfinanz", "c.nic.allfinanz", "d.nic.allfinanz"}, n, n, w{"de"}, "whois.nic.allfinanz", e, w{"https://rdap.centralnic.com/allfinanz/"}, t}
}

//go:noinline
func i47() {
	z[48] = Zone{"allfinanzberater", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[49] = Zone{"allfinanzberatung", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[50] = Zone{"allstate", r, x, 0x42, "Allstate Fire and Casualty Insurance Company", "https://www.allstate.com/landingpages/nic/nic-allstate.aspx", w{"a0.nic.allstate", "a2.nic.allstate", "b0.nic.allstate", "c0.nic.allstate"}, n, n, n, "whois.nic.allstate", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[51] = Zone{"ally", r, x, 0x42, "Ally Financial Inc.", "https://www.ally.com/learn/nic/", w{"a.nic.ally", "b.nic.ally", "c.nic.ally", "x.nic.ally", "y.nic.ally", "z.nic.ally"}, n, n, n, "whois.nic.ally", e, w{"https://rdap.afilias-srs.net/rdap/ally/"}, f}
	z[52] = Zone{"alsace", r, x, 0x4c0, "Region Grand Est", "https://www.icann.org/en/registry-agreements/details/alsace", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, w{"FR-A"}, w{"mul-Latn"}, "whois.nic.alsace", e, w{"https://rdap.nic.alsace/"}, t}
	z[53] = Zone{"alstom", r, x, 0x42, "ALSTOM", "http://nic.alstom/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.alstom", e, w{"https://rdap.nic.alstom/"}, t}
	z[54] = Zone{"alt", r, x, 0x108, "IANA", "https://datatracker.ietf.org/doc/draft-dnsop-alt-tld-25", n, n, n, n, e, e, n, t}
	z[55] = Zone{"am", r, z[1805:1812], 0xa0, e, "https://www.amnic.net/", w{"fork.sth.dnsnode.net", "ns-cdn.amnic.net", "ns-pch.amnic.net", "ns-pri.nic.am"}, n, n, n, "whois.amnic.net", e, n, t}
}

//go:noinline
func i55() {
	z[56] = Zone{"amazon", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.amazon/", w{"dns1.nic.amazon", "dns2.nic.amazon", "dns3.nic.amazon", "dns4.nic.amazon", "dnsa.nic.amazon", "dnsb.nic.amazon", "dnsc.nic.amazon", "dnsd.nic.amazon"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.amazon", e, w{"https://rdap.nominet.uk/amazon/"}, t}
	z[57] = Zone{"americanexpress", r, x, 0x42, "American Express Travel Related Services Company, Inc.", "https://about.americanexpress.com/newsroom/#media-contacts", w{"a.nic.americanexpress", "b.nic.americanexpress", "c.nic.americanexpress", "ns1.dns.nic.americanexpress", "ns2.dns.nic.americanexpress", "ns3.dns.nic.americanexpress"}, n, n, w{"es"}, "whois.nic.americanexpress", e, w{"https://rdap.nic.americanexpress/"}, t}
	z[58] = Zone{"americanfamily", r, x, 0x42, "AmFam, Inc.", "https://www.amfam.com/nic-amfam", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.americanfamily", e, w{"https://tld-rdap.verisign.com/americanfamily/v1/"}, t}
	z[59] = Zone{"amex", r, x, 0x42, "American Express Travel Related Services Company, Inc.", "https://about.americanexpress.com/newsroom/#media-contacts", w{"a.nic.amex", "b.nic.amex", "c.nic.amex", "ns1.dns.nic.amex", "ns2.dns.nic.amex", "ns3.dns.nic.amex"}, n, n, w{"es"}, "whois.nic.amex", e, w{"https://rdap.nic.amex/"}, t}
	z[60] = Zone{"amfam", r, x, 0x42, "AmFam, Inc.", "https://www.amfam.com/nic-amfam", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.amfam", e, w{"https://tld-rdap.verisign.com/amfam/v1/"}, t}
	z[61] = Zone{"amica", r, x, 0x42, "Amica Mutual Insurance Company", "https://nic.amica/", w{"a.nic.amica", "b.nic.amica", "c.nic.amica", "ns1.dns.nic.amica", "ns2.dns.nic.amica", "ns3.dns.nic.amica"}, n, n, n, "whois.nic.amica", e, w{"https://rdap.nic.amica/"}, f}
	z[62] = Zone{"amp", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[63] = Zone{"amsterdam", r, x, 0xc4, "Gemeente Amsterdam", "https://nic.amsterdam/", w{"ns1.dns.amsterdam", "ns3.dns.amsterdam", "ns4.dns.amsterdam"}, n, w{"Amsterdam", "NL-NH"}, n, "whois.nic.amsterdam", e, w{"https://rdap.nic.amsterdam/"}, f}
}

//go:noinline
func i63() {
	z[64] = Zone{"an", r, z[1812:1816], 0x10a8, e, "https://www.iana.org/domains/root/db/an.html", n, n, n, n, e, e, n, f}
	z[65] = Zone{"analytics", r, x, 0x40, "Campus IP LLC", "https://nic.analytics/", w{"a.nic.analytics", "b.nic.analytics", "c.nic.analytics", "ns1.dns.nic.analytics", "ns2.dns.nic.analytics", "ns3.dns.nic.analytics"}, n, n, w{"es"}, "whois.nic.analytics", e, w{"https://rdap.nic.analytics/"}, t}
	z[66] = Zone{"android", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[67] = Zone{"anquan", r, x, 0x40, "Beijing Qihu Keji Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/anquan", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, n, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/"}, f}
	z[68] = Zone{"ansons", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[69] = Zone{"anthem", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[70] = Zone{"antivirus", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[71] = Zone{"anz", r, x, 0x42, "Australia and New Zealand Banking Group Limited", "http://nic.anz/", w{"a.nic.anz", "b.nic.anz", "c.nic.anz", "x.nic.anz", "y.nic.anz", "z.nic.anz"}, n, n, n, "whois.nic.anz", e, w{"https://rdap.nic.anz/"}, f}
}

//go:noinline
func i71() {
	z[72] = Zone{"ao", r, z[1816:1822], 0xa0, e, "https://www.dns.ao/ao/", w{"ao01.dns.pt", "ao03.dns.pt", "h.dns.pt"}, n, n, n, e, e, n, f}
	z[73] = Zone{"aol", r, x, 0x42, "Oath Inc.", "https://nic.aol.com/", w{"a0.nic.aol", "a2.nic.aol", "b0.nic.aol", "c0.nic.aol"}, n, n, n, "whois.nic.aol", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[74] = Zone{"apartments", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.apartments", "v0n1.nic.apartments", "v0n2.nic.apartments", "v0n3.nic.apartments", "v2n0.nic.apartments", "v2n1.nic.apartments"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.apartments", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[75] = Zone{"app", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"ja", "mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[76] = Zone{"apple", r, x, 0x42, "Apple Inc.", "https://www.apple.com/legal/intellectual-property/tld/", w{"a0.nic.apple", "a2.nic.apple", "b0.nic.apple", "c0.nic.apple"}, n, n, n, "whois.nic.apple", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[77] = Zone{"aq", r, z[1822:1823], 0xa0, "Antarctica Network Information Centre Limited", "https://www.iana.org/domains/root/db/aq.html", w{"fork.sth.dnsnode.net", "ns1.anycast.dns.aq", "ns99.dns.net.nz"}, n, n, n, e, e, n, f}
	z[78] = Zone{"aquarelle", r, x, 0x42, "Aquarelle.com", "http://www.nic.aquarelle/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.aquarelle", e, w{"https://rdap.nic.aquarelle/"}, t}
	z[79] = Zone{"aquitaine", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i79() {
	z[80] = Zone{"ar", r, z[1823:1832], 0xa0, e, "https://nic.ar/", w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "d.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, "whois.nic.ar", e, w{"https://rdap.nic.ar/"}, t}
	z[81] = Zone{"arab", r, x, 0x40, "League of Arab States", "http://nic.arab/", w{"gtld.alpha.aridns.net.au", "gtld.beta.aridns.net.au", "gtld.delta.aridns.net.au", "gtld.gamma.aridns.net.au"}, n, n, w{"ar"}, "whois.nic.arab", e, w{"https://rdap.nic.arab/"}, t}
	z[82] = Zone{"aramco", r, x, 0x42, "Aramco Services Company", "http://www.nic.aramco/", w{"a.nic.aramco", "b.nic.aramco", "c.nic.aramco", "ns4.dns.nic.aramco", "ns5.dns.nic.aramco", "ns6.dns.nic.aramco"}, n, n, w{"ar", "es"}, "whois.nic.aramco", e, w{"https://rdap.nic.aramco/"}, t}
	z[83] = Zone{"archi", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.archi", "a2.nic.archi", "b0.nic.archi", "c0.nic.archi"}, n, n, n, "whois.nic.archi", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[84] = Zone{"architect", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[85] = Zone{"army", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.army", "v0n1.nic.army", "v0n2.nic.army", "v0n3.nic.army", "v2n0.nic.army", "v2n1.nic.army"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.army", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[86] = Zone{"arpa", r, z[1832:1847], 0x140, "IANA", "https://www.iana.org/domains/arpa", w{"a.ns.arpa", "b.ns.arpa", "c.ns.arpa", "d.ns.arpa", "e.ns.arpa", "f.ns.arpa", "g.ns.arpa", "h.ns.arpa", "i.ns.arpa", "k.ns.arpa", "l.ns.arpa", "m.ns.arpa"}, n, n, n, "whois.iana.org", e, n, f}
	z[87] = Zone{"art", r, x, 0x40, "UK Creative Ideas Limited", "https://nic.art/", w{"a.nic.art", "b.nic.art", "c.nic.art", "d.nic.art"}, n, n, w{"ar", "da", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "mul-Mymr", "pl", "ru", "sv", "th", "zh"}, "whois.nic.art", e, w{"https://rdap.centralnic.com/art/"}, t}
}

//go:noinline
func i87() {
	z[88] = Zone{"arte", r, x, 0x42, "Association Relative à la Télévision Européenne G.E.I.E.", "https://www.icann.org/en/registry-agreements/details/arte", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.arte", e, w{"https://tld-rdap.verisign.com/arte/v1/"}, t}
	z[89] = Zone{"as", r, x, 0xe0, e, "https://nic.as/", w{"ns1.asnic.biz", "ns2.asnic.info", "ns3.asnic.org", "ns4.asnic.uk", "ns5.asnic.us", "pch.nic.as"}, n, n, n, "whois.nic.as", e, n, t}
	z[90] = Zone{"asda", r, x, 0x42, "Wal-Mart Stores, Inc.", "https://www.icann.org/en/registry-agreements/details/asda", w{"a0.nic.asda", "a2.nic.asda", "b0.nic.asda", "c0.nic.asda"}, n, n, n, "whois.nic.asda", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[91] = Zone{"asia", r, x, 0x4040, "DotAsia Organisation Limited", "https://www.dot.asia/", w{"a0.asia.afilias-nst.info", "a2.asia.afilias-nst.info", "b0.asia.afilias-nst.asia", "b2.asia.afilias-nst.org", "c0.asia.afilias-nst.info", "d0.asia.afilias-nst.asia"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hr", "hu", "is", "it", "ja", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "tr", "uk", "zh"}, "whois.nic.asia", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[92] = Zone{"associates", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.associates", "v0n1.nic.associates", "v0n2.nic.associates", "v0n3.nic.associates", "v2n0.nic.associates", "v2n1.nic.associates"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.associates", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[93] = Zone{"astrium", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[94] = Zone{"at", r, z[1847:1852], 0xa0, e, "https://www.nic.at/", w{"d.ns.at", "j.ns.at", "n.ns.at", "ns1.univie.ac.at", "ns2.univie.ac.at", "ns9.univie.ac.at", "r.ns.at", "u.ns.at"}, n, n, w{"mul-Latn"}, "whois.nic.at", e, n, t}
	z[95] = Zone{"athleta", r, x, 0x42, "The Gap, Inc.", "https://www.nic.gap/", w{"a.nic.athleta", "b.nic.athleta", "c.nic.athleta", "ns1.dns.nic.athleta", "ns2.dns.nic.athleta", "ns3.dns.nic.athleta"}, n, n, w{"es"}, "whois.nic.athleta", e, w{"https://rdap.nic.athleta/"}, t}
}

//go:noinline
func i95() {
	z[96] = Zone{"attorney", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.attorney", "v0n1.nic.attorney", "v0n2.nic.attorney", "v0n3.nic.attorney", "v2n0.nic.attorney", "v2n1.nic.attorney"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.attorney", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[97] = Zone{"au", r, z[1852:1872], 0xa0, e, "https://www.auda.org.au/", w{"q.au", "r.au", "s.au", "t.au"}, n, n, n, "whois.auda.org.au", e, n, f}
	z[98] = Zone{"auction", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.auction", "v0n1.nic.auction", "v0n2.nic.auction", "v0n3.nic.auction", "v2n0.nic.auction", "v2n1.nic.auction"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.auction", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[99] = Zone{"audi", r, x, 0x42, "AUDI Aktiengesellschaft", "https://www.nic.audi/nic/web/en.html", w{"a0.nic.audi", "a2.nic.audi", "b0.nic.audi", "c0.nic.audi"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.audi", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[100] = Zone{"audible", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.audible/", w{"dns1.nic.audible", "dns2.nic.audible", "dns3.nic.audible", "dns4.nic.audible", "dnsa.nic.audible", "dnsb.nic.audible", "dnsc.nic.audible", "dnsd.nic.audible"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.audible", e, w{"https://rdap.nominet.uk/audible/"}, t}
	z[101] = Zone{"audio", r, x, 0x40, "XYZ.COM LLC", "https://nic.audio/", w{"a.nic.audio", "b.nic.audio", "c.nic.audio", "d.nic.audio"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.audio", e, w{"https://rdap.centralnic.com/audio/"}, t}
	z[102] = Zone{"auspost", r, x, 0x42, "Australian Postal Corporation", "http://nic.auspost/", w{"a.nic.auspost", "b.nic.auspost", "c.nic.auspost", "x.nic.auspost", "y.nic.auspost", "z.nic.auspost"}, n, n, n, "whois.nic.auspost", e, w{"https://rdap.nic.auspost/"}, f}
	z[103] = Zone{"author", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.author/", w{"dns1.nic.author", "dns2.nic.author", "dns3.nic.author", "dns4.nic.author", "dnsa.nic.author", "dnsb.nic.author", "dnsc.nic.author", "dnsd.nic.author"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.author", e, w{"https://rdap.nominet.uk/author/"}, t}
}

//go:noinline
func i103() {
	z[104] = Zone{"auto", r, x, 0x40, "XYZ.COM LLC", "https://nic.auto/", w{"a.nic.auto", "b.nic.auto", "c.nic.auto", "d.nic.auto"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.auto", e, w{"https://rdap.centralnic.com/auto/"}, t}
	z[105] = Zone{"autoinsurance", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[106] = Zone{"autos", r, x, 0x40, "XYZ.COM LLC", "https://nic.autos/", w{"a.nic.autos", "b.nic.autos", "c.nic.autos", "d.nic.autos"}, n, n, n, "whois.nic.autos", e, w{"https://rdap.centralnic.com/autos/"}, f}
	z[107] = Zone{"avery", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[108] = Zone{"avianca", r, x, 0x42, "Avianca Inc.", "https://www.icann.org/en/registry-agreements/details/avianca", w{"a0.nic.avianca", "a2.nic.avianca", "b0.nic.avianca", "c0.nic.avianca"}, n, n, n, "whois.nic.avianca", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[109] = Zone{"aw", r, z[1872:1873], 0xa0, e, "https://www.setarnet.aw/", w{"aw01.setarnet.aw", "aw02.setarnet.aw", "ns1.dns.aw", "ns3.dns.aw", "ns4.dns.aw"}, n, n, n, "whois.nic.aw", e, n, f}
	z[110] = Zone{"aws", r, x, 0x42, "AWS Registry LLC", "https://nic.aws/", w{"dns1.nic.aws", "dns2.nic.aws", "dns3.nic.aws", "dns4.nic.aws", "dnsa.nic.aws", "dnsb.nic.aws", "dnsc.nic.aws", "dnsd.nic.aws"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.aws", e, w{"https://rdap.nominet.uk/aws/"}, t}
	z[111] = Zone{"ax", r, x, 0xa0, e, "https://whois.ax/", w{"ns1.aland.net", "ns2.aland.net", "ns3.alcom.fi", "ns4.alcom.fi"}, n, n, n, "whois.ax", e, n, f}
}

//go:noinline
func i111() {
	z[112] = Zone{"axa", r, x, 0x42, "AXA Group Operations SAS", "https://nic.axa/", w{"a.nic.axa", "b.nic.axa", "c.nic.axa", "ns1.dns.nic.axa", "ns2.dns.nic.axa", "ns3.dns.nic.axa"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.axa", e, w{"https://rdap.nic.axa/"}, t}
	z[113] = Zone{"axis", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[114] = Zone{"az", r, z[1873:1894], 0xa0, e, "https://www.nic.az/", w{"az.hostmaster.ua", "ns3.intrans.az", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[115] = Zone{"azure", r, x, 0x42, "Microsoft Corporation", "https://nic.azure/", w{"dns1.nic.azure", "dns2.nic.azure", "dns3.nic.azure", "dns4.nic.azure", "dnsa.nic.azure", "dnsb.nic.azure", "dnsc.nic.azure", "dnsd.nic.azure"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/azure/"}, t}
	z[116] = Zone{"ba", r, z[1894:1907], 0xa0, e, "https://www.nic.ba/", w{"ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, "http://nic.ba/lat/menu/view/13", n, f}
	z[117] = Zone{"baby", r, x, 0x40, "XYZ.COM LLC", "https://nic.baby/", w{"a.nic.baby", "b.nic.baby", "c.nic.baby", "d.nic.baby"}, n, n, w{"ar", "es", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.baby", e, w{"https://rdap.centralnic.com/baby/"}, t}
	z[118] = Zone{"baidu", r, x, 0x42, "Baidu, Inc.", "https://www.icann.org/en/registry-agreements/details/baidu", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/baidu/"}, t}
	z[119] = Zone{"banamex", r, x, 0x42, "Citigroup Inc.", "https://www.citi.com/", w{"a.nic.banamex", "b.nic.banamex", "c.nic.banamex", "ns1.dns.nic.banamex", "ns2.dns.nic.banamex", "ns3.dns.nic.banamex"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.banamex", e, w{"https://rdap.nic.banamex/"}, t}
}

//go:noinline
func i119() {
	z[120] = Zone{"bananarepublic", r, x, 0x1042, "The Gap, Inc.", "https://www.nic.gap/", n, n, n, w{"es"}, "whois.nic.bananarepublic", e, w{"https://rdap.nic.bananarepublic/"}, t}
	z[121] = Zone{"band", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.band", "v0n1.nic.band", "v0n2.nic.band", "v0n3.nic.band", "v2n0.nic.band", "v2n1.nic.band"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.band", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[122] = Zone{"bank", r, x, 0x40, "fTLD Registry Services LLC", "https://www.register.bank/", w{"d.nic.bank", "e.nic.bank", "f.nic.bank", "w.nic.bank", "x.nic.bank", "y.nic.bank"}, n, n, n, "whois.nic.bank", e, w{"https://rdap.nic.bank/"}, f}
	z[123] = Zone{"banque", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[124] = Zone{"bar", r, x, 0x40, "Punto 2012 Sociedad Anonima Promotora de Inversion de Capital Variable", e, w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, w{"mul-Latn"}, "whois.nic.bar", e, w{"https://rdap.centralnic.com/bar/"}, t}
	z[125] = Zone{"barcelona", r, x, 0x40, "Municipi de Barcelona", "https://www.domini.barcelona/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.barcelona", e, w{"https://rdap.nic.barcelona/"}, t}
	z[126] = Zone{"barclaycard", r, x, 0x42, "Barclays Bank PLC", "https://nic.barclaycard/", w{"a0.nic.barclaycard", "a2.nic.barclaycard", "b0.nic.barclaycard", "c0.nic.barclaycard"}, n, n, n, "whois.nic.barclaycard", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[127] = Zone{"barclays", r, x, 0x42, "Barclays Bank PLC", "https://nic.barclays/", w{"a0.nic.barclays", "a2.nic.barclays", "b0.nic.barclays", "c0.nic.barclays"}, n, n, n, "whois.nic.barclays", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i127() {
	z[128] = Zone{"barefoot", r, x, 0x42, "Gallo Vineyards, Inc.", "https://www.icann.org/en/registry-agreements/details/barefoot", w{"a0.nic.barefoot", "a2.nic.barefoot", "b0.nic.barefoot", "c0.nic.barefoot"}, n, n, n, "whois.nic.barefoot", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[129] = Zone{"bargains", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.bargains", "v0n1.nic.bargains", "v0n2.nic.bargains", "v0n3.nic.bargains", "v2n0.nic.bargains", "v2n1.nic.bargains"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.bargains", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[130] = Zone{"baseball", r, x, 0x40, "MLB Advanced Media DH, LLC", "https://www.mlb.com/official-information/baseball-nic", w{"a.nic.baseball", "b.nic.baseball", "c.nic.baseball", "ns1.dns.nic.baseball", "ns2.dns.nic.baseball", "ns3.dns.nic.baseball"}, n, n, w{"es"}, "whois.nic.baseball", e, w{"https://rdap.nic.baseball/"}, t}
	z[131] = Zone{"basketball", r, x, 0x40, "Fédération Internationale de Basketball (FIBA)", "https://get.basketball/", w{"a.nic.basketball", "b.nic.basketball", "c.nic.basketball", "x.nic.basketball", "y.nic.basketball", "z.nic.basketball"}, n, n, n, "whois.nic.basketball", e, w{"https://rdap.centralnic.com/basketball/"}, f}
	z[132] = Zone{"bauhaus", r, x, 0x42, "Werkhaus GmbH", "https://nic.bauhaus/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, n, "whois.nic.bauhaus", e, w{"https://rdap.nic.bauhaus/"}, f}
	z[133] = Zone{"bayern", r, x, 0x4c0, "Bayern Connect GmbH", "https://nic.bayern/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net", "ari.alpha.tldns.godaddy", "ari.beta.tldns.godaddy", "ari.delta.tldns.godaddy", "ari.gamma.tldns.godaddy"}, n, w{"DE-BY"}, w{"de"}, "whois.nic.bayern", e, w{"https://rdap.nic.bayern/"}, t}
	z[134] = Zone{"bb", r, z[1907:1916], 0xa0, e, "https://whois.telecoms.gov.bb/", w{"ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb"}, n, n, n, e, "http://whois.telecoms.gov.bb/search_domain.php", n, f}
	z[135] = Zone{"bbb", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i135() {
	z[136] = Zone{"bbc", r, x, 0x42, "British Broadcasting Corporation", "http://nic.bbc/", w{"dns1.nic.bbc", "dns2.nic.bbc", "dns3.nic.bbc", "dns4.nic.bbc", "dnsa.nic.bbc", "dnsb.nic.bbc", "dnsc.nic.bbc", "dnsd.nic.bbc"}, n, n, n, "whois.nic.bbc", e, w{"https://rdap.nominet.uk/bbc/"}, f}
	z[137] = Zone{"bbt", r, x, 0x42, "BB&T Corporation", "http://www.nic.bbt/", w{"a0.nic.bbt", "a2.nic.bbt", "b0.nic.bbt", "c0.nic.bbt"}, n, n, n, "whois.nic.bbt", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[138] = Zone{"bbva", r, x, 0x42, "BANCO BILBAO VIZCAYA ARGENTARIA, S.A.", "http://nic.bbva/", w{"dns1.nic.bbva", "dns2.nic.bbva", "dns3.nic.bbva", "dns4.nic.bbva", "dnsa.nic.bbva", "dnsb.nic.bbva", "dnsc.nic.bbva", "dnsd.nic.bbva"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.bbva", e, w{"https://rdap.nominet.uk/bbva/"}, t}
	z[139] = Zone{"bcg", r, x, 0x42, "The Boston Consulting Group, Inc.", "https://www.bcg.com/topic/nic", w{"a0.nic.bcg", "a2.nic.bcg", "b0.nic.bcg", "c0.nic.bcg"}, n, n, n, "whois.nic.bcg", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[140] = Zone{"bcn", r, x, 0x40, "Municipi de Barcelona", "https://www.icann.org/en/registry-agreements/details/bcn", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.bcn", e, w{"https://rdap.nic.bcn/"}, t}
	z[141] = Zone{"bd", r, z[1916:1923], 0xa8, e, "http://www.bttb.net.bd/", w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, "http://whois.btcl.net.bd/", n, t}
	z[142] = Zone{"be", r, x, 0xa0, e, e, w{"a.nsset.be", "b.nsset.be", "c.nsset.be", "d.nsset.be", "y.nsset.be", "z.nsset.be"}, n, n, w{"mul-Latn"}, "whois.dns.be", e, n, t}
	z[143] = Zone{"beats", r, x, 0x42, "Beats Electronics, LLC", "https://www.icann.org/en/registry-agreements/details/beats", w{"a0.nic.beats", "a2.nic.beats", "b0.nic.beats", "c0.nic.beats"}, n, n, n, "whois.nic.beats", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i143() {
	z[144] = Zone{"beauty", r, x, 0x40, "XYZ.COM LLC", "https://nic.beauty/", w{"a.nic.beauty", "b.nic.beauty", "c.nic.beauty", "d.nic.beauty"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.beauty", e, w{"https://rdap.centralnic.com/beauty/"}, t}
	z[145] = Zone{"beer", r, x, 0x40, "Registry Services, LLC", "http://nic.beer/", w{"a.nic.beer", "b.nic.beer", "c.nic.beer", "x.nic.beer", "y.nic.beer", "z.nic.beer"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.beer", e, w{"https://rdap.nic.beer/"}, t}
	z[146] = Zone{"beknown", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[147] = Zone{"bentley", r, x, 0x42, "Bentley Motors Limited", "https://www.nic.bentley/", w{"dns1.nic.bentley", "dns2.nic.bentley", "dns3.nic.bentley", "dns4.nic.bentley", "dnsa.nic.bentley", "dnsb.nic.bentley", "dnsc.nic.bentley", "dnsd.nic.bentley"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.bentley", e, w{"https://rdap.nominet.uk/bentley/"}, t}
	z[148] = Zone{"berlin", r, x, 0xc4, "dotBERLIN GmbH & Co. KG", e, w{"a.dns.nic.berlin", "m.dns.nic.berlin", "n.dns.nic.berlin"}, n, w{"Berlin", "DE-BE"}, w{"mul-Cyrl", "mul-Latn"}, "whois.nic.berlin", e, w{"https://rdap.nic.berlin/v1/"}, t}
	z[149] = Zone{"best", r, x, 0x40, "BestTLD Pty Ltd", "https://nic.best/", w{"a.nic.best", "b.nic.best", "c.nic.best", "d.nic.best"}, n, n, w{"ar", "da", "de", "es", "fr", "ja", "ko", "nl", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.best", e, w{"https://rdap.centralnic.com/best/"}, t}
	z[150] = Zone{"bestbuy", r, x, 0x42, "BBY Solutions, Inc.", "http://nic.bestbuy/", w{"a0.nic.bestbuy", "a2.nic.bestbuy", "b0.nic.bestbuy", "c0.nic.bestbuy"}, n, n, n, "whois.nic.bestbuy", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[151] = Zone{"bet", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.bet", "a2.nic.bet", "b0.nic.bet", "c0.nic.bet"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.bet", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i151() {
	z[152] = Zone{"bf", r, z[1923:1924], 0xa0, e, e, w{"a.registre.bf", "dns-tld.ird.fr", "ns-bf.nic.fr", "pch.ns.registre.bf"}, n, n, n, "whois.registre.bf", e, n, t}
	z[153] = Zone{"bg", r, z[1924:1960], 0xa0, e, "https://www.register.bg/user/", w{"a.nic.bg", "b.nic.bg", "c.nic.bg", "d.nic.bg", "e.nic.bg", "p.nic.bg"}, n, n, w{"bg-BG", "ru-BG"}, "whois.register.bg", e, n, t}
	z[154] = Zone{"bh", r, z[1960:1971], 0xa0, e, e, w{"a.nic.bh", "b.nic.bh", "c.nic.bh", "d.nic.bh"}, n, n, n, "whois.nic.bh", e, n, f}
	z[155] = Zone{"bharti", r, x, 0x42, "Bharti Enterprises (Holding) Private Limited", "https://www.icann.org/en/registry-agreements/details/bharti", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, e, e, w{"https://tld-rdap.verisign.com/bharti/v1/"}, f}
	z[156] = Zone{"bi", r, z[1971:1981], 0xa0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, "whois1.nic.bi", e, n, f}
	z[157] = Zone{"bible", r, x, 0x40, "American Bible Society", "https://get.bible/", w{"a.nic.bible", "b.nic.bible", "c.nic.bible", "ns1.dns.nic.bible", "ns2.dns.nic.bible", "ns3.dns.nic.bible"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.bible", e, w{"https://rdap.nic.bible/"}, t}
	z[158] = Zone{"bid", r, x, 0x40, "dot Bid Limited", "http://nic.bid/", w{"a.nic.bid", "b.nic.bid", "c.nic.bid", "ns1.dns.nic.bid", "ns2.dns.nic.bid", "ns3.dns.nic.bid"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.bid", e, w{"https://rdap.nic.bid/"}, t}
	z[159] = Zone{"bike", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.bike", "v0n1.nic.bike", "v0n2.nic.bike", "v0n3.nic.bike", "v2n0.nic.bike", "v2n1.nic.bike"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.bike", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i159() {
	z[160] = Zone{"bing", r, x, 0x42, "Microsoft Corporation", "https://nic.bing/", w{"dns1.nic.bing", "dns2.nic.bing", "dns3.nic.bing", "dns4.nic.bing", "dnsa.nic.bing", "dnsb.nic.bing", "dnsc.nic.bing", "dnsd.nic.bing"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/bing/"}, t}
	z[161] = Zone{"bingo", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.bingo", "v0n1.nic.bingo", "v0n2.nic.bingo", "v0n3.nic.bingo", "v2n0.nic.bingo", "v2n1.nic.bingo"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.bingo", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[162] = Zone{"bio", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.bio", "a2.nic.bio", "b0.nic.bio", "c0.nic.bio"}, n, n, w{"de"}, "whois.nic.bio", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[163] = Zone{"biz", r, z[1981:1983], 0x40, "Registry Services, LLC", e, w{"a.gtld.biz", "b.gtld.biz", "c.gtld.biz", "w.gtld.biz", "x.gtld.biz", "y.gtld.biz"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "sv", "zh"}, "whois.nic.biz", e, w{"https://rdap.nic.biz/"}, t}
	z[164] = Zone{"bj", r, z[1983:1990], 0xa0, e, e, w{"ns-bj.afrinic.net", "ns-bj.nic.fr", "ns1.nic.bj", "ns2.nic.bj", "pch.nic.bj"}, n, n, n, "whois.nic.bj", e, n, f}
	z[165] = Zone{"black", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.black", "a2.nic.black", "b0.nic.black", "c0.nic.black"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.black", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[166] = Zone{"blackfriday", r, x, 0x40, "Registry Services, LLC", "https://uniregistry.link/extensions/blackfriday/", w{"a.nic.blackfriday", "b.nic.blackfriday", "c.nic.blackfriday", "x.nic.blackfriday", "y.nic.blackfriday", "z.nic.blackfriday"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.blackfriday", e, w{"https://rdap.nic.blackfriday/"}, t}
	z[167] = Zone{"blanco", r, x, 0x1040, e, e, n, n, n, w{"mul-Hang", "mul-Hani", "mul-Hano", "mul-Latn", "ru"}, "whois.nic.blanco", e, n, t}
}

//go:noinline
func i167() {
	z[168] = Zone{"blockbuster", r, x, 0x42, "Dish DBS Corporation", "https://www.dishtlds.com/blockbuster/", w{"a0.nic.blockbuster", "a2.nic.blockbuster", "b0.nic.blockbuster", "c0.nic.blockbuster"}, n, n, n, "whois.nic.blockbuster", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[169] = Zone{"blog", r, z[1990:2016], 0x40, "Knock Knock WHOIS There, LLC", "https://my.blog/", w{"a.nic.blog", "b.nic.blog", "c.nic.blog", "d.nic.blog"}, n, n, w{"zh"}, "whois.nic.blog", e, w{"https://rdap.centralnic.com/blog/"}, t}
	z[170] = Zone{"bloomberg", r, x, 0x42, "Bloomberg IP Holdings LLC", "https://www.icann.org/en/registry-agreements/details/bloomberg", w{"v0n0.nic.bloomberg", "v0n1.nic.bloomberg", "v0n2.nic.bloomberg", "v0n3.nic.bloomberg", "v2n0.nic.bloomberg", "v2n1.nic.bloomberg"}, n, n, n, "whois.nic.bloomberg", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[171] = Zone{"bloomingdales", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[172] = Zone{"blue", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.blue", "a2.nic.blue", "b0.nic.blue", "b2.nic.blue", "c0.nic.blue"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.blue", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[173] = Zone{"bm", r, z[2016:2021], 0xa0, e, "https://www.bermudanic.bm/", w{"a0.bm.afilias-nst.info", "a2.bm.afilias-nst.info", "b0.bm.afilias-nst.org", "b2.bm.afilias-nst.org", "c0.bm.afilias-nst.info", "d0.bm.afilias-nst.org"}, n, n, n, e, "https://www.bermudanic.bm/whois.htm", n, f}
	z[174] = Zone{"bms", r, x, 0x42, "Bristol-Myers Squibb Company", "http://nic.bms/", w{"a0.nic.bms", "a2.nic.bms", "b0.nic.bms", "c0.nic.bms"}, n, n, n, "whois.nic.bms", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[175] = Zone{"bmw", r, x, 0x42, "Bayerische Motoren Werke Aktiengesellschaft", "https://nic.bmw/", w{"a.nic.bmw", "b.nic.bmw", "c.nic.bmw", "d.nic.bmw"}, n, n, w{"de"}, "whois.nic.bmw", e, w{"https://rdap.centralnic.com/bmw/"}, t}
}

//go:noinline
func i175() {
	z[176] = Zone{"bn", r, z[2021:2026], 0xa0, e, "https://www.iana.org/domains/root/db/bn.html", w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, "whois.bnnic.bn", e, n, f}
	z[177] = Zone{"bnl", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.bnl", e, n, f}
	z[178] = Zone{"bnpparibas", r, x, 0x42, "BNP Paribas", "https://nic.bnpparibas/en/", w{"a0.nic.bnpparibas", "a2.nic.bnpparibas", "b0.nic.bnpparibas", "c0.nic.bnpparibas"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.bnpparibas", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[179] = Zone{"bo", r, z[2026:2035], 0xa0, e, e, w{"anycast.ns.nic.bo", "ns.dns.br", "ns.nic.bo", "ns2.nic.fr"}, n, n, w{"es"}, "whois.nic.bo", e, n, t}
	z[180] = Zone{"boats", r, x, 0x40, "XYZ.COM LLC", "https://nic.boats/", w{"a.nic.boats", "b.nic.boats", "c.nic.boats", "d.nic.boats"}, n, n, n, "whois.nic.boats", e, w{"https://rdap.centralnic.com/boats/"}, f}
	z[181] = Zone{"boehringer", r, x, 0x42, "Boehringer Ingelheim International GmbH", "https://www.icann.org/en/registry-agreements/details/boehringer", w{"a0.nic.boehringer", "a2.nic.boehringer", "b0.nic.boehringer", "c0.nic.boehringer"}, n, n, n, "whois.nic.boehringer", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[182] = Zone{"bofa", r, x, 0x42, "Bank of America Corporation", "https://www.icann.org/en/registry-agreements/details/bofa", w{"v0n0.nic.bofa", "v0n1.nic.bofa", "v0n2.nic.bofa", "v0n3.nic.bofa", "v2n0.nic.bofa", "v2n1.nic.bofa"}, n, n, n, "whois.nic.bofa", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[183] = Zone{"bom", r, x, 0x42, "Núcleo de Informação e Coordenação do Ponto BR - NIC.br", "https://www.icann.org/en/registry-agreements/details/bom", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
}

//go:noinline
func i183() {
	z[184] = Zone{"bond", r, x, 0x42, "ShortDot SA", "https://shortdot.bond/bond/", w{"a.nic.bond", "b.nic.bond", "c.nic.bond", "d.nic.bond"}, n, n, w{"cs", "da", "de", "es", "fr", "he", "is", "it", "ja", "ko", "lb", "lo", "mul-Grek", "mul-Latn", "pt", "ro", "ru", "tr", "uk", "zh"}, "whois.nic.bond", e, w{"https://rdap.centralnic.com/bond/"}, t}
	z[185] = Zone{"boo", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[186] = Zone{"book", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.book/", w{"dns1.nic.book", "dns2.nic.book", "dns3.nic.book", "dns4.nic.book", "dnsa.nic.book", "dnsb.nic.book", "dnsc.nic.book", "dnsd.nic.book"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.book", e, w{"https://rdap.nominet.uk/book/"}, t}
	z[187] = Zone{"booking", r, x, 0x42, "Booking.com B.V.", "https://nic.booking/", w{"a.nic.booking", "b.nic.booking", "c.nic.booking", "ns1.dns.nic.booking", "ns2.dns.nic.booking", "ns3.dns.nic.booking"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.booking", e, w{"https://rdap.nic.booking/"}, t}
	z[188] = Zone{"boots", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.boots", e, n, f}
	z[189] = Zone{"bosch", r, x, 0x42, "Robert Bosch GMBH", "https://nic.bosch/", w{"a0.nic.bosch", "a2.nic.bosch", "b0.nic.bosch", "c0.nic.bosch"}, n, n, n, "whois.nic.bosch", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[190] = Zone{"bostik", r, x, 0x42, "Bostik SA", "https://www.icann.org/en/registry-agreements/details/bostik", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.bostik", e, w{"https://rdap.nic.bostik/"}, t}
	z[191] = Zone{"boston", r, x, 0xc4, "Registry Services, LLC", "http://nic.boston/", w{"a.nic.boston", "b.nic.boston", "c.nic.boston", "x.nic.boston", "y.nic.boston", "z.nic.boston"}, n, w{"Boston"}, n, "whois.nic.boston", e, w{"https://rdap.nic.boston/"}, f}
}

//go:noinline
func i191() {
	z[192] = Zone{"bot", r, x, 0x40, "Amazon Registry Services, Inc.", "https://www.amazonregistry.com/bot", w{"dns1.nic.bot", "dns2.nic.bot", "dns3.nic.bot", "dns4.nic.bot", "dnsa.nic.bot", "dnsb.nic.bot", "dnsc.nic.bot", "dnsd.nic.bot"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.bot", e, w{"https://rdap.nominet.uk/bot/"}, t}
	z[193] = Zone{"boutique", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.boutique", "v0n1.nic.boutique", "v0n2.nic.boutique", "v0n3.nic.boutique", "v2n0.nic.boutique", "v2n1.nic.boutique"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.boutique", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[194] = Zone{"box", r, x, 0x40, "Intercap Registry Inc.", "https://whois.nic.box/", w{"a.nic.box", "b.nic.box", "c.nic.box", "d.nic.box"}, n, n, w{"zh", "zh-Hans", "zh-Hant"}, "whois.nic.box", e, w{"https://rdap.centralnic.com/box/"}, t}
	z[195] = Zone{"br", r, z[2035:2141], 0xa0, e, "https://nic.br/", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt-BR"}, "whois.registro.br", e, w{"https://rdap.registro.br/"}, t}
	z[196] = Zone{"bradesco", r, x, 0x42, "Banco Bradesco S.A.", "https://nic.bradesco/", w{"a0.nic.bradesco", "a2.nic.bradesco", "b0.nic.bradesco", "c0.nic.bradesco"}, n, n, n, "whois.nic.bradesco", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[197] = Zone{"bridgestone", r, x, 0x42, "Bridgestone Corporation", "https://nic.bridgestone/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.bridgestone", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[198] = Zone{"broadway", r, x, 0x40, "Celebrate Broadway, Inc.", "https://discover.broadway/", w{"dns1.nic.broadway", "dns2.nic.broadway", "dns3.nic.broadway", "dns4.nic.broadway", "dnsa.nic.broadway", "dnsb.nic.broadway", "dnsc.nic.broadway", "dnsd.nic.broadway"}, n, n, n, "whois.nic.broadway", e, w{"https://rdap.nominet.uk/broadway/"}, f}
	z[199] = Zone{"broker", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.broker", "v0n1.nic.broker", "v0n2.nic.broker", "v0n3.nic.broker", "v2n0.nic.broker", "v2n1.nic.broker"}, n, n, n, "whois.nic.broker", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i199() {
	z[200] = Zone{"brother", r, x, 0x42, "Brother Industries, Ltd.", "https://nic.brother/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.brother", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[201] = Zone{"brussels", r, x, 0xc4, "DNS.be vzw", "https://www.dnsbelgium.be/", w{"a.nsset.brussels", "b.nsset.brussels", "c.nsset.brussels", "d.nsset.brussels", "y.nsset.brussels", "z.nsset.brussels"}, n, w{"Brussels", "BE-BRU"}, w{"mul-Latn"}, "whois.nic.brussels", e, w{"https://rdap.nic.brussels/"}, t}
	z[202] = Zone{"bs", r, z[2141:2147], 0xa0, e, e, w{"anyns.dns.bs", "anyns.pch.net", "ns36.cdns.net"}, n, n, n, e, "http://www.nic.bs/cgi-bin/search.pl", n, t}
	z[203] = Zone{"bt", r, z[2147:2152], 0xa0, e, e, w{"auth00.ns.uu.net", "auth61.ns.uu.net", "ns.itu.ch", "ns1.druknet.bt", "ns2.druknet.bt", "ns3.druknet.bt", "phloem.uoregon.edu"}, n, n, n, e, "http://www.nic.bt/", n, f}
	z[204] = Zone{"budapest", r, x, 0x10c4, "Minds + Machines Group Limited", "https://www.icann.org/en/registry-agreements/details/budapest", n, n, w{"Budapest", "HU-BU"}, w{"hu"}, "whois.nic.budapest", e, w{"https://rdap.nominet.uk/budapest/"}, t}
	z[205] = Zone{"bugatti", r, x, 0x1042, "Bugatti International SA", e, n, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/bugatti/"}, t}
	z[206] = Zone{"buick", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[207] = Zone{"build", r, x, 0x40, "Plan Bee LLC", "https://about.build/", w{"a.nic.build", "b.nic.build", "c.nic.build", "d.nic.build"}, n, n, n, "whois.nic.build", e, w{"https://rdap.centralnic.com/build/"}, f}
}

//go:noinline
func i207() {
	z[208] = Zone{"builders", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.builders", "v0n1.nic.builders", "v0n2.nic.builders", "v0n3.nic.builders", "v2n0.nic.builders", "v2n1.nic.builders"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.builders", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[209] = Zone{"business", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.business", "v0n1.nic.business", "v0n2.nic.business", "v0n3.nic.business", "v2n0.nic.business", "v2n1.nic.business"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.business", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[210] = Zone{"buy", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.buy/", w{"dns1.nic.buy", "dns2.nic.buy", "dns3.nic.buy", "dns4.nic.buy", "dnsa.nic.buy", "dnsb.nic.buy", "dnsc.nic.buy", "dnsd.nic.buy"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh-CN", "zh-TW"}, "whois.nic.buy", e, w{"https://rdap.nominet.uk/buy/"}, t}
	z[211] = Zone{"buzz", r, x, 0x40, "DOTSTRATEGY CO.", e, w{"a.nic.buzz", "b.nic.buzz", "c.nic.buzz", "ns1.dns.nic.buzz", "ns2.dns.nic.buzz", "ns3.dns.nic.buzz"}, n, n, w{"es"}, "whois.nic.buzz", e, w{"https://rdap.nic.buzz/"}, t}
	z[212] = Zone{"bv", r, x, 0xa8, e, e, w{"nac.no", "nn.uninett.no", "server.nordu.net", "x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, "whois.norid.no", e, n, f}
	z[213] = Zone{"bw", r, z[2152:2156], 0xa0, e, e, w{"dns1.nic.net.bw", "master.btc.net.bw", "ns-bw.afrinic.net", "pch.nic.net.bw"}, n, n, n, "whois.nic.net.bw", e, n, f}
	z[214] = Zone{"bway", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[215] = Zone{"by", r, z[2156:2161], 0xa0, e, "https://domain.by/", w{"dns1.tld.becloudby.com", "dns2.tld.becloudby.com", "dns3.tld.becloudby.com", "dns4.tld.becloudby.com", "dns7.tld.becloudby.com"}, n, n, n, "whois.cctld.by", e, n, f}
}

//go:noinline
func i215() {
	z[216] = Zone{"bz", r, z[2161:2168], 0xe0, e, "https://www.belizenic.bz/", w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, "whois.afilias-grs.info", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[217] = Zone{"bzh", r, x, 0x4c0, "Association www.bzh", "https://www.pik.bzh/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, w{"FR-E"}, w{"fr"}, "whois.nic.bzh", e, w{"https://rdap.nic.bzh/"}, t}
	z[218] = Zone{"ca", r, z[2168:2184], 0xa0, e, e, w{"any.ca-servers.ca", "c.ca-servers.ca", "d.ca-servers.ca", "j.ca-servers.ca"}, n, n, w{"fr"}, "whois.cira.ca", e, w{"https://rdap.ca.fury.ca/rdap/"}, t}
	z[219] = Zone{"cab", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cab", "v0n1.nic.cab", "v0n2.nic.cab", "v0n3.nic.cab", "v2n0.nic.cab", "v2n1.nic.cab"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cab", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[220] = Zone{"cadillac", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[221] = Zone{"cafe", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.cafe", "v0n1.nic.cafe", "v0n2.nic.cafe", "v0n3.nic.cafe", "v2n0.nic.cafe", "v2n1.nic.cafe"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cafe", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[222] = Zone{"cal", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[223] = Zone{"call", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.call/", w{"dns1.nic.call", "dns2.nic.call", "dns3.nic.call", "dns4.nic.call", "dnsa.nic.call", "dnsb.nic.call", "dnsc.nic.call", "dnsd.nic.call"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.call", e, w{"https://rdap.nominet.uk/call/"}, t}
}

//go:noinline
func i223() {
	z[224] = Zone{"calvinklein", r, x, 0x42, "PVH gTLD Holdings LLC", "https://nic.calvinklein/", w{"a.nic.calvinklein", "b.nic.calvinklein", "c.nic.calvinklein", "ns1.dns.nic.calvinklein", "ns2.dns.nic.calvinklein", "ns3.dns.nic.calvinklein"}, n, n, w{"es"}, "whois.nic.calvinklein", e, w{"https://rdap.nic.calvinklein/"}, t}
	z[225] = Zone{"cam", r, x, 0x40, "Cam Connecting SARL", "https://nic.cam/", w{"a.nic.cam", "b.nic.cam", "c.nic.cam", "d.nic.cam"}, n, n, w{"ar", "az", "bg", "bs", "ca", "cs", "da", "de", "el", "es", "et", "fi", "fr", "he", "hr", "hu", "is", "it", "ja", "ka", "ko", "lb", "lt", "lv", "mk", "mul-Armn", "mul-Latn", "mul-Tavt", "mul-Thai", "nl", "no", "pl", "pt", "ro", "ro-MD", "ru", "sk", "sl", "sq", "sr", "sv", "tr", "uk"}, "whois.nic.cam", e, w{"https://rdap.centralnic.com/cam/"}, t}
	z[226] = Zone{"camera", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.camera", "v0n1.nic.camera", "v0n2.nic.camera", "v0n3.nic.camera", "v2n0.nic.camera", "v2n1.nic.camera"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.camera", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[227] = Zone{"camp", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.camp", "v0n1.nic.camp", "v0n2.nic.camp", "v0n3.nic.camp", "v2n0.nic.camp", "v2n1.nic.camp"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.camp", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[228] = Zone{"canalplus", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[229] = Zone{"cancerresearch", r, x, 0x1040, "Australian Cancer Research Foundation", e, n, n, n, n, "whois.nic.cancerresearch", e, w{"https://rdap.nic.cancerresearch/"}, f}
	z[230] = Zone{"canon", r, x, 0x42, "Canon Inc.", "https://nic.canon/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.canon", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[231] = Zone{"capetown", r, x, 0xc4, "ZA Central Registry NPC trading as ZA Central Registry", e, w{"coza1.dnsnode.net", "nsp.netnod.se"}, n, w{"Cape Town"}, n, "whois.nic.capetown", e, w{"https://rdap.nic.capetown/rdap/"}, t}
}

//go:noinline
func i231() {
	z[232] = Zone{"capital", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.capital", "v0n1.nic.capital", "v0n2.nic.capital", "v0n3.nic.capital", "v2n0.nic.capital", "v2n1.nic.capital"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.capital", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[233] = Zone{"capitalone", r, x, 0x42, "Capital One Financial Corporation", "https://www.icann.org/en/registry-agreements/details/capitalone", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, "whois.nic.capitalone", e, w{"https://tld-rdap.verisign.com/capitalone/v1/"}, f}
	z[234] = Zone{"car", r, x, 0x40, "XYZ.COM LLC", "https://nic.car/", w{"a.nic.car", "b.nic.car", "c.nic.car", "d.nic.car"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.car", e, w{"https://rdap.centralnic.com/car/"}, t}
	z[235] = Zone{"caravan", r, x, 0x42, "Caravan International, Inc.", "http://nic.caravan/", w{"a.nic.caravan", "b.nic.caravan", "c.nic.caravan", "ns1.dns.nic.caravan", "ns2.dns.nic.caravan", "ns3.dns.nic.caravan"}, n, n, w{"zh"}, "whois.nic.caravan", e, w{"https://rdap.nic.caravan/"}, t}
	z[236] = Zone{"cards", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cards", "v0n1.nic.cards", "v0n2.nic.cards", "v0n3.nic.cards", "v2n0.nic.cards", "v2n1.nic.cards"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cards", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[237] = Zone{"care", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.care", "v0n1.nic.care", "v0n2.nic.care", "v0n3.nic.care", "v2n0.nic.care", "v2n1.nic.care"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.care", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[238] = Zone{"career", r, x, 0x40, "dotCareer LLC", e, w{"dns1.nic.career", "dns2.nic.career", "dns3.nic.career", "dns4.nic.career", "dnsa.nic.career", "dnsb.nic.career", "dnsc.nic.career", "dnsd.nic.career"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.career", e, w{"https://rdap.nominet.uk/career/"}, t}
	z[239] = Zone{"careers", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.careers", "v0n1.nic.careers", "v0n2.nic.careers", "v0n3.nic.careers", "v2n0.nic.careers", "v2n1.nic.careers"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.careers", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i239() {
	z[240] = Zone{"carinsurance", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[241] = Zone{"cars", r, x, 0x40, "XYZ.COM LLC", "https://nic.cars/", w{"a.nic.cars", "b.nic.cars", "c.nic.cars", "d.nic.cars"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.cars", e, w{"https://rdap.centralnic.com/cars/"}, t}
	z[242] = Zone{"cartier", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.cartier", e, n, t}
	z[243] = Zone{"casa", r, x, 0x40, "Registry Services, LLC", "http://nic.casa/", w{"a.nic.casa", "b.nic.casa", "c.nic.casa", "x.nic.casa", "y.nic.casa", "z.nic.casa"}, n, n, w{"es", "it", "pt"}, "whois.nic.casa", e, w{"https://rdap.nic.casa/"}, t}
	z[244] = Zone{"case", r, x, 0x42, "Digity, LLC", "http://nic.case/", w{"a.nic.case", "b.nic.case", "c.nic.case", "d.nic.case"}, n, n, n, "whois.nic.case", e, w{"https://rdap.centralnic.com/case/"}, f}
	z[245] = Zone{"caseih", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.caseih", e, n, f}
	z[246] = Zone{"cash", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cash", "v0n1.nic.cash", "v0n2.nic.cash", "v0n3.nic.cash", "v2n0.nic.cash", "v2n1.nic.cash"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cash", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[247] = Zone{"casino", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.casino", "v0n1.nic.casino", "v0n2.nic.casino", "v0n3.nic.casino", "v2n0.nic.casino", "v2n1.nic.casino"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.casino", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i247() {
	z[248] = Zone{"cat", r, x, 0x4440, "Fundació puntCAT", "https://domini.cat/", w{"anyc1.irondns.net", "cat.pch.net", "ns.nic.cat", "ns1.nic.es", "nsc.nic.de", "switch.nic.cat"}, n, w{"ES-CT"}, w{"ca"}, "whois.nic.cat", e, w{"https://rdap.nic.cat/"}, t}
	z[249] = Zone{"catalonia", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[250] = Zone{"catering", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.catering", "v0n1.nic.catering", "v0n2.nic.catering", "v0n3.nic.catering", "v2n0.nic.catering", "v2n1.nic.catering"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.catering", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[251] = Zone{"catholic", r, x, 0x40, "Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)", "http://nic.catholic/", w{"a.nic.catholic", "b.nic.catholic", "c.nic.catholic", "x.nic.catholic", "y.nic.catholic", "z.nic.catholic"}, n, n, n, "whois.nic.catholic", e, w{"https://rdap.nic.catholic/"}, f}
	z[252] = Zone{"cba", r, x, 0x42, "COMMONWEALTH BANK OF AUSTRALIA", "http://nic.cba/", w{"a.nic.cba", "b.nic.cba", "c.nic.cba", "x.nic.cba", "y.nic.cba", "z.nic.cba"}, n, n, n, "whois.nic.cba", e, w{"https://rdap.nic.cba/"}, f}
	z[253] = Zone{"cbn", r, x, 0x42, "The Christian Broadcasting Network, Inc.", "http://nic.cbn/", w{"a.nic.cbn", "b.nic.cbn", "c.nic.cbn", "ns1.dns.nic.cbn", "ns2.dns.nic.cbn", "ns3.dns.nic.cbn"}, n, n, w{"zh"}, "whois.nic.cbn", e, w{"https://rdap.nic.cbn/"}, t}
	z[254] = Zone{"cbre", r, x, 0x42, "CBRE, Inc.", "http://nic.cbre/", w{"a.nic.cbre", "b.nic.cbre", "c.nic.cbre", "ns1.dns.nic.cbre", "ns2.dns.nic.cbre", "ns3.dns.nic.cbre"}, n, n, w{"es"}, "whois.nic.cbre", e, w{"https://rdap.nic.cbre/"}, t}
	z[255] = Zone{"cbs", r, x, 0x1042, "CBS Domains Inc.", "https://nic.cbs/", n, n, n, n, "whois.nic.cbs", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i255() {
	z[256] = Zone{"cc", r, z[2184:2188], 0xe0, e, e, w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, "ccwhois.verisign-grs.com", e, w{"https://tld-rdap.verisign.com/cc/v1/"}, t}
	z[257] = Zone{"cd", r, z[2188:2194], 0xe0, e, "https://www.nic.cd/nic.cd/", w{"ns-root-21.scpt-network.net", "ns-root-22.scpt-network.net", "ns-root-23.scpt-network.net"}, n, n, n, "whois.cd", e, n, f}
	z[258] = Zone{"ceb", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[259] = Zone{"center", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.center", "v0n1.nic.center", "v0n2.nic.center", "v0n3.nic.center", "v2n0.nic.center", "v2n1.nic.center"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.center", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[260] = Zone{"ceo", r, x, 0x40, "XYZ.COM LLC", e, w{"a.nic.ceo", "b.nic.ceo", "c.nic.ceo", "d.nic.ceo"}, n, n, w{"ar", "da", "de", "es", "fr", "ja", "ko", "mul-Latn", "nl", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.ceo", e, w{"https://rdap.centralnic.com/ceo/"}, t}
	z[261] = Zone{"cern", r, x, 0x42, "European Organization for Nuclear Research (\"CERN\")", "https://nic.cern/", w{"a0.nic.cern", "a2.nic.cern", "b0.nic.cern", "c0.nic.cern"}, n, n, n, "whois.nic.cern", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[262] = Zone{"cf", r, x, 0xa0, e, e, w{"a.ns.cf", "b.ns.cf", "c.ns.cf", "d.ns.cf"}, n, n, n, "whois.dot.cf", e, n, t}
	z[263] = Zone{"cfa", r, x, 0x42, "CFA Institute", "https://www.icann.org/en/registry-agreements/details/cfa", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, "whois.nic.cfa", e, w{"https://tld-rdap.verisign.com/cfa/v1/"}, f}
}

//go:noinline
func i263() {
	z[264] = Zone{"cfd", r, x, 0x40, "ShortDot SA", "https://shortdot.bond/cfd/", w{"a.nic.cfd", "b.nic.cfd", "c.nic.cfd", "d.nic.cfd"}, n, n, n, "whois.nic.cfd", e, w{"https://rdap.centralnic.com/cfd/"}, f}
	z[265] = Zone{"cg", r, x, 0xa0, e, e, w{"dns-fr.dnsafrica.net", "dns-za.dnsafrica.net", "sunic.sunet.se"}, n, n, n, e, "http://www.nic.cg/cgi-bin/whois.pl", n, f}
	z[266] = Zone{"ch", r, x, 0xa0, e, e, w{"a.nic.ch", "b.nic.ch", "d.nic.ch", "e.nic.ch", "f.nic.ch"}, n, n, n, "whois.nic.ch", e, w{"https://rdap.nic.ch/"}, t}
	z[267] = Zone{"chanel", r, x, 0x42, "Chanel International B.V.", "http://nic.chanel/", w{"a0.nic.chanel", "a2.nic.chanel", "b0.nic.chanel", "c0.nic.chanel"}, n, n, n, "whois.nic.chanel", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[268] = Zone{"changiairport", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[269] = Zone{"channel", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[270] = Zone{"charity", r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/work-with-us/contact/", w{"v0n0.nic.charity", "v0n1.nic.charity", "v0n2.nic.charity", "v0n3.nic.charity", "v2n0.nic.charity", "v2n1.nic.charity"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.charity", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[271] = Zone{"chartis", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i271() {
	z[272] = Zone{"chase", r, x, 0x42, "JPMorgan Chase Bank, National Association", "https://www.nic.chase/", w{"a.nic.chase", "b.nic.chase", "c.nic.chase", "ns4.dns.nic.chase", "ns5.dns.nic.chase", "ns6.dns.nic.chase"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.chase", e, w{"https://rdap.nic.chase/"}, t}
	z[273] = Zone{"chat", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.chat", "v0n1.nic.chat", "v0n2.nic.chat", "v0n3.nic.chat", "v2n0.nic.chat", "v2n1.nic.chat"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.chat", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[274] = Zone{"cheap", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cheap", "v0n1.nic.cheap", "v0n2.nic.cheap", "v0n3.nic.cheap", "v2n0.nic.cheap", "v2n1.nic.cheap"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cheap", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[275] = Zone{"chesapeake", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[276] = Zone{"chevrolet", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[277] = Zone{"chevy", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[278] = Zone{"chintai", r, x, 0x42, "CHINTAI Corporation", "http://nic.chintai/", w{"a.nic.chintai", "b.nic.chintai", "c.nic.chintai", "ns1.dns.nic.chintai", "ns2.dns.nic.chintai", "ns3.dns.nic.chintai"}, n, n, w{"ja"}, "whois.nic.chintai", e, w{"https://rdap.nic.chintai/"}, t}
	z[279] = Zone{"chk", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i279() {
	z[280] = Zone{"chloe", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.chloe", e, n, f}
	z[281] = Zone{"christmas", r, x, 0x40, "XYZ.COM LLC", "https://nic.christmas/", w{"a.nic.christmas", "b.nic.christmas", "c.nic.christmas", "d.nic.christmas"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.christmas", e, w{"https://rdap.centralnic.com/christmas/"}, t}
	z[282] = Zone{"chrome", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[283] = Zone{"chrysler", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[284] = Zone{"church", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.church", "v0n1.nic.church", "v0n2.nic.church", "v0n3.nic.church", "v2n0.nic.church", "v2n1.nic.church"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.church", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[285] = Zone{"ci", r, z[2194:2211], 0xa0, e, e, w{"any.nic.ci", "ci.hosting.nic.fr", "ns-ci.afrinic.net", "ns.nic.ci", "phloem.uoregon.edu"}, n, n, n, "whois.nic.ci", e, n, f}
	z[286] = Zone{"cimb", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[287] = Zone{"cipriani", r, x, 0x42, "Hotel Cipriani Srl", "https://www.icann.org/en/registry-agreements/details/cipriani", w{"a0.nic.cipriani", "a2.nic.cipriani", "b0.nic.cipriani", "c0.nic.cipriani"}, n, n, n, "whois.nic.cipriani", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i287() {
	z[288] = Zone{"circle", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.circle/", w{"dns1.nic.circle", "dns2.nic.circle", "dns3.nic.circle", "dns4.nic.circle", "dnsa.nic.circle", "dnsb.nic.circle", "dnsc.nic.circle", "dnsd.nic.circle"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.circle", e, w{"https://rdap.nominet.uk/circle/"}, t}
	z[289] = Zone{"cisco", r, x, 0x42, "Cisco Technology, Inc.", "https://www.nic.cisco/", w{"a.nic.cisco", "b.nic.cisco", "c.nic.cisco", "ns1.dns.nic.cisco", "ns2.dns.nic.cisco", "ns3.dns.nic.cisco"}, n, n, n, "whois.nic.cisco", e, w{"https://rdap.nic.cisco/"}, f}
	z[290] = Zone{"citadel", r, x, 0x42, "Citadel Domain LLC", "http://www.nic.citadel/", w{"a0.nic.citadel", "a2.nic.citadel", "b0.nic.citadel", "c0.nic.citadel"}, n, n, n, "whois.nic.citadel", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[291] = Zone{"citi", r, x, 0x42, "Citigroup Inc.", "https://www.citi.com/", w{"a.nic.citi", "b.nic.citi", "c.nic.citi", "ns1.dns.nic.citi", "ns2.dns.nic.citi", "ns3.dns.nic.citi"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.citi", e, w{"https://rdap.nic.citi/"}, t}
	z[292] = Zone{"citic", r, x, 0x42, "CITIC Group Corporation", "http://www.nic.citic/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.citic", e, w{"https://rdap.zdnsgtld.com/citic/"}, t}
	z[293] = Zone{"city", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.city", "v0n1.nic.city", "v0n2.nic.city", "v0n3.nic.city", "v2n0.nic.city", "v2n1.nic.city"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.city", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[294] = Zone{"cityeats", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/cityeats", n, n, n, n, "whois.nic.cityeats", e, w{"https://tld-rdap.verisign.com/cityeats/v1/"}, f}
	z[295] = Zone{"ck", r, z[2211:2219], 0xa8, "Telecom Cook Islands", "https://www.vodafone.co.ck/domain-admin-policy", w{"circa.mcs.vuw.ac.nz", "downstage.mcs.vuw.ac.nz", "parau.oyster.net.ck", "poiparau.oyster.net.ck"}, n, n, n, "whois.ck-nic.org.ck", e, n, f}
}

//go:noinline
func i295() {
	z[296] = Zone{"cl", r, x, 0xa0, e, e, w{"a.nic.cl", "b.nic.cl", "c.nic.cl", "cl-ns.anycast.pch.net", "cl1-tld.d-zone.ca", "cl1.dnsnode.net", "cl2-tld.d-zone.ca"}, n, n, w{"mul-Latn"}, "whois.nic.cl", e, n, t}
	z[297] = Zone{"claims", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.claims", "v0n1.nic.claims", "v0n2.nic.claims", "v0n3.nic.claims", "v2n0.nic.claims", "v2n1.nic.claims"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.claims", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[298] = Zone{"cleaning", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cleaning", "v0n1.nic.cleaning", "v0n2.nic.cleaning", "v0n3.nic.cleaning", "v2n0.nic.cleaning", "v2n1.nic.cleaning"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cleaning", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[299] = Zone{"click", r, x, 0x40, "Internet Naming Company LLC", "https://uniregistry.link/extensions/click/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.click", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[300] = Zone{"clinic", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.clinic", "v0n1.nic.clinic", "v0n2.nic.clinic", "v0n3.nic.clinic", "v2n0.nic.clinic", "v2n1.nic.clinic"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.clinic", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[301] = Zone{"clinique", r, x, 0x42, "The Estée Lauder Companies Inc.", "https://www.icann.org/en/registry-agreements/details/clinique", w{"a0.nic.clinique", "a2.nic.clinique", "b0.nic.clinique", "c0.nic.clinique"}, n, n, n, "whois.nic.clinique", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[302] = Zone{"clothing", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.clothing", "v0n1.nic.clothing", "v0n2.nic.clothing", "v0n3.nic.clothing", "v2n0.nic.clothing", "v2n1.nic.clothing"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.clothing", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[303] = Zone{"cloud", r, x, 0x40, "Aruba PEC S.p.A.", "https://www.get.cloud/home.aspx", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, "whois.nic.cloud", e, w{"https://rdap.registry.cloud/rdap/"}, f}
}

//go:noinline
func i303() {
	z[304] = Zone{"club", r, x, 0x40, "Registry Services, LLC", e, w{"a.nic.club", "b.nic.club", "c.nic.club", "ns1.dns.nic.club", "ns2.dns.nic.club", "ns3.dns.nic.club"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.club", e, w{"https://rdap.nic.club/"}, t}
	z[305] = Zone{"clubmed", r, x, 0x42, "Club Méditerranée S.A.", "https://ns.clubmed.com/ipm/nic_clubmed/index.html", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"ja", "ko", "mul-Deva", "mul-Hang", "mul-Hani", "mul-Hano"}, "whois.nic.clubmed", e, w{"https://tld-rdap.verisign.com/clubmed/v1/"}, t}
	z[306] = Zone{"cm", r, z[2219:2223], 0xa0, e, "https://antic.cm/", w{"auth02.ns.uu.net", "ns-cm.afrinic.net", "ns-cm.nic.fr", "ns.itu.ch", "ns1.nic.cm", "ns2.nic.cm", "phloem.uoregon.edu"}, n, n, n, "whois.netcom.cm", e, n, f}
	z[307] = Zone{"cn", r, z[2223:2265], 0xa0, e, "https://www.cnnic.net.cn/", w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "d.dns.cn", "e.dns.cn", "ns.cernet.net"}, n, n, w{"zh", "zh-Hans"}, "whois.cnnic.cn", e, n, t}
	z[308] = Zone{"co", r, z[2265:2272], 0xe0, e, "https://www.go.co/", w{"ns1.cctld.co", "ns2.cctld.co", "ns3.cctld.co", "ns4.cctld.co", "ns5.cctld.co", "ns6.cctld.co"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "mul-Latn", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.co", e, n, t}
	z[309] = Zone{"coach", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.coach", "v0n1.nic.coach", "v0n2.nic.coach", "v0n3.nic.coach", "v2n0.nic.coach", "v2n1.nic.coach"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.coach", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[310] = Zone{"codes", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.codes", "v0n1.nic.codes", "v0n2.nic.codes", "v0n3.nic.codes", "v2n0.nic.codes", "v2n1.nic.codes"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.codes", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[311] = Zone{"coffee", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.coffee", "v0n1.nic.coffee", "v0n2.nic.coffee", "v0n3.nic.coffee", "v2n0.nic.coffee", "v2n1.nic.coffee"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.coffee", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i311() {
	z[312] = Zone{"college", r, x, 0x40, "XYZ.COM LLC", "https://nic.college/", w{"a.nic.college", "b.nic.college", "c.nic.college", "d.nic.college"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.college", e, w{"https://rdap.centralnic.com/college/"}, t}
	z[313] = Zone{"cologne", r, x, 0xc4, "dotKoeln GmbH", e, w{"dns.ryce-rsp.com", "ns1.dns.business", "ns1.ryce-rsp.com"}, n, w{"Cologne", "Koeln"}, w{"de", "mul-Latn"}, "whois.ryce-rsp.com", e, w{"https://rdap.ryce-rsp.com/rdap/"}, t}
	z[314] = Zone{"com", r, z[2272:2304], 0x40, "VeriSign, Inc.", "https://yourdot.com/", w{"a.gtld-servers.net", "b.gtld-servers.net", "c.gtld-servers.net", "d.gtld-servers.net", "e.gtld-servers.net", "f.gtld-servers.net", "g.gtld-servers.net", "h.gtld-servers.net", "i.gtld-servers.net", "j.gtld-servers.net", "k.gtld-servers.net", "l.gtld-servers.net", "m.gtld-servers.net"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro-MD", "ru", "uk", "zh"}, "whois.verisign-grs.com", e, w{"https://rdap.verisign.com/com/v1/"}, t}
	z[315] = Zone{"comcast", r, x, 0x42, "Comcast IP Holdings I, LLC", "https://nic.comcast/", w{"dns1.nic.comcast", "dns2.nic.comcast", "dns3.nic.comcast", "dns4.nic.comcast", "dnsa.nic.comcast", "dnsb.nic.comcast", "dnsc.nic.comcast", "dnsd.nic.comcast"}, n, n, n, "whois.nic.comcast", e, w{"https://rdap.nominet.uk/comcast/"}, f}
	z[316] = Zone{"commbank", r, x, 0x42, "COMMONWEALTH BANK OF AUSTRALIA", "http://nic.commbank/", w{"a.nic.commbank", "b.nic.commbank", "c.nic.commbank", "x.nic.commbank", "y.nic.commbank", "z.nic.commbank"}, n, n, n, "whois.nic.commbank", e, w{"https://rdap.nic.commbank/"}, f}
	z[317] = Zone{"community", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.community", "v0n1.nic.community", "v0n2.nic.community", "v0n3.nic.community", "v2n0.nic.community", "v2n1.nic.community"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.community", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[318] = Zone{"company", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.company", "v0n1.nic.company", "v0n2.nic.company", "v0n3.nic.company", "v2n0.nic.company", "v2n1.nic.company"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.company", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[319] = Zone{"compare", r, x, 0x40, "Registry Services, LLC", "https://www.gocompare.com/", w{"a.nic.compare", "b.nic.compare", "c.nic.compare", "x.nic.compare", "y.nic.compare", "z.nic.compare"}, n, n, n, "whois.nic.compare", e, w{"https://rdap.nic.compare/"}, f}
}

//go:noinline
func i319() {
	z[320] = Zone{"computer", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.computer", "v0n1.nic.computer", "v0n2.nic.computer", "v0n3.nic.computer", "v2n0.nic.computer", "v2n1.nic.computer"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.computer", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[321] = Zone{"comsec", r, x, 0x42, "VeriSign, Inc.", "https://www.icann.org/en/registry-agreements/details/comsec", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.comsec", e, w{"https://tld-rdap.verisign.com/comsec/v1/"}, t}
	z[322] = Zone{"condos", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.condos", "v0n1.nic.condos", "v0n2.nic.condos", "v0n3.nic.condos", "v2n0.nic.condos", "v2n1.nic.condos"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.condos", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[323] = Zone{"connectors", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[324] = Zone{"construction", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.construction", "v0n1.nic.construction", "v0n2.nic.construction", "v0n3.nic.construction", "v2n0.nic.construction", "v2n1.nic.construction"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.construction", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[325] = Zone{"consulting", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.consulting", "v0n1.nic.consulting", "v0n2.nic.consulting", "v0n3.nic.consulting", "v2n0.nic.consulting", "v2n1.nic.consulting"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.consulting", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[326] = Zone{"contact", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.contact", "v0n1.nic.contact", "v0n2.nic.contact", "v0n3.nic.contact", "v2n0.nic.contact", "v2n1.nic.contact"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Cyrl", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.contact", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[327] = Zone{"contractors", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.contractors", "v0n1.nic.contractors", "v0n2.nic.contractors", "v0n3.nic.contractors", "v2n0.nic.contractors", "v2n1.nic.contractors"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.technology", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i327() {
	z[328] = Zone{"cooking", r, x, 0x40, "Registry Services, LLC", "http://nic.cooking/", w{"a.nic.cooking", "b.nic.cooking", "c.nic.cooking", "x.nic.cooking", "y.nic.cooking", "z.nic.cooking"}, n, n, w{"de", "es", "fr"}, "whois.nic.cooking", e, w{"https://rdap.nic.cooking/"}, t}
	z[329] = Zone{"cookingchannel", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/cookingchannel", n, n, n, n, "whois.nic.cookingchannel", e, w{"https://tld-rdap.verisign.com/cookingchannel/v1/"}, f}
	z[330] = Zone{"cool", r, z[2304:2305], 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cool", "v0n1.nic.cool", "v0n2.nic.cool", "v0n3.nic.cool", "v2n0.nic.cool", "v2n1.nic.cool"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cool", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[331] = Zone{"coop", r, x, 0x4040, "DotCooperation LLC", "https://identity.coop/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"ar", "es", "fr", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "mul-Mymr", "pl", "ru", "sv", "th", "zh"}, "whois.nic.coop", e, w{"https://rdap.registry.coop/rdap/"}, t}
	z[332] = Zone{"corp", r, x, 0x800, "ICANN", "https://features.icann.org/addressing-new-gtld-program-applications-corp-home-and-mail", n, n, n, n, e, e, n, t}
	z[333] = Zone{"corsica", r, x, 0x4c0, "Collectivité de Corse", "https://www.icann.org/en/registry-agreements/details/corsica", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, w{"FR-H"}, w{"mul-Latn"}, "whois.nic.corsica", e, w{"https://rdap.nic.corsica/"}, t}
	z[334] = Zone{"country", r, x, 0x40, "Internet Naming Company LLC", "https://uniregistry.link/extensions/country/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "mul-Cyrl", "pt"}, "whois.nic.country", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[335] = Zone{"coupon", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.coupon/", w{"a.nic.coupon", "b.nic.coupon", "c.nic.coupon", "ns1.dns.nic.coupon", "ns2.dns.nic.coupon", "ns3.dns.nic.coupon"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.coupon", e, w{"https://rdap.nic.coupon/"}, t}
}

//go:noinline
func i335() {
	z[336] = Zone{"coupons", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.coupons", "v0n1.nic.coupons", "v0n2.nic.coupons", "v0n3.nic.coupons", "v2n0.nic.coupons", "v2n1.nic.coupons"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.coupons", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[337] = Zone{"courses", r, x, 0x40, "Registry Services, LLC", "https://www.get.courses/", w{"a.nic.courses", "b.nic.courses", "c.nic.courses", "x.nic.courses", "y.nic.courses", "z.nic.courses"}, n, n, n, "whois.nic.courses", e, w{"https://rdap.nic.courses/"}, f}
	z[338] = Zone{"cpa", r, x, 0x40, "American Institute of Certified Public Accountants", "http://nic.cpa/", w{"a.nic.cpa", "b.nic.cpa", "c.nic.cpa", "x.nic.cpa", "y.nic.cpa", "z.nic.cpa"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.cpa", e, w{"http://rdap.nic.cpa/"}, t}
	z[339] = Zone{"cr", r, z[2305:2313], 0xa0, e, "https://www.nic.cr/", w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, "whois.nic.cr", e, w{"https://rdap.nic.cr/"}, f}
	z[340] = Zone{"credit", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.credit", "v0n1.nic.credit", "v0n2.nic.credit", "v0n3.nic.credit", "v2n0.nic.credit", "v2n1.nic.credit"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.credit", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[341] = Zone{"creditcard", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.creditcard", "v0n1.nic.creditcard", "v0n2.nic.creditcard", "v0n3.nic.creditcard", "v2n0.nic.creditcard", "v2n1.nic.creditcard"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.creditcard", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[342] = Zone{"creditunion", r, x, 0x40, "DotCooperation LLC", "https://identity.coop/creditunion/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns5.uniregistry.net", "ns6.uniregistry.info"}, n, n, n, "whois.nic.creditunion", e, w{"https://rdap.registry.coop/rdap/"}, t}
	z[343] = Zone{"cricket", r, x, 0x40, "dot Cricket Limited", "http://nic.cricket/", w{"a.nic.cricket", "b.nic.cricket", "c.nic.cricket", "ns1.dns.nic.cricket", "ns2.dns.nic.cricket", "ns3.dns.nic.cricket"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.cricket", e, w{"https://rdap.nic.cricket/"}, t}
}

//go:noinline
func i343() {
	z[344] = Zone{"crown", r, x, 0x42, "Crown Equipment Corporation", "http://nic.crown/", w{"a.ns.nic.crown", "b.ns.nic.crown"}, n, n, n, "whois.nic.crown", e, w{"https://rdap.crown.fury.ca/rdap/"}, f}
	z[345] = Zone{"crs", r, x, 0x42, "Federated Co-operatives Limited", "http://nic.crs/", w{"a0.nic.crs", "a2.nic.crs", "b0.nic.crs", "c0.nic.crs"}, n, n, n, "whois.nic.crs", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[346] = Zone{"cruise", r, x, 0x40, "Viking River Cruises (Bermuda) Ltd.", "http://nic.cruise/", w{"a0.nic.cruise", "a2.nic.cruise", "b0.nic.cruise", "c0.nic.cruise"}, n, n, n, "whois.nic.cruise", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[347] = Zone{"cruises", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.cruises", "v0n1.nic.cruises", "v0n2.nic.cruises", "v0n3.nic.cruises", "v2n0.nic.cruises", "v2n1.nic.cruises"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.cruises", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[348] = Zone{"csc", r, x, 0x1042, "Alliance-One Services, Inc.", e, n, n, n, n, "whois.nic.csc", e, w{"https://tld-rdap.verisign.com/csc/v1/"}, f}
	z[349] = Zone{"cu", r, z[2313:2324], 0xa0, e, e, w{"cu.cctld.authdns.ripe.net", "ns.ceniai.net.cu", "ns.dns.br", "ns2.ceniai.net.cu", "ns2.gip.net", "rip.psg.com"}, n, n, n, e, "http://www.nic.cu/", n, f}
	z[350] = Zone{"cuisinella", r, x, 0x42, "SCHMIDT GROUPE S.A.S.", "http://nic.cuisinella/", w{"a.nic.cuisinella", "b.nic.cuisinella", "c.nic.cuisinella", "x.nic.cuisinella", "y.nic.cuisinella", "z.nic.cuisinella"}, n, n, n, "whois.nic.cuisinella", e, w{"https://rdap.nic.cuisinella/"}, f}
	z[351] = Zone{"cv", r, z[2324:2332], 0xa0, e, e, w{"anyc.dnsnode.net", "c.dns.pt", "cv01.dns.pt", "ns.dns.cv"}, n, n, n, e, "http://www.dns.cv/", n, f}
}

//go:noinline
func i351() {
	z[352] = Zone{"cw", r, z[2332:2334], 0xa0, e, "https://en.wikipedia.org/wiki/.cw", w{"cw.cctld.authdns.ripe.net", "kadushi.curinfo.cw", "ns0.ja.net", "ns01-server.curinfo.cw", "ns1.dns.cw"}, n, n, n, e, e, n, f}
	z[353] = Zone{"cx", r, z[2334:2339], 0xa0, e, e, w{"ns.anycast.nic.cx", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.cx", e, n, f}
	z[354] = Zone{"cy", r, z[2339:2352], 0xa8, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, "http://www.nic.cy/nslookup/online_database.php", n, f}
	z[355] = Zone{"cymru", r, x, 0x440, "Nominet UK", "https://eincartrefarlein.cymru/", w{"dns1.nic.cymru", "dns2.nic.cymru", "dns3.nic.cymru", "dns4.nic.cymru", "dnsa.nic.cymru", "dnsb.nic.cymru", "dnsc.nic.cymru", "dnsd.nic.cymru"}, n, w{"GB-WLS"}, w{"cy"}, "whois.nic.cymru", e, w{"https://rdap.nominet.uk/cymru/"}, t}
	z[356] = Zone{"cyou", r, x, 0x40, "ShortDot SA", "https://shortdot.bond/cyou/", w{"a.nic.cyou", "b.nic.cyou", "c.nic.cyou", "d.nic.cyou"}, n, n, n, "whois.nic.cyou", e, w{"https://rdap.centralnic.com/cyou/"}, f}
	z[357] = Zone{"cz", r, z[2352:2353], 0xa0, e, "https://www.nic.cz/", w{"a.ns.nic.cz", "b.ns.nic.cz", "c.ns.nic.cz", "d.ns.nic.cz"}, n, n, n, "whois.nic.cz", e, w{"https://rdap.nic.cz/"}, t}
	z[358] = Zone{"dabur", r, x, 0x42, "Dabur India Limited", "http://nic.dabur/", w{"a0.nic.dabur", "a2.nic.dabur", "b0.nic.dabur", "c0.nic.dabur"}, n, n, n, "whois.nic.dabur", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[359] = Zone{"dad", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
}

//go:noinline
func i359() {
	z[360] = Zone{"dance", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.dance", "v0n1.nic.dance", "v0n2.nic.dance", "v0n3.nic.dance", "v2n0.nic.dance", "v2n1.nic.dance"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.dance", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[361] = Zone{"data", r, x, 0x42, "Dish DBS Corporation", "https://www.dishtlds.com/data/", w{"a0.nic.data", "a2.nic.data", "b0.nic.data", "c0.nic.data"}, n, n, n, "whois.nic.data", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[362] = Zone{"date", r, x, 0x40, "dot Date Limited", "http://nic.date/", w{"a.nic.date", "b.nic.date", "c.nic.date", "ns1.dns.nic.date", "ns2.dns.nic.date", "ns3.dns.nic.date"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.date", e, w{"https://rdap.nic.date/"}, t}
	z[363] = Zone{"dating", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.dating", "v0n1.nic.dating", "v0n2.nic.dating", "v0n3.nic.dating", "v2n0.nic.dating", "v2n1.nic.dating"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.dating", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[364] = Zone{"datsun", r, x, 0x42, "NISSAN MOTOR CO., LTD.", "https://www.nissan-global.com/EN/NIC/DATSUN/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[365] = Zone{"day", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[366] = Zone{"dclk", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[367] = Zone{"dds", r, x, 0x40, "Registry Services, LLC", "http://nic.dds/", w{"a.nic.dds", "b.nic.dds", "c.nic.dds", "x.nic.dds", "y.nic.dds", "z.nic.dds"}, n, n, w{"de", "es", "fr"}, "whois.nic.dds", e, w{"https://rdap.nic.dds/"}, t}
}

//go:noinline
func i367() {
	z[368] = Zone{"de", r, z[2353:2356], 0xa0, e, "https://www.denic.de/", w{"a.nic.de", "f.nic.de", "l.de.net", "n.de.net", "s.de.net", "z.nic.de"}, n, n, n, "whois.denic.de", e, w{"https://rdap.denic.de/"}, t}
	z[369] = Zone{"deal", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.deal/", w{"dns1.nic.deal", "dns2.nic.deal", "dns3.nic.deal", "dns4.nic.deal", "dnsa.nic.deal", "dnsb.nic.deal", "dnsc.nic.deal", "dnsd.nic.deal"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.deal", e, w{"https://rdap.nominet.uk/deal/"}, t}
	z[370] = Zone{"dealer", r, x, 0x40, "Intercap Registry Inc.", "http://nic.dealer/", w{"a.nic.dealer", "b.nic.dealer", "c.nic.dealer", "d.nic.dealer"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.dealer", e, w{"https://rdap.centralnic.com/dealer/"}, t}
	z[371] = Zone{"deals", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.deals", "v0n1.nic.deals", "v0n2.nic.deals", "v0n3.nic.deals", "v2n0.nic.deals", "v2n1.nic.deals"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.deals", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[372] = Zone{"degree", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.degree", "v0n1.nic.degree", "v0n2.nic.degree", "v0n3.nic.degree", "v2n0.nic.degree", "v2n1.nic.degree"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.degree", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[373] = Zone{"delivery", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.delivery", "v0n1.nic.delivery", "v0n2.nic.delivery", "v0n3.nic.delivery", "v2n0.nic.delivery", "v2n1.nic.delivery"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.delivery", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[374] = Zone{"dell", r, x, 0x42, "Dell Inc.", "https://nic.dell/", w{"a.nic.dell", "b.nic.dell", "c.nic.dell", "ns1.dns.nic.dell", "ns2.dns.nic.dell", "ns3.dns.nic.dell"}, n, n, w{"es"}, "whois.nic.dell", e, w{"https://rdap.nic.dell/"}, t}
	z[375] = Zone{"delmonte", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i375() {
	z[376] = Zone{"deloitte", r, x, 0x42, "Deloitte Touche Tohmatsu", "https://www.nic.deloitte/", w{"a.nic.deloitte", "b.nic.deloitte", "c.nic.deloitte", "d.nic.deloitte"}, n, n, w{"ar", "de", "es", "fr", "ja", "nl", "pt", "ru", "zh"}, "whois.nic.deloitte", e, w{"https://rdap.centralnic.com/deloitte/"}, t}
	z[377] = Zone{"delta", r, x, 0x42, "Delta Air Lines, Inc.", "http://www.nic.delta/", w{"a0.nic.delta", "a2.nic.delta", "b0.nic.delta", "c0.nic.delta"}, n, n, n, "whois.nic.delta", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[378] = Zone{"democrat", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.democrat", "v0n1.nic.democrat", "v0n2.nic.democrat", "v0n3.nic.democrat", "v2n0.nic.democrat", "v2n1.nic.democrat"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.democrat", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[379] = Zone{"dental", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.dental", "v0n1.nic.dental", "v0n2.nic.dental", "v0n3.nic.dental", "v2n0.nic.dental", "v2n1.nic.dental"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.dental", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[380] = Zone{"dentist", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.dentist", "v0n1.nic.dentist", "v0n2.nic.dentist", "v0n3.nic.dentist", "v2n0.nic.dentist", "v2n1.nic.dentist"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.dentist", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[381] = Zone{"desi", r, x, 0x40, "Desi Networks LLC", e, w{"dns1.emdns.uk", "dns2.emdns.uk", "dns3.emdns.uk", "dns4.emdns.uk", "dnsa.emdns.uk", "dnsb.emdns.uk", "dnsc.emdns.uk", "dnsd.emdns.uk"}, n, n, n, "whois.nic.desi", e, w{"https://rdap.nominet.uk/desi/"}, t}
	z[382] = Zone{"design", r, x, 0x40, "Registry Services, LLC", "https://your.design/", w{"a.nic.design", "b.nic.design", "c.nic.design", "x.nic.design", "y.nic.design", "z.nic.design"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.design", e, w{"https://rdap.nic.design/"}, t}
	z[383] = Zone{"deutschepost", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i383() {
	z[384] = Zone{"dev", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"ja", "mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[385] = Zone{"dhl", r, x, 0x42, "Deutsche Post AG", "https://www.nic.dhl/", w{"a.nic.dhl", "b.nic.dhl", "c.nic.dhl", "d.nic.dhl"}, n, n, w{"mul-Latn"}, e, e, w{"https://rdap.centralnic.com/dhl/"}, t}
	z[386] = Zone{"diamonds", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.diamonds", "v0n1.nic.diamonds", "v0n2.nic.diamonds", "v0n3.nic.diamonds", "v2n0.nic.diamonds", "v2n1.nic.diamonds"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.diamonds", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[387] = Zone{"diet", r, x, 0x40, "XYZ.COM LLC", "https://nic.diet/", w{"a.nic.diet", "b.nic.diet", "c.nic.diet", "d.nic.diet"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.diet", e, w{"https://rdap.centralnic.com/diet/"}, t}
	z[388] = Zone{"digikey", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[389] = Zone{"digital", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.digital", "v0n1.nic.digital", "v0n2.nic.digital", "v0n3.nic.digital", "v2n0.nic.digital", "v2n1.nic.digital"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.digital", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[390] = Zone{"direct", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.direct", "v0n1.nic.direct", "v0n2.nic.direct", "v0n3.nic.direct", "v2n0.nic.direct", "v2n1.nic.direct"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.direct", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[391] = Zone{"directory", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.directory", "v0n1.nic.directory", "v0n2.nic.directory", "v0n3.nic.directory", "v2n0.nic.directory", "v2n1.nic.directory"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.directory", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i391() {
	z[392] = Zone{"discount", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.discount", "v0n1.nic.discount", "v0n2.nic.discount", "v0n3.nic.discount", "v2n0.nic.discount", "v2n1.nic.discount"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.discount", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[393] = Zone{"discover", r, x, 0x42, "Discover Financial Services", "https://www.nic.discover//", w{"a.nic.discover", "b.nic.discover", "c.nic.discover", "ns1.dns.nic.discover", "ns2.dns.nic.discover", "ns3.dns.nic.discover"}, n, n, n, "whois.nic.discover", e, w{"https://rdap.nic.discover/"}, f}
	z[394] = Zone{"dish", r, x, 0x42, "Dish DBS Corporation", "https://www.dishtlds.com/dish/", w{"a0.nic.dish", "a2.nic.dish", "b0.nic.dish", "c0.nic.dish"}, n, n, n, "whois.nic.dish", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[395] = Zone{"diy", r, x, 0x42, "Internet Naming Company LLC", "https://www.icann.org/en/registry-agreements/details/diy", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.diy", e, w{"https://tld-rdap.verisign.com/diy/v1/"}, f}
	z[396] = Zone{"dj", r, x, 0xe0, e, e, w{"bow1.intnet.dj", "bow5.intnet.dj", "vps443605.ovh.net"}, n, n, n, e, "http://www.nic.dj/whois.php", n, f}
	z[397] = Zone{"dk", r, z[2356:2358], 0xa0, e, "https://www.dk-hostmaster.dk/da", w{"b.nic.dk", "c.nic.dk", "d.nic.dk", "h.nic.dk", "l.nic.dk", "s.nic.dk"}, n, n, n, "whois.punktum.dk", e, n, t}
	z[398] = Zone{"dm", r, z[2358:2364], 0xa0, e, e, w{"ns.blacknightsolutions.com", "ns1.uniregistry.net", "ns2.blacknightsolutions.com", "ns2.nic.dm", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns34.cdns.net", "ns4.uniregistry.info"}, n, n, n, "whois.dmdomains.dm", e, n, t}
	z[399] = Zone{"dnb", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i399() {
	z[400] = Zone{"dnp", r, x, 0x42, "Dai Nippon Printing Co., Ltd.", "https://nic.dnp/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.dnp", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[401] = Zone{"do", r, z[2364:2375], 0xa0, e, "https://www.nic.do/", w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, "whois.nic.do", e, n, f}
	z[402] = Zone{"docomo", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[403] = Zone{"docs", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[404] = Zone{"doctor", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.doctor", "v0n1.nic.doctor", "v0n2.nic.doctor", "v0n3.nic.doctor", "v2n0.nic.doctor", "v2n1.nic.doctor"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.doctor", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[405] = Zone{"dodge", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[406] = Zone{"dog", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.dog", "v0n1.nic.dog", "v0n2.nic.dog", "v0n3.nic.dog", "v2n0.nic.dog", "v2n1.nic.dog"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.dog", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[407] = Zone{"doha", r, x, 0x10c4, e, e, n, n, w{"Doha", "QA-DA"}, n, "whois.nic.doha", e, n, f}
}

//go:noinline
func i407() {
	z[408] = Zone{"domains", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.domains", "v0n1.nic.domains", "v0n2.nic.domains", "v0n3.nic.domains", "v2n0.nic.domains", "v2n1.nic.domains"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.domains", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[409] = Zone{"doosan", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.xn--cg4bki", e, n, f}
	z[410] = Zone{"dot", r, x, 0x40, "Dish DBS Corporation", "https://www.dishtlds.com/dot/", w{"a0.nic.dot", "a2.nic.dot", "b0.nic.dot", "c0.nic.dot"}, n, n, n, "whois.nic.dot", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[411] = Zone{"dotafrica", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[412] = Zone{"download", r, x, 0x40, "dot Support Limited", "http://nic.download/", w{"a.nic.download", "b.nic.download", "c.nic.download", "ns1.dns.nic.download", "ns2.dns.nic.download", "ns3.dns.nic.download"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.download", e, w{"https://rdap.nic.download/"}, t}
	z[413] = Zone{"drive", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[414] = Zone{"dstv", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[415] = Zone{"dtv", r, x, 0x42, "Dish DBS Corporation", "https://dishtlds.com/dtv/", w{"a0.nic.dtv", "a2.nic.dtv", "b0.nic.dtv", "c0.nic.dtv"}, n, n, n, "whois.nic.dtv", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i415() {
	z[416] = Zone{"dubai", r, x, 0xc4, "Dubai Smart Government Department", "http://nic.dubai/", w{"gtld.alpha.aridns.net.au", "gtld.beta.aridns.net.au", "gtld.delta.aridns.net.au", "gtld.gamma.aridns.net.au"}, n, w{"AE-DU", "Dubai"}, w{"ar"}, "whois.nic.dubai", e, w{"https://rdap.nic.dubai/"}, t}
	z[417] = Zone{"duck", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.duck", e, w{"https://tld-rdap.verisign.com/duck/v1/"}, f}
	z[418] = Zone{"dunlop", r, x, 0x42, "The Goodyear Tire & Rubber Company", "https://nic.dunlop/", w{"a0.nic.dunlop", "a2.nic.dunlop", "b0.nic.dunlop", "c0.nic.dunlop"}, n, n, n, "whois.nic.dunlop", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[419] = Zone{"duns", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.duns", e, n, f}
	z[420] = Zone{"dupont", r, x, 0x42, "DuPont Specialty Products USA, LLC", "https://www.dupont.com/nic.html", w{"a.nic.dupont", "b.nic.dupont", "c.nic.dupont", "ns4.dns.nic.dupont", "ns5.dns.nic.dupont", "ns6.dns.nic.dupont"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.dupont", e, w{"https://rdap.nic.dupont/"}, t}
	z[421] = Zone{"durban", r, x, 0xc4, "ZA Central Registry NPC trading as ZA Central Registry", e, w{"coza1.dnsnode.net", "nsp.netnod.se"}, n, w{"Durban"}, n, "whois.nic.durban", e, w{"https://rdap.nic.durban/rdap/"}, t}
	z[422] = Zone{"dvag", r, x, 0x42, "Deutsche Vermögensberatung Aktiengesellschaft DVAG", "https://nic.dvag/", w{"a.nic.dvag", "b.nic.dvag", "c.nic.dvag", "d.nic.dvag"}, n, n, w{"de"}, "whois.nic.dvag", e, w{"https://rdap.centralnic.com/dvag/"}, t}
	z[423] = Zone{"dvr", r, x, 0x42, "DISH Technologies L.L.C.", "https://www.icann.org/en/registry-agreements/details/dvr", w{"a0.nic.dvr", "a2.nic.dvr", "b0.nic.dvr", "c0.nic.dvr"}, n, n, n, "whois.nic.dvr", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i423() {
	z[424] = Zone{"dwg", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[425] = Zone{"dz", r, z[2375:2383], 0xa0, e, e, w{"ns-dz.afrinic.net", "ns1.nic.dz", "ns2.nic.dz", "ns3.nic.fr", "ns4.nic.dz", "ns5.nic.dz"}, n, n, n, "whois.nic.dz", e, n, f}
	z[426] = Zone{"earth", r, x, 0x40, "Interlink Systems Innovation Institute K.K.", "https://domain.earth/", w{"a.nic.earth", "b.nic.earth", "c.nic.earth", "ns1.dns.nic.earth", "ns2.dns.nic.earth", "ns3.dns.nic.earth"}, n, n, w{"ja", "ko", "zh"}, "whois.nic.earth", e, w{"https://rdap.nic.earth/"}, t}
	z[427] = Zone{"eat", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[428] = Zone{"ec", r, z[2383:2394], 0xa0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.ec", e, n, f}
	z[429] = Zone{"eco", r, x, 0x40, "Big Room Inc.", "https://bigroom.eco/", w{"a.ns.nic.eco", "b.ns.nic.eco"}, n, n, n, "whois.nic.eco", e, w{"https://rdap.eco.fury.ca/rdap/"}, f}
	z[430] = Zone{"ecom", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[431] = Zone{"edeka", r, x, 0x42, "EDEKA Verband kaufmännischer Genossenschaften e.V.", "https://www.nic.edeka/", w{"a0.nic.edeka", "a2.nic.edeka", "b0.nic.edeka", "c0.nic.edeka"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.edeka", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i431() {
	z[432] = Zone{"edu", r, x, 0x4040, e, e, w{"a.edu-servers.net", "b.edu-servers.net", "c.edu-servers.net", "d.edu-servers.net", "e.edu-servers.net", "f.edu-servers.net", "g.edu-servers.net", "h.edu-servers.net", "i.edu-servers.net", "j.edu-servers.net", "k.edu-servers.net", "l.edu-servers.net", "m.edu-servers.net"}, n, n, n, "whois.educause.edu", e, n, f}
	z[433] = Zone{"education", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.education", "v0n1.nic.education", "v0n2.nic.education", "v0n3.nic.education", "v2n0.nic.education", "v2n1.nic.education"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.education", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[434] = Zone{"ee", r, z[2394:2408], 0xa0, e, e, w{"b.tld.ee", "e.tld.ee", "ee.cert.ee", "ee.eenet.ee", "ns.tld.ee"}, n, n, n, "whois.tld.ee", e, n, f}
	z[435] = Zone{"eg", r, z[2408:2420], 0xa0, e, e, w{"frcu.eun.eg", "ns5.univie.ac.at", "rip.psg.com"}, n, n, n, e, "http://lookup.egregistry.eg/english.aspx", n, f}
	z[436] = Zone{"email", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.email", "v0n1.nic.email", "v0n2.nic.email", "v0n3.nic.email", "v2n0.nic.email", "v2n1.nic.email"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.email", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[437] = Zone{"emerck", r, x, 0x42, "Merck KGaA", "http://www.nic.emerck/darmstadt/germany/", w{"a0.nic.emerck", "a2.nic.emerck", "b0.nic.emerck", "c0.nic.emerck"}, n, n, w{"da", "de", "es", "hu", "is", "lt", "lv", "pl", "ru", "sv"}, "whois.nic.emerck", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[438] = Zone{"emerson", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[439] = Zone{"energy", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.energy", "v0n1.nic.energy", "v0n2.nic.energy", "v0n3.nic.energy", "v2n0.nic.energy", "v2n1.nic.energy"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.energy", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i439() {
	z[440] = Zone{"engineer", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.engineer", "v0n1.nic.engineer", "v0n2.nic.engineer", "v0n3.nic.engineer", "v2n0.nic.engineer", "v2n1.nic.engineer"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.engineer", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[441] = Zone{"engineering", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.engineering", "v0n1.nic.engineering", "v0n2.nic.engineering", "v0n3.nic.engineering", "v2n0.nic.engineering", "v2n1.nic.engineering"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.engineering", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[442] = Zone{"enterprises", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.enterprises", "v0n1.nic.enterprises", "v0n2.nic.enterprises", "v0n3.nic.enterprises", "v2n0.nic.enterprises", "v2n1.nic.enterprises"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.enterprises", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[443] = Zone{"epost", r, x, 0x1042, e, e, n, n, n, w{"mul-Latn"}, e, e, n, t}
	z[444] = Zone{"epson", r, x, 0x42, "Seiko Epson Corporation", "https://nic.epson/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.epson", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[445] = Zone{"equipment", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.equipment", "v0n1.nic.equipment", "v0n2.nic.equipment", "v0n3.nic.equipment", "v2n0.nic.equipment", "v2n1.nic.equipment"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.equipment", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[446] = Zone{"er", r, z[2420:2427], 0xa8, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[447] = Zone{"ericsson", r, x, 0x42, "Telefonaktiebolaget L M Ericsson", "https://nic.ericsson/", w{"a0.nic.ericsson", "a2.nic.ericsson", "b0.nic.ericsson", "c0.nic.ericsson"}, n, n, n, "whois.nic.ericsson", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i447() {
	z[448] = Zone{"erni", r, x, 0x42, "ERNI Group Holding AG", "https://www.nic.erni/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.erni", e, w{"https://rdap.nic.erni/"}, t}
	z[449] = Zone{"es", r, z[2427:2432], 0xa0, e, e, w{"a.nic.es", "c.nic.es", "g.nic.es", "h.nic.es"}, n, n, n, "whois.nic.es", e, n, t}
	z[450] = Zone{"esq", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[451] = Zone{"estate", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.estate", "v0n1.nic.estate", "v0n2.nic.estate", "v0n3.nic.estate", "v2n0.nic.estate", "v2n1.nic.estate"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.estate", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[452] = Zone{"esurance", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[453] = Zone{"et", r, z[2432:2440], 0xa0, e, "https://www.iana.org/domains/root/db/et.html", w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[454] = Zone{"etisalat", r, x, 0x1042, "Emirates Telecommunications Corporation (trading as Etisalat)", "https://nic.etisalat/", n, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/etisalat/"}, f}
	z[455] = Zone{"eu", r, z[2440:2443], 0xa0, e, e, w{"be.dns.eu", "si.dns.eu", "w.dns.eu", "x.dns.eu", "y.dns.eu"}, n, n, w{"mul-Cyrl", "mul-Grek", "mul-Latn"}, "whois.eu", e, n, t}
}

//go:noinline
func i455() {
	z[456] = Zone{"eurovision", r, x, 0x42, "European Broadcasting Union (EBU)", "https://www.ebu.ch/home", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.eurovision", e, w{"https://rdap.nic.eurovision/"}, t}
	z[457] = Zone{"eus", r, x, 0x440, "Puntueus Fundazioa", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.eus", e, w{"https://rdap.nic.eus/"}, t}
	z[458] = Zone{"events", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.events", "v0n1.nic.events", "v0n2.nic.events", "v0n3.nic.events", "v2n0.nic.events", "v2n1.nic.events"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.events", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[459] = Zone{"everbank", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.everbank", e, n, f}
	z[460] = Zone{"example", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc6761", n, n, n, n, e, e, n, t}
	z[461] = Zone{"exchange", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.exchange", "v0n1.nic.exchange", "v0n2.nic.exchange", "v0n3.nic.exchange", "v2n0.nic.exchange", "v2n1.nic.exchange"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.exchange", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[462] = Zone{"expert", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.expert", "v0n1.nic.expert", "v0n2.nic.expert", "v0n3.nic.expert", "v2n0.nic.expert", "v2n1.nic.expert"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.expert", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[463] = Zone{"exposed", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.exposed", "v0n1.nic.exposed", "v0n2.nic.exposed", "v0n3.nic.exposed", "v2n0.nic.exposed", "v2n1.nic.exposed"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.exposed", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i463() {
	z[464] = Zone{"express", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.express", "v0n1.nic.express", "v0n2.nic.express", "v0n3.nic.express", "v2n0.nic.express", "v2n1.nic.express"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.express", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[465] = Zone{"extraspace", r, x, 0x42, "Extra Space Storage LLC", "https://www.icann.org/en/registry-agreements/details/extraspace", w{"a0.nic.extraspace", "a2.nic.extraspace", "b0.nic.extraspace", "c0.nic.extraspace"}, n, n, n, "whois.nic.extraspace", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[466] = Zone{"fage", r, x, 0x42, "Fage International S.A.", "https://nic.fage/", w{"a0.nic.fage", "a2.nic.fage", "b0.nic.fage", "c0.nic.fage"}, n, n, n, "whois.nic.fage", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[467] = Zone{"fail", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.fail", "v0n1.nic.fail", "v0n2.nic.fail", "v0n3.nic.fail", "v2n0.nic.fail", "v2n1.nic.fail"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fail", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[468] = Zone{"fairwinds", r, x, 0x42, "FairWinds Partners, LLC", "https://nic.fairwinds/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.fairwinds", e, w{"https://tld-rdap.verisign.com/fairwinds/v1/"}, t}
	z[469] = Zone{"faith", r, x, 0x40, "dot Faith Limited", "http://nic.faith/", w{"a.nic.faith", "b.nic.faith", "c.nic.faith", "ns1.dns.nic.faith", "ns2.dns.nic.faith", "ns3.dns.nic.faith"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.faith", e, w{"https://rdap.nic.faith/"}, t}
	z[470] = Zone{"family", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.family", "v0n1.nic.family", "v0n2.nic.family", "v0n3.nic.family", "v2n0.nic.family", "v2n1.nic.family"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.family", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[471] = Zone{"fan", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.fan", "v0n1.nic.fan", "v0n2.nic.fan", "v0n3.nic.fan", "v2n0.nic.fan", "v2n1.nic.fan"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fan", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i471() {
	z[472] = Zone{"fans", r, x, 0x40, "ZDNS International Limited", "http://nic.fans/", w{"a.nic.fans", "b.nic.fans", "c.nic.fans", "d.nic.fans"}, n, n, w{"ar", "mul-Latn", "zh"}, "whois.nic.fans", e, w{"https://rdap.centralnic.com/fans/"}, t}
	z[473] = Zone{"farm", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.farm", "v0n1.nic.farm", "v0n2.nic.farm", "v0n3.nic.farm", "v2n0.nic.farm", "v2n1.nic.farm"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.farm", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[474] = Zone{"farmers", r, x, 0x42, "Farmers Insurance Exchange", "https://www.nic.farmers/", w{"a.nic.farmers", "b.nic.farmers", "c.nic.farmers", "ns1.dns.nic.farmers", "ns2.dns.nic.farmers", "ns3.dns.nic.farmers"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.farmers", e, w{"https://rdap.nic.farmers/"}, t}
	z[475] = Zone{"fashion", r, x, 0x40, "Registry Services, LLC", "http://nic.fashion/", w{"a.nic.fashion", "b.nic.fashion", "c.nic.fashion", "x.nic.fashion", "y.nic.fashion", "z.nic.fashion"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.fashion", e, w{"https://rdap.nic.fashion/"}, t}
	z[476] = Zone{"fast", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.fast/", w{"dns1.nic.fast", "dns2.nic.fast", "dns3.nic.fast", "dns4.nic.fast", "dnsa.nic.fast", "dnsb.nic.fast", "dnsc.nic.fast", "dnsd.nic.fast"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.fast", e, w{"https://rdap.nominet.uk/fast/"}, t}
	z[477] = Zone{"fedex", r, x, 0x42, "Federal Express Corporation", "https://www.fedex.com/en-us/shipping.html", w{"a0.nic.fedex", "a2.nic.fedex", "b0.nic.fedex", "c0.nic.fedex"}, n, n, n, "whois.nic.fedex", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[478] = Zone{"feedback", r, x, 0x40, "Top Level Spectrum, Inc.", "http://www.nic.feedback/", w{"a.nic.feedback", "b.nic.feedback", "c.nic.feedback", "d.nic.feedback"}, n, n, w{"ar", "mul-Latn"}, "whois.nic.feedback", e, w{"https://rdap.centralnic.com/feedback/"}, t}
	z[479] = Zone{"ferrari", r, x, 0x42, "Fiat Chrysler Automobiles N.V.", "https://www.icann.org/en/registry-agreements/details/ferrari", w{"a0.nic.ferrari", "a2.nic.ferrari", "b0.nic.ferrari", "c0.nic.ferrari"}, n, n, n, "whois.nic.ferrari", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i479() {
	z[480] = Zone{"ferrero", r, x, 0x42, "Ferrero Trading Lux S.A.", "https://nic.ferrero/", w{"a.nic.ferrero", "b.nic.ferrero", "c.nic.ferrero", "ns1.dns.nic.ferrero", "ns2.dns.nic.ferrero", "ns3.dns.nic.ferrero"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.ferrero", e, w{"https://rdap.nic.ferrero/"}, t}
	z[481] = Zone{"fi", r, x, 0xa0, e, e, w{"a.fi", "b.fi", "c.fi", "d.fi", "e.fi", "g.fi", "h.fi", "i.fi", "j.fi", "k.fi"}, n, n, n, "whois.fi", e, w{"https://rdap.fi/rdap/rdap/"}, t}
	z[482] = Zone{"fiat", r, x, 0x1042, "Fiat Chrysler Automobiles N.V.", "http://nic.fiat/", n, n, n, n, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/fiat/"}, f}
	z[483] = Zone{"fidelity", r, x, 0x42, "Fidelity Brokerage Services LLC", "https://www.fidelity.com/why-fidelity/nic", w{"a0.nic.fidelity", "a2.nic.fidelity", "b0.nic.fidelity", "c0.nic.fidelity"}, n, n, n, "whois.nic.fidelity", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[484] = Zone{"fido", r, x, 0x42, "Rogers Communications Canada Inc.", "https://www.fido.ca/nic", w{"a0.nic.fido", "a2.nic.fido", "b0.nic.fido", "c0.nic.fido"}, n, n, n, "whois.nic.fido", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[485] = Zone{"film", r, x, 0x40, "Motion Picture Domain Registry Pty Ltd", "http://nic.film/", w{"a.nic.film", "b.nic.film", "c.nic.film", "x.nic.film", "y.nic.film", "z.nic.film"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "he", "hi", "hu", "is", "it", "ja", "ko", "lt", "lv", "nl", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.film", e, w{"https://rdap.nic.film/"}, t}
	z[486] = Zone{"final", r, x, 0x40, "Núcleo de Informação e Coordenação do Ponto BR - NIC.br", "https://www.icann.org/en/registry-agreements/details/final", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
	z[487] = Zone{"finance", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.finance", "v0n1.nic.finance", "v0n2.nic.finance", "v0n3.nic.finance", "v2n0.nic.finance", "v2n1.nic.finance"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.finance", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i487() {
	z[488] = Zone{"financial", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.financial", "v0n1.nic.financial", "v0n2.nic.financial", "v0n3.nic.financial", "v2n0.nic.financial", "v2n1.nic.financial"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.financial", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[489] = Zone{"financialaid", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[490] = Zone{"finish", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[491] = Zone{"fire", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.fire/", w{"dns1.nic.fire", "dns2.nic.fire", "dns3.nic.fire", "dns4.nic.fire", "dnsa.nic.fire", "dnsb.nic.fire", "dnsc.nic.fire", "dnsd.nic.fire"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.fire", e, w{"https://rdap.nominet.uk/fire/"}, t}
	z[492] = Zone{"firestone", r, x, 0x42, "Bridgestone Licensing Services, Inc", "https://nic.firestone/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.firestone", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[493] = Zone{"firmdale", r, x, 0x42, "Firmdale Holdings Limited", "http://nic.firmdale/", w{"ns1.nic.firmdale", "ns2.nic.firmdale"}, n, n, n, "whois.nic.firmdale", e, w{"https://rdap.nic.firmdale/"}, f}
	z[494] = Zone{"fish", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.fish", "v0n1.nic.fish", "v0n2.nic.fish", "v0n3.nic.fish", "v2n0.nic.fish", "v2n1.nic.fish"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fish", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[495] = Zone{"fishing", r, x, 0x40, "Registry Services, LLC", "http://nic.fishing/", w{"a.nic.fishing", "b.nic.fishing", "c.nic.fishing", "x.nic.fishing", "y.nic.fishing", "z.nic.fishing"}, n, n, w{"de", "es", "fr"}, "whois.nic.fishing", e, w{"https://rdap.nic.fishing/"}, t}
}

//go:noinline
func i495() {
	z[496] = Zone{"fit", r, x, 0x40, "Registry Services, LLC", "http://nic.fit/", w{"a.nic.fit", "b.nic.fit", "c.nic.fit", "x.nic.fit", "y.nic.fit", "z.nic.fit"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.fit", e, w{"https://rdap.nic.fit/"}, t}
	z[497] = Zone{"fitness", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.fitness", "v0n1.nic.fitness", "v0n2.nic.fitness", "v0n3.nic.fitness", "v2n0.nic.fitness", "v2n1.nic.fitness"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fitness", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[498] = Zone{"fj", r, z[2443:2454], 0xa0, e, "https://www.domains.fj/index.php", w{"ns1.fj", "ns2.fj", "ns3.fj", "ns4.fj", "ns5.fj"}, n, n, n, "www.whois.fj", e, n, f}
	z[499] = Zone{"fk", r, z[2454:2460], 0xa8, e, e, w{"ns1.horizon.net.fk", "ns2.horizon.net.fk", "ns3.horizon.net.fk"}, n, n, n, e, "http://whois.marcaria.com/", n, f}
	z[500] = Zone{"flickr", r, x, 0x42, "Flickr, Inc.", "http://nic.flickr/", w{"a.nic.flickr", "b.nic.flickr", "c.nic.flickr", "ns1.dns.nic.flickr", "ns2.dns.nic.flickr", "ns3.dns.nic.flickr"}, n, n, w{"de", "es", "ja", "ko", "ru", "zh"}, "whois.nic.flickr", e, w{"https://rdap.nic.flickr/"}, t}
	z[501] = Zone{"flights", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.flights", "v0n1.nic.flights", "v0n2.nic.flights", "v0n3.nic.flights", "v2n0.nic.flights", "v2n1.nic.flights"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.flights", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[502] = Zone{"flir", r, x, 0x42, "FLIR Systems, Inc.", "http://nic.flir/", w{"a.nic.flir", "b.nic.flir", "c.nic.flir", "ns1.dns.nic.flir", "ns2.dns.nic.flir", "ns3.dns.nic.flir"}, n, n, n, "whois.nic.flir", e, w{"https://rdap.nic.flir/"}, f}
	z[503] = Zone{"florist", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.florist", "v0n1.nic.florist", "v0n2.nic.florist", "v0n3.nic.florist", "v2n0.nic.florist", "v2n1.nic.florist"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.florist", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i503() {
	z[504] = Zone{"flowers", r, x, 0x40, "XYZ.COM LLC", "https://nic.flowers/", w{"a.nic.flowers", "b.nic.flowers", "c.nic.flowers", "d.nic.flowers"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.flowers", e, w{"https://rdap.centralnic.com/flowers/"}, t}
	z[505] = Zone{"fls", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[506] = Zone{"flsmidth", r, x, 0x1042, e, e, n, n, n, n, "whois.ksregistry.net", e, n, f}
	z[507] = Zone{"fly", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[508] = Zone{"fm", r, z[2460:2466], 0xe0, e, "https://dot.fm/", w{"a.nic.fm", "b.nic.fm", "c.nic.fm", "d.nic.fm", "e.nic.fm", "f.nic.fm"}, n, n, n, "whois.nic.fm", e, w{"https://rdap.centralnic.com/fm/"}, t}
	z[509] = Zone{"fo", r, z[2466:2488], 0xa0, e, e, w{"a.nic.fo", "b.nic.fo", "c.nic.fo", "d.nic.fo"}, n, n, n, "whois.nic.fo", e, w{"https://rdap.centralnic.com/fo/"}, f}
	z[510] = Zone{"foo", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[511] = Zone{"food", r, x, 0x42, "Internet Naming Company LLC", "https://www.icann.org/en/registry-agreements/details/food", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.food", e, w{"https://tld-rdap.verisign.com/food/v1/"}, f}
}

//go:noinline
func i511() {
	z[512] = Zone{"foodnetwork", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/foodnetwork", n, n, n, n, "whois.nic.foodnetwork", e, w{"https://tld-rdap.verisign.com/foodnetwork/v1/"}, f}
	z[513] = Zone{"football", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.football", "v0n1.nic.football", "v0n2.nic.football", "v0n3.nic.football", "v2n0.nic.football", "v2n1.nic.football"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.football", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[514] = Zone{"ford", r, x, 0x42, "Ford Motor Company", "https://nic.ford/", w{"a.nic.ford", "b.nic.ford", "c.nic.ford", "ns1.dns.nic.ford", "ns2.dns.nic.ford", "ns3.dns.nic.ford"}, n, n, w{"da", "fr", "pl", "ru", "sv", "zh"}, "whois.nic.ford", e, w{"https://rdap.nic.ford/"}, t}
	z[515] = Zone{"forex", r, x, 0x40, "Dog Beach, LLC", "http://nic.forex/", w{"v0n0.nic.forex", "v0n1.nic.forex", "v0n2.nic.forex", "v0n3.nic.forex", "v2n0.nic.forex", "v2n1.nic.forex"}, n, n, n, "whois.nic.forex", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[516] = Zone{"forsale", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.forsale", "v0n1.nic.forsale", "v0n2.nic.forsale", "v0n3.nic.forsale", "v2n0.nic.forsale", "v2n1.nic.forsale"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.forsale", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[517] = Zone{"forum", r, x, 0x40, "Fegistry, LLC", "https://www.icann.org/en/registry-agreements/details/forum", w{"a.nic.forum", "b.nic.forum", "c.nic.forum", "d.nic.forum"}, n, n, w{"ar"}, "whois.nic.forum", e, w{"https://rdap.centralnic.com/forum/"}, t}
	z[518] = Zone{"foundation", r, x, 0x40, "Public Interest Registry", e, w{"v0n0.nic.foundation", "v0n1.nic.foundation", "v0n2.nic.foundation", "v0n3.nic.foundation", "v2n0.nic.foundation", "v2n1.nic.foundation"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.foundation", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[519] = Zone{"fox", r, x, 0x42, "FOX Registry, LLC", "https://www.nic.fox/", w{"a.nic.fox", "b.nic.fox", "c.nic.fox", "ns1.dns.nic.fox", "ns2.dns.nic.fox", "ns3.dns.nic.fox"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.fox", e, w{"https://rdap.nic.fox/"}, t}
}

//go:noinline
func i519() {
	z[520] = Zone{"fr", r, z[2488:2506], 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.fr", e, w{"https://rdap.nic.fr/"}, t}
	z[521] = Zone{"free", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.free/", w{"dns1.nic.free", "dns2.nic.free", "dns3.nic.free", "dns4.nic.free", "dnsa.nic.free", "dnsb.nic.free", "dnsc.nic.free", "dnsd.nic.free"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.free", e, w{"https://rdap.nominet.uk/free/"}, t}
	z[522] = Zone{"fresenius", r, x, 0x42, "Fresenius Immobilien-Verwaltungs-GmbH", "https://nic.fresenius/", w{"a.nic.fresenius", "b.nic.fresenius", "c.nic.fresenius", "d.nic.fresenius"}, n, n, w{"de"}, "whois.nic.fresenius", e, w{"https://rdap.centralnic.com/fresenius/"}, t}
	z[523] = Zone{"frl", r, x, 0x440, "FRLregistry B.V.", "https://nic.frl/", w{"a.nic.frl", "b.nic.frl", "c.nic.frl", "d.nic.frl"}, n, w{"NL-FY"}, w{"mul-Latn"}, "whois.nic.frl", e, w{"https://rdap.centralnic.com/frl/"}, t}
	z[524] = Zone{"frogans", r, x, 0x42, "OP3FT", "https://nic.frogans/en/main.html", w{"a0.nic.frogans", "a2.nic.frogans", "b0.nic.frogans", "c0.nic.frogans"}, n, n, w{"mul-Latn"}, "whois.nic.frogans", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[525] = Zone{"frontdoor", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/frontdoor", n, n, n, n, "whois.nic.frontdoor", e, w{"https://tld-rdap.verisign.com/frontdoor/v1/"}, f}
	z[526] = Zone{"frontier", r, x, 0x42, "Frontier Communications Corporation", "http://nic.frontier/", w{"a.nic.frontier", "b.nic.frontier", "c.nic.frontier", "ns1.dns.nic.frontier", "ns2.dns.nic.frontier", "ns3.dns.nic.frontier"}, n, n, n, "whois.nic.frontier", e, w{"https://rdap.nic.frontier/"}, f}
	z[527] = Zone{"ftr", r, x, 0x42, "Frontier Communications Corporation", "http://nic.ftr/", w{"a.nic.ftr", "b.nic.ftr", "c.nic.ftr", "ns1.dns.nic.ftr", "ns2.dns.nic.ftr", "ns3.dns.nic.ftr"}, n, n, n, "whois.nic.ftr", e, w{"https://rdap.nic.ftr/"}, f}
}

//go:noinline
func i527() {
	z[528] = Zone{"fujitsu", r, x, 0x42, "Fujitsu Limited", "https://nic.fujitsu/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[529] = Zone{"fujixerox", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.fujixerox", e, n, t}
	z[530] = Zone{"fun", r, x, 0x40, "Radix Technologies Inc.", "https://radix.website/dot-fun", w{"a.nic.fun", "b.nic.fun", "e.nic.fun", "f.nic.fun"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.fun", e, w{"https://rdap.centralnic.com/fun/"}, t}
	z[531] = Zone{"fund", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.fund", "v0n1.nic.fund", "v0n2.nic.fund", "v0n3.nic.fund", "v2n0.nic.fund", "v2n1.nic.fund"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fund", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[532] = Zone{"furniture", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.furniture", "v0n1.nic.furniture", "v0n2.nic.furniture", "v0n3.nic.furniture", "v2n0.nic.furniture", "v2n1.nic.furniture"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.furniture", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[533] = Zone{"futbol", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.futbol", "v0n1.nic.futbol", "v0n2.nic.futbol", "v0n3.nic.futbol", "v2n0.nic.futbol", "v2n1.nic.futbol"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.futbol", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[534] = Zone{"fyi", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.fyi", "v0n1.nic.fyi", "v0n2.nic.fyi", "v0n3.nic.fyi", "v2n0.nic.fyi", "v2n1.nic.fyi"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.fyi", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[535] = Zone{"ga", r, z[2506:2519], 0xa0, e, e, w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.dot.ga", e, n, t}
}

//go:noinline
func i535() {
	z[536] = Zone{"gal", r, x, 0x4c0, "Asociación puntoGAL", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, w{"ES-GA"}, w{"mul-Latn"}, "whois.nic.gal", e, w{"https://rdap.nic.gal/"}, t}
	z[537] = Zone{"gallery", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.gallery", "v0n1.nic.gallery", "v0n2.nic.gallery", "v0n3.nic.gallery", "v2n0.nic.gallery", "v2n1.nic.gallery"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gallery", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[538] = Zone{"gallo", r, x, 0x42, "Gallo Vineyards, Inc.", "https://www.icann.org/en/registry-agreements/details/gallo", w{"a0.nic.gallo", "a2.nic.gallo", "b0.nic.gallo", "c0.nic.gallo"}, n, n, n, "whois.nic.gallo", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[539] = Zone{"gallup", r, x, 0x42, "Gallup, Inc.", "http://nic.gallup/home.aspx", w{"a0.nic.gallup", "a2.nic.gallup", "b0.nic.gallup", "c0.nic.gallup"}, n, n, n, "whois.nic.gallup", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[540] = Zone{"game", r, x, 0x40, "XYZ.COM LLC", "https://nic.game/", w{"a.nic.game", "b.nic.game", "c.nic.game", "d.nic.game"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.game", e, w{"https://rdap.centralnic.com/game/"}, t}
	z[541] = Zone{"games", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.games", "v0n1.nic.games", "v0n2.nic.games", "v0n3.nic.games", "v2n0.nic.games", "v2n1.nic.games"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.games", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[542] = Zone{"gap", r, x, 0x42, "The Gap, Inc.", "https://www.nic.gap/", w{"a.nic.gap", "b.nic.gap", "c.nic.gap", "ns1.dns.nic.gap", "ns2.dns.nic.gap", "ns3.dns.nic.gap"}, n, n, w{"es"}, "whois.nic.gap", e, w{"https://rdap.nic.gap/"}, t}
	z[543] = Zone{"garden", r, x, 0x40, "Registry Services, LLC", "http://nic.garden/", w{"a.nic.garden", "b.nic.garden", "c.nic.garden", "x.nic.garden", "y.nic.garden", "z.nic.garden"}, n, n, w{"de", "es", "fr"}, "whois.nic.garden", e, w{"https://rdap.nic.garden/"}, t}
}

//go:noinline
func i543() {
	z[544] = Zone{"garnier", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[545] = Zone{"gay", r, x, 0x40, "Registry Services, LLC", "https://www.icann.org/en/registry-agreements/details/gay", w{"a.nic.gay", "b.nic.gay", "c.nic.gay", "x.nic.gay", "y.nic.gay", "z.nic.gay"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.gay", e, w{"https://rdap.nic.gay/"}, t}
	z[546] = Zone{"gb", r, z[2519:2520], 0xa8, e, e, w{"ns.uu.net", "ns0.ja.net", "ns4.ja.net"}, n, n, n, e, e, n, f}
	z[547] = Zone{"gbiz", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[548] = Zone{"gcc", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[549] = Zone{"gd", r, z[2520:2527], 0xa0, e, e, w{"a.nic.gd", "b.nic.gd", "c.nic.gd", "d.nic.gd"}, n, n, n, "whois.nic.gd", e, w{"https://rdap.centralnic.com/gd/"}, f}
	z[550] = Zone{"gdn", r, x, 0x42, "Joint Stock Company \"Navigation-information systems\"", "http://nic.gdn/", w{"ns1.nic.gdn", "ns3.nic.gdn", "ns4.nic.gdn"}, n, n, n, "whois.nic.gdn", e, w{"https://rdap.nic.gdn/"}, f}
	z[551] = Zone{"ge", r, z[2527:2535], 0xa0, e, "https://nic.ge/", w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, "whois.nic.ge", e, n, f}
}

//go:noinline
func i551() {
	z[552] = Zone{"gea", r, x, 0x42, "GEA Group Aktiengesellschaft", "https://nic.gea/", w{"a.dns.nic.gea", "m.dns.nic.gea", "n.dns.nic.gea"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.gea", e, w{"https://rdap.nic.gea/"}, t}
	z[553] = Zone{"gecompany", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[554] = Zone{"ged", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[555] = Zone{"gent", r, x, 0xc4, "Easyhost BV", e, w{"a.nic.gent", "b.nic.gent", "c.nic.gent", "d.nic.gent"}, n, w{"Gent", "Ghent"}, n, "whois.nic.gent", e, w{"https://rdap.centralnic.com/gent/"}, t}
	z[556] = Zone{"genting", r, x, 0x42, "Resorts World Inc Pte. Ltd.", "https://www.icann.org/en/registry-agreements/details/genting", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.genting", e, w{"https://tld-rdap.verisign.com/genting/v1/"}, t}
	z[557] = Zone{"george", r, x, 0x42, "Wal-Mart Stores, Inc.", "https://www.icann.org/en/registry-agreements/details/george", w{"a0.nic.george", "a2.nic.george", "b0.nic.george", "c0.nic.george"}, n, n, n, "whois.nic.george", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[558] = Zone{"gf", r, x, 0xa0, e, "https://www.dom-enic.com/", w{"ns1-fr.mediaserv.net", "ns1-gp.mediaserv.net", "ns1-mq.mediaserv.net"}, n, n, n, "whois.mediaserv.net", e, n, f}
	z[559] = Zone{"gg", r, z[2535:2546], 0xa0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "e.ci-servers.net"}, n, n, n, "whois.gg", e, n, f}
}

//go:noinline
func i559() {
	z[560] = Zone{"ggee", r, x, 0x42, "GMO Internet, Inc.", "https://nic.ggee/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.ggee", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[561] = Zone{"gh", r, z[2546:2552], 0xa8, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, "whois.nic.gh", e, n, f}
	z[562] = Zone{"gi", r, z[2552:2558], 0xa0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, "whois2.afilias-grs.net", e, n, f}
	z[563] = Zone{"gift", r, x, 0x40, "DotGift, LLC", "https://uniregistry.link/extensions/gift/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.uniregistry.net", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[564] = Zone{"gifts", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.gifts", "v0n1.nic.gifts", "v0n2.nic.gifts", "v0n3.nic.gifts", "v2n0.nic.gifts", "v2n1.nic.gifts"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gifts", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[565] = Zone{"gives", r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/work-with-us/contact/", w{"v0n0.nic.gives", "v0n1.nic.gives", "v0n2.nic.gives", "v0n3.nic.gives", "v2n0.nic.gives", "v2n1.nic.gives"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gives", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[566] = Zone{"giving", r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/work-with-us/contact/", w{"a0.nic.giving", "a2.nic.giving", "b0.nic.giving", "b2.nic.giving", "c0.nic.giving", "d0.nic.giving"}, n, n, n, "whois.nic.giving", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, f}
	z[567] = Zone{"gl", r, z[2558:2564], 0xa0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.gl", e, n, t}
}

//go:noinline
func i567() {
	z[568] = Zone{"glade", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.glade", e, w{"https://tld-rdap.verisign.com/glade/v1/"}, f}
	z[569] = Zone{"glass", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.glass", "v0n1.nic.glass", "v0n2.nic.glass", "v0n3.nic.glass", "v2n0.nic.glass", "v2n1.nic.glass"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.glass", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[570] = Zone{"gle", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[571] = Zone{"global", r, x, 0x40, "Identity Digital Limited", e, w{"a0.nic.global", "a2.nic.global", "b0.nic.global", "c0.nic.global"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "tr", "uk", "zh-CN", "zh-TW"}, "whois.nic.global", e, w{"https://rdap.donuts.co/rdap/"}, t}
	z[572] = Zone{"globalx", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[573] = Zone{"globo", r, x, 0x42, "Globo Comunicação e Participações S.A", "https://nic.globo/", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
	z[574] = Zone{"gm", r, x, 0xa0, e, e, w{"ns-gm.afrinic.net", "ns1.nic.gm", "ns2.nic.gm"}, n, n, n, e, "http://www.nic.gm/htmlpages/whois.htm", n, f}
	z[575] = Zone{"gmail", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
}

//go:noinline
func i575() {
	z[576] = Zone{"gmbh", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.gmbh", "v0n1.nic.gmbh", "v0n2.nic.gmbh", "v0n3.nic.gmbh", "v2n0.nic.gmbh", "v2n1.nic.gmbh"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gmbh", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[577] = Zone{"gmc", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[578] = Zone{"gmo", r, x, 0x42, "GMO Internet, Inc.", "https://nic.gmo/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[579] = Zone{"gmx", r, x, 0x42, "1&1 Mail & Media GmbH", "https://www.icann.org/en/registry-agreements/details/gmx", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"de"}, "whois.nic.gmx", e, w{"https://rdap.nic.gmx/"}, t}
	z[580] = Zone{"gn", r, z[2564:2570], 0xa8, e, e, w{"fork.sth.dnsnode.net", "ns-gn.afrinic.net", "ns-pch.gn", "ns1.gn"}, n, n, n, "whois.ande.gov.gn", e, n, f}
	z[581] = Zone{"godaddy", r, x, 0x42, "Go Daddy East, LLC", "https://www.godaddy.com/help/about-godaddy-domains-24085", w{"a.nic.godaddy", "b.nic.godaddy", "c.nic.godaddy", "x.nic.godaddy", "y.nic.godaddy", "z.nic.godaddy"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.godaddy", e, w{"https://rdap.nic.godaddy/"}, t}
	z[582] = Zone{"gold", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.gold", "v0n1.nic.gold", "v0n2.nic.gold", "v0n3.nic.gold", "v2n0.nic.gold", "v2n1.nic.gold"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gold", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[583] = Zone{"goldpoint", r, x, 0x42, "YODOBASHI CAMERA CO.,LTD.", "https://nic.goldpoint/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.goldpoint", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i583() {
	z[584] = Zone{"golf", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.golf", "v0n1.nic.golf", "v0n2.nic.golf", "v0n3.nic.golf", "v2n0.nic.golf", "v2n1.nic.golf"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.golf", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[585] = Zone{"goo", r, x, 0x42, "NTT DOCOMO, INC.", "https://nic.goo/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[586] = Zone{"goodhands", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[587] = Zone{"goodyear", r, x, 0x42, "The Goodyear Tire & Rubber Company", "https://nic.goodyear/", w{"a0.nic.goodyear", "a2.nic.goodyear", "b0.nic.goodyear", "c0.nic.goodyear"}, n, n, n, "whois.nic.goodyear", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[588] = Zone{"goog", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[589] = Zone{"google", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"ja", "mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[590] = Zone{"gop", r, x, 0x40, "Republican State Leadership Committee, Inc.", "https://www.join.gop/", w{"dns1.nic.gop", "dns2.nic.gop", "dns3.nic.gop", "dns4.nic.gop", "dnsa.nic.gop", "dnsb.nic.gop", "dnsc.nic.gop", "dnsd.nic.gop"}, n, n, w{"es"}, "whois.nic.gop", e, w{"https://rdap.nominet.uk/gop/"}, t}
	z[591] = Zone{"got", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.got/", w{"dns1.nic.got", "dns2.nic.got", "dns3.nic.got", "dns4.nic.got", "dnsa.nic.got", "dnsb.nic.got", "dnsc.nic.got", "dnsd.nic.got"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.got", e, w{"https://rdap.nominet.uk/got/"}, t}
}

//go:noinline
func i591() {
	z[592] = Zone{"gotv", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[593] = Zone{"gov", r, x, 0x4040, e, e, w{"a.ns.gov", "b.ns.gov", "c.ns.gov", "d.ns.gov"}, n, n, n, "whois.dotgov.gov", e, n, f}
	z[594] = Zone{"gp", r, z[2570:2581], 0xa0, e, "https://www.nic.gp/", w{"a.lactld.org", "gp.cctld.authdns.ripe.net", "ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, "whois.nic.gp", e, n, f}
	z[595] = Zone{"gq", r, x, 0xa0, e, e, w{"a.ns.gq", "b.ns.gq", "c.ns.gq", "d.ns.gq"}, n, n, n, "whois.dominio.gq", e, n, t}
	z[596] = Zone{"gr", r, z[2581:2586], 0xa0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "gr-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, "https://grweb.ics.forth.gr/", n, t}
	z[597] = Zone{"grainger", r, x, 0x42, "Grainger Registry Services, LLC", "http://nic.grainger/", w{"a.nic.grainger", "b.nic.grainger", "c.nic.grainger", "ns1.dns.nic.grainger", "ns2.dns.nic.grainger", "ns3.dns.nic.grainger"}, n, n, w{"es"}, "whois.nic.grainger", e, w{"https://rdap.nic.grainger/"}, t}
	z[598] = Zone{"graphics", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.graphics", "v0n1.nic.graphics", "v0n2.nic.graphics", "v0n3.nic.graphics", "v2n0.nic.graphics", "v2n1.nic.graphics"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.graphics", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[599] = Zone{"gratis", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.gratis", "v0n1.nic.gratis", "v0n2.nic.gratis", "v0n3.nic.gratis", "v2n0.nic.gratis", "v2n1.nic.gratis"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gratis", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i599() {
	z[600] = Zone{"gree", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[601] = Zone{"green", r, x, 0x40, "Identity Digital Limited", "http://nic.green/", w{"a0.nic.green", "a2.nic.green", "b0.nic.green", "c0.nic.green"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.green", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[602] = Zone{"gripe", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.gripe", "v0n1.nic.gripe", "v0n2.nic.gripe", "v0n3.nic.gripe", "v2n0.nic.gripe", "v2n1.nic.gripe"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.gripe", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[603] = Zone{"grocery", r, x, 0x40, "Wal-Mart Stores, Inc.", "https://www.icann.org/en/registry-agreements/details/grocery", w{"a0.nic.grocery", "a2.nic.grocery", "b0.nic.grocery", "c0.nic.grocery"}, n, n, n, "whois.nic.grocery", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[604] = Zone{"group", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.group", "v0n1.nic.group", "v0n2.nic.group", "v0n3.nic.group", "v2n0.nic.group", "v2n1.nic.group"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.group", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[605] = Zone{"gs", r, x, 0xa0, e, e, w{"ns.anycast.nic.gs", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.gs", e, n, f}
	z[606] = Zone{"gt", r, z[2586:2593], 0xa0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, "http://www.gt/", n, f}
	z[607] = Zone{"gu", r, z[2593:2600], 0xa0, e, "https://www.iana.org/domains/root/db/gu.html", w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, "http://gadao.gov.gu/domainsearch.htm", n, f}
}

//go:noinline
func i607() {
	z[608] = Zone{"guardian", r, x, 0x42, "The Guardian Life Insurance Company of America", "https://www.icann.org/en/registry-agreements/details/guardian", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, e, e, w{"https://tld-rdap.verisign.com/guardian/v1/"}, f}
	z[609] = Zone{"guardianlife", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[610] = Zone{"guardianmedia", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[611] = Zone{"gucci", r, x, 0x42, "Guccio Gucci S.p.a.", "http://nic.gucci/", w{"dns1.nic.gucci", "dns2.nic.gucci", "dns3.nic.gucci", "dns4.nic.gucci", "dnsa.nic.gucci", "dnsb.nic.gucci", "dnsc.nic.gucci", "dnsd.nic.gucci"}, n, n, w{"ar", "da", "de", "es", "fi", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.gucci", e, w{"https://rdap.nominet.uk/gucci/"}, t}
	z[612] = Zone{"guge", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[613] = Zone{"guide", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.guide", "v0n1.nic.guide", "v0n2.nic.guide", "v0n3.nic.guide", "v2n0.nic.guide", "v2n1.nic.guide"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.guide", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[614] = Zone{"guitars", r, x, 0x40, "XYZ.COM LLC", "https://nic.guitars/", w{"a.nic.guitars", "b.nic.guitars", "c.nic.guitars", "d.nic.guitars"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.guitars", e, w{"https://rdap.centralnic.com/guitars/"}, t}
	z[615] = Zone{"guru", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.guru", "v0n1.nic.guru", "v0n2.nic.guru", "v0n3.nic.guru", "v2n0.nic.guru", "v2n1.nic.guru"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.guru", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i615() {
	z[616] = Zone{"gw", r, x, 0xa0, e, e, w{"gw01.dns.pt", "gw03.dns.pt", "h.dns.pt"}, n, n, n, e, "http://nic.gw/en/whois/", n, f}
	z[617] = Zone{"gy", r, z[2600:2604], 0xa0, e, e, w{"a.lactld.org", "gy-ns.anycast.pch.net"}, n, n, n, "whois.registry.gy", e, n, f}
	z[618] = Zone{"hair", r, x, 0x40, "XYZ.COM LLC", "https://nic.hair/", w{"a.nic.hair", "b.nic.hair", "c.nic.hair", "d.nic.hair"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.hair", e, w{"https://rdap.centralnic.com/hair/"}, t}
	z[619] = Zone{"halal", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[620] = Zone{"hamburg", r, x, 0xc4, "Hamburg Top-Level-Domain GmbH", e, w{"a.dns.nic.hamburg", "m.dns.nic.hamburg", "n.dns.nic.hamburg"}, n, w{"Hamburg", "DE-HH"}, w{"mul-Cyrl", "mul-Latn"}, "whois.nic.hamburg", e, w{"https://rdap.nic.hamburg/v1/"}, t}
	z[621] = Zone{"hangout", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[622] = Zone{"haus", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.haus", "v0n1.nic.haus", "v0n2.nic.haus", "v0n3.nic.haus", "v2n0.nic.haus", "v2n1.nic.haus"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.haus", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[623] = Zone{"hbo", r, x, 0x42, "HBO Registry Services, Inc.", "http://nic.hbo/", w{"a.nic.hbo", "b.nic.hbo", "c.nic.hbo", "ns1.dns.nic.hbo", "ns2.dns.nic.hbo", "ns3.dns.nic.hbo"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.hbo", e, w{"https://rdap.nic.hbo/"}, t}
}

//go:noinline
func i623() {
	z[624] = Zone{"hdfc", r, x, 0x42, "HOUSING DEVELOPMENT FINANCE CORPORATION LIMITED", "http://nic.hdfc/", w{"a0.nic.hdfc", "a2.nic.hdfc", "b0.nic.hdfc", "c0.nic.hdfc"}, n, n, n, "whois.nic.hdfc", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[625] = Zone{"hdfcbank", r, x, 0x42, "HDFC Bank Limited", "https://www.icann.org/en/registry-agreements/details/hdfcbank", w{"a0.nic.hdfcbank", "a2.nic.hdfcbank", "b0.nic.hdfcbank", "c0.nic.hdfcbank"}, n, n, n, "whois.nic.hdfcbank", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[626] = Zone{"health", r, x, 0x40, "Registry Services, LLC", "https://www.icann.org/en/registry-agreements/details/health", w{"a.nic.health", "b.nic.health", "c.nic.health", "ns1.dns.nic.health", "ns2.dns.nic.health", "ns3.dns.nic.health"}, n, n, w{"es"}, "whois.nic.health", e, w{"https://rdap.nic.health/"}, t}
	z[627] = Zone{"healthcare", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.healthcare", "v0n1.nic.healthcare", "v0n2.nic.healthcare", "v0n3.nic.healthcare", "v2n0.nic.healthcare", "v2n1.nic.healthcare"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.healthcare", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[628] = Zone{"heart", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[629] = Zone{"heinz", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[630] = Zone{"help", r, x, 0x40, "Innovation service Limited", "https://uniregistry.link/extensions/help/", w{"a.nic.help", "b.nic.help", "c.nic.help", "d.nic.help"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.help", e, w{"https://rdap.centralnic.com/help/"}, t}
	z[631] = Zone{"helsinki", r, x, 0x40, "City of Helsinki", "https://www.hel.fi/Helsinki/en/administration/information/information/nic.helsinki/", w{"a0.nic.helsinki", "a2.nic.helsinki", "b0.nic.helsinki", "c0.nic.helsinki"}, n, n, w{"fi", "sv"}, "whois.nic.helsinki", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i631() {
	z[632] = Zone{"here", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[633] = Zone{"hermes", r, x, 0x42, "HERMES INTERNATIONAL", "http://www.nic.hermes/index.php", w{"a0.nic.hermes", "a2.nic.hermes", "b0.nic.hermes", "c0.nic.hermes"}, n, n, n, "whois.nic.hermes", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[634] = Zone{"hgtv", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/hgtv", n, n, n, n, "whois.nic.hgtv", e, w{"https://tld-rdap.verisign.com/hgtv/v1/"}, f}
	z[635] = Zone{"hilton", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[636] = Zone{"hiphop", r, x, 0x40, "Dot Hip Hop, LLC", "https://get.hiphop/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.hiphop", e, w{"https://rdap.registry.hiphop/rdap/"}, t}
	z[637] = Zone{"hisamitsu", r, x, 0x42, "Hisamitsu Pharmaceutical Co.,Inc.", "https://nic.hisamitsu/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[638] = Zone{"hitachi", r, x, 0x42, "Hitachi, Ltd.", "https://nic.hitachi/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[639] = Zone{"hiv", r, x, 0x40, "Internet Naming Company LLC", "https://uniregistry.link/extensions/hiv/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "mul-Cyrl", "pt", "zh"}, "whois.nic.hiv", e, w{"https://whois.uniregistry.net/rdap/"}, t}
}

//go:noinline
func i639() {
	z[640] = Zone{"hk", r, z[2604:2613], 0xa0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, "whois.hkirc.hk", e, n, t}
	z[641] = Zone{"hkt", r, x, 0x42, "PCCW-HKT DataCom Services Limited", "https://www.icann.org/en/registry-agreements/details/hkt", w{"a0.nic.hkt", "a2.nic.hkt", "b0.nic.hkt", "c0.nic.hkt"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.hkt", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[642] = Zone{"hm", r, x, 0xa0, e, "http://www.registry.hm/", w{"ns1.registry.hm", "ns2.registry.hm", "ns3.registry.hm"}, n, n, n, "whois.registry.hm", e, n, f}
	z[643] = Zone{"hn", r, z[2613:2619], 0xa0, e, e, w{"a.lactld.org", "ns1.anycastdns.cz", "ns2.anycastdns.cz", "pch-anycast.rds.org.hn"}, n, n, n, "whois.nic.hn", e, n, f}
	z[644] = Zone{"hockey", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.hockey", "v0n1.nic.hockey", "v0n2.nic.hockey", "v0n3.nic.hockey", "v2n0.nic.hockey", "v2n1.nic.hockey"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.hockey", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[645] = Zone{"holdings", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.holdings", "v0n1.nic.holdings", "v0n2.nic.holdings", "v0n3.nic.holdings", "v2n0.nic.holdings", "v2n1.nic.holdings"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.holdings", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[646] = Zone{"holiday", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.holiday", "v0n1.nic.holiday", "v0n2.nic.holiday", "v0n3.nic.holiday", "v2n0.nic.holiday", "v2n1.nic.holiday"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.holiday", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[647] = Zone{"home", r, x, 0x800, "ICANN", "https://features.icann.org/addressing-new-gtld-program-applications-corp-home-and-mail", n, n, n, n, e, e, n, t}
}

//go:noinline
func i647() {
	z[648] = Zone{"homedepot", r, x, 0x42, "Home Depot Product Authority, LLC", "https://www.icann.org/en/registry-agreements/details/homedepot", w{"a0.nic.homedepot", "a2.nic.homedepot", "b0.nic.homedepot", "c0.nic.homedepot"}, n, n, n, "whois.nic.homedepot", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[649] = Zone{"homegoods", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.homegoods/", w{"a.nic.homegoods", "b.nic.homegoods", "c.nic.homegoods", "ns1.dns.nic.homegoods", "ns2.dns.nic.homegoods", "ns3.dns.nic.homegoods"}, n, n, n, "whois.nic.homegoods", e, w{"https://rdap.nic.homegoods/"}, f}
	z[650] = Zone{"homes", r, x, 0x40, "XYZ.COM LLC", "https://nic.homes/", w{"a.nic.homes", "b.nic.homes", "c.nic.homes", "d.nic.homes"}, n, n, n, "whois.nic.homes", e, w{"https://rdap.centralnic.com/homes/"}, f}
	z[651] = Zone{"homesense", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.homesense/", w{"a.nic.homesense", "b.nic.homesense", "c.nic.homesense", "ns1.dns.nic.homesense", "ns2.dns.nic.homesense", "ns3.dns.nic.homesense"}, n, n, n, "whois.nic.homesense", e, w{"https://rdap.nic.homesense/"}, f}
	z[652] = Zone{"honda", r, x, 0x42, "Honda Motor Co., Ltd.", "https://nic.honda/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.honda", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[653] = Zone{"honeywell", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.honeywell", e, n, f}
	z[654] = Zone{"horse", r, x, 0x40, "Registry Services, LLC", "http://nic.horse/", w{"a.nic.horse", "b.nic.horse", "c.nic.horse", "x.nic.horse", "y.nic.horse", "z.nic.horse"}, n, n, w{"de", "es", "fr"}, "whois.nic.horse", e, w{"https://rdap.nic.horse/"}, t}
	z[655] = Zone{"hospital", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.hospital", "v0n1.nic.hospital", "v0n2.nic.hospital", "v0n3.nic.hospital", "v2n0.nic.hospital", "v2n1.nic.hospital"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.hospital", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i655() {
	z[656] = Zone{"host", r, x, 0x40, "Radix Technologies Inc.", e, w{"a.nic.host", "b.nic.host", "e.nic.host", "f.nic.host"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.host", e, w{"https://rdap.centralnic.com/host/"}, t}
	z[657] = Zone{"hosting", r, x, 0x40, "XYZ.COM LLC", "https://nic.hosting/", w{"a.nic.hosting", "b.nic.hosting", "c.nic.hosting", "d.nic.hosting"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.hosting", e, w{"https://rdap.centralnic.com/hosting/"}, t}
	z[658] = Zone{"hot", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.hot/", w{"dns1.nic.hot", "dns2.nic.hot", "dns3.nic.hot", "dns4.nic.hot", "dnsa.nic.hot", "dnsb.nic.hot", "dnsc.nic.hot", "dnsd.nic.hot"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.hot", e, w{"https://rdap.nominet.uk/hot/"}, t}
	z[659] = Zone{"hoteis", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[660] = Zone{"hotel", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[661] = Zone{"hoteles", r, x, 0x1040, "Travel Reservations SRL", "http://nic.hoteles/", n, n, n, w{"es"}, "whois.nic.hoteles", e, w{"https://rdap.nic.hoteles/"}, t}
	z[662] = Zone{"hotels", r, x, 0x40, "Booking.com B.V.", "https://www.icann.org/en/registry-agreements/details/hotels", w{"a.nic.hotels", "b.nic.hotels", "c.nic.hotels", "ns1.dns.nic.hotels", "ns2.dns.nic.hotels", "ns3.dns.nic.hotels"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.hotels", e, w{"https://rdap.nic.hotels/"}, t}
	z[663] = Zone{"hotmail", r, x, 0x42, "Microsoft Corporation", "https://nic.hotmail/", w{"dns1.nic.hotmail", "dns2.nic.hotmail", "dns3.nic.hotmail", "dns4.nic.hotmail", "dnsa.nic.hotmail", "dnsb.nic.hotmail", "dnsc.nic.hotmail", "dnsd.nic.hotmail"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/hotmail/"}, t}
}

//go:noinline
func i663() {
	z[664] = Zone{"house", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.house", "v0n1.nic.house", "v0n2.nic.house", "v0n3.nic.house", "v2n0.nic.house", "v2n1.nic.house"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.house", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[665] = Zone{"how", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"ja", "mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[666] = Zone{"hr", r, z[2619:2623], 0xa0, e, e, w{"hr-ns-1.carnet.hr", "n.dns.hr", "pch.carnet.hr"}, n, n, n, "whois.dns.hr", e, n, f}
	z[667] = Zone{"hsbc", r, x, 0x42, "HSBC Global Services (UK) Limited", "https://nic.hsbc/", w{"a.nic.hsbc", "b.nic.hsbc", "c.nic.hsbc", "ns1.dns.nic.hsbc", "ns2.dns.nic.hsbc", "ns3.dns.nic.hsbc"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.hsbc", e, w{"https://rdap.nic.hsbc/"}, t}
	z[668] = Zone{"ht", r, z[2623:2640], 0xa0, e, e, w{"a.lactld.org", "ns-ht.nic.fr", "pch.nic.ht"}, n, n, n, "whois.nic.ht", e, n, f}
	z[669] = Zone{"htc", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[670] = Zone{"hu", r, z[2640:2674], 0xa0, e, e, w{"a.hu", "b.hu", "c.hu", "d.hu", "e.hu", "f.hu"}, n, n, n, "whois.nic.hu", e, n, t}
	z[671] = Zone{"hughes", r, x, 0x42, "Hughes Satellite Systems Corporation", "https://www.icann.org/en/registry-agreements/details/hughes", w{"a0.nic.hughes", "a2.nic.hughes", "b0.nic.hughes", "c0.nic.hughes"}, n, n, n, "whois.nic.hughes", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i671() {
	z[672] = Zone{"hyatt", r, x, 0x42, "Hyatt GTLD, L.L.C.", "https://www.icann.org/en/registry-agreements/details/hyatt", w{"a.nic.hyatt", "b.nic.hyatt", "c.nic.hyatt", "ns1.dns.nic.hyatt", "ns2.dns.nic.hyatt", "ns3.dns.nic.hyatt"}, n, n, w{"es"}, "whois.nic.hyatt", e, w{"https://rdap.nic.hyatt/"}, t}
	z[673] = Zone{"hyundai", r, x, 0x42, "Hyundai Motor Company", "https://nic.hyundai/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ko"}, "whois.nic.hyundai", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[674] = Zone{"ibm", r, x, 0x42, "International Business Machines Corporation", "https://www.nic.ibm/", w{"a.nic.ibm", "b.nic.ibm", "c.nic.ibm", "x.nic.ibm", "y.nic.ibm", "z.nic.ibm"}, n, n, n, "whois.nic.ibm", e, w{"https://rdap.nic.ibm/"}, f}
	z[675] = Zone{"icbc", r, x, 0x42, "Industrial and Commercial Bank of China Limited", "http://nic.icbc/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, n, "whois.nic.icbc", e, w{"https://rdap.zdnsgtld.com/icbc/"}, f}
	z[676] = Zone{"ice", r, x, 0x42, "IntercontinentalExchange, Inc.", "https://www.icann.org/en/registry-agreements/details/ice", w{"a0.nic.ice", "a2.nic.ice", "b0.nic.ice", "c0.nic.ice"}, n, n, n, "whois.nic.ice", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[677] = Zone{"icu", r, x, 0x42, "ShortDot SA", "https://shortdot.bond/icu/", w{"a.nic.icu", "b.nic.icu", "c.nic.icu", "d.nic.icu"}, n, n, w{"ar", "cs", "da", "de", "es", "fr", "he", "is", "it", "ja", "ko", "lb", "lo", "mul-Grek", "mul-Latn", "nl", "pl", "pt", "ro", "ru", "sv", "tr", "uk", "zh"}, "whois.nic.icu", e, w{"https://rdap.centralnic.com/icu/"}, t}
	z[678] = Zone{"id", r, z[2674:2686], 0xa0, e, "https://pandi.id/", w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id", "ns4.apnic.net"}, n, n, n, "whois.id", e, w{"https://rdap.pandi.id/rdap/"}, t}
	z[679] = Zone{"idn", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i679() {
	z[680] = Zone{"ie", r, z[2686:2691], 0xa0, e, "https://www.weare.ie/", w{"a.ns.ie", "c.ns.ie", "d.ns.ie", "e.ns.ie", "g.ns.ie", "h.ns.ie", "i.ns.ie"}, n, n, n, "whois.weare.ie", e, n, t}
	z[681] = Zone{"ieee", r, x, 0x42, "IEEE Global LLC", "http://nic.ieee/", w{"dns1.nic.ieee", "dns2.nic.ieee", "dns3.nic.ieee", "dns4.nic.ieee", "dnsa.nic.ieee", "dnsb.nic.ieee", "dnsc.nic.ieee", "dnsd.nic.ieee"}, n, n, w{"es"}, "whois.nic.ieee", e, w{"https://rdap.nominet.uk/ieee/"}, t}
	z[682] = Zone{"ifm", r, x, 0x42, "ifm electronic gmbh", "https://www.icann.org/en/registry-agreements/details/ifm", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"de", "zh"}, "whois.nic.ifm", e, w{"https://rdap.nic.ifm/"}, t}
	z[683] = Zone{"iinet", r, x, 0x1042, e, e, n, n, n, n, "whois.aridnrs.net.au", e, n, f}
	z[684] = Zone{"ikano", r, x, 0x42, "Ikano S.A.", "https://group.ikano/nic-ikano/", w{"a.dns.nic.ikano", "m.dns.nic.ikano", "n.dns.nic.ikano"}, n, n, w{"mul-Cyrl", "mul-Latn", "mul-Thai"}, "whois.nic.ikano", e, w{"https://rdap.nic.ikano/v1/"}, t}
	z[685] = Zone{"il", r, z[2691:2699], 0xa8, e, "https://www.isoc.org.il/domain-name-registry", w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, w{"he"}, "whois.isoc.org.il", e, n, t}
	z[686] = Zone{"im", r, z[2699:2705], 0xa0, e, e, w{"barney.advsys.co.uk", "hoppy.iom.com", "ns4.ja.net", "pebbles.iom.com"}, n, n, n, "whois.nic.im", e, n, f}
	z[687] = Zone{"imamat", r, x, 0x42, "Fondation Aga Khan (Aga Khan Foundation)", "https://www.icann.org/en/registry-agreements/details/imamat", w{"a0.nic.imamat", "a2.nic.imamat", "b0.nic.imamat", "c0.nic.imamat"}, n, n, n, "whois.nic.imamat", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i687() {
	z[688] = Zone{"imdb", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.imdb/", w{"dns1.nic.imdb", "dns2.nic.imdb", "dns3.nic.imdb", "dns4.nic.imdb", "dnsa.nic.imdb", "dnsb.nic.imdb", "dnsc.nic.imdb", "dnsd.nic.imdb"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.imdb", e, w{"https://rdap.nominet.uk/imdb/"}, t}
	z[689] = Zone{"immo", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.immo", "v0n1.nic.immo", "v0n2.nic.immo", "v0n3.nic.immo", "v2n0.nic.immo", "v2n1.nic.immo"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.immo", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[690] = Zone{"immobilien", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.immobilien", "v0n1.nic.immobilien", "v0n2.nic.immobilien", "v0n3.nic.immobilien", "v2n0.nic.immobilien", "v2n1.nic.immobilien"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.immobilien", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[691] = Zone{"in", r, z[2705:2747], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.registry.in", e, w{"https://rdap.registry.in/"}, f}
	z[692] = Zone{"inc", r, x, 0x40, "Intercap Registry Inc.", "https://www.icann.org/en/registry-agreements/details/inc", w{"a.nic.inc", "b.nic.inc", "c.nic.inc", "d.nic.inc"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.inc", e, w{"https://rdap.centralnic.com/inc/"}, t}
	z[693] = Zone{"indians", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[694] = Zone{"industries", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.industries", "v0n1.nic.industries", "v0n2.nic.industries", "v0n3.nic.industries", "v2n0.nic.industries", "v2n1.nic.industries"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.industries", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[695] = Zone{"infiniti", r, x, 0x42, "NISSAN MOTOR CO., LTD.", "https://www.nissan-global.com/EN/NIC/INFINITI/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i695() {
	z[696] = Zone{"info", r, z[2747:2748], 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.info.afilias-nst.info", "a2.info.afilias-nst.info", "b0.info.afilias-nst.org", "b2.info.afilias-nst.org", "c0.info.afilias-nst.info", "d0.info.afilias-nst.org"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.afilias.net", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[697] = Zone{"infosys", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[698] = Zone{"infy", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[699] = Zone{"ing", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[700] = Zone{"ink", r, x, 0x40, "Registry Services, LLC", e, w{"a.nic.ink", "b.nic.ink", "c.nic.ink", "x.nic.ink", "y.nic.ink", "z.nic.ink"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.ink", e, w{"https://rdap.nic.ink/"}, t}
	z[701] = Zone{"institute", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.institute", "v0n1.nic.institute", "v0n2.nic.institute", "v0n3.nic.institute", "v2n0.nic.institute", "v2n1.nic.institute"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.institute", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[702] = Zone{"insurance", r, x, 0x40, "fTLD Registry Services LLC", "https://www.register.insurance/", w{"d.nic.insurance", "e.nic.insurance", "f.nic.insurance", "w.nic.insurance", "x.nic.insurance", "y.nic.insurance"}, n, n, n, "whois.nic.insurance", e, w{"https://rdap.nic.insurance/"}, f}
	z[703] = Zone{"insure", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.insure", "v0n1.nic.insure", "v0n2.nic.insure", "v0n3.nic.insure", "v2n0.nic.insure", "v2n1.nic.insure"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.insure", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i703() {
	z[704] = Zone{"int", r, z[2748:2749], 0x4040, e, e, w{"ns.uu.net", "ns0.ja.net", "sec2.authdns.ripe.net", "x.iana-servers.net", "y.iana-servers.net", "z.iana-servers.net"}, n, n, n, "whois.iana.org", e, n, f}
	z[705] = Zone{"intel", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.intel", e, n, f}
	z[706] = Zone{"international", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.international", "v0n1.nic.international", "v0n2.nic.international", "v0n3.nic.international", "v2n0.nic.international", "v2n1.nic.international"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.international", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[707] = Zone{"intuit", r, x, 0x42, "Intuit Administrative Services, Inc.", "http://www.nic.intuit/", w{"a.nic.intuit", "b.nic.intuit", "c.nic.intuit", "ns4.dns.nic.intuit", "ns5.dns.nic.intuit", "ns6.dns.nic.intuit"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.intuit", e, w{"https://rdap.nic.intuit/"}, t}
	z[708] = Zone{"invalid", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc6761", n, n, n, n, e, e, n, t}
	z[709] = Zone{"investments", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.investments", "v0n1.nic.investments", "v0n2.nic.investments", "v0n3.nic.investments", "v2n0.nic.investments", "v2n1.nic.investments"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.investments", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[710] = Zone{"io", r, z[2749:2751], 0xe0, e, "https://www.nic.io/", w{"a0.nic.io", "a2.nic.io", "b0.nic.io", "c0.nic.io"}, n, n, n, "whois.nic.io", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[711] = Zone{"ipiranga", r, x, 0x42, "Ipiranga Produtos de Petroleo S.A.", "http://www.nic.ipiranga/", w{"a.nic.ipiranga", "b.nic.ipiranga", "c.nic.ipiranga", "ns4.dns.nic.ipiranga", "ns5.dns.nic.ipiranga", "ns6.dns.nic.ipiranga"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.ipiranga", e, w{"https://rdap.nic.ipiranga/"}, t}
}

//go:noinline
func i711() {
	z[712] = Zone{"iq", r, z[2751:2761], 0xa0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, "whois.cmc.iq", e, n, f}
	z[713] = Zone{"ir", r, z[2761:2768], 0xa0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, w{"fa"}, "whois.nic.ir", e, n, t}
	z[714] = Zone{"ira", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[715] = Zone{"irish", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.irish", "v0n1.nic.irish", "v0n2.nic.irish", "v0n3.nic.irish", "v2n0.nic.irish", "v2n1.nic.irish"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.irish", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[716] = Zone{"is", r, x, 0xa0, e, e, w{"a.nic.is", "b.nic.is", "c.nic.is", "d.nic.is", "durinn.rhnet.is", "e.nic.is", "sunic.sunet.se"}, n, n, n, "whois.isnic.is", e, w{"https://rdap.isnic.is/rdap/"}, t}
	z[717] = Zone{"iselect", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.iselect", e, n, f}
	z[718] = Zone{"islam", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[719] = Zone{"ismaili", r, x, 0x42, "Fondation Aga Khan (Aga Khan Foundation)", "https://www.icann.org/en/registry-agreements/details/ismaili", w{"a0.nic.ismaili", "a2.nic.ismaili", "b0.nic.ismaili", "c0.nic.ismaili"}, n, n, n, "whois.nic.ismaili", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i719() {
	z[720] = Zone{"ist", r, x, 0xc4, "Istanbul Metropolitan Municipality", "https://www.icann.org/en/registry-agreements/details/ist", w{"a0.nic.ist", "a2.nic.ist", "b0.nic.ist", "c0.nic.ist"}, n, w{"Istanbul", "TR-34"}, w{"tr"}, "whois.nic.ist", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[721] = Zone{"istanbul", r, x, 0xc4, "Istanbul Metropolitan Municipality", "https://nic.istanbul/", w{"a0.nic.istanbul", "a2.nic.istanbul", "b0.nic.istanbul", "c0.nic.istanbul"}, n, w{"Istanbul", "TR-34"}, w{"tr"}, "whois.nic.istanbul", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[722] = Zone{"it", r, z[2768:3129], 0xa0, e, "https://www.nic.it/it", w{"a.dns.it", "d.dns.it", "dns.nic.it", "m.dns.it", "nameserver.cnr.it", "r.dns.it"}, n, n, n, "whois.nic.it", e, n, t}
	z[723] = Zone{"itau", r, x, 0x42, "Itau Unibanco Holding S.A.", "http://www.nic.itau/", w{"a.nic.itau", "b.nic.itau", "c.nic.itau", "ns4.dns.nic.itau", "ns5.dns.nic.itau", "ns6.dns.nic.itau"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.itau", e, w{"https://rdap.nic.itau/"}, t}
	z[724] = Zone{"itv", r, x, 0x42, "ITV Services Limited", "http://nic.itv/", w{"a0.nic.itv", "a2.nic.itv", "b0.nic.itv", "c0.nic.itv"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.itv", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[725] = Zone{"iveco", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.iveco", e, w{"https://rdap.afilias-srs.net/rdap/iveco/"}, f}
	z[726] = Zone{"iwc", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.iwc", e, n, f}
	z[727] = Zone{"jaguar", r, x, 0x42, "Jaguar Land Rover Ltd", "http://nic.jaguar/", w{"a0.nic.jaguar", "a2.nic.jaguar", "b0.nic.jaguar", "c0.nic.jaguar"}, n, n, n, "whois.nic.jaguar", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i727() {
	z[728] = Zone{"java", r, x, 0x42, "Oracle Corporation", "https://www.icann.org/en/registry-agreements/details/java", w{"a0.nic.java", "a2.nic.java", "b0.nic.java", "c0.nic.java"}, n, n, n, "whois.nic.java", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[729] = Zone{"jcb", r, x, 0x42, "JCB Co., Ltd.", "https://nic.jcb/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[730] = Zone{"jcp", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[731] = Zone{"je", r, z[3129:3135], 0xa0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "e.ci-servers.net"}, n, n, n, "whois.je", e, n, f}
	z[732] = Zone{"jeep", r, x, 0x42, "FCA US LLC.", "https://www.icann.org/en/registry-agreements/details/jeep", w{"a0.nic.jeep", "a2.nic.jeep", "b0.nic.jeep", "c0.nic.jeep"}, n, n, n, "whois.nic.jeep", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[733] = Zone{"jetzt", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.jetzt", "v0n1.nic.jetzt", "v0n2.nic.jetzt", "v0n3.nic.jetzt", "v2n0.nic.jetzt", "v2n1.nic.jetzt"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.jetzt", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[734] = Zone{"jewelry", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.jewelry", "v0n1.nic.jewelry", "v0n2.nic.jewelry", "v0n3.nic.jewelry", "v2n0.nic.jewelry", "v2n1.nic.jewelry"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.jewelry", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[735] = Zone{"jio", r, x, 0x42, "Reliance Industries Limited", "http://nic.jio/", w{"a0.nic.jio", "a2.nic.jio", "b0.nic.jio", "c0.nic.jio"}, n, n, n, "whois.nic.jio", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i735() {
	z[736] = Zone{"jlc", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.jlc", e, n, f}
	z[737] = Zone{"jll", r, x, 0x42, "Jones Lang LaSalle Incorporated", "https://www.icann.org/en/registry-agreements/details/jll", w{"a0.nic.jll", "a2.nic.jll", "b0.nic.jll", "c0.nic.jll"}, n, n, n, "whois.nic.jll", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[738] = Zone{"jm", r, z[3135:3141], 0xa8, e, e, w{"jm.cctld.authdns.ripe.net", "ns.jm", "ns.utechjamaica.edu.jm", "ns3-jm.fsl.org.jm", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[739] = Zone{"jmp", r, x, 0x42, "Matrix IP LLC", "https://nic.jmp/", w{"a.nic.jmp", "b.nic.jmp", "c.nic.jmp", "ns1.dns.nic.jmp", "ns2.dns.nic.jmp", "ns3.dns.nic.jmp"}, n, n, w{"es"}, "whois.nic.jmp", e, w{"https://rdap.nic.jmp/"}, t}
	z[740] = Zone{"jnj", r, x, 0x42, "Johnson & Johnson Services, Inc.", "http://www.nic.jnj/", w{"a.nic.jnj", "b.nic.jnj", "c.nic.jnj", "ns1.dns.nic.jnj", "ns2.dns.nic.jnj", "ns3.dns.nic.jnj"}, n, n, w{"es"}, "whois.nic.jnj", e, w{"https://rdap.nic.jnj/"}, t}
	z[741] = Zone{"jo", r, z[3141:3149], 0xa0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, "http://www.dns.jo/Whois.aspx", n, f}
	z[742] = Zone{"jobs", r, x, 0x4040, "Employ Media LLC", e, w{"dns1.nic.jobs", "dns2.nic.jobs", "dns3.nic.jobs", "dns4.nic.jobs", "dnsa.nic.jobs", "dnsb.nic.jobs", "dnsc.nic.jobs", "dnsd.nic.jobs"}, n, n, n, "whois.nic.jobs", e, w{"https://rdap.nominet.uk/jobs/"}, f}
	z[743] = Zone{"joburg", r, x, 0xc4, "ZA Central Registry NPC trading as ZA Central Registry", e, w{"coza1.dnsnode.net", "nsp.netnod.se"}, n, w{"Johannesburg"}, n, "whois.nic.joburg", e, w{"https://rdap.nic.joburg/rdap/"}, t}
}

//go:noinline
func i743() {
	z[744] = Zone{"jot", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.jot/", w{"dns1.nic.jot", "dns2.nic.jot", "dns3.nic.jot", "dns4.nic.jot", "dnsa.nic.jot", "dnsb.nic.jot", "dnsc.nic.jot", "dnsd.nic.jot"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.jot", e, w{"https://rdap.nominet.uk/jot/"}, t}
	z[745] = Zone{"joy", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.joy/", w{"dns1.nic.joy", "dns2.nic.joy", "dns3.nic.joy", "dns4.nic.joy", "dnsa.nic.joy", "dnsb.nic.joy", "dnsc.nic.joy", "dnsd.nic.joy"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.joy", e, w{"https://rdap.nominet.uk/joy/"}, t}
	z[746] = Zone{"jp", r, z[3149:3252], 0xa0, e, e, w{"a.dns.jp", "b.dns.jp", "c.dns.jp", "d.dns.jp", "e.dns.jp", "f.dns.jp", "g.dns.jp", "h.dns.jp"}, n, n, w{"ja-JP"}, "whois.jprs.jp", e, n, t}
	z[747] = Zone{"jpmorgan", r, x, 0x42, "JPMorgan Chase Bank, National Association", "https://www.nic.jpmorgan/", w{"a.nic.jpmorgan", "b.nic.jpmorgan", "c.nic.jpmorgan", "ns4.dns.nic.jpmorgan", "ns5.dns.nic.jpmorgan", "ns6.dns.nic.jpmorgan"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.jpmorgan", e, w{"https://rdap.nic.jpmorgan/"}, t}
	z[748] = Zone{"jpmorganchase", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[749] = Zone{"jprs", r, x, 0x42, "Japan Registry Services Co., Ltd.", "https://nic.jprs/", w{"tld1.nic.jprs", "tld2.nic.jprs", "tld3.nic.jprs", "tld4.nic.jprs", "tld5.nic.jprs"}, n, n, w{"ja"}, e, e, w{"https://rdap.nic.jprs/rdap/"}, t}
	z[750] = Zone{"juegos", r, x, 0x40, "Dog Beach, LLC", "https://uniregistry.link/extensions/juegos/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.uniregistry.net", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[751] = Zone{"juniper", r, x, 0x42, "JUNIPER NETWORKS, INC.", "http://nic.juniper/", w{"a0.nic.juniper", "a2.nic.juniper", "b0.nic.juniper", "c0.nic.juniper"}, n, n, n, "whois.nic.juniper", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i751() {
	z[752] = Zone{"justforu", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[753] = Zone{"kaufen", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.kaufen", "v0n1.nic.kaufen", "v0n2.nic.kaufen", "v0n3.nic.kaufen", "v2n0.nic.kaufen", "v2n1.nic.kaufen"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.kaufen", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[754] = Zone{"kddi", r, x, 0x42, "KDDI CORPORATION", "https://nic.kddi/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.kddi", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[755] = Zone{"ke", r, z[3252:3261], 0xa8, e, "https://kenic.or.ke/", w{"kenic.anycastdns.cz", "mzizi.kenic.or.ke", "ns-ke.afrinic.net", "ns.anycast.kenic.or.ke", "ns2ke.dns.business"}, n, n, n, "whois.kenic.or.ke", e, n, f}
	z[756] = Zone{"kerastase", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[757] = Zone{"kerryhotels", r, x, 0x42, "Kerry Trading Co. Limited", "https://www.icann.org/en/registry-agreements/details/kerryhotels", w{"a0.nic.kerryhotels", "a2.nic.kerryhotels", "b0.nic.kerryhotels", "c0.nic.kerryhotels"}, n, n, n, "whois.nic.kerryhotels", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[758] = Zone{"kerrylogisitics", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[759] = Zone{"kerrylogistics", r, x, 0, "Kerry Trading Co. Limited", e, w{"a0.nic.kerrylogistics", "a2.nic.kerrylogistics", "b0.nic.kerrylogistics", "c0.nic.kerrylogistics"}, n, n, n, "whois.nic.kerrylogistics", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i759() {
	z[760] = Zone{"kerryproperties", r, x, 0x42, "Kerry Trading Co. Limited", "https://www.icann.org/en/registry-agreements/details/kerryproperties", w{"a0.nic.kerryproperties", "a2.nic.kerryproperties", "b0.nic.kerryproperties", "c0.nic.kerryproperties"}, n, n, n, "whois.nic.kerryproperties", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[761] = Zone{"ketchup", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[762] = Zone{"kfh", r, x, 0x42, "Kuwait Finance House", "https://nic.kfh/", w{"a.nic.kfh", "b.nic.kfh", "c.nic.kfh", "d.nic.kfh"}, n, n, n, "whois.nic.kfh", e, w{"https://rdap.centralnic.com/kfh/"}, f}
	z[763] = Zone{"kg", r, z[3261:3267], 0xa0, e, e, w{"kg.cctld.authdns.ripe.net", "ns.kg", "ns2.kg"}, n, n, n, "whois.kg", e, w{"http://rdap.cctld.kg/"}, f}
	z[764] = Zone{"kh", r, z[3267:3274], 0xa0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[765] = Zone{"ki", r, z[3274:3286], 0xa0, e, e, w{"ns.cocca.fr", "pch.nic.ki"}, n, n, n, "whois.nic.ki", e, n, f}
	z[766] = Zone{"kia", r, x, 0x42, "KIA MOTORS CORPORATION", "https://nic.kia/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ko"}, "whois.nic.kia", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[767] = Zone{"kid", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i767() {
	z[768] = Zone{"kids", r, x, 0x40, "DotKids Foundation Limited", "https://nic.kids/", w{"a0.nic.kids", "a2.nic.kids", "b0.nic.kids", "c0.nic.kids"}, n, n, w{"be", "bg", "bs-Cyrl", "cnr", "da", "de", "es", "fi", "fr", "hr", "hu", "is", "it", "ko", "lt", "lv", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.kids", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[769] = Zone{"kiehls", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[770] = Zone{"kim", r, x, 0x40, "Identity Digital Limited", "http://nic.kim/", w{"a0.nic.kim", "a2.nic.kim", "b0.nic.kim", "c0.nic.kim"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.kim", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[771] = Zone{"kinder", r, x, 0x1042, "Ferrero Trading Lux S.A.", "https://nic.kinder/", n, n, n, w{"da", "de", "es", "fi", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.kinder", e, w{"https://rdap.nic.kinder/"}, t}
	z[772] = Zone{"kindle", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.kindle/", w{"dns1.nic.kindle", "dns2.nic.kindle", "dns3.nic.kindle", "dns4.nic.kindle", "dnsa.nic.kindle", "dnsb.nic.kindle", "dnsc.nic.kindle", "dnsd.nic.kindle"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.kindle", e, w{"https://rdap.nominet.uk/kindle/"}, t}
	z[773] = Zone{"kitchen", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.kitchen", "v0n1.nic.kitchen", "v0n2.nic.kitchen", "v0n3.nic.kitchen", "v2n0.nic.kitchen", "v2n1.nic.kitchen"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.kitchen", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[774] = Zone{"kiwi", r, x, 0xe0, "DOT KIWI LIMITED", "https://hello.kiwi/", w{"a.ns.nic.kiwi", "b.ns.nic.kiwi"}, n, n, w{"mul-Latn"}, "whois.nic.kiwi", e, w{"https://rdap.kiwi.fury.ca/rdap/"}, t}
	z[775] = Zone{"km", r, z[3286:3300], 0xa0, e, e, w{"dns1.nic.km", "dns2.nic.km", "ns-km.afrinic.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i775() {
	z[776] = Zone{"kn", r, z[3300:3306], 0xa0, e, "https://www.nic.kn/", w{"ns1.anycastdns.cz", "ns2.anycastdns.cz", "pch.nic.kn"}, n, n, n, "whois.nic.kn", e, n, f}
	z[777] = Zone{"koeln", r, x, 0xc4, "dotKoeln GmbH", e, w{"dns.ryce-rsp.com", "ns1.dns.business", "ns1.ryce-rsp.com"}, n, w{"Cologne", "Koeln"}, w{"de", "mul-Latn"}, "whois.ryce-rsp.com", e, w{"https://rdap.ryce-rsp.com/rdap/"}, t}
	z[778] = Zone{"komatsu", r, x, 0x42, "Komatsu Ltd.", "https://nic.komatsu/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.komatsu", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[779] = Zone{"konami", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[780] = Zone{"kone", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[781] = Zone{"kosher", r, x, 0x40, "Kosher Marketing Assets LLC", "https://nic.kosher/", w{"a0.nic.kosher", "a2.nic.kosher", "b0.nic.kosher", "c0.nic.kosher"}, n, n, n, "whois.nic.kosher", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[782] = Zone{"kp", r, z[3306:3310], 0xa0, e, e, w{"ns1.kptc.kp", "ns2.kptc.kp"}, n, n, n, e, e, n, f}
	z[783] = Zone{"kpmg", r, x, 0x42, "KPMG International Cooperative (KPMG International Genossenschaft)", "https://www.nic.kpmg/index.html", w{"a.nic.kpmg", "b.nic.kpmg", "c.nic.kpmg", "ns1.dns.nic.kpmg", "ns2.dns.nic.kpmg", "ns3.dns.nic.kpmg"}, n, n, w{"de", "es", "ja", "ko", "zh"}, "whois.nic.kpmg", e, w{"https://rdap.nic.kpmg/"}, t}
}

//go:noinline
func i783() {
	z[784] = Zone{"kpn", r, x, 0x42, "Koninklijke KPN N.V.", "https://nic.kpn/", w{"a.nic.kpn", "b.nic.kpn", "c.nic.kpn", "d.nic.kpn"}, n, n, n, e, e, w{"https://rdap.centralnic.com/kpn/"}, f}
	z[785] = Zone{"kr", r, z[3310:3339], 0xa0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, w{"ko-KR"}, "whois.kr", e, n, t}
	z[786] = Zone{"krd", r, x, 0x440, "KRG Department of Information Technology", "http://nic.krd/", w{"a.nic.krd", "b.nic.krd", "c.nic.krd", "x.nic.krd", "y.nic.krd", "z.nic.krd"}, n, n, n, "whois.nic.krd", e, w{"https://rdap.nic.krd/"}, f}
	z[787] = Zone{"kred", r, x, 0x50, "KredTLD Pty Ltd", "https://www.nic.kred/", w{"a.nic.kred", "b.nic.kred", "c.nic.kred", "d.nic.kred"}, n, n, w{"ar", "da", "de", "es", "fr", "ja", "ko", "nl", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.kred", e, w{"https://rdap.centralnic.com/kred/"}, t}
	z[788] = Zone{"kuokgroup", r, x, 0x42, "Kerry Trading Co. Limited", "https://www.icann.org/en/registry-agreements/details/kuokgroup", w{"a0.nic.kuokgroup", "a2.nic.kuokgroup", "b0.nic.kuokgroup", "c0.nic.kuokgroup"}, n, n, n, "whois.nic.kuokgroup", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[789] = Zone{"kw", r, z[3339:3344], 0xa0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, "http://www.kw/", n, f}
	z[790] = Zone{"ky", r, z[3344:3349], 0xa0, e, "https://uniregistry.ky/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, "whois.kyregistry.ky", e, w{"https://whois.kyregistry.ky/rdap/"}, f}
	z[791] = Zone{"kyknet", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i791() {
	z[792] = Zone{"kyoto", r, x, 0xc4, "Academic Institution: Kyoto Jyoho Gakuen", "https://nic.kyoto/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"Kyoto", "JP-26"}, w{"ja"}, "whois.nic.kyoto", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[793] = Zone{"kz", r, z[3349:3355], 0xa0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, "whois.nic.kz", e, n, f}
	z[794] = Zone{"la", r, x, 0xa4, e, e, w{"ns1.nic.la", "ns2.nic.la", "ns3.nic.la", "ns4.nic.la"}, n, w{"Los Angeles"}, w{"ja", "lo", "mul-Grek", "mul-Latn", "mul-Tavt", "th", "zh"}, "whois.nic.la", e, n, t}
	z[795] = Zone{"lacaixa", r, x, 0x42, "Fundación Bancaria Caixa d’Estalvis i Pensions de Barcelona, “la Caixa”", "http://www.nic.lacaixa/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.lacaixa", e, w{"https://rdap.nic.lacaixa/"}, t}
	z[796] = Zone{"ladbrokes", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.ladbrokes", e, n, t}
	z[797] = Zone{"lamborghini", r, x, 0x42, "Automobili Lamborghini S.p.A.", "http://nic.lamborghini/", w{"a0.nic.lamborghini", "a2.nic.lamborghini", "b0.nic.lamborghini", "c0.nic.lamborghini"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.lamborghini", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[798] = Zone{"lamer", r, x, 0x42, "The Estée Lauder Companies Inc.", "https://www.icann.org/en/registry-agreements/details/lamer", w{"a0.nic.lamer", "a2.nic.lamer", "b0.nic.lamer", "c0.nic.lamer"}, n, n, n, "whois.nic.lamer", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[799] = Zone{"lancaster", r, x, 0x42, "LANCASTER", "https://nic.lancaster/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.lancaster", e, w{"https://rdap.nic.lancaster/"}, t}
}

//go:noinline
func i799() {
	z[800] = Zone{"lancia", r, x, 0x1042, "Fiat Chrysler Automobiles N.V.", "http://nic.lancia/", n, n, n, n, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/lancia/"}, f}
	z[801] = Zone{"lancome", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.lancome", e, n, f}
	z[802] = Zone{"land", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.land", "v0n1.nic.land", "v0n2.nic.land", "v0n3.nic.land", "v2n0.nic.land", "v2n1.nic.land"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.land", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[803] = Zone{"landrover", r, x, 0x42, "Jaguar Land Rover Ltd", "http://nic.landrover/", w{"a0.nic.landrover", "a2.nic.landrover", "b0.nic.landrover", "c0.nic.landrover"}, n, n, n, "whois.nic.landrover", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[804] = Zone{"lanxess", r, x, 0x42, "LANXESS Corporation", "https://nic.lanxess/", w{"a.nic.lanxess", "b.nic.lanxess", "c.nic.lanxess", "ns1.dns.nic.lanxess", "ns2.dns.nic.lanxess", "ns3.dns.nic.lanxess"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.lanxess", e, w{"https://rdap.nic.lanxess/"}, t}
	z[805] = Zone{"lasalle", r, x, 0x42, "Jones Lang LaSalle Incorporated", "https://www.lasalle.com/tld", w{"a0.nic.lasalle", "a2.nic.lasalle", "b0.nic.lasalle", "c0.nic.lasalle"}, n, n, n, "whois.nic.lasalle", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[806] = Zone{"lat", r, x, 0x40, "XYZ.COM LLC", "https://nic.lat/", w{"a.nic.lat", "b.nic.lat", "c.nic.lat", "d.nic.lat"}, n, n, w{"es"}, "whois.nic.lat", e, w{"https://rdap.centralnic.com/lat/"}, t}
	z[807] = Zone{"latino", r, x, 0x40, "Dish DBS Corporation", "https://www.dishtlds.com/latino/", w{"a0.nic.latino", "a2.nic.latino", "b0.nic.latino", "c0.nic.latino"}, n, n, n, "whois.nic.latino", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i807() {
	z[808] = Zone{"latrobe", r, x, 0x42, "La Trobe University", "https://www.icann.org/en/registry-agreements/details/latrobe", w{"a.nic.latrobe", "b.nic.latrobe", "c.nic.latrobe", "x.nic.latrobe", "y.nic.latrobe", "z.nic.latrobe"}, n, n, n, "whois.nic.latrobe", e, w{"https://rdap.nic.latrobe/"}, f}
	z[809] = Zone{"law", r, x, 0x40, "Registry Services, LLC", "https://nic.law/", w{"a.nic.law", "b.nic.law", "c.nic.law", "x.nic.law", "y.nic.law", "z.nic.law"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.law", e, w{"https://rdap.nic.law/"}, t}
	z[810] = Zone{"lawyer", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.lawyer", "v0n1.nic.lawyer", "v0n2.nic.lawyer", "v0n3.nic.lawyer", "v2n0.nic.lawyer", "v2n1.nic.lawyer"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.lawyer", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[811] = Zone{"lb", r, z[3355:3360], 0xa8, e, e, w{"fork.sth.dnsnode.net", "nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, "whois.lbdr.org.lb", e, n, f}
	z[812] = Zone{"lc", r, z[3360:3369], 0xa0, e, "https://www.nic.lc/", w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, "whois.afilias-grs.info", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[813] = Zone{"lds", r, x, 0x42, "IRI Domain Management, LLC", "https://www.icann.org/en/registry-agreements/details/lds", w{"a0.nic.lds", "a2.nic.lds", "b0.nic.lds", "c0.nic.lds"}, n, n, n, "whois.nic.lds", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[814] = Zone{"lease", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.lease", "v0n1.nic.lease", "v0n2.nic.lease", "v0n3.nic.lease", "v2n0.nic.lease", "v2n1.nic.lease"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.lease", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[815] = Zone{"leclerc", r, x, 0x42, "A.C.D. LEC Association des Centres Distributeurs Edouard Leclerc", "https://www.nic.leclerc/fr", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.leclerc", e, w{"https://rdap.nic.leclerc/"}, t}
}

//go:noinline
func i815() {
	z[816] = Zone{"lefrak", r, x, 0x42, "LeFrak Organization, Inc.", "http://nic.lefrak/", w{"a0.nic.lefrak", "a2.nic.lefrak", "b0.nic.lefrak", "c0.nic.lefrak"}, n, n, n, "whois.nic.lefrak", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[817] = Zone{"legal", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.legal", "v0n1.nic.legal", "v0n2.nic.legal", "v0n3.nic.legal", "v2n0.nic.legal", "v2n1.nic.legal"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.legal", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[818] = Zone{"lego", r, x, 0x42, "LEGO Juris A/S", "https://nic.lego/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"mul-Arab", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Latn", "ru"}, "whois.nic.lego", e, w{"https://tld-rdap.verisign.com/lego/v1/"}, t}
	z[819] = Zone{"lexus", r, x, 0x42, "TOYOTA MOTOR CORPORATION", "http://www.nic.lexus/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.lexus", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[820] = Zone{"lgbt", r, x, 0x40, "Identity Digital Limited", "http://nic.lgbt/", w{"a0.nic.lgbt", "a2.nic.lgbt", "b0.nic.lgbt", "c0.nic.lgbt"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.lgbt", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[821] = Zone{"li", r, x, 0xa0, e, e, w{"a.nic.li", "b.nic.li", "d.nic.li", "e.nic.li", "f.nic.li"}, n, n, n, "whois.nic.li", e, w{"https://rdap.nic.li/"}, t}
	z[822] = Zone{"liaison", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.liaison", e, n, f}
	z[823] = Zone{"lidl", r, x, 0x42, "Schwarz Domains und Services GmbH & Co. KG", "https://nic.lidl/", w{"a.nic.lidl", "b.nic.lidl", "c.nic.lidl", "d.nic.lidl"}, n, n, w{"mul-Cyrl", "mul-Grek", "mul-Latn"}, "whois.nic.lidl", e, w{"https://rdap.centralnic.com/lidl/"}, t}
}

//go:noinline
func i823() {
	z[824] = Zone{"life", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.life", "v0n1.nic.life", "v0n2.nic.life", "v0n3.nic.life", "v2n0.nic.life", "v2n1.nic.life"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.life", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[825] = Zone{"lifeinsurance", r, x, 0x40, "American Council of Life Insurers", "http://nic.lifeinsurance/", w{"a.nic.lifeinsurance", "b.nic.lifeinsurance", "c.nic.lifeinsurance", "ns1.dns.nic.lifeinsurance", "ns2.dns.nic.lifeinsurance", "ns3.dns.nic.lifeinsurance"}, n, n, w{"es"}, "whois.nic.lifeinsurance", e, w{"https://rdap.nic.lifeinsurance/"}, t}
	z[826] = Zone{"lifestyle", r, x, 0x40, "Internet Naming Company LLC", "https://www.icann.org/en/registry-agreements/details/lifestyle", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.lifestyle", e, w{"https://tld-rdap.verisign.com/lifestyle/v1/"}, f}
	z[827] = Zone{"lighting", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.lighting", "v0n1.nic.lighting", "v0n2.nic.lighting", "v0n3.nic.lighting", "v2n0.nic.lighting", "v2n1.nic.lighting"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.lighting", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[828] = Zone{"like", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.like/", w{"dns1.nic.like", "dns2.nic.like", "dns3.nic.like", "dns4.nic.like", "dnsa.nic.like", "dnsb.nic.like", "dnsc.nic.like", "dnsd.nic.like"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.like", e, w{"https://rdap.nominet.uk/like/"}, t}
	z[829] = Zone{"lilly", r, x, 0x42, "Eli Lilly and Company", "https://www.lilly.com/lilly-global-top-level-domain", w{"a.nic.lilly", "b.nic.lilly", "c.nic.lilly", "ns1.dns.nic.lilly", "ns2.dns.nic.lilly", "ns3.dns.nic.lilly"}, n, n, w{"es"}, "whois.nic.lilly", e, w{"https://rdap.nic.lilly/"}, t}
	z[830] = Zone{"limited", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.limited", "v0n1.nic.limited", "v0n2.nic.limited", "v0n3.nic.limited", "v2n0.nic.limited", "v2n1.nic.limited"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.limited", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[831] = Zone{"limo", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.limo", "v0n1.nic.limo", "v0n2.nic.limo", "v0n3.nic.limo", "v2n0.nic.limo", "v2n1.nic.limo"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.limo", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i831() {
	z[832] = Zone{"lincoln", r, x, 0x42, "Ford Motor Company", "https://nic.lincoln/", w{"a.nic.lincoln", "b.nic.lincoln", "c.nic.lincoln", "ns1.dns.nic.lincoln", "ns2.dns.nic.lincoln", "ns3.dns.nic.lincoln"}, n, n, w{"da", "fr", "pl", "ru", "sv", "zh"}, "whois.nic.lincoln", e, w{"https://rdap.nic.lincoln/"}, t}
	z[833] = Zone{"linde", r, x, 0x1042, "Linde Aktiengesellschaft", "https://www.icann.org/en/registry-agreements/details/linde", n, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.linde", e, w{"https://tld-rdap.verisign.com/linde/v1/"}, t}
	z[834] = Zone{"link", r, x, 0x40, "Nova Registry Ltd", "https://nova.link/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.uniregistry.net", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[835] = Zone{"lipsy", r, x, 0x42, "Lipsy Ltd", "http://nic.lipsy/", w{"a0.nic.lipsy", "a2.nic.lipsy", "b0.nic.lipsy", "c0.nic.lipsy"}, n, n, n, "whois.nic.lipsy", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[836] = Zone{"live", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.live", "v0n1.nic.live", "v0n2.nic.live", "v0n3.nic.live", "v2n0.nic.live", "v2n1.nic.live"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.live", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[837] = Zone{"livestrong", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[838] = Zone{"living", r, x, 0x40, "Internet Naming Company LLC", "https://www.icann.org/en/registry-agreements/details/living", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, e, e, w{"https://tld-rdap.verisign.com/living/v1/"}, f}
	z[839] = Zone{"lixil", r, x, 0x1042, "LIXIL Group Corporation", e, n, n, n, w{"ja"}, "whois.nic.lixil", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i839() {
	z[840] = Zone{"lk", r, z[3369:3383], 0xa0, e, e, w{"c.nic.lk", "d.nic.lk", "m.nic.lk", "ns1.ac.lk", "p.nic.lk", "pendragon.cs.purdue.edu", "t.nic.lk"}, n, n, n, "whois.nic.lk", e, n, f}
	z[841] = Zone{"llc", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.llc", "a2.nic.llc", "b0.nic.llc", "c0.nic.llc"}, n, n, n, "whois.nic.llc", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[842] = Zone{"llp", r, x, 0x40, "Intercap Registry Inc.", "https://uniregistry.link/extensions/llp/", w{"a.nic.llp", "b.nic.llp", "c.nic.llp", "d.nic.llp"}, n, n, w{"es", "ja"}, "whois.nic.llp", e, w{"https://rdap.centralnic.com/llp/"}, t}
	z[843] = Zone{"loan", r, x, 0x40, "dot Loan Limited", "http://nic.loan/", w{"a.nic.loan", "b.nic.loan", "c.nic.loan", "ns1.dns.nic.loan", "ns2.dns.nic.loan", "ns3.dns.nic.loan"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.loan", e, w{"https://rdap.nic.loan/"}, t}
	z[844] = Zone{"loans", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.loans", "v0n1.nic.loans", "v0n2.nic.loans", "v0n3.nic.loans", "v2n0.nic.loans", "v2n1.nic.loans"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.loans", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[845] = Zone{"local", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[846] = Zone{"localhost", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc6761", n, w{"127.0.0.1", "::1"}, n, n, e, e, n, t}
	z[847] = Zone{"locker", r, x, 0x42, "Orange Domains LLC", "https://dishtlds.com/locker/", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.locker", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i847() {
	z[848] = Zone{"locus", r, x, 0x42, "Locus Analytics LLC", "http://nic.locus/", w{"dns1.nic.locus", "dns2.nic.locus", "dns3.nic.locus", "dns4.nic.locus", "dnsa.nic.locus", "dnsb.nic.locus", "dnsc.nic.locus", "dnsd.nic.locus"}, n, n, n, "whois.nic.locus", e, w{"https://rdap.nominet.uk/locus/"}, f}
	z[849] = Zone{"loft", r, x, 0x1042, "Annco, Inc.", "http://www.nic.loft/", n, n, n, n, "whois.nic.loft", e, w{"https://rdap.nic.loft/"}, f}
	z[850] = Zone{"lol", r, x, 0x40, "XYZ.COM LLC", "https://nic.lol/", w{"a.nic.lol", "b.nic.lol", "c.nic.lol", "d.nic.lol"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.lol", e, w{"https://rdap.centralnic.com/lol/"}, t}
	z[851] = Zone{"london", r, x, 0xc4, "Dot London Domains Limited", "http://nic.london/", w{"a.nic.london", "b.nic.london", "c.nic.london", "d.nic.london"}, n, w{"London", "GB-LND", "GB-BDG", "GB-BNE", "GB-BEX", "GB-BEN", "GB-BRY", "GB-CMD", "GB-CRY", "GB-EAL", "GB-ENF", "GB-GRE", "GB-HCK", "GB-HMF", "GB-HRY", "GB-HRW", "GB-HAV", "GB-HIL", "GB-HNS", "GB-ISL", "GB-KEC", "GB-KTT", "GB-LBH", "GB-LEW", "GB-MRT", "GB-NWM", "GB-RDB", "GB-RIC", "GB-SWK", "GB-STN", "GB-TWH", "GB-WFT", "GB-WND", "GB-WSM"}, w{"de", "es", "fr", "it"}, "whois.nic.london", e, w{"https://rdap.centralnic.com/london/"}, t}
	z[852] = Zone{"loreal", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[853] = Zone{"lotte", r, x, 0x42, "Lotte Holdings Co., Ltd.", "https://nic.lotte/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.lotte", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[854] = Zone{"lotto", r, x, 0x40, "Identity Digital Limited", "http://nic.lotto/", w{"a0.nic.lotto", "a2.nic.lotto", "b0.nic.lotto", "c0.nic.lotto"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.lotto", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[855] = Zone{"love", r, x, 0x40, "Merchant Law Group LLP", "https://get.love/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.love", e, w{"https://rdap.registry.love/rdap/"}, t}
}

//go:noinline
func i855() {
	z[856] = Zone{"lpl", r, x, 0x42, "LPL Holdings, Inc.", "https://www.nic.lpl/", w{"a.nic.lpl", "b.nic.lpl", "c.nic.lpl", "d.nic.lpl"}, n, n, w{"es"}, "whois.nic.lpl", e, w{"https://rdap.centralnic.com/lpl/"}, t}
	z[857] = Zone{"lplfinancial", r, x, 0x42, "LPL Holdings, Inc.", "https://www.nic.lplfinancial/", w{"a.nic.lplfinancial", "b.nic.lplfinancial", "c.nic.lplfinancial", "d.nic.lplfinancial"}, n, n, w{"es"}, "whois.nic.lplfinancial", e, w{"https://rdap.centralnic.com/lplfinancial/"}, t}
	z[858] = Zone{"lr", r, z[3383:3389], 0xa8, e, e, w{"fork.sth.dnsnode.net", "ns-lr.afrinic.net", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[859] = Zone{"ls", r, z[3389:3395], 0xa0, e, "https://www.nic.ls/", w{"ls-ns.anycast.pch.net", "ns-ls.afrinic.net", "ns1.nic.ls", "ns2.nic.ls"}, n, n, n, "whois.nic.ls", e, n, f}
	z[860] = Zone{"lt", r, z[3395:3396], 0xa0, e, e, w{"a.tld.lt", "b.tld.lt", "c.tld.lt", "d.tld.lt", "e.tld.lt", "f.tld.lt"}, n, n, n, "whois.domreg.lt", e, n, t}
	z[861] = Zone{"ltd", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.ltd", "v0n1.nic.ltd", "v0n2.nic.ltd", "v0n3.nic.ltd", "v2n0.nic.ltd", "v2n1.nic.ltd"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.ltd", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[862] = Zone{"ltda", r, x, 0x40, "InterNetX, Corp", "http://nic.ltda/", w{"a0.nic.ltda", "a2.nic.ltda", "b0.nic.ltda", "c0.nic.ltda"}, n, n, n, "whois.nic.ltda", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[863] = Zone{"lu", r, x, 0xa0, e, e, w{"g.dns.lu", "i.dns.lu", "j.dns.lu", "k.dns.lu", "ns1.dns.lu", "p.dns.lu"}, n, n, n, "whois.dns.lu", e, n, t}
}

//go:noinline
func i863() {
	z[864] = Zone{"lundbeck", r, x, 0x42, "H. Lundbeck A/S", "https://www.lundbeck.com/global/nic", w{"a0.nic.lundbeck", "a2.nic.lundbeck", "b0.nic.lundbeck", "c0.nic.lundbeck"}, n, n, n, "whois.nic.lundbeck", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[865] = Zone{"lupin", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[866] = Zone{"luxe", r, x, 0x40, "Registry Services, LLC", "https://nic.luxe/", w{"a.nic.luxe", "b.nic.luxe", "c.nic.luxe", "x.nic.luxe", "y.nic.luxe", "z.nic.luxe"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "he", "it", "no", "ru", "sv", "zh"}, "whois.nic.luxe", e, w{"https://rdap.nic.luxe/"}, t}
	z[867] = Zone{"luxury", r, x, 0x40, "Luxury Partners, LLC", "https://nic.luxury/", w{"a.nic.luxury", "b.nic.luxury", "c.nic.luxury", "d.nic.luxury"}, n, n, n, "whois.nic.luxury", e, w{"https://rdap.centralnic.com/luxury/"}, f}
	z[868] = Zone{"lv", r, z[3396:3405], 0xa0, e, e, w{"a.nic.lv", "b.nic.lv", "d.nic.lv", "n.nic.lv", "nu.nic.lv", "sunic.sunet.se"}, n, n, n, "whois.nic.lv", e, n, t}
	z[869] = Zone{"ly", r, z[3405:3414], 0xa0, e, "https://nic.ly/", w{"dns.lttnet.net", "dns1.lttnet.net", "ns-ly.afrinic.net", "pch.ltt.ly", "phloem.uoregon.edu"}, n, n, n, "whois.nic.ly", e, n, f}
	z[870] = Zone{"ma", r, z[3414:3420], 0xa0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, "whois.registre.ma", e, n, f}
	z[871] = Zone{"macys", r, x, 0x1042, "Macys, Inc.", "https://www.icann.org/en/registry-agreements/details/macys", n, n, n, n, "whois.nic.macys", e, w{"https://tld-rdap.verisign.com/macys/v1/"}, f}
}

//go:noinline
func i871() {
	z[872] = Zone{"madrid", r, x, 0xc4, "Comunidad de Madrid", "https://dominio.madrid/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, w{"Madrid", "ES-MD"}, w{"mul-Latn"}, "whois.nic.madrid", e, w{"https://rdap.nic.madrid/"}, t}
	z[873] = Zone{"maif", r, x, 0x42, "Mutuelle Assurance Instituteur France (MAIF)", "http://nic.maif/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"mul-Latn"}, e, e, w{"https://tld-rdap.verisign.com/maif/v1/"}, t}
	z[874] = Zone{"mail", r, x, 0x800, "ICANN", "https://features.icann.org/addressing-new-gtld-program-applications-corp-home-and-mail", n, n, n, n, e, e, n, t}
	z[875] = Zone{"maison", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.maison", "v0n1.nic.maison", "v0n2.nic.maison", "v0n3.nic.maison", "v2n0.nic.maison", "v2n1.nic.maison"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.maison", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[876] = Zone{"makeup", r, x, 0x42, "XYZ.COM LLC", "https://nic.makeup/", w{"a.nic.makeup", "b.nic.makeup", "c.nic.makeup", "d.nic.makeup"}, n, n, w{"ja", "ko", "mul-Grek", "mul-Latn", "ru", "zh"}, "whois.nic.makeup", e, w{"https://rdap.centralnic.com/makeup/"}, t}
	z[877] = Zone{"man", r, x, 0x42, "MAN SE", "https://nic.man/en/index.jsp", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"ar", "ca", "cs", "da", "de", "el", "es", "fi", "fr", "he", "hr", "hu", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "tr", "uk", "zh"}, "whois.nic.man", e, w{"https://rdap.nic.man/"}, t}
	z[878] = Zone{"management", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.management", "v0n1.nic.management", "v0n2.nic.management", "v0n3.nic.management", "v2n0.nic.management", "v2n1.nic.management"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.management", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[879] = Zone{"mango", r, x, 0x42, "PUNTO FA S.L.", "https://shop.mango.com/preHome.faces", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.mango", e, w{"https://rdap.nic.mango/"}, t}
}

//go:noinline
func i879() {
	z[880] = Zone{"map", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[881] = Zone{"market", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.market", "v0n1.nic.market", "v0n2.nic.market", "v0n3.nic.market", "v2n0.nic.market", "v2n1.nic.market"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.market", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[882] = Zone{"marketing", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.marketing", "v0n1.nic.marketing", "v0n2.nic.marketing", "v0n3.nic.marketing", "v2n0.nic.marketing", "v2n1.nic.marketing"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.marketing", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[883] = Zone{"markets", r, x, 0x40, "Dog Beach, LLC", "http://nic.markets/", w{"v0n0.nic.markets", "v0n1.nic.markets", "v0n2.nic.markets", "v0n3.nic.markets", "v2n0.nic.markets", "v2n1.nic.markets"}, n, n, n, "whois.nic.markets", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[884] = Zone{"marriott", r, x, 0x42, "Marriott Worldwide Corporation", "https://www.marriott.com/marriott/gtld.mi", w{"a0.nic.marriott", "a2.nic.marriott", "b0.nic.marriott", "c0.nic.marriott"}, n, n, n, "whois.nic.marriott", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[885] = Zone{"marshalls", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.marshalls/", w{"a.nic.marshalls", "b.nic.marshalls", "c.nic.marshalls", "ns1.dns.nic.marshalls", "ns2.dns.nic.marshalls", "ns3.dns.nic.marshalls"}, n, n, n, "whois.nic.marshalls", e, w{"https://rdap.nic.marshalls/"}, f}
	z[886] = Zone{"maserati", r, x, 0x1042, "Fiat Chrysler Automobiles N.V.", "https://www.icann.org/en/registry-agreements/details/maserati", n, n, n, n, "whois.nic.maserati", e, w{"https://rdap.afilias-srs.net/rdap/maserati/"}, f}
	z[887] = Zone{"matrix", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i887() {
	z[888] = Zone{"mattel", r, x, 0x42, "Mattel Sites, Inc.", "https://www.icann.org/en/registry-agreements/details/mattel", w{"a.nic.mattel", "b.nic.mattel", "c.nic.mattel", "ns1.dns.nic.mattel", "ns2.dns.nic.mattel", "ns3.dns.nic.mattel"}, n, n, w{"es"}, "whois.nic.mattel", e, w{"https://rdap.nic.mattel/"}, t}
	z[889] = Zone{"maybelline", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[890] = Zone{"mba", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.mba", "v0n1.nic.mba", "v0n2.nic.mba", "v0n3.nic.mba", "v2n0.nic.mba", "v2n1.nic.mba"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.mba", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[891] = Zone{"mc", r, z[3420:3422], 0xa0, e, "https://www.nic.mc/", w{"mc.cctld.authdns.ripe.net", "ns1.nic.mc", "ns2.nic.mc", "ns3.nic.mc"}, n, n, n, e, e, n, f}
	z[892] = Zone{"mcd", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[893] = Zone{"mcdonalds", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[894] = Zone{"mckinsey", r, x, 0x42, "McKinsey Holdings, Inc.", "https://www.mckinsey.com/nic", w{"a0.nic.mckinsey", "a2.nic.mckinsey", "b0.nic.mckinsey", "c0.nic.mckinsey"}, n, n, n, "whois.nic.mckinsey", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[895] = Zone{"md", r, x, 0xa0, e, e, w{"nsa.tld.md", "nsb.tld.md", "nsc.dns.md", "nsf.dns.md", "nsr.dns.md"}, n, n, n, "whois.nic.md", e, n, f}
}

//go:noinline
func i895() {
	z[896] = Zone{"me", r, z[3422:3430], 0xe0, e, "https://domain.me/", w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, "whois.nic.me", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[897] = Zone{"med", r, x, 0x40, "Medistry LLC", "http://www.nic.med/", w{"dns1.nic.med", "dns2.nic.med", "dns3.nic.med", "dns4.nic.med", "dnsa.nic.med", "dnsb.nic.med", "dnsc.nic.med", "dnsd.nic.med"}, n, n, n, "whois.nic.med", e, w{"https://rdap.nominet.uk/med/"}, f}
	z[898] = Zone{"media", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.media", "v0n1.nic.media", "v0n2.nic.media", "v0n3.nic.media", "v2n0.nic.media", "v2n1.nic.media"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.media", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[899] = Zone{"medical", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[900] = Zone{"meet", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[901] = Zone{"melbourne", r, x, 0xc4, "The Crown in right of the State of Victoria, represented by its Department of State Development, Business and Innovation", "https://www.live.melbourne/", w{"a.nic.melbourne", "b.nic.melbourne", "c.nic.melbourne", "x.nic.melbourne", "y.nic.melbourne", "z.nic.melbourne"}, n, w{"Melbourne"}, n, "whois.nic.melbourne", e, w{"https://rdap.nic.melbourne/"}, f}
	z[902] = Zone{"meme", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[903] = Zone{"memorial", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.memorial", "v0n1.nic.memorial", "v0n2.nic.memorial", "v0n3.nic.memorial", "v2n0.nic.memorial", "v2n1.nic.memorial"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.memorial", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i903() {
	z[904] = Zone{"men", r, x, 0x40, "Exclusive Registry Limited", "http://nic.men/", w{"a.nic.men", "b.nic.men", "c.nic.men", "x.nic.men", "y.nic.men", "z.nic.men"}, n, n, n, "whois.nic.men", e, w{"https://rdap.nic.men/"}, t}
	z[905] = Zone{"menu", r, x, 0x40, "Dot Menu Registry, LLC", e, w{"a.nic.menu", "b.nic.menu", "c.nic.menu", "x.nic.menu", "y.nic.menu", "z.nic.menu"}, n, n, n, "whois.nic.menu", e, w{"https://rdap.nic.menu/"}, f}
	z[906] = Zone{"meo", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[907] = Zone{"merck", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[908] = Zone{"merckmsd", r, x, 0x42, "MSD Registry Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/merckmsd", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, e, e, w{"https://tld-rdap.verisign.com/merckmsd/v1/"}, f}
	z[909] = Zone{"metlife", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.metlife", e, n, f}
	z[910] = Zone{"mf", r, x, 0x10a0, e, "https://en.wikipedia.org/wiki/.mf", n, n, n, n, e, e, n, f}
	z[911] = Zone{"mg", r, z[3430:3441], 0xa0, e, e, w{"dns-tld.ird.fr", "ns-mg.afrinic.net", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, "whois.nic.mg", e, n, f}
}

//go:noinline
func i911() {
	z[912] = Zone{"mh", r, x, 0xa8, e, e, w{"ns.amarshallinc.com", "ns.ntamar.net"}, n, n, n, e, e, n, f}
	z[913] = Zone{"miami", r, x, 0xc4, "Registry Services, LLC", "http://nic.miami/", w{"a.nic.miami", "b.nic.miami", "c.nic.miami", "x.nic.miami", "y.nic.miami", "z.nic.miami"}, n, w{"Miami"}, w{"es", "fr"}, "whois.nic.miami", e, w{"https://rdap.nic.miami/"}, t}
	z[914] = Zone{"microsoft", r, x, 0x42, "Microsoft Corporation", "https://nic.microsoft/", w{"dns1.nic.microsoft", "dns2.nic.microsoft", "dns3.nic.microsoft", "dns4.nic.microsoft", "dnsa.nic.microsoft", "dnsb.nic.microsoft", "dnsc.nic.microsoft", "dnsd.nic.microsoft"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/microsoft/"}, t}
	z[915] = Zone{"mih", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[916] = Zone{"mii", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[917] = Zone{"mil", r, x, 0x4040, e, e, w{"con1.nipr.mil", "con2.nipr.mil", "eur1.nipr.mil", "eur2.nipr.mil", "pac1.nipr.mil", "pac2.nipr.mil"}, n, n, n, "is-1.nic.mil", e, n, f}
	z[918] = Zone{"mini", r, x, 0x42, "Bayerische Motoren Werke Aktiengesellschaft", "https://nic.mini/", w{"a.nic.mini", "b.nic.mini", "c.nic.mini", "d.nic.mini"}, n, n, w{"de"}, "whois.nic.mini", e, w{"https://rdap.centralnic.com/mini/"}, t}
	z[919] = Zone{"mint", r, x, 0x42, "Intuit Administrative Services, Inc.", "http://www.nic.mint/", w{"a.nic.mint", "b.nic.mint", "c.nic.mint", "ns4.dns.nic.mint", "ns5.dns.nic.mint", "ns6.dns.nic.mint"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.mint", e, w{"https://rdap.nic.mint/"}, t}
}

//go:noinline
func i919() {
	z[920] = Zone{"mit", r, x, 0x42, "Massachusetts Institute of Technology", "https://nic.mit/", w{"a0.nic.mit", "a2.nic.mit", "b0.nic.mit", "c0.nic.mit"}, n, n, n, "whois.nic.mit", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[921] = Zone{"mitek", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[922] = Zone{"mitsubishi", r, x, 0x42, "Mitsubishi Corporation", "https://nic.mitsubishi/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[923] = Zone{"mk", r, z[3441:3449], 0xa0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, "whois.marnet.mk", e, n, f}
	z[924] = Zone{"ml", r, z[3449:3456], 0xa0, e, e, w{"a.nic.ml", "b.nic.ml", "c.nic.ml", "d.nic.ml"}, n, n, n, "whois.nic.ml", e, n, t}
	z[925] = Zone{"mlb", r, x, 0x42, "MLB Advanced Media DH, LLC", "https://www.mlb.com/official-information/mlb-nic", w{"a.nic.mlb", "b.nic.mlb", "c.nic.mlb", "ns1.dns.nic.mlb", "ns2.dns.nic.mlb", "ns3.dns.nic.mlb"}, n, n, w{"es"}, "whois.nic.mlb", e, w{"https://rdap.nic.mlb/"}, t}
	z[926] = Zone{"mls", r, x, 0x40, "The Canadian Real Estate Association", "https://www.icann.org/en/registry-agreements/details/mls", w{"a.ns.nic.mls", "b.ns.nic.mls"}, n, n, w{"fr"}, "whois.nic.mls", e, w{"https://rdap.mls.fury.ca/rdap/"}, t}
	z[927] = Zone{"mm", r, z[3456:3464], 0xa8, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, "whois.registry.gov.mm", e, n, f}
}

//go:noinline
func i927() {
	z[928] = Zone{"mma", r, x, 0x42, "MMA IARD", "http://nic.mma/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.mma", e, w{"https://rdap.nic.mma/"}, t}
	z[929] = Zone{"mn", r, z[3464:3467], 0xa0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "ns1.magic.mn", "ns2.magic.mn", "ns3.magic.mn", "ns4.magic.mn"}, n, n, n, "whois.afilias-grs.info", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[930] = Zone{"mnet", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[931] = Zone{"mo", r, z[3467:3473], 0xa0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, "whois.monic.mo", e, n, f}
	z[932] = Zone{"mobi", r, x, 0x4040, "Identity Digital Limited", "http://nic.mobi/", w{"a0.mobi.afilias-nst.info", "a2.mobi.afilias-nst.info", "b0.mobi.afilias-nst.org", "b2.mobi.afilias-nst.org", "c0.mobi.afilias-nst.info", "d0.mobi.afilias-nst.org"}, n, n, w{"zh-CN"}, "whois.dotmobiregistry.net", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[933] = Zone{"mobile", r, x, 0x40, "Dish DBS Corporation", "https://www.dishtlds.com/mobile/", w{"a0.nic.mobile", "a2.nic.mobile", "b0.nic.mobile", "c0.nic.mobile"}, n, n, n, "whois.nic.mobile", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[934] = Zone{"mobily", r, x, 0x1040, e, e, n, n, n, n, "whois.nic.mobily", e, n, f}
	z[935] = Zone{"moda", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.moda", "v0n1.nic.moda", "v0n2.nic.moda", "v0n3.nic.moda", "v2n0.nic.moda", "v2n1.nic.moda"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.moda", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i935() {
	z[936] = Zone{"moe", r, x, 0x40, "Interlink Systems Innovation Institute K.K.", e, w{"a.nic.moe", "b.nic.moe", "c.nic.moe", "ns1.dns.nic.moe", "ns2.dns.nic.moe", "ns3.dns.nic.moe"}, n, n, w{"ja"}, "whois.nic.moe", e, w{"https://rdap.nic.moe/"}, t}
	z[937] = Zone{"moi", r, x, 0x42, "Amazon Registry Services, Inc.", "https://bienvenue.moi/", w{"dns1.nic.moi", "dns2.nic.moi", "dns3.nic.moi", "dns4.nic.moi", "dnsa.nic.moi", "dnsb.nic.moi", "dnsc.nic.moi", "dnsd.nic.moi"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.moi", e, w{"https://rdap.nominet.uk/moi/"}, t}
	z[938] = Zone{"mom", r, x, 0x40, "XYZ.COM LLC", "https://nic.mom/", w{"a.nic.mom", "b.nic.mom", "c.nic.mom", "d.nic.mom"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.mom", e, w{"https://rdap.centralnic.com/mom/"}, t}
	z[939] = Zone{"monash", r, x, 0x42, "Monash University", "http://nic.monash/", w{"a.nic.monash", "b.nic.monash", "c.nic.monash", "x.nic.monash", "y.nic.monash", "z.nic.monash"}, n, n, n, "whois.nic.monash", e, w{"https://rdap.nic.monash/"}, f}
	z[940] = Zone{"money", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.money", "v0n1.nic.money", "v0n2.nic.money", "v0n3.nic.money", "v2n0.nic.money", "v2n1.nic.money"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.money", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[941] = Zone{"monster", r, x, 0x42, "XYZ.COM LLC", "https://nic.monster/", w{"a.nic.monster", "b.nic.monster", "c.nic.monster", "d.nic.monster"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.monster", e, w{"https://rdap.centralnic.com/monster/"}, t}
	z[942] = Zone{"montblanc", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[943] = Zone{"mopar", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
}

//go:noinline
func i943() {
	z[944] = Zone{"mormon", r, x, 0x42, "IRI Domain Management, LLC", "https://www.icann.org/en/registry-agreements/details/mormon", w{"a0.nic.mormon", "a2.nic.mormon", "b0.nic.mormon", "c0.nic.mormon"}, n, n, n, "whois.nic.mormon", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[945] = Zone{"mortgage", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.mortgage", "v0n1.nic.mortgage", "v0n2.nic.mortgage", "v0n3.nic.mortgage", "v2n0.nic.mortgage", "v2n1.nic.mortgage"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.mortgage", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[946] = Zone{"moscow", r, x, 0xc4, "Foundation for Assistance for Internet Technologies and Infrastructure Development (FAITID)", "http://www.nic.moscow/", w{"a.dns.flexireg.ru", "b.dns.flexireg.net", "c.dns.flexireg.org", "d.dns.flexireg.domains"}, n, w{"Moscow", "RU-MOW"}, n, "whois.nic.moscow", e, w{"https://flexireg.net/moscow/rdap/"}, f}
	z[947] = Zone{"moto", r, x, 0x42, "Motorola Trademark Holdings, LLC", "http://www.nic.moto/", w{"a.nic.moto", "b.nic.moto", "c.nic.moto", "ns4.dns.nic.moto", "ns5.dns.nic.moto", "ns6.dns.nic.moto"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.moto", e, w{"https://rdap.nic.moto/"}, t}
	z[948] = Zone{"motorcycles", r, x, 0x40, "XYZ.COM LLC", "https://nic.motorcycles/", w{"a.nic.motorcycles", "b.nic.motorcycles", "c.nic.motorcycles", "d.nic.motorcycles"}, n, n, n, "whois.nic.motorcycles", e, w{"https://rdap.centralnic.com/motorcycles/"}, f}
	z[949] = Zone{"mov", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[950] = Zone{"movie", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.movie", "v0n1.nic.movie", "v0n2.nic.movie", "v0n3.nic.movie", "v2n0.nic.movie", "v2n1.nic.movie"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.movie", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[951] = Zone{"movistar", r, x, 0x1042, e, e, n, n, n, n, "whois-fe.movistar.tango.knipp.de", e, n, f}
}

//go:noinline
func i951() {
	z[952] = Zone{"mozaic", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[953] = Zone{"mp", r, z[3473:3477], 0xa0, e, "https://get.mp/", w{"ns1.nic.mp", "ns2.nic.mp", "ns3.nic.mp", "ns4.nic.mp"}, n, n, n, "whois.nic.mp", e, n, f}
	z[954] = Zone{"mq", r, x, 0xa0, e, "https://www.dom-enic.com/", w{"ns1-fr.mediaserv.net", "ns1-gp.mediaserv.net", "ns1-mq.mediaserv.net"}, n, n, n, "whois.mediaserv.net", e, n, f}
	z[955] = Zone{"mr", r, z[3477:3481], 0xa0, e, e, w{"ns-mr.afrinic.net", "ns-mr.nic.fr", "ns1.nic.mr", "ns2.nic.mr", "ns3.nic.mr"}, n, n, n, "whois.nic.mr", e, n, f}
	z[956] = Zone{"mrmuscle", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[957] = Zone{"mrporter", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[958] = Zone{"ms", r, z[3481:3485], 0xe0, e, e, w{"a.lactld.org", "ms-ns.anycast.pch.net", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.ms", e, n, f}
	z[959] = Zone{"msd", r, x, 0x42, "MSD Registry Holdings, Inc.", "https://www.nic.msd/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, n, e, e, w{"https://tld-rdap.verisign.com/msd/v1/"}, f}
}

//go:noinline
func i959() {
	z[960] = Zone{"mt", r, z[3485:3490], 0xa0, e, e, w{"a.ns.mt", "b.ns.mt", "f.ns.mt", "p.ns.mt"}, n, n, n, e, "https://www.nic.org.mt/dotmt/", n, f}
	z[961] = Zone{"mtn", r, x, 0x42, "MTN Dubai Limited", "http://www.nic.mtn/", w{"dns1.nic.mtn", "dns2.nic.mtn", "dns3.nic.mtn", "dns4.nic.mtn", "dnsa.nic.mtn", "dnsb.nic.mtn", "dnsc.nic.mtn", "dnsd.nic.mtn"}, n, n, n, "whois.nic.mtn", e, w{"https://rdap.nominet.uk/mtn/"}, f}
	z[962] = Zone{"mtpc", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.mtpc", e, n, f}
	z[963] = Zone{"mtr", r, x, 0x42, "MTR Corporation Limited", "https://www.icann.org/en/registry-agreements/details/mtr", w{"ns1.nic.mtr", "ns2.nic.mtr", "ns3.nic.mtr", "ns4.nic.mtr"}, n, n, n, "whois.nic.mtr", e, w{"https://whois.nic.mtr/rdap/"}, f}
	z[964] = Zone{"mu", r, z[3490:3499], 0xa0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, "whois.nic.mu", e, n, f}
	z[965] = Zone{"multichoice", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[966] = Zone{"museum", r, x, 0x4040, "Museum Domain Management Association (MuseDoma)", "https://welcome.museum/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.museum", e, w{"https://rdap.nic.museum/"}, t}
	z[967] = Zone{"music", r, x, 0x40, "DotMusic Limited", "https://nic.music/", w{"a.nic.music", "b.nic.music", "c.nic.music", "d.nic.music"}, n, n, n, "whois.nic.music", e, w{"http://rdap.centralnic.com/music/"}, f}
}

//go:noinline
func i967() {
	z[968] = Zone{"mutual", r, x, 0x1042, "Northwestern Mutual MU TLD Registry, LLC", "https://www.nic.mutual/", n, n, n, w{"es"}, "whois.nic.mutual", e, w{"https://rdap.nic.mutual/"}, t}
	z[969] = Zone{"mutualfunds", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[970] = Zone{"mutuelle", r, x, 0x1040, e, e, n, n, n, n, "whois-mutuelle.nic.fr", e, n, f}
	z[971] = Zone{"mv", r, z[3499:3513], 0xa0, e, e, w{"mv-ns.anycast.pch.net", "ns.dhivehinet.net.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[972] = Zone{"mw", r, z[3513:3523], 0xa0, e, "http://nic.mw/", w{"chambo.sdnp.org.mw", "domwe.sdn.mw", "mw-e.ns.nic.cz", "mw.cctld.authdns.ripe.net", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, "whois.nic.mw", e, n, f}
	z[973] = Zone{"mx", r, z[3523:3528], 0xa0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, "whois.mx", e, n, f}
	z[974] = Zone{"my", r, z[3528:3536], 0xa0, e, "https://mynic.my/", w{"a.mynic.centralnic-dns.com", "a.nic.my", "a1.nic.my", "b.mynic.centralnic-dns.com", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, w{"ms-MY"}, "whois.mynic.my", e, n, t}
	z[975] = Zone{"mz", r, z[3536:3543], 0xa8, e, "https://www.domains.co.mz/", w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, "whois.nic.mz", e, n, f}
}

//go:noinline
func i975() {
	z[976] = Zone{"mzansimagic", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[977] = Zone{"na", r, z[3543:3561], 0xa0, e, e, w{"anyc2.irondns.net", "etld-1.anycast.net", "na-ns.anycast.pch.net", "na.anycastdns.cz"}, n, n, n, "whois.na-nic.com.na", e, n, f}
	z[978] = Zone{"nab", r, x, 0x42, "National Australia Bank Limited", "https://www.nab.com.au/nic", w{"a0.nic.nab", "a2.nic.nab", "b0.nic.nab", "c0.nic.nab"}, n, n, n, "whois.nic.nab", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[979] = Zone{"nadex", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.nadex", e, n, f}
	z[980] = Zone{"nagoya", r, x, 0xc4, "GMO Registry, Inc.", e, w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"Nagoya-shi"}, w{"ja"}, "whois.nic.nagoya", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[981] = Zone{"name", r, z[3561:3564], 0x40, "VeriSign, Inc.", e, w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.name", e, w{"https://tld-rdap.verisign.com/name/v1/"}, t}
	z[982] = Zone{"naspers", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[983] = Zone{"nationwide", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.nationwide", e, w{"https://tld-rdap.verisign.com/nationwide/v1/"}, f}
}

//go:noinline
func i983() {
	z[984] = Zone{"natura", r, x, 0x42, "NATURA COSMÉTICOS S.A.", "https://nic.natura/", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
	z[985] = Zone{"navy", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.navy", "v0n1.nic.navy", "v0n2.nic.navy", "v0n3.nic.navy", "v2n0.nic.navy", "v2n1.nic.navy"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.navy", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[986] = Zone{"nba", r, x, 0x42, "NBA REGISTRY, LLC", "http://www.nic.nba/", w{"a.nic.nba", "b.nic.nba", "c.nic.nba", "ns1.dns.nic.nba", "ns2.dns.nic.nba", "ns3.dns.nic.nba"}, n, n, w{"es"}, "whois.nic.nba", e, w{"https://rdap.nic.nba/"}, t}
	z[987] = Zone{"nc", r, z[3564:3567], 0xa0, e, e, w{"any-ns1.nc", "nc.cctld.authdns.ripe.net", "ns1.nc", "ns2.nc"}, n, n, n, "whois.nc", e, n, f}
	z[988] = Zone{"ne", r, z[3567:3572], 0xa0, e, e, w{"bow.rain.fr", "ne.cctld.authdns.ripe.net", "ns-ne.afrinic.net", "ns.intnet.ne"}, n, n, n, e, e, n, f}
	z[989] = Zone{"nec", r, x, 0x42, "NEC Corporation", "https://nic.nec/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.nec", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[990] = Zone{"net", r, z[3572:3583], 0x40, "VeriSign, Inc.", "https://www.verisign.com/en_US/domain-names/net-domain-names/index.xhtml", w{"a.gtld-servers.net", "b.gtld-servers.net", "c.gtld-servers.net", "d.gtld-servers.net", "e.gtld-servers.net", "f.gtld-servers.net", "g.gtld-servers.net", "h.gtld-servers.net", "i.gtld-servers.net", "j.gtld-servers.net", "k.gtld-servers.net", "l.gtld-servers.net", "m.gtld-servers.net"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro-MD", "ru", "uk", "zh"}, "whois.verisign-grs.com", e, w{"https://rdap.verisign.com/net/v1/"}, t}
	z[991] = Zone{"netaporter", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i991() {
	z[992] = Zone{"netbank", r, x, 0x42, "COMMONWEALTH BANK OF AUSTRALIA", "http://nic.netbank/", w{"a.nic.netbank", "b.nic.netbank", "c.nic.netbank", "x.nic.netbank", "y.nic.netbank", "z.nic.netbank"}, n, n, n, "whois.nic.netbank", e, w{"https://rdap.nic.netbank/"}, f}
	z[993] = Zone{"netflix", r, x, 0x42, "Netflix, Inc.", "https://about.netflix.com/en/nic", w{"a.nic.netflix", "b.nic.netflix", "c.nic.netflix", "ns4.dns.nic.netflix", "ns5.dns.nic.netflix", "ns6.dns.nic.netflix"}, n, n, n, "whois.nic.netflix", e, w{"https://rdap.nic.netflix/"}, f}
	z[994] = Zone{"network", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.network", "v0n1.nic.network", "v0n2.nic.network", "v0n3.nic.network", "v2n0.nic.network", "v2n1.nic.network"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.network", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[995] = Zone{"neustar", r, x, 0x42, "NeuStar, Inc.", "http://nic.neustar/", w{"a.nic.neustar", "b.nic.neustar", "c.nic.neustar", "ns4.dns.nic.neustar", "ns5.dns.nic.neustar", "ns6.dns.nic.neustar"}, n, n, w{"es"}, "whois.nic.neustar", e, w{"https://rdap.nic.neustar/"}, t}
	z[996] = Zone{"new", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[997] = Zone{"newholland", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.newholland", e, n, f}
	z[998] = Zone{"news", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.news", "v0n1.nic.news", "v0n2.nic.news", "v0n3.nic.news", "v2n0.nic.news", "v2n1.nic.news"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.news", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[999] = Zone{"next", r, x, 0x42, "Next plc", "http://nic.next/", w{"a0.nic.next", "a2.nic.next", "b0.nic.next", "c0.nic.next"}, n, n, n, "whois.nic.next", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i999() {
	z[1000] = Zone{"nextdirect", r, x, 0x42, "Next plc", "http://nic.nextdirect/", w{"a0.nic.nextdirect", "a2.nic.nextdirect", "b0.nic.nextdirect", "c0.nic.nextdirect"}, n, n, n, "whois.nic.nextdirect", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1001] = Zone{"nexus", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1002] = Zone{"nf", r, z[3583:3595], 0xa0, e, e, w{"ns.anycast.nic.nf", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.nf", e, n, f}
	z[1003] = Zone{"nfl", r, x, 0x42, "NFL Reg Ops LLC", "http://www.nic.nfl/", w{"a.nic.nfl", "b.nic.nfl", "c.nic.nfl", "ns4.dns.nic.nfl", "ns5.dns.nic.nfl", "ns6.dns.nic.nfl"}, n, n, w{"es"}, "whois.nic.nfl", e, w{"https://rdap.nic.nfl/"}, t}
	z[1004] = Zone{"ng", r, z[3595:3605], 0xa0, e, "https://web4africa.ng/ng-domains/", w{"ns2.nic.net.ng", "ns3.nic.net.ng", "ns4.nic.net.ng", "ns5.nic.net.ng", "nsa.nic.net.ng"}, n, n, n, "whois.nic.net.ng", e, n, f}
	z[1005] = Zone{"ngo", r, x, 0x40, "Public Interest Registry", "https://pir.org/products/ngo-ong-domain/", w{"a0.nic.ngo", "a2.nic.ngo", "b0.nic.ngo", "b2.nic.ngo", "c0.nic.ngo", "d0.nic.ngo"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.ngo", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[1006] = Zone{"nhk", r, x, 0x42, "Japan Broadcasting Corporation (NHK)", "https://nic.nhk/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "und"}, "whois.nic.nhk", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1007] = Zone{"ni", r, z[3605:3620], 0xa8, e, "https://www.nic.ni/", w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, "http://www.nic.ni/", n, f}
}

//go:noinline
func i1007() {
	z[1008] = Zone{"nico", r, x, 0x42, "DWANGO Co., Ltd.", "https://nic.nico/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.nico", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1009] = Zone{"nike", r, x, 0x42, "NIKE, Inc.", "https://www.nike.com/", w{"a.nic.nike", "b.nic.nike", "c.nic.nike", "ns1.dns.nic.nike", "ns2.dns.nic.nike", "ns3.dns.nic.nike"}, n, n, w{"es"}, "whois.nic.nike", e, w{"https://rdap.nic.nike/"}, t}
	z[1010] = Zone{"nikon", r, x, 0x42, "NIKON CORPORATION", "https://www.nic.nikon/", w{"a0.nic.nikon", "a2.nic.nikon", "b0.nic.nikon", "c0.nic.nikon"}, n, n, n, "whois.nic.nikon", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1011] = Zone{"ninja", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.ninja", "v0n1.nic.ninja", "v0n2.nic.ninja", "v0n3.nic.ninja", "v2n0.nic.ninja", "v2n1.nic.ninja"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.ninja", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1012] = Zone{"nissan", r, x, 0x42, "NISSAN MOTOR CO., LTD.", "https://www.nissan-global.com/EN/NIC/NISSAN/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1013] = Zone{"nissay", r, x, 0x42, "Nippon Life Insurance Company", "https://www.nissay.co.jp/sorry/gtld/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"ja", "mul-Hang", "mul-Hani", "mul-Hano"}, "whois.nic.nissay", e, w{"https://tld-rdap.verisign.com/nissay/v1/"}, t}
	z[1014] = Zone{"nl", r, z[3620:3626], 0xa0, e, e, w{"ns1.dns.nl", "ns3.dns.nl", "ns4.dns.nl"}, n, n, n, "whois.domain-registry.nl", e, w{"https://rdap.sidn.nl/"}, f}
	z[1015] = Zone{"no", r, z[3626:3636], 0xa0, e, e, w{"i.nic.no", "njet.norid.no", "not.norid.no", "x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, "whois.norid.no", e, w{"https://rdap.norid.no/"}, t}
}

//go:noinline
func i1015() {
	z[1016] = Zone{"nokia", r, x, 0x42, "Nokia Corporation", "http://nic.nokia/", w{"a.nic.nokia", "b.nic.nokia", "c.nic.nokia", "d.nic.nokia"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.nokia", e, w{"https://rdap.centralnic.com/nokia/"}, t}
	z[1017] = Zone{"northlandinsurance", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1018] = Zone{"northwesternmutual", r, x, 0x1042, "Northwestern Mutual Registry, LLC", "https://www.nic.northwesternmutual/", n, n, n, w{"es"}, "whois.nic.northwesternmutual", e, w{"https://rdap.nic.northwesternmutual/"}, t}
	z[1019] = Zone{"norton", r, x, 0x42, "NortonLifeLock Inc.", "https://www.icann.org/en/registry-agreements/details/norton", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "bs-Cyrl", "cnr-Cyrl", "da", "de", "el", "es", "fi", "fr", "he", "hi", "hr", "hu", "is", "it", "ja", "ko", "ku", "lo", "lt", "lv", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "nl", "no", "pl", "pt", "ro", "ru", "sr", "sv", "th", "uk", "und-Beng", "yi", "zh"}, "whois.nic.norton", e, w{"https://tld-rdap.verisign.com/norton/v1/"}, t}
	z[1020] = Zone{"now", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.now/", w{"dns1.nic.now", "dns2.nic.now", "dns3.nic.now", "dns4.nic.now", "dnsa.nic.now", "dnsb.nic.now", "dnsc.nic.now", "dnsd.nic.now"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.now", e, w{"https://rdap.nominet.uk/now/"}, t}
	z[1021] = Zone{"nowruz", r, x, 0x42, "Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.", "http://nic.nowruz/", w{"a.ns.nic.nowruz", "b.ns.nic.nowruz", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, w{"Iran"}, n, "whois.nic.nowruz", e, w{"https://api.rdap.nic.nowruz/"}, f}
	z[1022] = Zone{"nowtv", r, x, 0x42, "Starbucks (HK) Limited", "https://www.icann.org/en/registry-agreements/details/nowtv", w{"a0.nic.nowtv", "a2.nic.nowtv", "b0.nic.nowtv", "c0.nic.nowtv"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.nowtv", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1023] = Zone{"np", r, z[3636:3782], 0xa8, e, e, w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, "http://register.mos.com.np/np-whois-lookup", n, f}
}

//go:noinline
func i1023() {
	z[1024] = Zone{"nr", r, z[3782:3789], 0xa0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr", "phloem.uoregon.edu"}, n, n, n, e, "http://www.cenpac.net.nr/dns/whois.html", n, f}
	z[1025] = Zone{"nra", r, x, 0x42, "NRA Holdings Company, INC.", "https://www.icann.org/en/registry-agreements/details/nra", w{"a0.nic.nra", "a2.nic.nra", "b0.nic.nra", "c0.nic.nra"}, n, n, n, "whois.nic.nra", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1026] = Zone{"nrw", r, x, 0x440, "Minds + Machines GmbH", "https://nic.nrw/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net", "ari.alpha.tldns.godaddy", "ari.beta.tldns.godaddy", "ari.delta.tldns.godaddy", "ari.gamma.tldns.godaddy"}, n, w{"DE-NW"}, w{"de"}, "whois.nic.nrw", e, w{"https://rdap.nic.nrw/"}, t}
	z[1027] = Zone{"ntt", r, x, 0x42, "NIPPON TELEGRAPH AND TELEPHONE CORPORATION", "https://www.icann.org/en/registry-agreements/details/ntt", w{"tld1.nic.ntt", "tld2.nic.ntt", "tld3.nic.ntt", "tld5.nic.ntt"}, n, n, w{"ja", "ja-JP"}, e, e, w{"https://rdap.nic.ntt/rdap/"}, t}
	z[1028] = Zone{"nu", r, z[3789:3794], 0xe0, e, "https://www.nic.nu/", w{"a.ns.nu", "c.ns.nu", "d.ns.nu", "m.ns.nu", "y.ns.nu", "z.ns.nu"}, n, n, w{"und-Latn"}, "whois.iis.nu", e, n, t}
	z[1029] = Zone{"nyc", r, x, 0xc4, "The City of New York by and through the New York City Department of Information Technology & Telecommunications", e, w{"a.nic.nyc", "b.nic.nyc", "c.nic.nyc", "ns1.dns.nic.nyc", "ns2.dns.nic.nyc", "ns3.dns.nic.nyc"}, n, w{"New York City"}, w{"bn", "es", "fr", "ru", "yi", "zh-Hani"}, "whois.nic.nyc", e, w{"https://rdap.nic.nyc/"}, t}
	z[1030] = Zone{"nz", r, z[3794:3809], 0xa0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, w{"mul-Latn"}, "whois.irs.net.nz", e, n, t}
	z[1031] = Zone{"obi", r, x, 0x42, "OBI Group Holding SE & Co. KGaA", "http://nic.obi/", w{"a0.nic.obi", "a2.nic.obi", "b0.nic.obi", "c0.nic.obi"}, n, n, n, "whois.nic.obi", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i1031() {
	z[1032] = Zone{"observer", r, x, 0x42, "Fegistry, LLC", "https://identity.digital/", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.observer", e, w{"https://rdap.registry.click/rdap/"}, f}
	z[1033] = Zone{"off", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.off", e, w{"https://tld-rdap.verisign.com/off/v1/"}, f}
	z[1034] = Zone{"office", r, x, 0x42, "Microsoft Corporation", "https://nic.office/", w{"a.nic.office", "b.nic.office", "c.nic.office", "ns1.dns.nic.office", "ns2.dns.nic.office", "ns3.dns.nic.office"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.office", e, w{"https://rdap.nic.office/"}, t}
	z[1035] = Zone{"okinawa", r, x, 0xc4, "BRregistry, Inc.", e, w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"Okinawa", "JP-47"}, w{"ja"}, "whois.nic.okinawa", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1036] = Zone{"olayan", r, x, 0x42, "Competrol (Luxembourg) Sarl", "http://nic.olayan/", w{"a.nic.olayan", "b.nic.olayan", "c.nic.olayan", "x.nic.olayan", "y.nic.olayan", "z.nic.olayan"}, n, n, w{"ar"}, "whois.nic.olayan", e, w{"https://rdap.nic.olayan/"}, t}
	z[1037] = Zone{"olayangroup", r, x, 0x42, "Competrol (Luxembourg) Sarl", "http://nic.olayangroup/", w{"a.nic.olayangroup", "b.nic.olayangroup", "c.nic.olayangroup", "x.nic.olayangroup", "y.nic.olayangroup", "z.nic.olayangroup"}, n, n, w{"ar"}, "whois.nic.olayangroup", e, w{"https://rdap.nic.olayangroup/"}, t}
	z[1038] = Zone{"oldnavy", r, x, 0x1042, "The Gap, Inc.", "https://www.nic.gap/", n, n, n, w{"es"}, "whois.nic.oldnavy", e, w{"https://rdap.nic.oldnavy/"}, t}
	z[1039] = Zone{"ollo", r, x, 0x42, "Dish DBS Corporation", "https://dishtlds.com/ollo/", w{"a0.nic.ollo", "a2.nic.ollo", "b0.nic.ollo", "c0.nic.ollo"}, n, n, n, "whois.nic.ollo", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i1039() {
	z[1040] = Zone{"olympus", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1041] = Zone{"om", r, z[3809:3822], 0xa0, e, "https://www.registry.om/", w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, "whois.registry.om", e, n, f}
	z[1042] = Zone{"omega", r, x, 0x42, "The Swatch Group Ltd", "https://nic.omega/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.omega", e, w{"https://tld-rdap.verisign.com/omega/v1/"}, t}
	z[1043] = Zone{"one", r, x, 0x40, "One.com A/S", "https://get.one/", w{"a.nic.one", "b.nic.one", "c.nic.one", "x.nic.one", "y.nic.one", "z.nic.one"}, n, n, n, "whois.nic.one", e, w{"https://rdap.nic.one/"}, f}
	z[1044] = Zone{"ong", r, x, 0x40, "Public Interest Registry", "https://pir.org/products/ngo-ong-domain/", w{"a0.nic.ong", "a2.nic.ong", "b0.nic.ong", "b2.nic.ong", "c0.nic.ong", "d0.nic.ong"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.ong", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[1045] = Zone{"onion", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc7686", n, n, n, n, e, e, n, t}
	z[1046] = Zone{"onl", r, x, 0x40, "iRegistry GmbH", "https://nic.onl/", w{"a0.nic.onl", "a2.nic.onl", "b0.nic.onl", "c0.nic.onl"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.onl", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1047] = Zone{"online", r, x, 0x40, "Radix Technologies Inc.", "https://radix.website/dot-online", w{"a.nic.online", "b.nic.online", "e.nic.online", "f.nic.online"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.online", e, w{"https://rdap.centralnic.com/online/"}, t}
}

//go:noinline
func i1047() {
	z[1048] = Zone{"onyourside", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.onyourside", e, w{"https://tld-rdap.verisign.com/onyourside/v1/"}, f}
	z[1049] = Zone{"ooo", r, x, 0x40, "INFIBEAM AVENUES LIMITED", e, w{"a.nic.ooo", "b.nic.ooo", "c.nic.ooo", "d.nic.ooo"}, n, n, w{"ar", "be", "es", "fr", "he", "hr", "ja", "ko", "lo", "mk", "mul-Cans", "mul-Cyrl", "mul-Grek", "mul-Latn", "mul-Mymr", "mul-Ogam", "mul-Phnx", "mul-Runr", "pl", "ru", "sr", "sv", "th", "uk", "zh"}, "whois.nic.ooo", e, w{"https://rdap.centralnic.com/ooo/"}, t}
	z[1050] = Zone{"open", r, x, 0x42, "American Express Travel Related Services Company, Inc.", "https://about.americanexpress.com/newsroom/#media-contacts", w{"a.nic.open", "b.nic.open", "c.nic.open", "ns1.dns.nic.open", "ns2.dns.nic.open", "ns3.dns.nic.open"}, n, n, w{"es"}, "whois.nic.open", e, w{"https://rdap.nic.open/"}, t}
	z[1051] = Zone{"oracle", r, x, 0x42, "Oracle Corporation", "https://www.icann.org/en/registry-agreements/details/oracle", w{"a0.nic.oracle", "a2.nic.oracle", "b0.nic.oracle", "c0.nic.oracle"}, n, n, n, "whois.nic.oracle", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1052] = Zone{"orange", r, x, 0x42, "Orange Brand Services Limited", "https://www.orange.com/en/nic/domains", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"mul-Arab", "mul-Latn"}, "whois.nic.orange", e, w{"https://tld-rdap.verisign.com/orange/v1/"}, t}
	z[1053] = Zone{"org", r, z[3822:3829], 0x40, "Public Interest Registry", "https://pir.org/products/org-domain/", w{"a0.org.afilias-nst.info", "a2.org.afilias-nst.info", "b0.org.afilias-nst.org", "b2.org.afilias-nst.org", "c0.org.afilias-nst.info", "d0.org.afilias-nst.org"}, n, n, w{"be", "bg", "bs", "cnr", "da", "de", "es", "fi", "hi", "hr", "hu", "is", "it", "ja", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.publicinterestregistry.org", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[1054] = Zone{"organic", r, x, 0x40, "Identity Digital Limited", "http://nic.organic/", w{"a0.nic.organic", "a2.nic.organic", "b0.nic.organic", "c0.nic.organic"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.organic", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1055] = Zone{"orientexpress", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
}

//go:noinline
func i1055() {
	z[1056] = Zone{"origins", r, x, 0x42, "The Estée Lauder Companies Inc.", "https://www.icann.org/en/registry-agreements/details/origins", w{"a0.nic.origins", "a2.nic.origins", "b0.nic.origins", "c0.nic.origins"}, n, n, n, "whois.nic.origins", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1057] = Zone{"osaka", r, x, 0xc4, "Osaka Registry Co., Ltd.", "https://domain.osaka/", w{"a.nic.osaka", "b.nic.osaka", "c.nic.osaka", "ns1.dns.nic.osaka", "ns2.dns.nic.osaka", "ns3.dns.nic.osaka"}, n, w{"Osaka", "JP-27"}, w{"ja"}, "whois.nic.osaka", e, w{"https://rdap.nic.osaka/"}, t}
	z[1058] = Zone{"otsuka", r, x, 0x42, "Otsuka Holdings Co., Ltd.", "https://nic.otsuka/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.otsuka", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1059] = Zone{"ott", r, x, 0x42, "Dish DBS Corporation", "https://dishtlds.com/ott/", w{"a0.nic.ott", "a2.nic.ott", "b0.nic.ott", "c0.nic.ott"}, n, n, n, "whois.nic.ott", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1060] = Zone{"overheidnl", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1061] = Zone{"ovh", r, x, 0x40, "MédiaBC", e, w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"fr"}, "whois.nic.ovh", e, w{"https://rdap.nic.ovh/"}, t}
	z[1062] = Zone{"pa", r, z[3829:3840], 0xa0, e, "http://www.nic.pa/", w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, "http://www.nic.pa/", n, f}
	z[1063] = Zone{"page", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"ja", "mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
}

//go:noinline
func i1063() {
	z[1064] = Zone{"pamperedchef", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[1065] = Zone{"panasonic", r, x, 0x42, "Panasonic Holdings Corporation", "https://nic.panasonic/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1066] = Zone{"panerai", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.panerai", e, n, f}
	z[1067] = Zone{"paris", r, x, 0xc4, "City of Paris", e, w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr", "h.ext.nic.fr"}, n, w{"Paris", "FR-J"}, w{"fr"}, "whois.nic.paris", e, w{"https://rdap.nic.paris/"}, t}
	z[1068] = Zone{"pars", r, x, 0x40, "Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.", "http://nic.pars/", w{"a.ns.nic.pars", "b.ns.nic.pars", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, w{"Iran"}, n, "whois.nic.pars", e, w{"https://api.rdap.nic.pars/"}, f}
	z[1069] = Zone{"partners", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.partners", "v0n1.nic.partners", "v0n2.nic.partners", "v0n3.nic.partners", "v2n0.nic.partners", "v2n1.nic.partners"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.partners", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1070] = Zone{"parts", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.parts", "v0n1.nic.parts", "v0n2.nic.parts", "v0n3.nic.parts", "v2n0.nic.parts", "v2n1.nic.parts"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.parts", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1071] = Zone{"party", r, x, 0x40, "Blue Sky Registry Limited", "http://nic.party/", w{"a.nic.party", "b.nic.party", "c.nic.party", "ns1.dns.nic.party", "ns2.dns.nic.party", "ns3.dns.nic.party"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.party", e, w{"https://rdap.nic.party/"}, t}
}

//go:noinline
func i1071() {
	z[1072] = Zone{"passagens", r, x, 0x1042, "Travel Reservations SRL", "https://www.icann.org/en/registry-agreements/details/passagens", n, n, n, w{"es"}, "whois.nic.passagens", e, w{"https://rdap.nic.passagens/"}, t}
	z[1073] = Zone{"patagonia", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1074] = Zone{"patch", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1075] = Zone{"pay", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.pay/", w{"dns1.nic.pay", "dns2.nic.pay", "dns3.nic.pay", "dns4.nic.pay", "dnsa.nic.pay", "dnsb.nic.pay", "dnsc.nic.pay", "dnsd.nic.pay"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.pay", e, w{"https://rdap.nominet.uk/pay/"}, t}
	z[1076] = Zone{"payu", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1077] = Zone{"pccw", r, x, 0x42, "PCCW Enterprises Limited", "https://www.icann.org/en/registry-agreements/details/pccw", w{"a0.nic.pccw", "a2.nic.pccw", "b0.nic.pccw", "c0.nic.pccw"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.pccw", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1078] = Zone{"pe", r, z[3840:3849], 0xa0, e, "https://punto.pe/", w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, "kero.yachay.pe", e, n, t}
	z[1079] = Zone{"persiangulf", r, x, 0x100c0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1079() {
	z[1080] = Zone{"pet", r, x, 0x40, "Identity Digital Limited", "http://nic.pet/", w{"a0.nic.pet", "a2.nic.pet", "b0.nic.pet", "c0.nic.pet"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.pet", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1081] = Zone{"pets", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1082] = Zone{"pf", r, z[3849:3854], 0xa0, e, e, w{"ns1.mana.pf", "ns2.mana.pf"}, n, n, n, "whois.registry.pf", e, n, f}
	z[1083] = Zone{"pfizer", r, x, 0x42, "Pfizer Inc.", "https://www.nic.pfizer/", w{"a.nic.pfizer", "b.nic.pfizer", "c.nic.pfizer", "ns1.dns.nic.pfizer", "ns2.dns.nic.pfizer", "ns3.dns.nic.pfizer"}, n, n, w{"es"}, "whois.nic.pfizer", e, w{"https://rdap.nic.pfizer/"}, t}
	z[1084] = Zone{"pg", r, z[3854:3860], 0xa0, e, e, w{"munnari.oz.au", "ns.uu.net", "ns1.tiare.net.pg", "ns1.unitech.ac.pg", "ns2.tiare.net.pg"}, n, n, n, e, e, n, f}
	z[1085] = Zone{"ph", r, z[3860:3869], 0xa0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, "http://www.dot.ph/whois", n, f}
	z[1086] = Zone{"pharmacy", r, x, 0x40, "National Association of Boards of Pharmacy", "https://nic.pharmacy/", w{"dns1.nic.pharmacy", "dns2.nic.pharmacy", "dns3.nic.pharmacy", "dns4.nic.pharmacy", "dnsa.nic.pharmacy", "dnsb.nic.pharmacy", "dnsc.nic.pharmacy", "dnsd.nic.pharmacy"}, n, n, w{"es"}, "whois.nic.pharmacy", e, w{"https://rdap.nominet.uk/pharmacy/"}, t}
	z[1087] = Zone{"phd", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
}

//go:noinline
func i1087() {
	z[1088] = Zone{"philips", r, x, 0x42, "Koninklijke Philips N.V.", "http://nic.philips/", w{"a.nic.philips", "b.nic.philips", "c.nic.philips", "x.nic.philips", "y.nic.philips", "z.nic.philips"}, n, n, n, "whois.nic.philips", e, w{"https://rdap.nic.philips/"}, f}
	z[1089] = Zone{"phone", r, x, 0x40, "Dish DBS Corporation", "https://www.dishtlds.com/phone/", w{"a0.nic.phone", "a2.nic.phone", "b0.nic.phone", "c0.nic.phone"}, n, n, n, "whois.nic.phone", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1090] = Zone{"photo", r, x, 0x40, "Registry Services, LLC", "https://nic.photo/", w{"a.nic.photo", "b.nic.photo", "c.nic.photo", "x.nic.photo", "y.nic.photo", "z.nic.photo"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.photo", e, w{"https://rdap.nic.photo/"}, t}
	z[1091] = Zone{"photography", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.photography", "v0n1.nic.photography", "v0n2.nic.photography", "v0n3.nic.photography", "v2n0.nic.photography", "v2n1.nic.photography"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.photography", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1092] = Zone{"photos", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.photos", "v0n1.nic.photos", "v0n2.nic.photos", "v0n3.nic.photos", "v2n0.nic.photos", "v2n1.nic.photos"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.photos", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1093] = Zone{"physio", r, x, 0x40, "PhysBiz Pty Ltd", "http://nic.physio/", w{"a.nic.physio", "b.nic.physio", "c.nic.physio", "x.nic.physio", "y.nic.physio", "z.nic.physio"}, n, n, n, "whois.nic.physio", e, w{"https://rdap.nic.physio/"}, f}
	z[1094] = Zone{"piaget", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.piaget", e, n, f}
	z[1095] = Zone{"pics", r, x, 0x40, "XYZ.COM LLC", "https://nic.pics/", w{"a.nic.pics", "b.nic.pics", "c.nic.pics", "d.nic.pics"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.pics", e, w{"https://rdap.centralnic.com/pics/"}, t}
}

//go:noinline
func i1095() {
	z[1096] = Zone{"pictet", r, x, 0x42, "Pictet Europe S.A.", "https://www.icann.org/en/registry-agreements/details/pictet", w{"v0n0.nic.pictet", "v0n1.nic.pictet", "v0n2.nic.pictet", "v0n3.nic.pictet", "v2n0.nic.pictet", "v2n1.nic.pictet"}, n, n, n, "whois.nic.pictet", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1097] = Zone{"pictures", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.pictures", "v0n1.nic.pictures", "v0n2.nic.pictures", "v0n3.nic.pictures", "v2n0.nic.pictures", "v2n1.nic.pictures"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.pictures", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1098] = Zone{"pid", r, x, 0x40, "Top Level Spectrum, Inc.", "http://www.nic.pid/", w{"a.nic.pid", "b.nic.pid", "c.nic.pid", "d.nic.pid"}, n, n, w{"ar"}, "whois.nic.pid", e, w{"https://rdap.centralnic.com/pid/"}, t}
	z[1099] = Zone{"pin", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.pin/", w{"dns1.nic.pin", "dns2.nic.pin", "dns3.nic.pin", "dns4.nic.pin", "dnsa.nic.pin", "dnsb.nic.pin", "dnsc.nic.pin", "dnsd.nic.pin"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.pin", e, w{"https://rdap.nominet.uk/pin/"}, t}
	z[1100] = Zone{"ping", r, x, 0x42, "Ping Registry Provider, Inc.", "https://nic.ping/", w{"a.nic.ping", "b.nic.ping", "c.nic.ping", "ns1.dns.nic.ping", "ns2.dns.nic.ping", "ns3.dns.nic.ping"}, n, n, n, "whois.nic.ping", e, w{"https://rdap.nic.ping/"}, f}
	z[1101] = Zone{"pink", r, x, 0x40, "Identity Digital Limited", "http://nic.pink/", w{"a0.nic.pink", "a2.nic.pink", "b0.nic.pink", "b2.nic.pink", "c0.nic.pink"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.pink", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1102] = Zone{"pioneer", r, x, 0x42, "Pioneer Corporation", "https://nic.pioneer/", w{"dns1.nic.pioneer", "dns2.nic.pioneer", "dns3.nic.pioneer", "dns4.nic.pioneer", "dnsa.nic.pioneer", "dnsb.nic.pioneer", "dnsc.nic.pioneer", "dnsd.nic.pioneer"}, n, n, w{"ja"}, "whois.nic.pioneer", e, w{"https://rdap.nominet.uk/pioneer/"}, t}
	z[1103] = Zone{"piperlime", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1103() {
	z[1104] = Zone{"pitney", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1105] = Zone{"pizza", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.pizza", "v0n1.nic.pizza", "v0n2.nic.pizza", "v0n3.nic.pizza", "v2n0.nic.pizza", "v2n1.nic.pizza"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.pizza", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1106] = Zone{"pk", r, z[3869:3884], 0xa0, e, e, w{"root-c1.pknic.pk", "root-c2.pknic.pk", "root-e.pknic.pk", "root-s.pknic.pk"}, n, n, n, "whois.pknic.net.pk", e, n, f}
	z[1107] = Zone{"pl", r, z[3884:4052], 0xa0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, w{"be-PL", "bg-PL", "ca-PL", "cs-PL", "da-PL", "de-PL", "el-PL", "eo-PL", "es-PL", "et-PL", "fi-PL", "fr-PL", "ga-PL", "he-PL", "hr-PL", "hu-PL", "is-PL", "it-PL", "lb-PL", "lt-PL", "lv-PL", "mk-PL", "mt-PL", "nl-PL", "no-PL", "pl-PL", "pt-PL", "ro-PL", "ru-PL", "sk-PL", "sl-PL", "sq-PL", "sr-PL", "sv-PL", "tr-PL", "uk-PL"}, "whois.dns.pl", e, n, t}
	z[1108] = Zone{"place", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.place", "v0n1.nic.place", "v0n2.nic.place", "v0n3.nic.place", "v2n0.nic.place", "v2n1.nic.place"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.place", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1109] = Zone{"play", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1110] = Zone{"playstation", r, x, 0x42, "Sony Interactive Entertainment Inc.", "https://nic.playstation/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.playstation", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1111] = Zone{"plumbing", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.plumbing", "v0n1.nic.plumbing", "v0n2.nic.plumbing", "v0n3.nic.plumbing", "v2n0.nic.plumbing", "v2n1.nic.plumbing"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.plumbing", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1111() {
	z[1112] = Zone{"plus", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.plus", "v0n1.nic.plus", "v0n2.nic.plus", "v0n3.nic.plus", "v2n0.nic.plus", "v2n1.nic.plus"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.plus", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1113] = Zone{"pm", r, x, 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.pm", e, w{"https://rdap.nic.pm/"}, t}
	z[1114] = Zone{"pn", r, z[4052:4060], 0xa0, e, "https://nic.pn/", w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, "https://nic.pn/whois.html", w{"https://rdap.nominet.uk/pn/"}, t}
	z[1115] = Zone{"pnc", r, x, 0x42, "PNC Domain Co., LLC", "https://www.pnc.com/en/security-privacy/nic.html", w{"a0.nic.pnc", "a2.nic.pnc", "b0.nic.pnc", "c0.nic.pnc"}, n, n, n, "whois.nic.pnc", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1116] = Zone{"pohl", r, x, 0x42, "Deutsche Vermögensberatung Aktiengesellschaft DVAG", "https://www.icann.org/en/registry-agreements/details/pohl", w{"a.nic.pohl", "b.nic.pohl", "c.nic.pohl", "d.nic.pohl"}, n, n, w{"de"}, "whois.nic.pohl", e, w{"https://rdap.centralnic.com/pohl/"}, t}
	z[1117] = Zone{"poker", r, x, 0x40, "Identity Digital Limited", "http://nic.poker/", w{"a0.nic.poker", "a2.nic.poker", "b0.nic.poker", "c0.nic.poker"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.poker", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1118] = Zone{"politie", r, x, 0x42, "Politie Nederland", "https://nic.politie/", w{"ns1.dns.politie", "ns3.dns.politie", "ns4.dns.politie"}, n, n, n, "whois.nic.politie", e, w{"https://rdap.nic.politie/"}, f}
	z[1119] = Zone{"polo", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1119() {
	z[1120] = Zone{"porn", r, x, 0x41, "ICM Registry PN LLC", "https://nic.porn/", w{"a.nic.porn", "b.nic.porn", "c.nic.porn", "x.nic.porn", "y.nic.porn", "z.nic.porn"}, n, n, w{"ar", "be", "bg", "bs", "cnr", "da", "de", "es", "fr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.porn", e, w{"https://rdap.nic.porn/"}, t}
	z[1121] = Zone{"post", r, x, 0x4040, "Universal Postal Union", "https://en.wikipedia.org/wiki/.post", w{"a0.post.afilias-nst.info", "a2.post.afilias-nst.info", "b0.post.afilias-nst.org", "b2.post.afilias-nst.org", "c0.post.afilias-nst.info", "d0.post.afilias-nst.org"}, n, n, n, "whois.dotpostregistry.net", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1122] = Zone{"pr", r, z[4060:4112], 0xa0, e, "https://domains.pr/", w{"a.lactld.org", "a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org", "pr-dns.denic.de"}, n, n, n, "whois.nic.pr", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1123] = Zone{"pramerica", r, x, 0x42, "Prudential Financial, Inc.", "http://nic.pramerica/nicpramerica/nic%20pramerica%20static%20page%20.html", w{"a.nic.pramerica", "b.nic.pramerica", "c.nic.pramerica", "ns1.dns.nic.pramerica", "ns2.dns.nic.pramerica", "ns3.dns.nic.pramerica"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.pramerica", e, w{"https://rdap.nic.pramerica/"}, t}
	z[1124] = Zone{"praxi", r, x, 0x42, "Praxi S.p.A.", "http://nic.praxi/", w{"a.nic.praxi", "b.nic.praxi", "c.nic.praxi", "ns1.dns.nic.praxi", "ns2.dns.nic.praxi", "ns3.dns.nic.praxi"}, n, n, n, "whois.nic.praxi", e, w{"https://rdap.nic.praxi/"}, f}
	z[1125] = Zone{"press", r, x, 0x40, "Radix Technologies Inc.", e, w{"a.nic.press", "b.nic.press", "e.nic.press", "f.nic.press"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.press", e, w{"https://rdap.centralnic.com/press/"}, t}
	z[1126] = Zone{"prime", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.prime/", w{"dns1.nic.prime", "dns2.nic.prime", "dns3.nic.prime", "dns4.nic.prime", "dnsa.nic.prime", "dnsb.nic.prime", "dnsc.nic.prime", "dnsd.nic.prime"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.prime", e, w{"https://rdap.nominet.uk/prime/"}, t}
	z[1127] = Zone{"pro", r, z[4112:4148], 0x40, "Identity Digital Limited", "http://nic.pro/", w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.pro", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1127() {
	z[1128] = Zone{"prod", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1129] = Zone{"productions", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.productions", "v0n1.nic.productions", "v0n2.nic.productions", "v0n3.nic.productions", "v2n0.nic.productions", "v2n1.nic.productions"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.productions", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1130] = Zone{"prof", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1131] = Zone{"progressive", r, x, 0x42, "Progressive Casualty Insurance Company", "https://www.progressive.com/support/gtld/", w{"a0.nic.progressive", "a2.nic.progressive", "b0.nic.progressive", "c0.nic.progressive"}, n, n, n, "whois.nic.progressive", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1132] = Zone{"promo", r, x, 0x40, "Identity Digital Limited", "http://nic.promo/", w{"a0.nic.promo", "a2.nic.promo", "b0.nic.promo", "c0.nic.promo"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.promo", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1133] = Zone{"properties", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.properties", "v0n1.nic.properties", "v0n2.nic.properties", "v0n3.nic.properties", "v2n0.nic.properties", "v2n1.nic.properties"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.properties", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1134] = Zone{"property", r, x, 0x40, "Digital Property Infrastructure Limited", "https://nic.property/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.property", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[1135] = Zone{"protection", r, x, 0x40, "XYZ.COM LLC", "https://nic.protection/", w{"a.nic.protection", "b.nic.protection", "e.nic.protection", "f.nic.protection"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.protection", e, w{"https://rdap.centralnic.com/protection/"}, t}
}

//go:noinline
func i1135() {
	z[1136] = Zone{"pru", r, x, 0x42, "Prudential Financial, Inc.", "http://nic.pru/nicpru/nic-%20pru.html", w{"a.nic.pru", "b.nic.pru", "c.nic.pru", "ns1.dns.nic.pru", "ns2.dns.nic.pru", "ns3.dns.nic.pru"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.pru", e, w{"https://rdap.nic.pru/"}, t}
	z[1137] = Zone{"prudential", r, x, 0x42, "Prudential Financial, Inc.", "http://nic.prudential/nicprudential/nic-%20prudential%20static%20page.html", w{"a.nic.prudential", "b.nic.prudential", "c.nic.prudential", "ns1.dns.nic.prudential", "ns2.dns.nic.prudential", "ns3.dns.nic.prudential"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.prudential", e, w{"https://rdap.nic.prudential/"}, t}
	z[1138] = Zone{"ps", r, z[4148:4157], 0xa0, e, e, w{"dns1.gov.ps", "fork.sth.dnsnode.net", "ns1.pnina.ps", "ps-ns.anycast.pch.net", "ps.cctld.authdns.ripe.net", "rip.psg.com"}, n, n, n, "whois.pnina.ps", e, n, f}
	z[1139] = Zone{"pt", r, z[4157:4166], 0xa0, e, e, w{"a.dns.pt", "b.dns.pt", "c.dns.pt", "d.dns.pt", "e.dns.pt", "g.dns.pt", "h.dns.pt", "ns.dns.br", "ns2.nic.fr"}, n, n, n, "whois.dns.pt", e, n, t}
	z[1140] = Zone{"pub", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.pub", "v0n1.nic.pub", "v0n2.nic.pub", "v0n3.nic.pub", "v2n0.nic.pub", "v2n1.nic.pub"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.pub", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1141] = Zone{"pw", r, z[4166:4177], 0xa0, e, e, w{"ns1.nic.pw", "ns2.nic.pw", "ns5.nic.pw", "ns6.nic.pw"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.pw", e, w{"https://rdap.centralnic.com/pw/"}, t}
	z[1142] = Zone{"pwc", r, x, 0x42, "PricewaterhouseCoopers LLP", "http://www.nic.pwc/", w{"a0.nic.pwc", "a2.nic.pwc", "b0.nic.pwc", "c0.nic.pwc"}, n, n, n, "whois.nic.pwc", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1143] = Zone{"py", r, z[4177:4185], 0xa8, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, "http://www.nic.py/consulta-datos.php", n, f}
}

//go:noinline
func i1143() {
	z[1144] = Zone{"qa", r, z[4185:4193], 0xa0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, "whois.registry.qa", e, n, f}
	z[1145] = Zone{"qpon", r, x, 0x40, "dotQPON LLC", e, w{"a.nic.qpon", "b.nic.qpon", "c.nic.qpon", "d.nic.qpon"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "it", "ja", "ko", "lt", "lv", "nl", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.qpon", e, w{"https://rdap.centralnic.com/qpon/"}, t}
	z[1146] = Zone{"qtel", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1147] = Zone{"quebec", r, x, 0xc4, "PointQuébec Inc", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, w{"Quebec", "CA-QC"}, w{"fr"}, "whois.nic.quebec", e, w{"https://rdap.nic.quebec/"}, t}
	z[1148] = Zone{"quest", r, x, 0x42, "XYZ.COM LLC", "https://nic.quest/", w{"a-cnic.nic.quest", "b-cnic.nic.quest", "c-cnic.nic.quest", "d-cnic.nic.quest"}, n, n, w{"ja", "ko", "mul-Grek", "mul-Latn", "ru", "zh"}, "whois.nic.quest", e, w{"https://rdap.centralnic.com/quest/"}, t}
	z[1149] = Zone{"qvc", r, x, 0x1042, "QVC, Inc.", e, n, n, n, n, "whois.nic.qvc", e, w{"https://rdap.nic.qvc/"}, f}
	z[1150] = Zone{"racing", r, x, 0x40, "Premier Registry Limited", "http://nic.racing/", w{"a.nic.racing", "b.nic.racing", "c.nic.racing", "ns1.dns.nic.racing", "ns2.dns.nic.racing", "ns3.dns.nic.racing"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.racing", e, w{"https://rdap.nic.racing/"}, t}
	z[1151] = Zone{"radio", r, x, 0x40, "European Broadcasting Union (EBU)", "https://www.nic.radio/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"el", "he", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Latn", "zh"}, "whois.nic.radio", e, w{"https://rdap.nic.radio/"}, t}
}

//go:noinline
func i1151() {
	z[1152] = Zone{"raid", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.raid", e, w{"https://tld-rdap.verisign.com/raid/v1/"}, f}
	z[1153] = Zone{"ram", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1154] = Zone{"re", r, z[4193:4196], 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.re", e, w{"https://rdap.nic.re/"}, t}
	z[1155] = Zone{"read", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.read/", w{"dns1.nic.read", "dns2.nic.read", "dns3.nic.read", "dns4.nic.read", "dnsa.nic.read", "dnsb.nic.read", "dnsc.nic.read", "dnsd.nic.read"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.read", e, w{"https://rdap.nominet.uk/read/"}, t}
	z[1156] = Zone{"realestate", r, x, 0x40, "dotRealEstate LLC", "https://home.realestate/", w{"dns1.nic.realestate", "dns2.nic.realestate", "dns3.nic.realestate", "dns4.nic.realestate", "dnsa.nic.realestate", "dnsb.nic.realestate", "dnsc.nic.realestate", "dnsd.nic.realestate"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.realestate", e, w{"https://rdap.nominet.uk/realestate/"}, t}
	z[1157] = Zone{"realtor", r, x, 0x40, "Real Estate Domains LLC", "http://nic.realtor/", w{"dns1.nic.realtor", "dns2.nic.realtor", "dns3.nic.realtor", "dns4.nic.realtor", "dnsa.nic.realtor", "dnsb.nic.realtor", "dnsc.nic.realtor", "dnsd.nic.realtor"}, n, n, n, "whois.nic.realtor", e, w{"https://rdap.nominet.uk/realtor/"}, f}
	z[1158] = Zone{"realty", r, x, 0x40, "Internet Naming Company LLC", "https://identity.digital/", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, w{"ar"}, "whois.nic.realty", e, w{"https://rdap.registry.click/rdap/"}, t}
	z[1159] = Zone{"recipes", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.recipes", "v0n1.nic.recipes", "v0n2.nic.recipes", "v0n3.nic.recipes", "v2n0.nic.recipes", "v2n1.nic.recipes"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.recipes", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1159() {
	z[1160] = Zone{"red", r, x, 0x40, "Identity Digital Limited", "http://nic.red/", w{"a0.nic.red", "a2.nic.red", "b0.nic.red", "c0.nic.red"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.red", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1161] = Zone{"redken", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1162] = Zone{"redstone", r, x, 0x42, "Redstone Haute Couture Co., Ltd.", "http://nic.redstone/", w{"a0.nic.redstone", "a2.nic.redstone", "b0.nic.redstone", "c0.nic.redstone"}, n, n, n, "whois.nic.redstone", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1163] = Zone{"redumbrella", r, x, 0x42, "Travelers TLD, LLC", "https://nic.redumbrella/", w{"a0.nic.redumbrella", "a2.nic.redumbrella", "b0.nic.redumbrella", "c0.nic.redumbrella"}, n, n, n, "whois.nic.redumbrella", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1164] = Zone{"rehab", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.rehab", "v0n1.nic.rehab", "v0n2.nic.rehab", "v0n3.nic.rehab", "v2n0.nic.rehab", "v2n1.nic.rehab"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.rehab", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1165] = Zone{"reise", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.reise", "v0n1.nic.reise", "v0n2.nic.reise", "v0n3.nic.reise", "v2n0.nic.reise", "v2n1.nic.reise"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.reise", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1166] = Zone{"reisen", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.reisen", "v0n1.nic.reisen", "v0n2.nic.reisen", "v0n3.nic.reisen", "v2n0.nic.reisen", "v2n1.nic.reisen"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.reisen", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1167] = Zone{"reit", r, x, 0x50, "National Association of Real Estate Investment Trusts, Inc.", "https://nic.reit/", w{"a.nic.reit", "b.nic.reit", "c.nic.reit", "d.nic.reit"}, n, n, n, "whois.nic.reit", e, w{"https://rdap.centralnic.com/reit/"}, f}
}

//go:noinline
func i1167() {
	z[1168] = Zone{"reliance", r, x, 0x42, "Reliance Industries Limited", "http://nic.reliance/", w{"a0.nic.reliance", "a2.nic.reliance", "b0.nic.reliance", "c0.nic.reliance"}, n, n, n, "whois.nic.reliance", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1169] = Zone{"ren", r, x, 0x42, "ZDNS International Limited", "http://nic.ren/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.ren", e, w{"https://rdap.zdnsgtld.com/ren/"}, t}
	z[1170] = Zone{"rent", r, x, 0x40, "XYZ.COM LLC", "https://nic.rent/", w{"a.nic.rent", "b.nic.rent", "c.nic.rent", "d.nic.rent"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.rent", e, w{"https://rdap.centralnic.com/rent/"}, t}
	z[1171] = Zone{"rentals", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.rentals", "v0n1.nic.rentals", "v0n2.nic.rentals", "v0n3.nic.rentals", "v2n0.nic.rentals", "v2n1.nic.rentals"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.rentals", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1172] = Zone{"repair", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.repair", "v0n1.nic.repair", "v0n2.nic.repair", "v0n3.nic.repair", "v2n0.nic.repair", "v2n1.nic.repair"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.repair", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1173] = Zone{"report", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.report", "v0n1.nic.report", "v0n2.nic.report", "v0n3.nic.report", "v2n0.nic.report", "v2n1.nic.report"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.report", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1174] = Zone{"republican", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.republican", "v0n1.nic.republican", "v0n2.nic.republican", "v0n3.nic.republican", "v2n0.nic.republican", "v2n1.nic.republican"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.republican", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1175] = Zone{"rest", r, x, 0x40, "Punto 2012 Sociedad Anonima Promotora de Inversion de Capital Variable", e, w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, w{"mul-Latn"}, "whois.nic.rest", e, w{"https://rdap.centralnic.com/rest/"}, t}
}

//go:noinline
func i1175() {
	z[1176] = Zone{"restaurant", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.restaurant", "v0n1.nic.restaurant", "v0n2.nic.restaurant", "v0n3.nic.restaurant", "v2n0.nic.restaurant", "v2n1.nic.restaurant"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.restaurant", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1177] = Zone{"retirement", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1178] = Zone{"review", r, x, 0x40, "dot Review Limited", "http://nic.review/", w{"a.nic.review", "b.nic.review", "c.nic.review", "ns1.dns.nic.review", "ns2.dns.nic.review", "ns3.dns.nic.review"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.review", e, w{"https://rdap.nic.review/"}, t}
	z[1179] = Zone{"reviews", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.reviews", "v0n1.nic.reviews", "v0n2.nic.reviews", "v0n3.nic.reviews", "v2n0.nic.reviews", "v2n1.nic.reviews"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.reviews", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1180] = Zone{"rexroth", r, x, 0x42, "Robert Bosch GMBH", "https://apps.boschrexroth.com/microsites/nic_rexroth/index.html", w{"a0.nic.rexroth", "a2.nic.rexroth", "b0.nic.rexroth", "c0.nic.rexroth"}, n, n, n, "whois.nic.rexroth", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1181] = Zone{"rich", r, x, 0x40, "iRegistry GmbH", "https://nic.rich/", w{"a0.nic.rich", "a2.nic.rich", "b0.nic.rich", "c0.nic.rich"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.rich", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1182] = Zone{"richardli", r, x, 0x42, "Pacific Century Asset Management (HK) Limited", "https://www.icann.org/en/registry-agreements/details/richardli", w{"a0.nic.richardli", "a2.nic.richardli", "b0.nic.richardli", "c0.nic.richardli"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.richardli", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1183] = Zone{"ricoh", r, x, 0x42, "Ricoh Company, Ltd.", "https://nic.ricoh/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.ricoh", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i1183() {
	z[1184] = Zone{"rightathome", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.rightathome", e, n, f}
	z[1185] = Zone{"ril", r, x, 0x42, "Reliance Industries Limited", "http://nic.ril/", w{"a0.nic.ril", "a2.nic.ril", "b0.nic.ril", "c0.nic.ril"}, n, n, n, "whois.nic.ril", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1186] = Zone{"rio", r, x, 0xc4, "Empresa Municipal de Informática SA - IPLANRIO", "https://nic.rio/", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, w{"Rio de Janeiro", "BR-RJ"}, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
	z[1187] = Zone{"rip", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.rip", "v0n1.nic.rip", "v0n2.nic.rip", "v0n3.nic.rip", "v2n0.nic.rip", "v2n1.nic.rip"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.rip", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1188] = Zone{"rmit", r, x, 0x1042, "Royal Melbourne Institute of Technology", e, n, n, n, n, "whois.nic.rmit", e, w{"https://rdap.nic.rmit/"}, f}
	z[1189] = Zone{"ro", r, z[4196:4213], 0xa0, e, e, w{"dns-at.rotld.ro", "dns-c.rotld.ro", "dns-ro.denic.de", "primary.rotld.ro", "sec-dns-a.rotld.ro", "sec-dns-b.rotld.ro"}, n, n, n, "whois.rotld.ro", e, n, f}
	z[1190] = Zone{"rocher", r, x, 0x1042, "Ferrero Trading Lux S.A.", "https://nic.rocher/", n, n, n, w{"da", "de", "es", "fi", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.rocher", e, w{"https://rdap.nic.rocher/"}, t}
	z[1191] = Zone{"rocks", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.rocks", "v0n1.nic.rocks", "v0n2.nic.rocks", "v0n3.nic.rocks", "v2n0.nic.rocks", "v2n1.nic.rocks"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.rocks", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1191() {
	z[1192] = Zone{"rockwool", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1193] = Zone{"rodeo", r, x, 0x40, "Registry Services, LLC", "http://nic.rodeo/", w{"a.nic.rodeo", "b.nic.rodeo", "c.nic.rodeo", "x.nic.rodeo", "y.nic.rodeo", "z.nic.rodeo"}, n, n, w{"de", "es", "fr"}, "whois.nic.rodeo", e, w{"https://rdap.nic.rodeo/"}, t}
	z[1194] = Zone{"rogers", r, x, 0x42, "Rogers Communications Canada Inc.", "https://www.rogers.com/nic", w{"a0.nic.rogers", "a2.nic.rogers", "b0.nic.rogers", "c0.nic.rogers"}, n, n, n, "whois.nic.rogers", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1195] = Zone{"roma", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1196] = Zone{"room", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.room/", w{"dns1.nic.room", "dns2.nic.room", "dns3.nic.room", "dns4.nic.room", "dnsa.nic.room", "dnsb.nic.room", "dnsc.nic.room", "dnsd.nic.room"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.room", e, w{"https://rdap.nominet.uk/room/"}, t}
	z[1197] = Zone{"root", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1198] = Zone{"rs", r, z[4213:4219], 0xa0, e, e, w{"a.nic.rs", "b.nic.rs", "f.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, w{"cnr-Latn", "cs", "de", "hr", "hu", "rmn-Latn", "ro", "sk", "sl", "sq", "sr", "sr-Latn"}, "whois.rnids.rs", e, n, t}
	z[1199] = Zone{"rsvp", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
}

//go:noinline
func i1199() {
	z[1200] = Zone{"ru", r, z[4219:4247], 0xa0, e, e, w{"a.dns.ripn.net", "b.dns.ripn.net", "d.dns.ripn.net", "e.dns.ripn.net", "f.dns.ripn.net"}, n, n, w{"ru-RU"}, "whois.tcinet.ru", e, n, t}
	z[1201] = Zone{"rugby", r, x, 0x40, "World Rugby Strategic Developments Limited", "http://nic.rugby/", w{"a.nic.rugby", "b.nic.rugby", "c.nic.rugby", "x.nic.rugby", "y.nic.rugby", "z.nic.rugby"}, n, n, n, "whois.nic.rugby", e, w{"https://rdap.centralnic.com/rugby/"}, f}
	z[1202] = Zone{"ruhr", r, x, 0x440, "dotSaarland GmbH", e, w{"a.nic.ruhr", "b.nic.ruhr", "c.nic.ruhr", "d.nic.ruhr"}, n, w{"DE-NW"}, w{"de"}, "whois.nic.ruhr", e, w{"https://rdap.centralnic.com/ruhr/"}, t}
	z[1203] = Zone{"run", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.run", "v0n1.nic.run", "v0n2.nic.run", "v0n3.nic.run", "v2n0.nic.run", "v2n1.nic.run"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.run", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1204] = Zone{"rw", r, z[4247:4257], 0xa0, e, "https://ricta.org.rw/", w{"ans.dnsstudy.africa", "fork.sth.dnsnode.net", "ns-rw.afrinic.net", "ns1.ricta.org.rw", "ns2.ricta.org.rw", "ns3.ricta.org.rw", "pch.ricta.org.rw"}, n, n, n, "whois.ricta.org.rw", e, n, f}
	z[1205] = Zone{"rwe", r, x, 0x42, "RWE AG", "http://nic.rwe/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.rwe", e, w{"https://tld-rdap.verisign.com/rwe/v1/"}, t}
	z[1206] = Zone{"ryukyu", r, x, 0x440, "BRregistry, Inc.", "https://www.brregistry.com/geotlds/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"JP-46", "JP-47"}, w{"ja"}, "whois.nic.ryukyu", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1207] = Zone{"sa", r, z[4257:4265], 0xa0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, "whois.nic.net.sa", e, n, f}
}

//go:noinline
func i1207() {
	z[1208] = Zone{"saarland", r, x, 0x4c0, "dotSaarland GmbH", e, w{"a.nic.saarland", "b.nic.saarland", "c.nic.saarland", "d.nic.saarland"}, n, w{"DE-SL"}, w{"de"}, "whois.nic.saarland", e, w{"https://rdap.centralnic.com/saarland/"}, t}
	z[1209] = Zone{"safe", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.safe/", w{"dns1.nic.safe", "dns2.nic.safe", "dns3.nic.safe", "dns4.nic.safe", "dnsa.nic.safe", "dnsb.nic.safe", "dnsc.nic.safe", "dnsd.nic.safe"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.safe", e, w{"https://rdap.nominet.uk/safe/"}, t}
	z[1210] = Zone{"safety", r, x, 0x40, "Safety Registry Services, LLC.", "http://nic.safety/", w{"a.nic.safety", "b.nic.safety", "c.nic.safety", "ns1.dns.nic.safety", "ns2.dns.nic.safety", "ns3.dns.nic.safety"}, n, n, w{"es"}, "whois.nic.safety", e, w{"https://rdap.nic.safety/"}, t}
	z[1211] = Zone{"safeway", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1212] = Zone{"sakura", r, x, 0x42, "SAKURA Internet Inc.", "https://www.icann.org/en/registry-agreements/details/sakura", w{"tld1.nic.sakura", "tld2.nic.sakura", "tld3.nic.sakura", "tld5.nic.sakura"}, n, n, w{"ja"}, e, e, w{"https://rdap.nic.sakura/rdap/"}, t}
	z[1213] = Zone{"sale", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.sale", "v0n1.nic.sale", "v0n2.nic.sale", "v0n3.nic.sale", "v2n0.nic.sale", "v2n1.nic.sale"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.sale", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1214] = Zone{"salon", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.salon", "v0n1.nic.salon", "v0n2.nic.salon", "v0n3.nic.salon", "v2n0.nic.salon", "v2n1.nic.salon"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.salon", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1215] = Zone{"samsclub", r, x, 0x42, "Wal-Mart Stores, Inc.", "https://www.icann.org/en/registry-agreements/details/samsclub", w{"a0.nic.samsclub", "a2.nic.samsclub", "b0.nic.samsclub", "c0.nic.samsclub"}, n, n, n, "whois.nic.samsclub", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i1215() {
	z[1216] = Zone{"samsung", r, x, 0x42, "SAMSUNG SDS CO., LTD", "http://nic.samsung/", w{"ns1.samsung.centralnic-dns.com", "ns2.samsung.centralnic-dns.com", "ns3.samsung.centralnic-dns.com", "ns4.samsung.centralnic-dns.com", "tld01.akam.net", "tld02.akam.net", "tld03.akam.net", "tld04.akam.net"}, n, n, w{"ko"}, "whois.nic.samsung", e, w{"https://nic.samsung:8443/rdap/"}, t}
	z[1217] = Zone{"sandvik", r, x, 0x42, "Sandvik AB", "https://www.nic.sandvik/en/", w{"a.nic.sandvik", "b.nic.sandvik", "c.nic.sandvik", "x.nic.sandvik", "y.nic.sandvik", "z.nic.sandvik"}, n, n, w{"da", "fr", "pl", "ru", "sv", "zh"}, "whois.nic.sandvik", e, w{"https://rdap.nic.sandvik/"}, t}
	z[1218] = Zone{"sandvikcoromant", r, x, 0x42, "Sandvik AB", "https://www.sandvik.coromant.com/en-gb/pages/nic-sandvikcoromant.aspx", w{"a.nic.sandvikcoromant", "b.nic.sandvikcoromant", "c.nic.sandvikcoromant", "x.nic.sandvikcoromant", "y.nic.sandvikcoromant", "z.nic.sandvikcoromant"}, n, n, w{"da", "fr", "pl", "ru", "sv", "zh"}, "whois.nic.sandvikcoromant", e, w{"https://rdap.nic.sandvikcoromant/"}, t}
	z[1219] = Zone{"sanofi", r, x, 0x42, "Sanofi", "http://nic.sanofi/", w{"a0.nic.sanofi", "a2.nic.sanofi", "b0.nic.sanofi", "c0.nic.sanofi"}, n, n, n, "whois.nic.sanofi", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1220] = Zone{"sap", r, x, 0x42, "SAP AG", "https://www.icann.org/en/registry-agreements/details/sap", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"ar", "ca", "cs", "da", "de", "el", "es", "fi", "fr", "he", "hr", "hu", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "tr", "uk", "zh"}, "whois.nic.sap", e, w{"https://rdap.nic.sap/"}, t}
	z[1221] = Zone{"sapo", r, x, 0x1042, e, e, n, n, n, n, e, e, n, f}
	z[1222] = Zone{"sapphire", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1223] = Zone{"sarl", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.sarl", "v0n1.nic.sarl", "v0n2.nic.sarl", "v0n3.nic.sarl", "v2n0.nic.sarl", "v2n1.nic.sarl"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.sarl", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1223() {
	z[1224] = Zone{"sas", r, x, 0x42, "Research IP LLC", "https://nic.sas/", w{"a.nic.sas", "b.nic.sas", "c.nic.sas", "ns1.dns.nic.sas", "ns2.dns.nic.sas", "ns3.dns.nic.sas"}, n, n, w{"es"}, "whois.nic.sas", e, w{"https://rdap.nic.sas/"}, t}
	z[1225] = Zone{"save", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.save/", w{"dns1.nic.save", "dns2.nic.save", "dns3.nic.save", "dns4.nic.save", "dnsa.nic.save", "dnsb.nic.save", "dnsc.nic.save", "dnsd.nic.save"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.save", e, w{"https://rdap.nominet.uk/save/"}, t}
	z[1226] = Zone{"saxo", r, x, 0x42, "Saxo Bank A/S", "http://nic.saxo/", w{"a.nic.saxo", "b.nic.saxo", "c.nic.saxo", "x.nic.saxo", "y.nic.saxo", "z.nic.saxo"}, n, n, n, "whois.nic.saxo", e, w{"https://rdap.nic.saxo/"}, f}
	z[1227] = Zone{"sb", r, z[4265:4270], 0xa0, e, e, w{"ns1.anycastdns.cz", "ns2.anycastdns.cz", "pch.nic.sb"}, n, n, n, "whois.nic.net.sb", e, n, f}
	z[1228] = Zone{"sbi", r, x, 0x42, "STATE BANK OF INDIA", "http://nic.sbi/", w{"a0.nic.sbi", "a2.nic.sbi", "b0.nic.sbi", "c0.nic.sbi"}, n, n, n, "whois.nic.sbi", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1229] = Zone{"sbs", r, x, 0x42, "ShortDot SA", "https://shortdot.bond/sbs/", w{"a.nic.sbs", "b.nic.sbs", "c.nic.sbs", "d.nic.sbs"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.sbs", e, w{"https://rdap.centralnic.com/sbs/"}, t}
	z[1230] = Zone{"sc", r, z[4270:4275], 0xe0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "ns1.nic.sc"}, n, n, w{"zh-CN"}, "whois2.afilias-grs.net", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1231] = Zone{"sca", r, x, 0x1042, "SVENSKA CELLULOSA AKTIEBOLAGET SCA (publ)", "http://nic.sca/", n, n, n, w{"ar", "ru", "sv", "zh"}, "whois.nic.sca", e, w{"https://tld-rdap.verisign.com/sca/v1/"}, t}
}

//go:noinline
func i1231() {
	z[1232] = Zone{"scb", r, x, 0x42, "The Siam Commercial Bank Public Company Limited (\"SCB\")", "https://www.nic.scb/", w{"a.nic.scb", "c.nic.scb", "rz.nic.scb"}, n, n, w{"th"}, "whois.nic.scb", e, w{"https://rdap.nic.scb/"}, t}
	z[1233] = Zone{"schaeffler", r, x, 0x42, "Schaeffler Technologies AG & Co. KG", "https://www.nic.schaeffler/en/", w{"a.dns.nic.schaeffler", "m.dns.nic.schaeffler", "n.dns.nic.schaeffler"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.afilias-srs.net", e, w{"https://rdap.nic.schaeffler/"}, t}
	z[1234] = Zone{"schmidt", r, x, 0x42, "SCHMIDT GROUPE S.A.S.", "http://nic.schmidt/", w{"a.nic.schmidt", "b.nic.schmidt", "c.nic.schmidt", "x.nic.schmidt", "y.nic.schmidt", "z.nic.schmidt"}, n, n, n, "whois.nic.schmidt", e, w{"https://rdap.nic.schmidt/"}, f}
	z[1235] = Zone{"scholarships", r, x, 0x40, "Scholarships.com, LLC", "http://nic.scholarships/", w{"a0.nic.scholarships", "a2.nic.scholarships", "b0.nic.scholarships", "c0.nic.scholarships"}, n, n, n, "whois.nic.scholarships", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1236] = Zone{"school", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.school", "v0n1.nic.school", "v0n2.nic.school", "v0n3.nic.school", "v2n0.nic.school", "v2n1.nic.school"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.school", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1237] = Zone{"schule", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.schule", "v0n1.nic.schule", "v0n2.nic.schule", "v0n3.nic.schule", "v2n0.nic.schule", "v2n1.nic.schule"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.schule", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1238] = Zone{"schwarz", r, x, 0x42, "Schwarz Domains und Services GmbH & Co. KG", "https://nic.schwarz/", w{"a.nic.schwarz", "b.nic.schwarz", "c.nic.schwarz", "d.nic.schwarz"}, n, n, w{"mul-Cyrl", "mul-Grek", "mul-Latn"}, "whois.nic.schwarz", e, w{"https://rdap.centralnic.com/schwarz/"}, t}
	z[1239] = Zone{"schwarzgroup", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1239() {
	z[1240] = Zone{"science", r, x, 0x40, "dot Science Limited", "http://nic.science/", w{"a.nic.science", "b.nic.science", "c.nic.science", "ns1.dns.nic.science", "ns2.dns.nic.science", "ns3.dns.nic.science"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.science", e, w{"https://rdap.nic.science/"}, t}
	z[1241] = Zone{"scjohnson", r, x, 0x1042, "Johnson Shareholdings, Inc.", e, n, n, n, n, "whois.nic.scjohnson", e, w{"https://tld-rdap.verisign.com/scjohnson/v1/"}, f}
	z[1242] = Zone{"scor", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.scor", e, n, f}
	z[1243] = Zone{"scot", r, x, 0x440, "Dot Scot Registry Limited", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, w{"GB-SCT"}, w{"mul-Latn"}, "whois.nic.scot", e, w{"https://rdap.nic.scot/"}, t}
	z[1244] = Zone{"sd", r, z[4275:4283], 0xa0, e, e, w{"ans1.canar.sd", "ans1.sis.sd", "ans2.canar.sd", "ans2.sis.sd", "ns-sd.afrinic.net", "pch.sis.sd", "sd.cctld.authdns.ripe.net"}, n, n, n, "whois.nic.sd", e, w{"https://rdap.nic.sd/"}, f}
	z[1245] = Zone{"se", r, z[4283:4315], 0xa0, e, e, w{"a.ns.se", "b.ns.se", "c.ns.se", "f.ns.se", "g.ns.se", "i.ns.se", "m.ns.se", "x.ns.se", "y.ns.se", "z.ns.se"}, n, n, w{"mul-Latn", "yi"}, "whois.iis.se", e, n, t}
	z[1246] = Zone{"search", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1247] = Zone{"seat", r, x, 0x42, "SEAT, S.A. (Sociedad Unipersonal)", "https://www.icann.org/en/registry-agreements/details/seat", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Latn"}, "whois.nic.seat", e, w{"https://rdap.nic.seat/"}, t}
}

//go:noinline
func i1247() {
	z[1248] = Zone{"secure", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.secure/", w{"dns1.nic.secure", "dns2.nic.secure", "dns3.nic.secure", "dns4.nic.secure", "dnsa.nic.secure", "dnsb.nic.secure", "dnsc.nic.secure", "dnsd.nic.secure"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.secure", e, w{"https://rdap.nominet.uk/secure/"}, t}
	z[1249] = Zone{"security", r, x, 0x40, "XYZ.COM LLC", "https://nic.security/", w{"a.nic.security", "b.nic.security", "e.nic.security", "f.nic.security"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.security", e, w{"https://rdap.centralnic.com/security/"}, t}
	z[1250] = Zone{"seek", r, x, 0x42, "Seek Limited", "http://www.nic.seek/", w{"a.nic.seek", "b.nic.seek", "c.nic.seek", "x.nic.seek", "y.nic.seek", "z.nic.seek"}, n, n, n, "whois.nic.seek", e, w{"https://rdap.nic.seek/"}, f}
	z[1251] = Zone{"select", r, x, 0x42, "Registry Services, LLC", "https://www.be.select/", w{"a.nic.select", "b.nic.select", "c.nic.select", "x.nic.select", "y.nic.select", "z.nic.select"}, n, n, n, "whois.nic.select", e, w{"https://rdap.nic.select/"}, f}
	z[1252] = Zone{"sener", r, x, 0x42, "Sener Ingeniería y Sistemas, S.A.", "http://nic.sener/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://tld-rdap.verisign.com/sener/v1/"}, t}
	z[1253] = Zone{"services", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.services", "v0n1.nic.services", "v0n2.nic.services", "v0n3.nic.services", "v2n0.nic.services", "v2n1.nic.services"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.services", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1254] = Zone{"ses", r, x, 0x1042, "SES", "http://nic.ses/", n, n, n, n, "whois.nic.ses", e, w{"https://tld-rdap.verisign.com/ses/v1/"}, f}
	z[1255] = Zone{"seven", r, x, 0x42, "Seven West Media Ltd", "http://nic.seven/", w{"a.nic.seven", "b.nic.seven", "c.nic.seven", "x.nic.seven", "y.nic.seven", "z.nic.seven"}, n, n, n, "whois.nic.seven", e, w{"https://rdap.nic.seven/"}, f}
}

//go:noinline
func i1255() {
	z[1256] = Zone{"sew", r, x, 0x42, "SEW-EURODRIVE GmbH & Co KG", "http://nic.sew/", w{"a0.nic.sew", "a2.nic.sew", "b0.nic.sew", "c0.nic.sew"}, n, n, n, "whois.nic.sew", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1257] = Zone{"sex", r, x, 0x41, "ICM Registry SX LLC", "https://nic.sex/", w{"a.nic.sex", "b.nic.sex", "c.nic.sex", "x.nic.sex", "y.nic.sex", "z.nic.sex"}, n, n, w{"ar", "be", "bg", "bs", "cnr", "da", "de", "es", "fr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.sex", e, w{"https://rdap.nic.sex/"}, t}
	z[1258] = Zone{"sexy", r, x, 0x41, "Internet Naming Company LLC", "https://nic.sexy/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "zh"}, "whois.nic.sexy", e, w{"https://whois.uniregistry.net/rdap/"}, t}
	z[1259] = Zone{"sfr", r, x, 0x42, "Societe Francaise du Radiotelephone - SFR", "http://nic.sfr/", w{"a.nic.sfr", "b.nic.sfr", "c.nic.sfr", "d.nic.sfr"}, n, n, w{"mul-Latn"}, "whois.nic.sfr", e, w{"https://rdap.centralnic.com/sfr/"}, t}
	z[1260] = Zone{"sg", r, z[4315:4321], 0xa0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "ns4.apnic.net", "pch.sgzones.sg"}, n, n, n, "whois.sgnic.sg", e, n, t}
	z[1261] = Zone{"sh", r, z[4321:4328], 0xa0, e, "https://www.nic.sh/", w{"a0.nic.sh", "a2.nic.sh", "b0.nic.sh", "c0.nic.sh"}, n, n, n, "whois.nic.sh", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1262] = Zone{"shangrila", r, x, 0x42, "Shangri‐La International Hotel Management Limited", "https://www.icann.org/en/registry-agreements/details/shangrila", w{"a0.nic.shangrila", "a2.nic.shangrila", "b0.nic.shangrila", "c0.nic.shangrila"}, n, n, n, "whois.nic.shangrila", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1263] = Zone{"sharp", r, x, 0x42, "Sharp Corporation", "https://nic.sharp/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i1263() {
	z[1264] = Zone{"shaw", r, x, 0x42, "Shaw Cablesystems G.P.", "https://www.shaw.ca/terms-of-use/", w{"a0.nic.shaw", "a2.nic.shaw", "b0.nic.shaw", "c0.nic.shaw"}, n, n, n, "whois.afilias-srs.net", e, w{"https://rdap.afilias-srs.net/rdap/shaw/"}, f}
	z[1265] = Zone{"shell", r, x, 0x42, "Shell Information Technology International Inc", "https://www.nic.shell/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.shell", e, w{"https://tld-rdap.verisign.com/shell/v1/"}, t}
	z[1266] = Zone{"shia", r, x, 0x40, "Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.", "http://nic.shia/", w{"a.ns.nic.shia", "b.ns.nic.shia", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, w{"Iran"}, n, "whois.nic.shia", e, w{"https://api.rdap.nic.shia/"}, f}
	z[1267] = Zone{"shiksha", r, x, 0x40, "Identity Digital Limited", "http://nic.shiksha/", w{"a0.nic.shiksha", "a2.nic.shiksha", "b0.nic.shiksha", "c0.nic.shiksha"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.shiksha", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1268] = Zone{"shoes", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.shoes", "v0n1.nic.shoes", "v0n2.nic.shoes", "v0n3.nic.shoes", "v2n0.nic.shoes", "v2n1.nic.shoes"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.shoes", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1269] = Zone{"shop", r, x, 0x40, "GMO Registry, Inc.", "https://get.shop/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ar", "de", "es", "fr", "ja", "ko", "nl", "pl", "pt", "ru", "tr", "zh", "zh-Hans", "zh-Hant"}, "whois.nic.shop", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1270] = Zone{"shopping", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.shopping", "v0n1.nic.shopping", "v0n2.nic.shopping", "v0n3.nic.shopping", "v2n0.nic.shopping", "v2n1.nic.shopping"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.shopping", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1271] = Zone{"shopyourway", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1271() {
	z[1272] = Zone{"shouji", r, x, 0x42, "Beijing Qihu Keji Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/shouji", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, n, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/"}, f}
	z[1273] = Zone{"show", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.show", "v0n1.nic.show", "v0n2.nic.show", "v0n3.nic.show", "v2n0.nic.show", "v2n1.nic.show"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.show", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1274] = Zone{"showtime", r, x, 0x1042, "CBS Domains Inc.", "https://nic.showtime/", n, n, n, n, "whois.nic.showtime", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1275] = Zone{"shriram", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[1276] = Zone{"si", r, z[4328:4335], 0xa0, e, e, w{"b.dns.si", "f.dns.si", "h.dns.si", "i.dns.si", "k.dns.si", "l.dns.si"}, n, n, w{"mul-Latn"}, "whois.register.si", e, n, t}
	z[1277] = Zone{"silk", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.silk/", w{"dns1.nic.silk", "dns2.nic.silk", "dns3.nic.silk", "dns4.nic.silk", "dnsa.nic.silk", "dnsb.nic.silk", "dnsc.nic.silk", "dnsd.nic.silk"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.silk", e, w{"https://rdap.nominet.uk/silk/"}, t}
	z[1278] = Zone{"sina", r, x, 0x42, "Sina Corporation", "https://www.icann.org/en/registry-agreements/details/sina", w{"a0.nic.sina", "a2.nic.sina", "b0.nic.sina", "c0.nic.sina"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.sina", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1279] = Zone{"singles", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.singles", "v0n1.nic.singles", "v0n2.nic.singles", "v0n3.nic.singles", "v2n0.nic.singles", "v2n1.nic.singles"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.singles", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1279() {
	z[1280] = Zone{"site", r, x, 0x40, "Radix Technologies Inc.", "https://radix.website/dot-site", w{"a.nic.site", "b.nic.site", "e.nic.site", "f.nic.site"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.site", e, w{"https://rdap.centralnic.com/site/"}, t}
	z[1281] = Zone{"sj", r, x, 0xa8, e, e, w{"nac.no", "nn.uninett.no", "server.nordu.net", "x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[1282] = Zone{"sk", r, z[4335:4336], 0xa0, e, e, w{"a.tld.sk", "b.tld.sk", "c.tld.sk", "e.tld.sk", "f.tld.sk", "g.tld.sk", "h.tld.sk"}, n, n, n, "whois.sk-nic.sk", e, n, f}
	z[1283] = Zone{"ski", r, x, 0x40, "Identity Digital Limited", "http://nic.ski/", w{"a0.nic.ski", "a2.nic.ski", "b0.nic.ski", "c0.nic.ski"}, n, n, w{"de"}, "whois.nic.ski", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1284] = Zone{"skin", r, x, 0x40, "XYZ.COM LLC", "https://nic.skin/", w{"a.nic.skin", "b.nic.skin", "c.nic.skin", "d.nic.skin"}, n, n, w{"ja", "ko", "mul-Grek", "mul-Latn", "ru", "zh"}, "whois.nic.skin", e, w{"https://rdap.centralnic.com/skin/"}, t}
	z[1285] = Zone{"skolkovo", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1286] = Zone{"sky", r, x, 0x42, "Sky International AG", "https://www.skygroup.sky/nic-sky", w{"dns1.nic.sky", "dns2.nic.sky", "dns3.nic.sky", "dns4.nic.sky", "dnsa.nic.sky", "dnsb.nic.sky", "dnsc.nic.sky", "dnsd.nic.sky"}, n, n, w{"mul-Arab", "mul-Latn"}, "whois.nic.sky", e, w{"https://rdap.nominet.uk/sky/"}, t}
	z[1287] = Zone{"skydrive", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1287() {
	z[1288] = Zone{"skype", r, x, 0x42, "Microsoft Corporation", "https://nic.skype/", w{"a.nic.skype", "b.nic.skype", "c.nic.skype", "ns1.dns.nic.skype", "ns2.dns.nic.skype", "ns3.dns.nic.skype"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.skype", e, w{"https://rdap.nic.skype/"}, t}
	z[1289] = Zone{"sl", r, z[4336:4341], 0xa0, e, "https://www.nic.sl/", w{"ns1.neoip.com", "ns2.neoip.com"}, n, n, n, "whois.nic.sl", e, n, f}
	z[1290] = Zone{"sling", r, x, 0x42, "DISH Technologies L.L.C.", "https://www.icann.org/en/registry-agreements/details/sling", w{"a0.nic.sling", "a2.nic.sling", "b0.nic.sling", "c0.nic.sling"}, n, n, n, "whois.nic.sling", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1291] = Zone{"sm", r, x, 0xa0, e, e, w{"dns.intelcom.sm", "dns.omniway.sm", "ns3.telecomitalia.sm", "sm.cctld.authdns.ripe.net"}, n, n, n, "whois.nic.sm", e, n, f}
	z[1292] = Zone{"smart", r, x, 0x42, "Smart Communications, Inc. (SMART)", "https://www.icann.org/en/registry-agreements/details/smart", w{"a.nic.smart", "b.nic.smart", "c.nic.smart", "d.nic.smart"}, n, n, n, "whois.nic.smart", e, w{"https://rdap.centralnic.com/smart/"}, f}
	z[1293] = Zone{"smile", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.smile/", w{"dns1.nic.smile", "dns2.nic.smile", "dns3.nic.smile", "dns4.nic.smile", "dnsa.nic.smile", "dnsb.nic.smile", "dnsc.nic.smile", "dnsd.nic.smile"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.smile", e, w{"https://rdap.nominet.uk/smile/"}, t}
	z[1294] = Zone{"sn", r, z[4341:4348], 0xa8, e, e, w{"dns-tld.ird.fr", "ns-sn.nic.fr", "ns.ucad.sn", "ns1.sonatel.sn", "sn.cctld.authdns.ripe.net"}, n, n, n, "whois.nic.sn", e, n, f}
	z[1295] = Zone{"sncf", r, x, 0x42, "Société Nationale SNCF", "https://www.icann.org/en/registry-agreements/details/sncf", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.sncf", e, w{"https://rdap.nic.sncf/"}, t}
}

//go:noinline
func i1295() {
	z[1296] = Zone{"so", r, z[4348:4352], 0xa0, e, e, w{"d.nic.so", "e.nic.so"}, n, n, n, "whois.nic.so", e, n, f}
	z[1297] = Zone{"soccer", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.soccer", "v0n1.nic.soccer", "v0n2.nic.soccer", "v0n3.nic.soccer", "v2n0.nic.soccer", "v2n1.nic.soccer"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.soccer", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1298] = Zone{"social", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.social", "v0n1.nic.social", "v0n2.nic.social", "v0n3.nic.social", "v2n0.nic.social", "v2n1.nic.social"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.social", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1299] = Zone{"softbank", r, x, 0x42, "SoftBank Group Corp.", "https://nic.softbank/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.softbank", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1300] = Zone{"software", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.software", "v0n1.nic.software", "v0n2.nic.software", "v0n3.nic.software", "v2n0.nic.software", "v2n1.nic.software"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.software", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1301] = Zone{"sohu", r, x, 0x42, "Sohu.com Limited", "https://www.icann.org/en/registry-agreements/details/sohu", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.sohu", e, w{"https://rdap.zdnsgtld.com/sohu/"}, t}
	z[1302] = Zone{"solar", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.solar", "v0n1.nic.solar", "v0n2.nic.solar", "v0n3.nic.solar", "v2n0.nic.solar", "v2n1.nic.solar"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.solar", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1303] = Zone{"solutions", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.solutions", "v0n1.nic.solutions", "v0n2.nic.solutions", "v0n3.nic.solutions", "v2n0.nic.solutions", "v2n1.nic.solutions"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.solutions", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1303() {
	z[1304] = Zone{"song", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.song/", w{"a.nic.song", "b.nic.song", "c.nic.song", "ns1.dns.nic.song", "ns2.dns.nic.song", "ns3.dns.nic.song"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.song", e, w{"https://rdap.nic.song/"}, t}
	z[1305] = Zone{"sony", r, x, 0x42, "Sony Corporation", "https://nic.sony/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.sony", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1306] = Zone{"soy", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1307] = Zone{"spa", r, x, 0x50, "Asia Spa and Wellness Promotion Council Limited", "https://www.nic.spa/", w{"a0.nic.spa", "a2.nic.spa", "b0.nic.spa", "c0.nic.spa"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.spa", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1308] = Zone{"space", r, z[4352:4353], 0x40, "Radix Technologies Inc.", "https://radix.website/dot-space", w{"a.nic.space", "b.nic.space", "e.nic.space", "f.nic.space"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.space", e, w{"https://rdap.centralnic.com/space/"}, t}
	z[1309] = Zone{"spiegel", r, x, 0x1042, e, e, n, n, n, n, "whois.ksregistry.net", e, n, f}
	z[1310] = Zone{"sport", r, x, 0x40, "SportAccord", "https://nic.sport/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"el", "he", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Latn", "zh"}, "whois.nic.sport", e, w{"https://rdap.nic.sport/"}, t}
	z[1311] = Zone{"sports", r, x, 0x10040, e, "https://gtldresult.icann.org/applicationstatus/applicationdetails/604", n, n, n, n, e, e, n, t}
}

//go:noinline
func i1311() {
	z[1312] = Zone{"spot", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.spot/", w{"dns1.nic.spot", "dns2.nic.spot", "dns3.nic.spot", "dns4.nic.spot", "dnsa.nic.spot", "dnsb.nic.spot", "dnsc.nic.spot", "dnsd.nic.spot"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.spot", e, w{"https://rdap.nominet.uk/spot/"}, t}
	z[1313] = Zone{"spreadbetting", r, x, 0x1040, e, "https://www.icann.org/en/registry-agreements/details/spreadbetting", n, n, n, n, "whois.nic.spreadbetting", e, w{"https://rdap.centralnic.com/spreadbetting/"}, f}
	z[1314] = Zone{"sr", r, x, 0xe0, e, e, w{"ns1.sr.net", "ns2.sr.net"}, n, n, n, e, "http://www.register.sr/", n, f}
	z[1315] = Zone{"srl", r, x, 0x40, "InterNetX, Corp", "https://www.internetx.info/en/", w{"a0.nic.srl", "a2.nic.srl", "b0.nic.srl", "c0.nic.srl"}, n, n, n, "whois.nic.srl", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1316] = Zone{"srt", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[1317] = Zone{"ss", r, z[4353:4360], 0, e, e, w{"b.ns.tznic.or.tz", "ns-ss.afrinic.net", "pch.nic.ss", "ssnic.anycastdns.cz"}, n, n, n, "whois.nic.ss", e, n, t}
	z[1318] = Zone{"st", r, z[4360:4373], 0xa0, e, e, w{"dns-st.bahnhof.net", "ns1.bahnhof.net", "southeast-2.dns-au.st", "west-2.dns-us.st"}, n, n, n, "whois.nic.st", e, n, f}
	z[1319] = Zone{"stada", r, x, 0x42, "STADA Arzneimittel AG", "https://nic.stada/", w{"a0.nic.stada", "a2.nic.stada", "b0.nic.stada", "c0.nic.stada"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.stada", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1319() {
	z[1320] = Zone{"staples", r, x, 0x42, "Staples, Inc.", "https://www.icann.org/en/registry-agreements/details/staples", w{"a.nic.staples", "b.nic.staples", "c.nic.staples", "ns1.dns.nic.staples", "ns2.dns.nic.staples", "ns3.dns.nic.staples"}, n, n, w{"es"}, "whois.nic.staples", e, w{"https://rdap.nic.staples/"}, t}
	z[1321] = Zone{"star", r, x, 0x42, "Star India Private Limited", "https://www.icann.org/en/registry-agreements/details/star", w{"a0.nic.star", "a2.nic.star", "b0.nic.star", "c0.nic.star"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.star", e, w{"https://rdap.afilias-srs.net/rdap/star/"}, t}
	z[1322] = Zone{"starhub", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.starhub", e, n, f}
	z[1323] = Zone{"statebank", r, x, 0x42, "STATE BANK OF INDIA", "http://nic.statebank/", w{"a0.nic.statebank", "a2.nic.statebank", "b0.nic.statebank", "c0.nic.statebank"}, n, n, n, "whois.nic.statebank", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1324] = Zone{"statefarm", r, x, 0x42, "State Farm Mutual Automobile Insurance Company", "https://nic.statefarm/", w{"a.nic.statefarm", "b.nic.statefarm", "c.nic.statefarm", "ns1.dns.nic.statefarm", "ns2.dns.nic.statefarm", "ns3.dns.nic.statefarm"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.statefarm", e, w{"https://rdap.nic.statefarm/"}, t}
	z[1325] = Zone{"statoil", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.statoil", e, n, f}
	z[1326] = Zone{"stc", r, x, 0x42, "Saudi Telecom Company", "https://nic.stc/", w{"a.nic.stc", "b.nic.stc", "c.nic.stc", "d.nic.stc"}, n, n, n, "whois.nic.stc", e, w{"https://rdap.centralnic.com/stc/"}, f}
	z[1327] = Zone{"stcgroup", r, x, 0x42, "Saudi Telecom Company", "https://nic.stcgroup/", w{"a.nic.stcgroup", "b.nic.stcgroup", "c.nic.stcgroup", "d.nic.stcgroup"}, n, n, n, "whois.nic.stcgroup", e, w{"https://rdap.centralnic.com/stcgroup/"}, f}
}

//go:noinline
func i1327() {
	z[1328] = Zone{"stockholm", r, x, 0xc4, "Stockholms kommun", "http://nic.stockholm/", w{"a0.nic.stockholm", "a2.nic.stockholm", "b0.nic.stockholm", "c0.nic.stockholm"}, n, w{"Stockholm", "SE-AB"}, w{"sv"}, "whois.nic.stockholm", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1329] = Zone{"storage", r, x, 0x40, "XYZ.COM LLC", "https://nic.storage/", w{"a.nic.storage", "b.nic.storage", "c.nic.storage", "d.nic.storage"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.storage", e, w{"https://rdap.centralnic.com/storage/"}, t}
	z[1330] = Zone{"store", r, x, 0x40, "Radix Technologies Inc.", "https://radix.website/dot-store", w{"a.nic.store", "b.nic.store", "e.nic.store", "f.nic.store"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.store", e, w{"https://rdap.centralnic.com/store/"}, t}
	z[1331] = Zone{"stream", r, x, 0x40, "dot Stream Limited", "http://nic.stream/", w{"a.nic.stream", "b.nic.stream", "c.nic.stream", "ns1.dns.nic.stream", "ns2.dns.nic.stream", "ns3.dns.nic.stream"}, n, n, w{"da", "de", "no", "sv"}, "whois.nic.stream", e, w{"https://rdap.nic.stream/"}, t}
	z[1332] = Zone{"stroke", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1333] = Zone{"studio", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.studio", "v0n1.nic.studio", "v0n2.nic.studio", "v0n3.nic.studio", "v2n0.nic.studio", "v2n1.nic.studio"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.studio", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1334] = Zone{"study", r, x, 0x40, "Registry Services, LLC", "https://www.get.study/", w{"a.nic.study", "b.nic.study", "c.nic.study", "x.nic.study", "y.nic.study", "z.nic.study"}, n, n, n, "whois.nic.study", e, w{"https://rdap.nic.study/"}, f}
	z[1335] = Zone{"style", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.style", "v0n1.nic.style", "v0n2.nic.style", "v0n3.nic.style", "v2n0.nic.style", "v2n1.nic.style"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.style", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1335() {
	z[1336] = Zone{"su", r, z[4373:4425], 0xe0, e, e, w{"a.dns.ripn.net", "b.dns.ripn.net", "d.dns.ripn.net", "e.dns.ripn.net", "f.dns.ripn.net"}, n, n, n, "whois.tcinet.ru", e, n, t}
	z[1337] = Zone{"sucks", r, x, 0x40, "Vox Populi Registry Ltd.", "https://www.get.sucks/", w{"a.nic.sucks", "b.nic.sucks", "c.nic.sucks", "x.nic.sucks", "y.nic.sucks", "z.nic.sucks"}, n, n, n, "whois.nic.sucks", e, w{"https://rdap.nic.sucks/"}, f}
	z[1338] = Zone{"supersport", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1339] = Zone{"supplies", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.supplies", "v0n1.nic.supplies", "v0n2.nic.supplies", "v0n3.nic.supplies", "v2n0.nic.supplies", "v2n1.nic.supplies"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.supplies", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1340] = Zone{"supply", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.supply", "v0n1.nic.supply", "v0n2.nic.supply", "v0n3.nic.supply", "v2n0.nic.supply", "v2n1.nic.supply"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.supply", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1341] = Zone{"support", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.support", "v0n1.nic.support", "v0n2.nic.support", "v0n3.nic.support", "v2n0.nic.support", "v2n1.nic.support"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.support", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1342] = Zone{"surf", r, x, 0x40, "Registry Services, LLC", "http://nic.surf/", w{"a.nic.surf", "b.nic.surf", "c.nic.surf", "x.nic.surf", "y.nic.surf", "z.nic.surf"}, n, n, w{"de", "es", "fr"}, "whois.nic.surf", e, w{"https://rdap.nic.surf/"}, t}
	z[1343] = Zone{"surgery", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.surgery", "v0n1.nic.surgery", "v0n2.nic.surgery", "v0n3.nic.surgery", "v2n0.nic.surgery", "v2n1.nic.surgery"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.surgery", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1343() {
	z[1344] = Zone{"suzuki", r, x, 0x42, "SUZUKI MOTOR CORPORATION", "https://nic.suzuki/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.suzuki", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1345] = Zone{"sv", r, z[4425:4430], 0xa0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "sir.red.sv"}, n, n, n, e, "http://www.svnet.org.sv/", n, f}
	z[1346] = Zone{"svr", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1347] = Zone{"swatch", r, x, 0x42, "The Swatch Group Ltd", "https://nic.swatch/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.swatch", e, w{"https://tld-rdap.verisign.com/swatch/v1/"}, t}
	z[1348] = Zone{"swiftcover", r, x, 0x1042, "Swiftcover Insurance Services Limited", e, n, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.swiftcover", e, w{"https://rdap.nic.swiftcover/"}, t}
	z[1349] = Zone{"swiss", r, x, 0x50, "Swiss Confederation", "https://www.nic.swiss/nic/de/home.html", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net", "g.nic.swiss", "ns15.rcode0.net", "u.nic.swiss"}, n, n, w{"mul-Latn"}, "whois.nic.swiss", e, w{"https://rdap.nic.swiss/"}, t}
	z[1350] = Zone{"sx", r, x, 0xa0, e, "https://en.wikipedia.org/wiki/.sx", w{"ns1.ns.sx", "ns2.ns.sx"}, n, n, n, "whois.sx", e, n, f}
	z[1351] = Zone{"sy", r, z[4430:4436], 0xa0, e, "https://nic.sy/", w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, "whois.tld.sy", e, n, f}
}

//go:noinline
func i1351() {
	z[1352] = Zone{"sydney", r, x, 0xc4, "State of New South Wales, Department of Premier and Cabinet", "https://www.iconic.sydney/", w{"a.nic.sydney", "b.nic.sydney", "c.nic.sydney", "x.nic.sydney", "y.nic.sydney", "z.nic.sydney"}, n, w{"Sydney"}, n, "whois.nic.sydney", e, w{"https://rdap.nic.sydney/"}, f}
	z[1353] = Zone{"symantec", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.symantec", e, n, t}
	z[1354] = Zone{"systems", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.systems", "v0n1.nic.systems", "v0n2.nic.systems", "v0n3.nic.systems", "v2n0.nic.systems", "v2n1.nic.systems"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.systems", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1355] = Zone{"sz", r, z[4436:4439], 0xa0, e, e, w{"ns1.sispa.org.sz", "rip.psg.com", "sz.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
	z[1356] = Zone{"tab", r, x, 0x42, "Tabcorp Holdings Limited", "http://nic.tab/", w{"a.nic.tab", "b.nic.tab", "c.nic.tab", "x.nic.tab", "y.nic.tab", "z.nic.tab"}, n, n, n, "whois.nic.tab", e, w{"https://rdap.nic.tab/"}, f}
	z[1357] = Zone{"taipei", r, x, 0xc4, "Taipei City Government", "http://nic.taipei/", w{"a.nic.taipei", "b.nic.taipei", "c.nic.taipei", "ns1.dns.nic.taipei", "ns2.dns.nic.taipei", "ns3.dns.nic.taipei"}, n, w{"Taipei", "TW-TPQ"}, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.taipei", e, w{"https://rdap.nic.taipei/"}, t}
	z[1358] = Zone{"talk", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.talk/", w{"dns1.nic.talk", "dns2.nic.talk", "dns3.nic.talk", "dns4.nic.talk", "dnsa.nic.talk", "dnsb.nic.talk", "dnsc.nic.talk", "dnsd.nic.talk"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.talk", e, w{"https://rdap.nominet.uk/talk/"}, t}
	z[1359] = Zone{"taobao", r, x, 0x42, "Alibaba Group Holding Limited", "http://nic.taobao/", w{"a0.nic.taobao", "a2.nic.taobao", "b0.nic.taobao", "c0.nic.taobao"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh", "zh-CN", "zh-TW"}, "whois.nic.taobao", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1359() {
	z[1360] = Zone{"target", r, x, 0x42, "Target Domain Holdings, LLC", "https://nic.target/", w{"a.nic.target", "b.nic.target", "c.nic.target", "ns1.dns.nic.target", "ns2.dns.nic.target", "ns3.dns.nic.target"}, n, n, w{"es"}, "whois.nic.target", e, w{"https://rdap.nic.target/"}, t}
	z[1361] = Zone{"tata", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1362] = Zone{"tatamotors", r, x, 0x42, "Tata Motors Ltd", "https://www.icann.org/en/registry-agreements/details/tatamotors", w{"a0.nic.tatamotors", "a2.nic.tatamotors", "b0.nic.tatamotors", "c0.nic.tatamotors"}, n, n, n, "whois.nic.tatamotors", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1363] = Zone{"tatar", r, x, 0x440, "Limited Liability Company \"Coordination Center of Regional Domain of Tatarstan Republic\"", "https://nic.tatar/", w{"a.dns.ripn.net", "b.dns.ripn.net", "d.dns.ripn.net", "e.dns.ripn.net", "f.dns.ripn.net"}, n, w{"RU-TA"}, n, "whois.nic.tatar", e, w{"https://whois.nic.tatar/rdap/"}, f}
	z[1364] = Zone{"tattoo", r, x, 0x40, "Registry Services, LLC", "https://nic.tattoo/", w{"a.nic.tattoo", "b.nic.tattoo", "c.nic.tattoo", "x.nic.tattoo", "y.nic.tattoo", "z.nic.tattoo"}, n, n, w{"de", "es", "fr", "it", "ja", "mul-Cyrl", "pt", "und-Cyrl", "zh", "zh-Hani"}, "whois.nic.tattoo", e, w{"http://rdap.nic.tattoo/"}, t}
	z[1365] = Zone{"tax", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.tax", "v0n1.nic.tax", "v0n2.nic.tax", "v0n3.nic.tax", "v2n0.nic.tax", "v2n1.nic.tax"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tax", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1366] = Zone{"taxi", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.taxi", "v0n1.nic.taxi", "v0n2.nic.taxi", "v0n3.nic.taxi", "v2n0.nic.taxi", "v2n1.nic.taxi"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.taxi", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1367] = Zone{"tc", r, z[4439:4445], 0xa0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc", "root8.zone.tc"}, n, n, n, "whois.nic.tc", e, n, f}
}

//go:noinline
func i1367() {
	z[1368] = Zone{"tci", r, x, 0x42, "Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.", "http://nic.tci/", w{"a.ns.nic.tci", "b.ns.nic.tci", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.tci", e, w{"https://api.rdap.nic.tci/"}, f}
	z[1369] = Zone{"td", r, z[4445:4449], 0xa0, e, "https://nic.td/", w{"anycastdns1.nic.td", "anycastdns2.nic.td", "ns-td.afrinic.net", "pch.nic.td"}, n, n, n, "whois.nic.td", e, n, f}
	z[1370] = Zone{"tdk", r, x, 0x42, "TDK Corporation", "http://nic.tdk/", w{"a.nic.tdk", "b.nic.tdk", "c.nic.tdk", "ns1.dns.nic.tdk", "ns2.dns.nic.tdk", "ns3.dns.nic.tdk"}, n, n, w{"ja"}, "whois.nic.tdk", e, w{"https://rdap.nic.tdk/"}, t}
	z[1371] = Zone{"team", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.team", "v0n1.nic.team", "v0n2.nic.team", "v0n3.nic.team", "v2n0.nic.team", "v2n1.nic.team"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.team", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1372] = Zone{"tech", r, x, 0x40, "Radix Technologies Inc.", "https://radix.website/dot-tech", w{"a.nic.tech", "b.nic.tech", "e.nic.tech", "f.nic.tech"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.tech", e, w{"https://rdap.centralnic.com/tech/"}, t}
	z[1373] = Zone{"technology", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.technology", "v0n1.nic.technology", "v0n2.nic.technology", "v0n3.nic.technology", "v2n0.nic.technology", "v2n1.nic.technology"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.technology", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1374] = Zone{"tel", r, x, 0x4040, "Telnames Limited", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.tel", e, w{"https://rdap.nic.tel/"}, t}
	z[1375] = Zone{"telecity", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.telecity", e, n, f}
}

//go:noinline
func i1375() {
	z[1376] = Zone{"telefonica", r, x, 0x1042, e, e, n, n, n, n, "whois-fe.telefonica.tango.knipp.de", e, n, f}
	z[1377] = Zone{"temasek", r, x, 0x42, "Temasek Holdings (Private) Limited", "https://www.temasek.com.sg/en/site-services/nictemasek", w{"a0.nic.temasek", "a2.nic.temasek", "b0.nic.temasek", "c0.nic.temasek"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.temasek", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1378] = Zone{"tennis", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.tennis", "v0n1.nic.tennis", "v0n2.nic.tennis", "v0n3.nic.tennis", "v2n0.nic.tennis", "v2n1.nic.tennis"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tennis", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1379] = Zone{"terra", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1380] = Zone{"test", r, x, 0x2008, "IANA", "https://datatracker.ietf.org/doc/rfc6761", n, n, n, n, e, e, n, t}
	z[1381] = Zone{"teva", r, x, 0x42, "Teva Pharmaceutical Industries Limited", "http://nic.teva/", w{"a.nic.teva", "b.nic.teva", "c.nic.teva", "ns1.dns.nic.teva", "ns2.dns.nic.teva", "ns3.dns.nic.teva"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.teva", e, w{"https://rdap.nic.teva/"}, t}
	z[1382] = Zone{"tf", r, z[4449:4463], 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.tf", e, w{"https://rdap.nic.tf/"}, t}
	z[1383] = Zone{"tg", r, x, 0xa0, e, "https://www.nic.tg/", w{"ns1.admin.net", "ns1.nic.tg", "ns2.admin.net", "ns2.nic.tg", "ns3.admin.net", "ns4.admin.net", "ns5.admin.net", "tld.cafe.tg"}, n, n, n, "whois.nic.tg", e, n, f}
}

//go:noinline
func i1383() {
	z[1384] = Zone{"th", r, z[4463:4470], 0xa8, e, "https://www.thnic.co.th/", w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, w{"th-TH"}, "whois.thnic.co.th", e, n, t}
	z[1385] = Zone{"thai", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1386] = Zone{"thd", r, x, 0x42, "Home Depot Product Authority, LLC", "https://www.icann.org/en/registry-agreements/details/thd", w{"a0.nic.thd", "a2.nic.thd", "b0.nic.thd", "c0.nic.thd"}, n, n, n, "whois.nic.thd", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1387] = Zone{"theater", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.theater", "v0n1.nic.theater", "v0n2.nic.theater", "v0n3.nic.theater", "v2n0.nic.theater", "v2n1.nic.theater"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.theater", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1388] = Zone{"theatre", r, x, 0x40, "XYZ.COM LLC", "https://nic.theatre/", w{"a.nic.theatre", "b.nic.theatre", "c.nic.theatre", "d.nic.theatre"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.theatre", e, w{"https://rdap.centralnic.com/theatre/"}, t}
	z[1389] = Zone{"theguardian", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1390] = Zone{"thehartford", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1391] = Zone{"tiaa", r, x, 0x42, "Teachers Insurance and Annuity Association of America", "https://www.icann.org/en/registry-agreements/details/tiaa", w{"v0n0.nic.tiaa", "v0n1.nic.tiaa", "v0n2.nic.tiaa", "v0n3.nic.tiaa", "v2n0.nic.tiaa", "v2n1.nic.tiaa"}, n, n, n, "whois.nic.tiaa", e, w{"https://tld-rdap.verisign.com/tiaa/v1/"}, f}
}

//go:noinline
func i1391() {
	z[1392] = Zone{"tickets", r, x, 0x40, "XYZ.COM LLC", "https://nic.tickets/", w{"a.nic.tickets", "b.nic.tickets", "c.nic.tickets", "d.nic.tickets"}, n, n, w{"ar", "da", "de", "es", "fr", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "mul-Mymr", "pl", "ru", "sv", "th", "zh"}, "whois.nic.tickets", e, w{"https://rdap.centralnic.com/tickets/"}, t}
	z[1393] = Zone{"tienda", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.tienda", "v0n1.nic.tienda", "v0n2.nic.tienda", "v0n3.nic.tienda", "v2n0.nic.tienda", "v2n1.nic.tienda"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tienda", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1394] = Zone{"tiffany", r, x, 0x1042, "Tiffany and Company", "http://nic.tiffany/", n, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.tiffany", e, w{"https://tld-rdap.verisign.com/tiffany/v1/"}, t}
	z[1395] = Zone{"tiia", r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1396] = Zone{"tips", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.tips", "v0n1.nic.tips", "v0n2.nic.tips", "v0n3.nic.tips", "v2n0.nic.tips", "v2n1.nic.tips"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tips", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1397] = Zone{"tires", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.tires", "v0n1.nic.tires", "v0n2.nic.tires", "v0n3.nic.tires", "v2n0.nic.tires", "v2n1.nic.tires"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tires", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1398] = Zone{"tirol", r, x, 0x440, "punkt Tirol GmbH", "http://www.nic.tirol/", w{"dns.ryce-rsp.com", "ns1.dns.business", "ns1.ryce-rsp.com"}, n, w{"AT-7"}, w{"mul-Latn"}, "whois.nic.tirol", e, w{"https://rdap.ryce-rsp.com/rdap/"}, t}
	z[1399] = Zone{"tj", r, z[4470:4491], 0xa0, e, e, w{"ns1.nic.tj", "ns2.tojikiston.com", "phloem.uoregon.edu", "tj.cctld.authdns.ripe.net"}, n, n, n, e, "http://www.nic.tj/whois.html", n, f}
}

//go:noinline
func i1399() {
	z[1400] = Zone{"tjmaxx", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.tjmaxx/", w{"a.nic.tjmaxx", "b.nic.tjmaxx", "c.nic.tjmaxx", "ns1.dns.nic.tjmaxx", "ns2.dns.nic.tjmaxx", "ns3.dns.nic.tjmaxx"}, n, n, n, "whois.nic.tjmaxx", e, w{"https://rdap.nic.tjmaxx/"}, f}
	z[1401] = Zone{"tjx", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.tjx/", w{"a.nic.tjx", "b.nic.tjx", "c.nic.tjx", "ns1.dns.nic.tjx", "ns2.dns.nic.tjx", "ns3.dns.nic.tjx"}, n, n, n, "whois.nic.tjx", e, w{"https://rdap.nic.tjx/"}, f}
	z[1402] = Zone{"tk", r, x, 0xa0, e, e, w{"a.ns.tk", "b.ns.tk", "c.ns.tk", "d.ns.tk"}, n, n, n, "whois.dot.tk", e, n, t}
	z[1403] = Zone{"tkmaxx", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.tkmaxx/", w{"a.nic.tkmaxx", "b.nic.tkmaxx", "c.nic.tkmaxx", "ns1.dns.nic.tkmaxx", "ns2.dns.nic.tkmaxx", "ns3.dns.nic.tkmaxx"}, n, n, n, "whois.nic.tkmaxx", e, w{"https://rdap.nic.tkmaxx/"}, f}
	z[1404] = Zone{"tl", r, z[4491:4498], 0xa0, e, e, w{"ns.anycast.nic.tl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, "whois.nic.tl", e, n, f}
	z[1405] = Zone{"tm", r, z[4498:4506], 0xa0, e, "https://nic.tm/", w{"ns-a1.tm", "ns-a2.tm", "ns-a3.tm", "ns-a4.tm", "ns-d1.tm", "ns-l1.tm", "ns-y1.tm"}, n, n, n, "whois.nic.tm", e, n, t}
	z[1406] = Zone{"tmall", r, x, 0x42, "Alibaba Group Holding Limited", "http://nic.tmall/", w{"a0.nic.tmall", "a2.nic.tmall", "b0.nic.tmall", "c0.nic.tmall"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh", "zh-CN", "zh-TW"}, "whois.nic.tmall", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1407] = Zone{"tn", r, z[4506:4525], 0xa0, e, e, w{"ns-tn.afrinic.net", "ns1.ati.tn", "ns2.ati.tn", "ns2.nic.fr", "pch.ati.tn", "rip.psg.com"}, n, n, n, "whois.ati.tn", e, n, f}
}

//go:noinline
func i1407() {
	z[1408] = Zone{"to", r, x, 0xa0, e, "https://www.tonic.to/", w{"cd5.tonic.to", "cd6.tonic.to", "cd7.tonic.to", "cd8.tonic.to", "frankfurt.tonic.to", "helsinki.tonic.to", "newyork.tonic.to", "singapore.tonic.to", "sydney.tonic.to", "tonic.to"}, n, n, n, "whois.tonic.to", e, n, t}
	z[1409] = Zone{"today", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.today", "v0n1.nic.today", "v0n2.nic.today", "v0n3.nic.today", "v2n0.nic.today", "v2n1.nic.today"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.today", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1410] = Zone{"tokyo", r, x, 0xc4, "GMO Registry, Inc.", e, w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"Tokyo", "JP-13"}, w{"ja"}, "whois.nic.tokyo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1411] = Zone{"tools", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.tools", "v0n1.nic.tools", "v0n2.nic.tools", "v0n3.nic.tools", "v2n0.nic.tools", "v2n1.nic.tools"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tools", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1412] = Zone{"top", r, x, 0x40, ".TOP Registry", "https://www.nic.top/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"ar", "de", "es", "fr", "ja", "ru", "tw", "zh-Hans", "zh-Hant"}, "whois.nic.top", e, w{"https://rdap.zdnsgtld.com/top/"}, t}
	z[1413] = Zone{"toray", r, x, 0x42, "Toray Industries, Inc.", "https://nic.toray/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.toray", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1414] = Zone{"toshiba", r, x, 0x42, "TOSHIBA Corporation", "https://www.icann.org/en/registry-agreements/details/toshiba", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"es", "fr", "ja", "ko", "pl", "pt"}, "whois.nic.toshiba", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1415] = Zone{"total", r, x, 0x42, "TotalEnergies SE", "https://nic.total/", w{"d.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, w{"mul-Latn"}, "whois.nic.total", e, w{"https://rdap.nic.total/"}, t}
}

//go:noinline
func i1415() {
	z[1416] = Zone{"tour", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1417] = Zone{"tours", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.tours", "v0n1.nic.tours", "v0n2.nic.tours", "v0n3.nic.tours", "v2n0.nic.tours", "v2n1.nic.tours"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.tours", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1418] = Zone{"town", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.town", "v0n1.nic.town", "v0n2.nic.town", "v0n3.nic.town", "v2n0.nic.town", "v2n1.nic.town"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.town", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1419] = Zone{"toyota", r, x, 0x42, "TOYOTA MOTOR CORPORATION", "http://www.nic.toyota/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.toyota", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1420] = Zone{"toys", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.toys", "v0n1.nic.toys", "v0n2.nic.toys", "v0n3.nic.toys", "v2n0.nic.toys", "v2n1.nic.toys"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.toys", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1421] = Zone{"tp", r, z[4525:4528], 0x10a0, e, e, n, n, n, n, e, e, n, f}
	z[1422] = Zone{"tr", r, z[4528:4549], 0xa8, e, "https://www.iana.org/domains/root/db/tr.html", w{"ns41.nic.tr", "ns42.nic.tr", "ns61.nic.tr", "ns71.nic.tr", "ns72.nic.tr", "ns73.nic.tr", "ns74.nic.tr"}, n, n, n, "whois.trabis.gov.tr", e, n, t}
	z[1423] = Zone{"trade", r, x, 0x40, "Elite Registry Limited", "http://nic.trade/", w{"a.nic.trade", "b.nic.trade", "c.nic.trade", "ns1.dns.nic.trade", "ns2.dns.nic.trade", "ns3.dns.nic.trade"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.trade", e, w{"https://rdap.nic.trade/"}, t}
}

//go:noinline
func i1423() {
	z[1424] = Zone{"tradershotels", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1425] = Zone{"trading", r, x, 0x40, "Dog Beach, LLC", "http://nic.trading/", w{"v0n0.nic.trading", "v0n1.nic.trading", "v0n2.nic.trading", "v0n3.nic.trading", "v2n0.nic.trading", "v2n1.nic.trading"}, n, n, n, "whois.nic.trading", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1426] = Zone{"training", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.training", "v0n1.nic.training", "v0n2.nic.training", "v0n3.nic.training", "v2n0.nic.training", "v2n1.nic.training"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.training", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1427] = Zone{"transformers", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1428] = Zone{"translations", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1429] = Zone{"transunion", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1430] = Zone{"travel", r, x, 0x4040, "Dog Beach, LLC", "https://www.travel.domains/", w{"v0n0.nic.travel", "v0n1.nic.travel", "v0n2.nic.travel", "v0n3.nic.travel", "v2n0.nic.travel", "v2n1.nic.travel"}, n, n, n, "whois.nic.travel", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1431] = Zone{"travelchannel", r, x, 0x1042, "Lifestyle Domain Holdings, Inc.", "https://www.icann.org/en/registry-agreements/details/travelchannel", n, n, n, n, "whois.nic.travelchannel", e, w{"https://tld-rdap.verisign.com/travelchannel/v1/"}, f}
}

//go:noinline
func i1431() {
	z[1432] = Zone{"travelers", r, x, 0x42, "Travelers TLD, LLC", "https://nic.travelers/", w{"a0.nic.travelers", "a2.nic.travelers", "b0.nic.travelers", "c0.nic.travelers"}, n, n, n, "whois.nic.travelers", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1433] = Zone{"travelersinsurance", r, x, 0x42, "Travelers TLD, LLC", "https://nic.travelersinsurance/", w{"a0.nic.travelersinsurance", "a2.nic.travelersinsurance", "b0.nic.travelersinsurance", "c0.nic.travelersinsurance"}, n, n, n, "whois.nic.travelersinsurance", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1434] = Zone{"travelguard", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1435] = Zone{"trust", r, x, 0x40, "Internet Naming Company LLC", "https://uniregistry.link/extensions/trust/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, "whois.nic.trust", e, w{"https://whois.uniregistry.net/rdap/"}, f}
	z[1436] = Zone{"trv", r, x, 0x42, "Travelers TLD, LLC", "https://nic.trv/", w{"a0.nic.trv", "a2.nic.trv", "b0.nic.trv", "c0.nic.trv"}, n, n, n, "whois.nic.trv", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1437] = Zone{"tt", r, z[4549:4581], 0xa0, e, e, w{"a.lactld.org", "pch.nic.tt", "ripe.nic.tt"}, n, n, n, e, "http://www.nic.tt/cgi-bin/search.pl", n, f}
	z[1438] = Zone{"tube", r, x, 0x40, "Latin American Telecom LLC", "http://nic.tube/", w{"a.nic.tube", "b.nic.tube", "c.nic.tube", "ns1.dns.nic.tube", "ns2.dns.nic.tube", "ns3.dns.nic.tube"}, n, n, n, "whois.nic.tube", e, w{"https://rdap.nic.tube/"}, f}
	z[1439] = Zone{"tui", r, x, 0x42, "TUI AG", "https://nic.tui/", w{"a.nic.tui", "b.nic.tui", "c.nic.tui", "d.nic.tui"}, n, n, w{"de"}, "whois.nic.tui", e, w{"https://rdap.centralnic.com/tui/"}, t}
}

//go:noinline
func i1439() {
	z[1440] = Zone{"tunes", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.tunes/", w{"dns1.nic.tunes", "dns2.nic.tunes", "dns3.nic.tunes", "dns4.nic.tunes", "dnsa.nic.tunes", "dnsb.nic.tunes", "dnsc.nic.tunes", "dnsd.nic.tunes"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.tunes", e, w{"https://rdap.nominet.uk/tunes/"}, t}
	z[1441] = Zone{"tushu", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.tushu/", w{"dns1.nic.tushu", "dns2.nic.tushu", "dns3.nic.tushu", "dns4.nic.tushu", "dnsa.nic.tushu", "dnsb.nic.tushu", "dnsc.nic.tushu", "dnsd.nic.tushu"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.tushu", e, w{"https://rdap.nominet.uk/tushu/"}, t}
	z[1442] = Zone{"tv", r, x, 0xe0, e, e, w{"a.nic.tv", "b.nic.tv", "c.nic.tv", "d.nic.tv"}, n, n, n, "whois.nic.tv", e, w{"https://rdap.nic.tv/"}, t}
	z[1443] = Zone{"tvs", r, x, 0x42, "T V SUNDRAM IYENGAR  & SONS LIMITED", "http://nic.tvs/", w{"a0.nic.tvs", "a2.nic.tvs", "b0.nic.tvs", "c0.nic.tvs"}, n, n, n, "whois.nic.tvs", e, w{"https://rdap.afilias-srs.net/rdap/tvs/"}, f}
	z[1444] = Zone{"tw", r, z[4581:4594], 0xa0, e, e, w{"a.dns.tw", "anytld.apnic.net", "b.dns.tw", "c.dns.tw", "d.dns.tw", "e.dns.tw", "f.dns.tw", "g.dns.tw", "h.dns.tw", "ns.twnic.net"}, n, n, w{"de", "fr", "ja", "ko", "th", "zh", "zh-Hant"}, "whois.twnic.net.tw", e, n, t}
	z[1445] = Zone{"tz", r, z[4594:4606], 0xa8, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, "whois.tznic.or.tz", e, w{"https://whois.tznic.or.tz/rdap/"}, f}
	z[1446] = Zone{"ua", r, z[4606:4672], 0xa0, e, "https://www.nic.net.ua/", w{"bg.ns.ua", "ho1.ns.ua", "in1.ns.ua", "nn.ns.ua", "pch.ns.ua", "rcz.ns.ua"}, n, n, w{"mul-Cyrl"}, "whois.ua", e, n, t}
	z[1447] = Zone{"ubank", r, x, 0x42, "National Australia Bank Limited", "https://www.ubank.com.au/nic", w{"a0.nic.ubank", "a2.nic.ubank", "b0.nic.ubank", "c0.nic.ubank"}, n, n, n, "whois.nic.ubank", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i1447() {
	z[1448] = Zone{"ubs", r, x, 0x42, "UBS AG", "https://www.icann.org/en/registry-agreements/details/ubs", w{"a0.nic.ubs", "a2.nic.ubs", "b0.nic.ubs", "c0.nic.ubs"}, n, n, n, "whois.nic.ubs", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1449] = Zone{"uconnect", r, x, 0x1042, e, e, n, n, n, n, "whois.afilias-srs.net", e, n, f}
	z[1450] = Zone{"ug", r, z[4672:4681], 0xa0, e, e, w{"anycast.eahd.or.ug", "ns-ug.afrinic.net", "ns.icann.org", "root.eahd.or.ug", "ug.cctld.authdns.ripe.net"}, n, n, n, "whois.co.ug", e, n, f}
	z[1451] = Zone{"uk", r, z[4681:4691], 0xa0, e, "https://www.nominet.uk/", w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, "whois.nic.uk", e, w{"https://rdap.nominet.uk/uk/"}, f}
	z[1452] = Zone{"ultrabook", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1453] = Zone{"um", r, x, 0x10a0, e, e, n, n, n, n, e, e, n, f}
	z[1454] = Zone{"ummah", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1455] = Zone{"unicom", r, x, 0x42, "China United Network Communications Corporation Limited", "https://www.icann.org/en/registry-agreements/details/unicom", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"es", "ru", "zh-Hans", "zh-Hant"}, "whois.nic.unicom", e, w{"https://rdap.zdnsgtld.com/unicom/"}, t}
}

//go:noinline
func i1455() {
	z[1456] = Zone{"unicorn", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1457] = Zone{"university", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.university", "v0n1.nic.university", "v0n2.nic.university", "v0n3.nic.university", "v2n0.nic.university", "v2n1.nic.university"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.university", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1458] = Zone{"uno", r, x, 0x40, "Radix Technologies Inc.", e, w{"a.nic.uno", "b.nic.uno", "c.nic.uno", "d.nic.uno"}, n, n, w{"ar", "es", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.uno", e, w{"https://rdap.centralnic.com/uno/"}, t}
	z[1459] = Zone{"uol", r, x, 0x42, "UBN INTERNET LTDA.", "https://nic.uol/", w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, w{"pt"}, "whois.gtlds.nic.br", e, w{"https://rdap.gtlds.nic.br/"}, t}
	z[1460] = Zone{"ups", r, x, 0x42, "UPS Market Driver, Inc.", "https://www.ups.com/lasso/login?reasonCode=-1", w{"a0.nic.ups", "a2.nic.ups", "b0.nic.ups", "c0.nic.ups"}, n, n, n, "whois.nic.ups", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1461] = Zone{"us", r, z[4691:4751], 0xa0, e, e, w{"b.cctld.us", "f.cctld.us", "k.cctld.us", "w.cctld.us", "x.cctld.us", "y.cctld.us"}, n, n, n, "whois.nic.us", e, n, f}
	z[1462] = Zone{"uy", r, z[4751:4757], 0xa0, e, e, w{"a.lactld.org", "a.nic.uy", "b.nic.uy", "d.nic.uy", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy"}, n, n, n, "whois.nic.org.uy", e, n, f}
	z[1463] = Zone{"uz", r, z[4757:4771], 0xa0, e, e, w{"ns1.uz", "ns2.uz", "ns3.uz", "ns4.uz", "ns5.uz", "ns6.uz", "ns7.uz", "ns8.uz"}, n, n, n, "whois.cctld.uz", e, w{"http://cctld.uz:9000/"}, f}
}

//go:noinline
func i1463() {
	z[1464] = Zone{"va", r, x, 0xa0, e, e, w{"a.nic.va", "b.nic.va", "c.nic.va", "dns.nic.it", "osiris.namex.it", "seth.namex.it", "va.cctld.authdns.ripe.net"}, n, n, n, "whois.iana.org", e, n, f}
	z[1465] = Zone{"vacations", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.vacations", "v0n1.nic.vacations", "v0n2.nic.vacations", "v0n3.nic.vacations", "v2n0.nic.vacations", "v2n1.nic.vacations"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.vacations", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1466] = Zone{"vana", r, x, 0x42, "Internet Naming Company LLC", "https://www.icann.org/en/registry-agreements/details/vana", w{"ns01.trs-dns.com", "ns01.trs-dns.info", "ns01.trs-dns.net", "ns01.trs-dns.org"}, n, n, n, "whois.nic.vana", e, w{"https://tld-rdap.verisign.com/vana/v1/"}, f}
	z[1467] = Zone{"vanguard", r, x, 0x42, "The Vanguard Group, Inc.", "https://www.icann.org/en/registry-agreements/details/vanguard", w{"a0.nic.vanguard", "a2.nic.vanguard", "b0.nic.vanguard", "c0.nic.vanguard"}, n, n, n, "whois.nic.vanguard", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1468] = Zone{"vanish", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1469] = Zone{"vc", r, z[4771:4774], 0xe0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, "whois2.afilias-grs.net", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1470] = Zone{"ve", r, z[4774:4785], 0xa8, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, w{"es"}, "whois.nic.ve", e, n, t}
	z[1471] = Zone{"vegas", r, x, 0xc4, "Dot Vegas, Inc.", "https://nic.vegas/", w{"a0.nic.vegas", "a2.nic.vegas", "b0.nic.vegas", "c0.nic.vegas"}, n, w{"Las Vegas"}, n, "whois.nic.vegas", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
}

//go:noinline
func i1471() {
	z[1472] = Zone{"ventures", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.ventures", "v0n1.nic.ventures", "v0n2.nic.ventures", "v0n3.nic.ventures", "v2n0.nic.ventures", "v2n1.nic.ventures"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.ventures", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1473] = Zone{"verisign", r, x, 0x42, "VeriSign, Inc.", "https://www.icann.org/en/registry-agreements/details/verisign", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.verisign", e, w{"https://tld-rdap.verisign.com/verisign/v1/"}, t}
	z[1474] = Zone{"xn--vermgensberater-ctb" /* vermögensberater */, r, x, 0x48, "Deutsche Vermögensberatung Aktiengesellschaft DVAG", "https://www.icann.org/en/registry-agreements/details/xn--vermgensberater-ctb", w{"a.nic.xn--vermgensberater-ctb", "b.nic.xn--vermgensberater-ctb", "c.nic.xn--vermgensberater-ctb", "d.nic.xn--vermgensberater-ctb"}, n, n, w{"de"}, "whois.nic.xn--vermgensberater-ctb", e, w{"https://rdap.centralnic.com/xn--vermgensberater-ctb/"}, t}
	z[1475] = Zone{"xn--vermgensberatung-pwb" /* vermögensberatung */, r, x, 0x48, "Deutsche Vermögensberatung Aktiengesellschaft DVAG", "https://www.icann.org/en/registry-agreements/details/xn--vermgensberatung-pwb", w{"a.nic.xn--vermgensberatung-pwb", "b.nic.xn--vermgensberatung-pwb", "c.nic.xn--vermgensberatung-pwb", "d.nic.xn--vermgensberatung-pwb"}, n, n, w{"de"}, "whois.nic.xn--vermgensberatung-pwb", e, w{"https://rdap.centralnic.com/xn--vermgensberatung-pwb/"}, t}
	z[1476] = Zone{"versicherung", r, x, 0x40, "tldbox GmbH", e, w{"a.dns.nic.versicherung", "m.dns.nic.versicherung", "n.dns.nic.versicherung"}, n, n, w{"mul-Latn"}, "whois.nic.versicherung", e, w{"https://rdap.nic.versicherung/v1/"}, t}
	z[1477] = Zone{"vet", r, x, 0x40, "Dog Beach, LLC", e, w{"v0n0.nic.vet", "v0n1.nic.vet", "v0n2.nic.vet", "v0n3.nic.vet", "v2n0.nic.vet", "v2n1.nic.vet"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.vet", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1478] = Zone{"vg", r, z[4785:4786], 0xa0, e, e, w{"a.nic.vg", "b.nic.vg", "c.nic.vg", "d.nic.vg"}, w{"88.198.29.97"}, n, n, "whois.nic.vg", e, w{"https://rdap.centralnic.com/vg/"}, f}
	z[1479] = Zone{"vi", r, z[4786:4792], 0xa8, e, "https://secure.nic.vi/", w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, "https://secure.nic.vi/whois-lookup/", n, f}
}

//go:noinline
func i1479() {
	z[1480] = Zone{"viajes", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.viajes", "v0n1.nic.viajes", "v0n2.nic.viajes", "v0n3.nic.viajes", "v2n0.nic.viajes", "v2n1.nic.viajes"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.viajes", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1481] = Zone{"video", r, x, 0x40, "Dog Beach, LLC", "https://identity.digital/", w{"v0n0.nic.video", "v0n1.nic.video", "v0n2.nic.video", "v0n3.nic.video", "v2n0.nic.video", "v2n1.nic.video"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.video", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1482] = Zone{"vig", r, x, 0x42, "VIENNA INSURANCE GROUP AG Wiener Versicherung Gruppe", "https://nic.vig/en/home.html", w{"a0.nic.vig", "a2.nic.vig", "b0.nic.vig", "c0.nic.vig"}, n, n, n, "whois.nic.vig", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1483] = Zone{"viking", r, x, 0x42, "Viking River Cruises (Bermuda) Ltd.", "http://nic.viking/", w{"a0.nic.viking", "a2.nic.viking", "b0.nic.viking", "c0.nic.viking"}, n, n, n, "whois.nic.viking", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1484] = Zone{"villas", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.villas", "v0n1.nic.villas", "v0n2.nic.villas", "v0n3.nic.villas", "v2n0.nic.villas", "v2n1.nic.villas"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.villas", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1485] = Zone{"vin", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.vin", "v0n1.nic.vin", "v0n2.nic.vin", "v0n3.nic.vin", "v2n0.nic.vin", "v2n1.nic.vin"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.vin", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1486] = Zone{"vip", r, x, 0x40, "Registry Services, LLC", "http://nic.vip/", w{"a.nic.vip", "b.nic.vip", "c.nic.vip", "x.nic.vip", "y.nic.vip", "z.nic.vip"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "he", "it", "no", "ru", "sv", "zh"}, "whois.nic.vip", e, w{"https://rdap.nic.vip/"}, t}
	z[1487] = Zone{"virgin", r, x, 0x42, "Virgin Enterprises Limited", "http://nic.virgin/", w{"dns1.nic.virgin", "dns2.nic.virgin", "dns3.nic.virgin", "dns4.nic.virgin", "dnsa.nic.virgin", "dnsb.nic.virgin", "dnsc.nic.virgin", "dnsd.nic.virgin"}, n, n, n, "whois.nic.virgin", e, w{"https://rdap.nominet.uk/virgin/"}, f}
}

//go:noinline
func i1487() {
	z[1488] = Zone{"visa", r, x, 0x42, "Visa Worldwide Pte. Limited", "https://usa.visa.com/legal/visa-nic.html", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.visa", e, w{"https://tld-rdap.verisign.com/visa/v1/"}, t}
	z[1489] = Zone{"vision", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.vision", "v0n1.nic.vision", "v0n2.nic.vision", "v0n3.nic.vision", "v2n0.nic.vision", "v2n1.nic.vision"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.vision", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1490] = Zone{"vista", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.vista", e, n, f}
	z[1491] = Zone{"vistaprint", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.vistaprint", e, n, f}
	z[1492] = Zone{"viva", r, x, 0x42, "Saudi Telecom Company", "https://nic.viva/", w{"a.nic.viva", "b.nic.viva", "c.nic.viva", "d.nic.viva"}, n, n, n, "whois.nic.viva", e, w{"https://rdap.centralnic.com/viva/"}, f}
	z[1493] = Zone{"vivo", r, x, 0x42, "Telefonica Brasil S.A.", "https://www.vivo.com.br/para-voce", w{"a.nic.vivo", "b.nic.vivo", "c.nic.vivo", "ns1.dns.nic.vivo", "ns2.dns.nic.vivo", "ns3.dns.nic.vivo"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.vivo", e, w{"https://rdap.nic.vivo/"}, t}
	z[1494] = Zone{"vlaanderen", r, x, 0x440, "DNS.be vzw", "https://www.dnsbelgium.be/", w{"a.nsset.vlaanderen", "b.nsset.vlaanderen", "c.nsset.vlaanderen", "d.nsset.vlaanderen", "y.nsset.vlaanderen", "z.nsset.vlaanderen"}, n, w{"BE-VLG"}, w{"mul-Latn"}, "whois.nic.vlaanderen", e, w{"https://rdap.nic.vlaanderen/"}, t}
	z[1495] = Zone{"vn", r, z[4792:4804], 0xa0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, "http://www.vnnic.vn/en/domain", n, t}
}

//go:noinline
func i1495() {
	z[1496] = Zone{"vodka", r, x, 0x40, "Registry Services, LLC", "http://nic.vodka/", w{"a.nic.vodka", "b.nic.vodka", "c.nic.vodka", "x.nic.vodka", "y.nic.vodka", "z.nic.vodka"}, n, n, w{"de", "es", "fr"}, "whois.nic.vodka", e, w{"https://rdap.nic.vodka/"}, t}
	z[1497] = Zone{"volkswagen", r, x, 0x1042, "Volkswagen Group of America Inc.", "http://nic.volkswagen/", n, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.volkswagen", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1498] = Zone{"volvo", r, x, 0x42, "Volvo Holding Sverige Aktiebolag", "https://www.nic.volvo/en/", w{"a0.nic.volvo", "a2.nic.volvo", "b0.nic.volvo", "c0.nic.volvo"}, n, n, n, "whois.nic.volvo", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1499] = Zone{"vons", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1500] = Zone{"vote", r, x, 0x40, "Monolith Registry LLC", "http://nic.vote/", w{"a0.nic.vote", "a2.nic.vote", "b0.nic.vote", "c0.nic.vote"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.vote", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1501] = Zone{"voting", r, x, 0x40, "Valuetainment Corp.", e, w{"a.nic.voting", "b.nic.voting", "c.nic.voting", "x.nic.voting", "y.nic.voting", "z.nic.voting"}, n, n, w{"mul-Cyrl", "mul-Latn"}, "whois.nic.voting", e, w{"https://rdap.nic.voting/"}, t}
	z[1502] = Zone{"voto", r, x, 0x40, "Monolith Registry LLC", "http://nic.voto/", w{"a0.nic.voto", "a2.nic.voto", "b0.nic.voto", "c0.nic.voto"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.voto", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1503] = Zone{"voyage", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.voyage", "v0n1.nic.voyage", "v0n2.nic.voyage", "v0n3.nic.voyage", "v2n0.nic.voyage", "v2n1.nic.voyage"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.voyage", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
}

//go:noinline
func i1503() {
	z[1504] = Zone{"vu", r, z[4804:4809], 0xa0, e, e, w{"ns1.tldns.vu", "ns2.tldns.vu", "ns3.tldns.vu", "ns4.tldns.vu"}, n, n, n, "whois.dnrs.vu", e, n, f}
	z[1505] = Zone{"vuelos", r, x, 0x1040, "Travel Reservations SRL", "http://nic.vuelos/", n, n, n, w{"es"}, "whois.nic.vuelos", e, w{"https://rdap.nic.vuelos/"}, t}
	z[1506] = Zone{"wales", r, x, 0x4c0, "Nominet UK", "https://ourhomeonline.wales/", w{"dns1.nic.wales", "dns2.nic.wales", "dns3.nic.wales", "dns4.nic.wales", "dnsa.nic.wales", "dnsb.nic.wales", "dnsc.nic.wales", "dnsd.nic.wales"}, n, w{"GB-WLS"}, w{"cy"}, "whois.nic.wales", e, w{"https://rdap.nominet.uk/wales/"}, t}
	z[1507] = Zone{"walmart", r, x, 0x42, "Wal-Mart Stores, Inc.", "https://www.icann.org/en/registry-agreements/details/walmart", w{"a0.nic.walmart", "a2.nic.walmart", "b0.nic.walmart", "c0.nic.walmart"}, n, n, n, "whois.nic.walmart", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1508] = Zone{"walter", r, x, 0x42, "Sandvik AB", "https://www.walter-tools.com/en-gb/company/mission_facts/legal_notice/pages/nic-walter.aspx", w{"a.nic.walter", "b.nic.walter", "c.nic.walter", "x.nic.walter", "y.nic.walter", "z.nic.walter"}, n, n, w{"da", "fr", "pl", "ru", "sv", "zh"}, "whois.nic.walter", e, w{"https://rdap.nic.walter/"}, t}
	z[1509] = Zone{"wang", r, x, 0x40, "Zodiac Wang Limited", e, w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/wang/"}, t}
	z[1510] = Zone{"wanggou", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.wanggou/", w{"dns1.nic.wanggou", "dns2.nic.wanggou", "dns3.nic.wanggou", "dns4.nic.wanggou", "dnsa.nic.wanggou", "dnsb.nic.wanggou", "dnsc.nic.wanggou", "dnsd.nic.wanggou"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.wanggou", e, w{"https://rdap.nominet.uk/wanggou/"}, t}
	z[1511] = Zone{"warman", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.warman", e, n, f}
}

//go:noinline
func i1511() {
	z[1512] = Zone{"watch", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.watch", "v0n1.nic.watch", "v0n2.nic.watch", "v0n3.nic.watch", "v2n0.nic.watch", "v2n1.nic.watch"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.watch", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1513] = Zone{"watches", r, x, 0x40, "Identity Digital Limited", "https://identity.digital/", w{"a0.nic.watches", "a2.nic.watches", "b0.nic.watches", "c0.nic.watches"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "pl", "pt", "ru", "sv", "zh-Hans", "zh-TW"}, "whois.nic.watches", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1514] = Zone{"weather", r, x, 0x40, "International Business Machines Corporation", "http://www.nic.weather/", w{"a.nic.weather", "b.nic.weather", "c.nic.weather", "ns4.dns.nic.weather", "ns5.dns.nic.weather", "ns6.dns.nic.weather"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.weather", e, w{"https://rdap.nic.weather/"}, t}
	z[1515] = Zone{"weatherchannel", r, x, 0x42, "International Business Machines Corporation", "http://www.nic.weatherchannel/", w{"a.nic.weatherchannel", "b.nic.weatherchannel", "c.nic.weatherchannel", "ns4.dns.nic.weatherchannel", "ns5.dns.nic.weatherchannel", "ns6.dns.nic.weatherchannel"}, n, n, n, "whois.nic.weatherchannel", e, w{"https://rdap.nic.weatherchannel/"}, f}
	z[1516] = Zone{"web", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1517] = Zone{"webcam", r, x, 0x40, "dot Webcam Limited", "http://nic.webcam/", w{"a.nic.webcam", "b.nic.webcam", "c.nic.webcam", "ns1.dns.nic.webcam", "ns2.dns.nic.webcam", "ns3.dns.nic.webcam"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.webcam", e, w{"https://rdap.nic.webcam/"}, t}
	z[1518] = Zone{"weber", r, x, 0x42, "Saint-Gobain Weber SA", "https://nic.weber/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"mul-Arab", "mul-Cyrl", "mul-Hang", "mul-Hani", "mul-Latn"}, "whois.nic.weber", e, w{"https://tld-rdap.verisign.com/weber/v1/"}, t}
	z[1519] = Zone{"webjet", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1519() {
	z[1520] = Zone{"webs", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1521] = Zone{"website", r, x, 0x40, "Radix Technologies Inc.", e, w{"a.nic.website", "b.nic.website", "e.nic.website", "f.nic.website"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.website", e, w{"https://rdap.centralnic.com/website/"}, t}
	z[1522] = Zone{"wed", r, x, 0x40, e, "https://www.icann.org/resources/pages/ebero-2013-04-02-en", w{"dns1.emdns.uk", "dns2.emdns.uk", "dns3.emdns.uk", "dns4.emdns.uk", "dnsa.emdns.uk", "dnsb.emdns.uk", "dnsc.emdns.uk", "dnsd.emdns.uk"}, n, n, n, "whois.nic.wed", e, w{"https://rdap.nominet.uk/wed/"}, f}
	z[1523] = Zone{"wedding", r, x, 0x40, "Registry Services, LLC", "http://nic.wedding/", w{"a.nic.wedding", "b.nic.wedding", "c.nic.wedding", "x.nic.wedding", "y.nic.wedding", "z.nic.wedding"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.wedding", e, w{"https://rdap.nic.wedding/"}, t}
	z[1524] = Zone{"weibo", r, x, 0x40, "Sina Corporation", "https://www.icann.org/en/registry-agreements/details/weibo", w{"a0.nic.weibo", "a2.nic.weibo", "b0.nic.weibo", "c0.nic.weibo"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.weibo", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1525] = Zone{"weir", r, x, 0x42, "Weir Group IP Limited", "http://nic.weir/", w{"a0.nic.weir", "a2.nic.weir", "b0.nic.weir", "c0.nic.weir"}, n, n, n, "whois.nic.weir", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1526] = Zone{"wf", r, x, 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.wf", e, w{"https://rdap.nic.wf/"}, t}
	z[1527] = Zone{"whoswho", r, x, 0x40, "Who's Who Registry", "https://internet.whoswho/", w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"de", "es", "fr", "ja", "ko", "pl", "pt", "ru", "zh"}, "whois.nic.whoswho", e, w{"https://rdap.nic.whoswho/"}, t}
}

//go:noinline
func i1527() {
	z[1528] = Zone{"wien", r, x, 0xc4, "punkt.wien GmbH", e, w{"dns.ryce-rsp.com", "ns1.dns.business", "ns1.ryce-rsp.com"}, n, w{"Vienna", "AT-9"}, w{"mul-Latn"}, "whois.nic.wien", e, w{"https://rdap.ryce-rsp.com/rdap/"}, t}
	z[1529] = Zone{"wiki", r, x, 0x40, "Registry Services, LLC", e, w{"a.nic.wiki", "b.nic.wiki", "c.nic.wiki", "x.nic.wiki", "y.nic.wiki", "z.nic.wiki"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Cyrl", "mul-Grek", "mul-Latn", "th", "zh"}, "whois.nic.wiki", e, w{"https://rdap.nic.wiki/"}, t}
	z[1530] = Zone{"williamhill", r, x, 0x42, "William Hill Organization Limited", "https://nic.williamhill/", w{"a.nic.williamhill", "b.nic.williamhill", "c.nic.williamhill", "ns1.dns.nic.williamhill", "ns2.dns.nic.williamhill", "ns3.dns.nic.williamhill"}, n, n, n, "whois.nic.williamhill", e, w{"https://rdap.nic.williamhill/"}, f}
	z[1531] = Zone{"wilmar", r, x, 0x10042, e, e, n, n, n, n, e, e, n, f}
	z[1532] = Zone{"win", r, x, 0x40, "First Registry Limited", "http://nic.win/", w{"a.nic.win", "b.nic.win", "c.nic.win", "ns1.dns.nic.win", "ns2.dns.nic.win", "ns3.dns.nic.win"}, n, n, w{"da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.win", e, w{"https://rdap.nic.win/"}, t}
	z[1533] = Zone{"windows", r, x, 0x42, "Microsoft Corporation", "https://nic.windows/", w{"dns1.nic.windows", "dns2.nic.windows", "dns3.nic.windows", "dns4.nic.windows", "dnsa.nic.windows", "dnsb.nic.windows", "dnsc.nic.windows", "dnsd.nic.windows"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/windows/"}, t}
	z[1534] = Zone{"wine", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.wine", "v0n1.nic.wine", "v0n2.nic.wine", "v0n3.nic.wine", "v2n0.nic.wine", "v2n1.nic.wine"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.wine", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1535] = Zone{"winners", r, x, 0x42, "The TJX Companies, Inc.", "http://nic.winners/", w{"a.nic.winners", "b.nic.winners", "c.nic.winners", "ns1.dns.nic.winners", "ns2.dns.nic.winners", "ns3.dns.nic.winners"}, n, n, n, "whois.nic.winners", e, w{"https://rdap.nic.winners/"}, f}
}

//go:noinline
func i1535() {
	z[1536] = Zone{"wme", r, x, 0x42, "William Morris Endeavor Entertainment, LLC", "https://nic.wme/", w{"a.nic.wme", "b.nic.wme", "c.nic.wme", "d.nic.wme"}, n, n, w{"de"}, "whois.nic.wme", e, w{"https://rdap.centralnic.com/wme/"}, t}
	z[1537] = Zone{"wolterskluwer", r, x, 0x42, "Wolters Kluwer N.V.", "http://nic.wolterskluwer/", w{"a0.nic.wolterskluwer", "a2.nic.wolterskluwer", "b0.nic.wolterskluwer", "c0.nic.wolterskluwer"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.wolterskluwer", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1538] = Zone{"woodside", r, x, 0x42, "Woodside Petroleum Limited", "http://nic.woodside/", w{"a.nic.woodside", "b.nic.woodside", "c.nic.woodside", "x.nic.woodside", "y.nic.woodside", "z.nic.woodside"}, n, n, n, "whois.nic.woodside", e, w{"https://rdap.nic.woodside/"}, f}
	z[1539] = Zone{"work", r, x, 0x40, "Registry Services, LLC", "http://nic.work/", w{"a.nic.work", "b.nic.work", "c.nic.work", "x.nic.work", "y.nic.work", "z.nic.work"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.work", e, w{"https://rdap.nic.work/"}, t}
	z[1540] = Zone{"works", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.works", "v0n1.nic.works", "v0n2.nic.works", "v0n3.nic.works", "v2n0.nic.works", "v2n1.nic.works"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.works", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1541] = Zone{"world", r, x, 0x40, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.world", "v0n1.nic.world", "v0n2.nic.world", "v0n3.nic.world", "v2n0.nic.world", "v2n1.nic.world"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.world", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1542] = Zone{"wow", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.wow/", w{"dns1.nic.wow", "dns2.nic.wow", "dns3.nic.wow", "dns4.nic.wow", "dnsa.nic.wow", "dnsb.nic.wow", "dnsc.nic.wow", "dnsd.nic.wow"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.wow", e, w{"https://rdap.nominet.uk/wow/"}, t}
	z[1543] = Zone{"ws", r, z[4809:4814], 0xe0, e, "https://www.worldsite.ws/", w{"a.dns.ws", "ns2.dns.ws", "ns5.dns.ws", "s.dns.ws", "us3.dns.ws", "us4.dns.ws"}, w{"64.70.19.203"}, n, n, "whois.website.ws", e, n, t}
}

//go:noinline
func i1543() {
	z[1544] = Zone{"wtc", r, x, 0x42, "World Trade Centers Association, Inc.", "http://nic.wtc/", w{"a.nic.wtc", "b.nic.wtc", "c.nic.wtc", "x.nic.wtc", "y.nic.wtc", "z.nic.wtc"}, n, n, n, "whois.nic.wtc", e, w{"https://rdap.nic.wtc/"}, f}
	z[1545] = Zone{"wtf", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.wtf", "v0n1.nic.wtf", "v0n2.nic.wtf", "v0n3.nic.wtf", "v2n0.nic.wtf", "v2n1.nic.wtf"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.wtf", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1546] = Zone{"xbox", r, x, 0x42, "Microsoft Corporation", "https://nic.xbox/", w{"dns1.nic.xbox", "dns2.nic.xbox", "dns3.nic.xbox", "dns4.nic.xbox", "dnsa.nic.xbox", "dnsb.nic.xbox", "dnsc.nic.xbox", "dnsd.nic.xbox"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, e, e, w{"https://rdap.nominet.uk/xbox/"}, t}
	z[1547] = Zone{"xerox", r, x, 0x42, "Xerox DNHC LLC", "http://nic.xerox/", w{"a.nic.xerox", "b.nic.xerox", "c.nic.xerox", "x.nic.xerox", "y.nic.xerox", "z.nic.xerox"}, n, n, n, "whois.nic.xerox", e, w{"https://rdap.nic.xerox/"}, f}
	z[1548] = Zone{"xfinity", r, x, 0x42, "Comcast IP Holdings I, LLC", "https://nic.xfinity/", w{"dns1.nic.xfinity", "dns2.nic.xfinity", "dns3.nic.xfinity", "dns4.nic.xfinity", "dnsa.nic.xfinity", "dnsb.nic.xfinity", "dnsc.nic.xfinity", "dnsd.nic.xfinity"}, n, n, n, "whois.nic.xfinity", e, w{"https://rdap.nominet.uk/xfinity/"}, f}
	z[1549] = Zone{"xihuan", r, x, 0x40, "Beijing Qihu Keji Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xihuan", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, n, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/"}, f}
	z[1550] = Zone{"xin", r, x, 0x40, "Elegant Leader Limited", "https://nic.xin/", w{"a0.nic.xin", "a2.nic.xin", "b0.nic.xin", "c0.nic.xin"}, n, n, n, "whois.nic.xin", e, w{"https://rdap.identitydigital.services/rdap/"}, f}
	z[1551] = Zone{"xperia", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.xperia", e, n, f}
}

//go:noinline
func i1551() {
	z[1552] = Zone{"xxx", r, x, 0x41, "ICM Registry LLC", "https://nic.xxx/", w{"a.nic.xxx", "b.nic.xxx", "c.nic.xxx", "x.nic.xxx", "y.nic.xxx", "z.nic.xxx"}, n, n, w{"be", "bg", "bs", "cnr", "da", "de", "es", "fr", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.xxx", e, w{"https://rdap.nic.xxx/"}, t}
	z[1553] = Zone{"xyz", r, x, 0x40, "XYZ.COM LLC", e, w{"generationxyz.nic.xyz", "x.nic.xyz", "y.nic.xyz", "z.nic.xyz"}, n, n, w{"ar", "he", "ja", "ko", "lo", "mul-Grek", "mul-Latn", "ru", "th", "zh"}, "whois.nic.xyz", e, w{"https://rdap.centralnic.com/xyz/"}, t}
	z[1554] = Zone{"yachts", r, x, 0x40, "XYZ.COM LLC", "https://nic.yachts/", w{"a.nic.yachts", "b.nic.yachts", "c.nic.yachts", "d.nic.yachts"}, n, n, n, "whois.nic.yachts", e, w{"https://rdap.centralnic.com/yachts/"}, f}
	z[1555] = Zone{"yahoo", r, x, 0x42, "Oath Inc.", "https://gtldyahoo.tumblr.com/", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.yahoo", e, w{"https://tld-rdap.verisign.com/yahoo/v1/"}, t}
	z[1556] = Zone{"yamaxun", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.yamaxun/", w{"dns1.nic.yamaxun", "dns2.nic.yamaxun", "dns3.nic.yamaxun", "dns4.nic.yamaxun", "dnsa.nic.yamaxun", "dnsb.nic.yamaxun", "dnsc.nic.yamaxun", "dnsd.nic.yamaxun"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.yamaxun", e, w{"https://rdap.nominet.uk/yamaxun/"}, t}
	z[1557] = Zone{"yandex", r, x, 0x42, "Yandex Europe B.V.", "https://nic.yandex/", w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, "whois.nic.yandex", e, w{"https://rdap.nic.yandex/"}, f}
	z[1558] = Zone{"ye", r, z[4814:4822], 0xa8, e, e, w{"ns1.yemen.net.ye", "ns2.yemen.net.ye", "sah1.ye", "sah2.ye", "tld1.ye", "tld2.ye"}, n, n, n, "whois.y.net.ye", e, n, f}
	z[1559] = Zone{"yodobashi", r, x, 0x42, "YODOBASHI CAMERA CO.,LTD.", "https://nic.yodobashi/", w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, n, w{"ja"}, "whois.nic.gmo", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
}

//go:noinline
func i1559() {
	z[1560] = Zone{"yoga", r, x, 0x40, "Registry Services, LLC", "http://nic.yoga/", w{"a.nic.yoga", "b.nic.yoga", "c.nic.yoga", "x.nic.yoga", "y.nic.yoga", "z.nic.yoga"}, n, n, w{"de", "es", "fr", "zh"}, "whois.nic.yoga", e, w{"https://rdap.nic.yoga/"}, t}
	z[1561] = Zone{"yokohama", r, x, 0xc4, "GMO Registry, Inc.", e, w{"a.gmoregistry.net", "b.gmoregistry.net", "k.gmoregistry.net", "l.gmoregistry.net"}, n, w{"Yokohama"}, w{"ja"}, "whois.nic.yokohama", e, w{"https://rdap.gmoregistry.net/rdap/"}, t}
	z[1562] = Zone{"you", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.you/", w{"dns1.nic.you", "dns2.nic.you", "dns3.nic.you", "dns4.nic.you", "dnsa.nic.you", "dnsb.nic.you", "dnsc.nic.you", "dnsd.nic.you"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.you", e, w{"https://rdap.nominet.uk/you/"}, t}
	z[1563] = Zone{"youtube", r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1564] = Zone{"yt", r, x, 0xa0, e, e, w{"d.nic.fr", "e.ext.nic.fr", "f.ext.nic.fr", "g.ext.nic.fr"}, n, n, n, "whois.nic.yt", e, w{"https://rdap.nic.yt/"}, t}
	z[1565] = Zone{"yu", r, z[4822:4828], 0x10a0, e, e, n, n, n, n, e, e, n, f}
	z[1566] = Zone{"yun", r, x, 0x40, "Beijing Qihu Keji Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/yun", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, n, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/"}, f}
	z[1567] = Zone{"za", r, z[4828:4855], 0xa8, e, e, w{"nsza.is.co.za", "za-ns.anycast.pch.net", "za1.dnsnode.net"}, n, n, n, "whois.nic.za", e, n, f}
}

//go:noinline
func i1567() {
	z[1568] = Zone{"zappos", r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.zappos/", w{"dns1.nic.zappos", "dns2.nic.zappos", "dns3.nic.zappos", "dns4.nic.zappos", "dnsa.nic.zappos", "dnsb.nic.zappos", "dnsc.nic.zappos", "dnsd.nic.zappos"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.zappos", e, w{"https://rdap.nominet.uk/zappos/"}, t}
	z[1569] = Zone{"zara", r, x, 0x42, "Industria de Diseño Textil, S.A. (INDITEX, S.A.)", "https://nic.zara/", w{"a0.nic.zara", "a2.nic.zara", "b0.nic.zara", "c0.nic.zara"}, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.zara", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1570] = Zone{"zero", r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.zero/", w{"a.nic.zero", "b.nic.zero", "c.nic.zero", "ns1.dns.nic.zero", "ns2.dns.nic.zero", "ns3.dns.nic.zero"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.zero", e, w{"https://rdap.nic.zero/"}, t}
	z[1571] = Zone{"zip", r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1572] = Zone{"zippo", r, x, 0x1042, e, e, n, n, n, n, "whois.nic.zippo", e, n, f}
	z[1573] = Zone{"zm", r, z[4855:4863], 0xa0, e, e, w{"gransy.nic.zm", "ns-zm.afrinic.net", "pch.nic.zm"}, n, n, n, "whois.zicta.zm", e, n, f}
	z[1574] = Zone{"zone", r, x, 0x40, "Binky Moon, LLC", e, w{"v0n0.nic.zone", "v0n1.nic.zone", "v0n2.nic.zone", "v0n3.nic.zone", "v2n0.nic.zone", "v2n1.nic.zone"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.zone", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1575] = Zone{"zuerich", r, x, 0xc4, "Kanton Zürich (Canton of Zurich)", "https://nic.zuerich/", w{"a.nic.zuerich", "b.nic.zuerich", "c.nic.zuerich", "d.nic.zuerich"}, n, w{"Zurich", "CH-ZH"}, w{"de", "de-CH"}, "whois.nic.zuerich", e, w{"https://rdap.centralnic.com/zuerich/"}, t}
}

//go:noinline
func i1575() {
	z[1576] = Zone{"zulu", r, x, 0x10040, e, e, n, n, n, n, e, e, n, f}
	z[1577] = Zone{"zw", r, z[4863:4867], 0xa0, e, e, w{"ns1.liquidtelecom.net", "ns1zim.telone.co.zw", "ns2.liquidtelecom.net", "ns2zim.telone.co.zw", "zw-ns.anycast.pch.net"}, n, n, n, e, "http://www.zispa.org.zw/", n, f}
	z[1578] = Zone{"xn--jxalpdlp" /* δοκιμή */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1579] = Zone{"xn--qxam" /* ελ */, r, x, 0xa0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, "https://grweb.ics.forth.gr/", n, t}
	z[1580] = Zone{"xn--qxa6a" /* ευ */, r, x, 0xa0, e, e, w{"be.dns.eu", "si.dns.eu", "w.dns.eu", "x.dns.eu", "y.dns.eu"}, n, n, n, "whois.eu", e, n, t}
	z[1581] = Zone{"xn--90ae" /* бг */, r, x, 0xa0, e, e, w{"a.nic.bg", "b.nic.bg", "c.nic.bg", "d.nic.bg", "e.nic.bg", "p.nic.bg"}, n, n, w{"bg-BG"}, "whois.imena.bg", e, n, t}
	z[1582] = Zone{"xn--90ais" /* бел */, r, x, 0xa0, e, e, w{"dns1.tld.becloudby.com", "dns2.tld.becloudby.com", "dns3.tld.becloudby.com", "dns4.tld.becloudby.com", "dns7.tld.becloudby.com"}, n, n, n, "whois.cctld.by", e, n, t}
	z[1583] = Zone{"xn--d1acj3b" /* дети */, r, x, 0x40, "The Foundation for Network Initiatives “The Smart Internet”", e, w{"a.dns.ripn.net", "b.dns.ripn.net", "d.dns.ripn.net", "e.dns.ripn.net", "f.dns.ripn.net"}, n, n, w{"ru"}, "whois.nic.xn--d1acj3b", e, w{"https://whois.nic.xn--/rdap/-vofi1a3h"}, t}
}

//go:noinline
func i1583() {
	z[1584] = Zone{"xn--e1a4c" /* ею */, r, x, 0xa0, e, e, w{"be.dns.eu", "si.dns.eu", "w.dns.eu", "x.dns.eu", "y.dns.eu"}, n, n, w{"mul-Cyrl"}, "whois.eu", e, n, t}
	z[1585] = Zone{"xn--80akhbyknj4f" /* испытание */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1586] = Zone{"xn--80aqecdr1a" /* католик */, r, x, 0x40, "Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)", "http://nic.xn--80aqecdr1a/", w{"a.nic.xn--80aqecdr1a", "b.nic.xn--80aqecdr1a", "c.nic.xn--80aqecdr1a", "x.nic.xn--80aqecdr1a", "y.nic.xn--80aqecdr1a", "z.nic.xn--80aqecdr1a"}, n, n, w{"mul-Cyrl"}, "whois.nic.xn--80aqecdr1a", e, w{"https://rdap.nic.xn--/-7sbygceu5a"}, t}
	z[1587] = Zone{"xn--j1aef" /* ком */, r, x, 0x40, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--j1aef", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--j1aef", e, w{"https://tld-rdap.verisign.com/xn--j1aef/v1/"}, t}
	z[1588] = Zone{"xn--d1alf" /* мкд */, r, x, 0xa0, e, "https://www.iana.org/domains/root/db/xn--d1alf.html", w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, "whois.marnet.mk", e, n, t}
	z[1589] = Zone{"xn--l1acc" /* мон */, r, x, 0xa0, "Datacom Co.,Ltd", "https://www.mon.mn/", w{"ns1.idn.mn", "ns2.idn.mn", "ns3.idn.mn"}, n, n, w{"mn"}, e, e, n, t}
	z[1590] = Zone{"xn--80adxhks" /* москва */, r, x, 0xc4, "Foundation for Assistance for Internet Technologies and Infrastructure Development (FAITID)", e, w{"a.dns.flexireg.ru", "b.dns.flexireg.net", "c.dns.flexireg.org", "d.dns.flexireg.domains"}, n, w{"Moscow"}, w{"ru"}, "whois.nic.xn--80adxhks", e, w{"https://flexireg.net/xn--80adxhks/rdap/"}, t}
	z[1591] = Zone{"xn--80asehdb" /* онлайн */, r, x, 0x40, "CORE Association", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Cyrl"}, "whois.nic.xn--80asehdb", e, w{"https://rdap.nic.xn--/-7sb1agjdc"}, t}
}

//go:noinline
func i1591() {
	z[1592] = Zone{"xn--c1avg" /* орг */, r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/domain-products/opr/", w{"a0.nic.xn--c1avg", "a2.nic.xn--c1avg", "b0.nic.xn--c1avg", "c0.nic.xn--c1avg"}, n, n, w{"be", "bg", "bs", "mk", "ru", "sr", "sr-ME", "uk"}, "whois.nic.xn--c1avg", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[1593] = Zone{"xn--p1acf" /* рус */, r, z[4867:5067], 0x40, "Rusnames Limited", "https://www.webnames.ru/", w{"ns1.anycastdns.cz", "ns1.nic.xn--p1acf", "ns2.anycastdns.cz", "ns2.nic.xn--p1acf"}, n, n, w{"mul-Cyrl"}, "whois.nic.xn--p1acf", e, w{"https://api.rdap.nic.xn--/-4tbdh"}, t}
	z[1594] = Zone{"xn--p1ai" /* рф */, r, z[5067:5070], 0xa0, e, "https://www.iana.org/domains/root/db/xn--p1ai.html", w{"a.dns.ripn.net", "b.dns.ripn.net", "d.dns.ripn.net", "e.dns.ripn.net", "f.dns.ripn.net"}, n, n, w{"mul-Cyrl", "ru-RU"}, "whois.tcinet.ru", e, n, t}
	z[1595] = Zone{"xn--80aswg" /* сайт */, r, x, 0x40, "CORE Association", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Cyrl"}, "whois.nic.xn--80aswg", e, w{"https://rdap.nic.xn--/-7sb1a4ah"}, t}
	z[1596] = Zone{"xn--90a3ac" /* срб */, r, z[5070:5075], 0xa0, e, e, w{"a.nic.rs", "b.nic.rs", "f.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, w{"bg", "cnr-Cyrl", "mk", "ru", "rue-Cyrl", "sr-Cyrl", "uk"}, "whois.rnids.rs", e, n, t}
	z[1597] = Zone{"xn--j1amh" /* укр */, r, x, 0xa0, e, "http://uanic.net/", w{"dns.tci.net.ua", "dns1.u-registry.com", "dns2.u-registry.net", "dns3.dotukr.com", "tier1.num.net.ua", "ukr.ns.ua"}, n, n, w{"crh-UA", "mul-Cyrl-UA", "ru-UA", "uk-UA"}, "whois.dotukr.com", e, n, t}
	z[1598] = Zone{"xn--80ao21a" /* қаз */, r, x, 0xa0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, "whois.nic.kz", e, n, t}
	z[1599] = Zone{"xn--y9a3aq" /* հայ */, r, x, 0xa0, e, e, w{"fork.sth.dnsnode.net", "ns-cdn.amnic.net", "ns-pch.amnic.net", "ns-pri.nic.am"}, n, n, n, "whois.amnic.net", e, n, t}
}

//go:noinline
func i1599() {
	z[1600] = Zone{"xn--deba0ad" /* טעסט */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1601] = Zone{"xn--4dbrk0ce" /* ישראל */, r, x, 0xa0, e, e, w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, "whois.isoc.org.il", e, n, t}
	z[1602] = Zone{"xn--9dbq2a" /* קום */, r, x, 0, "VeriSign Sarl", e, w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--9dbq2a", e, w{"https://tld-rdap.verisign.com/xn--9dbq2a/v1/"}, t}
	z[1603] = Zone{"xn--hgbk6aj7f53bba" /* آزمایشی */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1604] = Zone{"xn--kgbechtv" /* إختبار */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1605] = Zone{"xn--mgbca7dzdo" /* ابوظبي */, r, x, 0x40, "Abu Dhabi Systems and Information Centre", "https://www.icann.org/en/registry-agreements/details/xn--mgbca7dzdo", w{"gtld.alpha.aridns.net.au", "gtld.beta.aridns.net.au", "gtld.delta.aridns.net.au", "gtld.gamma.aridns.net.au"}, n, n, w{"ar"}, "whois.nic.xn--mgbca7dzdo", e, w{"https://rdap.nic.xn--/-ymcda3fteqa"}, t}
	z[1606] = Zone{"xn--mgbaakc7dvf" /* اتصالات */, r, x, 0x1040, "Emirates Telecommunications Corporation (trading as Etisalat)", "https://nic.xn--mgbaakc7dvf/", n, n, n, w{"ar"}, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/xn--mgbaakc7dvf/"}, t}
	z[1607] = Zone{"xn--mgba3a3ejt" /* ارامكو */, r, x, 0x40, "Aramco Services Company", "http://www.nic.xn--mgba3a3ejt/", w{"a.nic.xn--mgba3a3ejt", "b.nic.xn--mgba3a3ejt", "c.nic.xn--mgba3a3ejt", "ns4.dns.nic.xn--mgba3a3ejt", "ns5.dns.nic.xn--mgba3a3ejt", "ns6.dns.nic.xn--mgba3a3ejt"}, n, n, w{"ar", "es"}, "whois.nic.xn--mgba3a3ejt", e, w{"https://rdap.nic.xn--/-ymca3b0flw"}, t}
}

//go:noinline
func i1607() {
	z[1608] = Zone{"xn--mgbayh7gpa" /* الاردن */, r, x, 0xa0, e, e, w{"amra.nic.gov.jo", "cloudns.nic.net.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, "http://idn.jo/whois_a.aspx", n, t}
	z[1609] = Zone{"xn--mgbcpq6gpa1a" /* البحرين */, r, x, 0xa0, e, e, w{"a.nic.xn--mgbcpq6gpa1a", "b.nic.xn--mgbcpq6gpa1a", "c.nic.xn--mgbcpq6gpa1a", "d.nic.xn--mgbcpq6gpa1a"}, n, n, n, e, e, n, t}
	z[1610] = Zone{"xn--lgbbat1ad8j" /* الجزائر */, r, x, 0xa0, e, "https://www.iana.org/domains/root/db/xn--lgbbat1ad8j.html", w{"idn1.nic.dz", "idn2.nic.dz"}, n, n, n, "whois.nic.dz", e, n, t}
	z[1611] = Zone{"xn--mgberp4a5d4ar" /* السعودية */, r, x, 0xa0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, w{"ar", "ar-SA"}, "whois.nic.net.sa", e, n, t}
	z[1612] = Zone{"xn--mgba7c0bbn0a" /* العليان */, r, x, 0x40, "Competrol (Luxembourg) Sarl", "http://nic.xn--mgba7c0bbn0a/", w{"a.nic.xn--mgba7c0bbn0a", "b.nic.xn--mgba7c0bbn0a", "c.nic.xn--mgba7c0bbn0a", "x.nic.xn--mgba7c0bbn0a", "y.nic.xn--mgba7c0bbn0a", "z.nic.xn--mgba7c0bbn0a"}, n, n, w{"ar"}, "whois.nic.xn--mgba7c0bbn0a", e, w{"https://rdap.nic.xn--/-ymca5e9bbp4a"}, t}
	z[1613] = Zone{"xn--mgbc0a9azcg" /* المغرب */, r, x, 0xa0, e, "https://www.iana.org/domains/root/db/xn--mgbc0a9azcg.html", w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, "whois.iam.net.ma", e, n, t}
	z[1614] = Zone{"xn--mgbaam7a8h" /* امارات */, r, x, 0xa0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "nsext-pch.aedns.ae"}, n, n, w{"ar-AE"}, "whois.aeda.net.ae", e, n, t}
	z[1615] = Zone{"xn--mgba3a4f16a" /* ایران */, r, x, 0xa0, e, e, w{"a.nic.ir", "b.nic.ir"}, n, n, n, "whois.nic.ir", e, n, t}
}

//go:noinline
func i1615() {
	z[1616] = Zone{"xn--mgbbh1a" /* بارت */, r, z[5075:5076], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-MH"}, w{"ks-IN"}, "whois.registry.in", e, n, t}
	z[1617] = Zone{"xn--mgbab2bd" /* بازار */, r, x, 0x40, "CORE Association", e, w{"anycast10.irondns.net", "anycast23.irondns.net", "anycast24.irondns.net", "anycast9.irondns.net"}, n, n, w{"mul-Arab"}, "whois.nic.xn--mgbab2bd", e, w{"https://rdap.nic.xn--/-ymcac1ce"}, t}
	z[1618] = Zone{"xn--ngbe9e0a" /* بيتك */, r, x, 0x40, "Kuwait Finance House", "https://nic.xn--ngbe9e0a/", w{"a.nic.xn--ngbe9e0a", "b.nic.xn--ngbe9e0a", "c.nic.xn--ngbe9e0a", "d.nic.xn--ngbe9e0a"}, n, n, w{"ar"}, "whois.nic.xn--ngbe9e0a", e, w{"https://rdap.centralnic.com/xn--ngbe9e0a/"}, t}
	z[1619] = Zone{"xn--mgbbh1a71e" /* بھارت */, r, z[5076:5077], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-JK", "IN-TG", "IN-UP", "IN-BR", "IN-JH", "IN-WB", "IN-DL"}, w{"ur-IN"}, "whois.registry.in", e, n, t}
	z[1620] = Zone{"xn--pgbs0dh" /* تونس */, r, x, 0xa0, e, e, w{"ns-tn.afrinic.net", "ns1.ati.tn", "ns2.ati.tn", "ns2.nic.fr", "pch.ati.tn", "rip.psg.com"}, n, n, n, "whois.ati.tn", e, n, t}
	z[1621] = Zone{"xn--mgbpl2fh" /* سودان */, r, x, 0xa0, e, e, w{"ans1.sis.sd", "pch.sis.sd"}, n, n, n, e, e, n, t}
	z[1622] = Zone{"xn--ogbpf8fl" /* سورية */, r, x, 0xa0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, "whois.tld.sy", e, n, t}
	z[1623] = Zone{"xn--ngbc5azd" /* شبكة */, r, x, 0x40, "International Domain Registry Pty. Ltd.", "https://www.dotshabaka.com/", w{"a.nic.xn--ngbc5azd", "b.nic.xn--ngbc5azd", "c.nic.xn--ngbc5azd", "x.nic.xn--ngbc5azd", "y.nic.xn--ngbc5azd", "z.nic.xn--ngbc5azd"}, n, n, w{"mul-Arab"}, "whois.nic.xn--ngbc5azd", e, w{"https://rdap.nic.xn--/-0mcd6b8d"}, t}
}

//go:noinline
func i1623() {
	z[1624] = Zone{"xn--mgbtx2b" /* عراق */, r, x, 0xa0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, "whois.cmc.iq", e, n, t}
	z[1625] = Zone{"xn--ngbrx" /* عرب */, r, x, 0x40, "League of Arab States", "http://nic.xn--ngbrx/", w{"gtld.alpha.aridns.net.au", "gtld.beta.aridns.net.au", "gtld.delta.aridns.net.au", "gtld.gamma.aridns.net.au"}, n, n, w{"ar"}, "whois.nic.xn--ngbrx", e, w{"https://rdap.nic.xn--/-0mc0a5a"}, t}
	z[1626] = Zone{"xn--mgb9awbf" /* عمان */, r, x, 0xa0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, "whois.registry.om", e, n, t}
	z[1627] = Zone{"xn--ygbi2ammx" /* فلسطين */, r, x, 0xa0, e, e, w{"dns1.gov.ps", "dns3.gov.ps", "idn.pnina.ps", "ns1.pnina.ps"}, n, n, n, "whois.pnina.ps", e, n, t}
	z[1628] = Zone{"xn--wgbl6a" /* قطر */, r, x, 0xa0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, "whois.registry.qa", e, n, t}
	z[1629] = Zone{"xn--mgbi4ecexp" /* كاثوليك */, r, x, 0x40, "Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)", "http://nic.xn--mgbi4ecexp/", w{"a.nic.xn--mgbi4ecexp", "b.nic.xn--mgbi4ecexp", "c.nic.xn--mgbi4ecexp", "x.nic.xn--mgbi4ecexp", "y.nic.xn--mgbi4ecexp", "z.nic.xn--mgbi4ecexp"}, n, n, w{"ar"}, "whois.nic.xn--mgbi4ecexp", e, w{"https://rdap.nic.xn--/-ymcm8gcf1ar"}, t}
	z[1630] = Zone{"xn--fhbei" /* كوم */, r, x, 0x40, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--fhbei", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--fhbei", e, w{"https://tld-rdap.verisign.com/xn--fhbei/v1/"}, t}
	z[1631] = Zone{"xn--pgb3ceoj" /* كيوتل */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i1631() {
	z[1632] = Zone{"xn--wgbh1c" /* مصر */, r, x, 0xa0, e, e, w{"ns1.dotmasr.eg", "ns2.dotmasr.eg", "ns3.dotmasr.eg", "ns4.dotmasr.eg"}, n, n, w{"ar-EG"}, "whois.dotmasr.eg", e, n, t}
	z[1633] = Zone{"xn--mgbx4cd0ab" /* مليسيا */, r, x, 0x40, e, e, w{"a.mynic.centralnic-dns.com", "a.nic.my", "a1.nic.my", "b.mynic.centralnic-dns.com", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, "whois.mynic.my", e, n, t}
	z[1634] = Zone{"xn--mgbb9fbpob" /* موبايلي */, r, x, 0x1040, e, e, n, n, n, n, "whois.nic.xn--mgbb9fbpob", e, n, t}
	z[1635] = Zone{"xn--mgbah1a3hjkrd" /* موريتانيا */, r, x, 0xa0, e, e, w{"ns-mr.afrinic.net", "ns-mr.nic.fr", "ns1.nic.mr", "ns2.nic.mr", "ns3.nic.mr"}, n, n, n, "whois.nic.mr", e, n, t}
	z[1636] = Zone{"xn--mgbv6cfpo" /* موزايك */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1637] = Zone{"xn--4gbrim" /* موقع */, r, x, 0x40, "Helium TLDs Ltd", "https://www.icann.org/en/registry-agreements/details/xn--4gbrim", w{"a.nic.xn--4gbrim", "b.nic.xn--4gbrim", "c.nic.xn--4gbrim", "d.nic.xn--4gbrim"}, n, n, n, "whois.nic.xn--4gbrim", e, w{"https://rdap.centralnic.com/xn--4gbrim/"}, t}
	z[1638] = Zone{"xn--mgbt3dhd" /* همراه */, r, x, 0x40, "Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.", "https://www.icann.org/en/registry-agreements/details/xn--mgbt3dhd", w{"a.ns.nic.xn--mgbt3dhd", "b.ns.nic.xn--mgbt3dhd", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, w{"Iran"}, n, "whois.nic.xn--mgbt3dhd", e, w{"https://api.rdap.nic.xn--/-ymc3a9ejd"}, t}
	z[1639] = Zone{"xn--mgbai9azgqp6j" /* پاکستان */, r, x, 0, e, e, w{"ns.ntc.net.pk", "ns1.ntc.net.pk", "ns2.ntc.net.pk"}, n, n, n, e, e, n, t}
}

//go:noinline
func i1639() {
	z[1640] = Zone{"xn--mgbgu82a" /* ڀارت */, r, z[5077:5078], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-RJ", "IN-GJ", "IN-MP"}, w{"sd-IN"}, "whois.registry.in", e, n, t}
	z[1641] = Zone{"xn--11b4c3d" /* कॉम */, r, x, 0x40, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--11b4c3d", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--11b4c3d", e, w{"https://tld-rdap.verisign.com/xn--11b4c3d/v1/"}, t}
	z[1642] = Zone{"xn--c2br7g" /* नेट */, r, x, 0x40, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--c2br7g", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--c2br7g", e, w{"https://tld-rdap.verisign.com/xn--c2br7g/v1/"}, t}
	z[1643] = Zone{"xn--11b5bs3a9aj6g" /* परीक्षा */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1644] = Zone{"xn--h2brj9c" /* भारत */, r, z[5078:5079], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN"}, w{"hi-IN"}, "whois.registry.in", e, n, t}
	z[1645] = Zone{"xn--h2breg3eve" /* भारतम् */, r, z[5079:5080], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, w{"sa-IN"}, "whois.registry.in", e, n, t}
	z[1646] = Zone{"xn--h2brj9c8c" /* भारोत */, r, z[5080:5081], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-AS", "IN-BR", "IN-JH", "IN-MZ", "IN-OR", "IN-TR", "IN-WB"}, w{"sat-Olck-IN"}, "whois.registry.in", e, n, t}
	z[1647] = Zone{"xn--i1b6b1a6a2e" /* संगठन */, r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/domain-products/%E0%A4%B8%E0%A4%82%E0%A4%97%E0%A4%A0%E0%A4%A8/", w{"a0.nic.xn--i1b6b1a6a2e", "a2.nic.xn--i1b6b1a6a2e", "b0.nic.xn--i1b6b1a6a2e", "c0.nic.xn--i1b6b1a6a2e"}, n, n, w{"hi"}, "whois.nic.xn--i1b6b1a6a2e", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
}

//go:noinline
func i1647() {
	z[1648] = Zone{"xn--54b7fta0cc" /* বাংলা */, r, x, 0xa0, e, e, w{"bayanno.btcl.net.bd", "bd-ns.anycast.pch.net", "ekushey.btcl.net.bd"}, n, n, n, e, e, n, t}
	z[1649] = Zone{"xn--45brj9c" /* ভারত */, r, z[5081:5083], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-WB", "IN-TR", "IN-AS"}, w{"bn-IN"}, "whois.registry.in", e, n, t}
	z[1650] = Zone{"xn--45br5cyl" /* ভাৰত */, r, z[5083:5084], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-AS", "IN-AR", "IN-NL"}, w{"as-IN"}, "whois.registry.in", e, n, t}
	z[1651] = Zone{"xn--s9brj9c" /* ਭਾਰਤ */, r, z[5084:5085], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-PB"}, w{"pa-Guru-IN"}, "whois.registry.in", e, n, t}
	z[1652] = Zone{"xn--gecrj9c" /* ભારત */, r, z[5085:5086], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-GJ"}, w{"gu-IN"}, "whois.registry.in", e, n, t}
	z[1653] = Zone{"xn--3hcrj9c" /* ଭାରତ */, r, z[5086:5087], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-OR", "IN-AP", "IN-MP", "IN-JH", "IN-WB", "IN-CT"}, w{"or-IN"}, "whois.registry.in", e, n, t}
	z[1654] = Zone{"xn--xkc2dl3a5ee0h" /* இந்தியா */, r, z[5087:5088], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-TN", "IN-PY"}, w{"ta-IN"}, "whois.registry.in", e, n, t}
	z[1655] = Zone{"xn--xkc2al3hye2a" /* இலங்கை */, r, x, 0xa0, e, "https://www.iana.org/domains/root/db/xn--xkc2al3hye2a.html", w{"lk.communitydns.net", "nic.lk-anycast.pch.net", "ns-c.nic.lk", "ns-d.nic.lk", "ns-l.nic.lk", "ns-t.nic.lk", "ns1.ac.lk", "ns3.ac.lk"}, n, n, n, "whois.nic.lk", e, n, t}
}

//go:noinline
func i1655() {
	z[1656] = Zone{"xn--clchc0ea0b2g2a9gcd" /* சிங்கப்பூர் */, r, x, 0xa0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "ns4.apnic.net", "pch.sgzones.sg"}, n, n, n, "whois.sgnic.sg", e, n, t}
	z[1657] = Zone{"xn--hlcj6aya9esc7a" /* பரிட்சை */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1658] = Zone{"xn--fpcrj9c3d" /* భారత్ */, r, z[5088:5089], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-AP", "IN-TG"}, w{"te-IN"}, "whois.registry.in", e, n, t}
	z[1659] = Zone{"xn--2scrj9c" /* ಭಾರತ */, r, z[5089:5090], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-KA", "IN-MH", "IN-AP", "IN-TN", "IN-TG", "IN-KL", "IN-GA"}, w{"kn-IN"}, "whois.registry.in", e, n, t}
	z[1660] = Zone{"xn--rvc1e0am3e" /* ഭാരതം */, r, z[5090:5091], 0xa0, e, "https://www.registry.in/", w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, w{"IN-KL", "IN-LD", "IN-PY"}, w{"ml-IN"}, "whois.registry.in", e, n, t}
	z[1661] = Zone{"xn--fzc2c9e2c" /* ලංකා */, r, x, 0xa0, e, e, w{"lk.communitydns.net", "nic.lk-anycast.pch.net", "ns-c.nic.lk", "ns-d.nic.lk", "ns-l.nic.lk", "ns-t.nic.lk", "ns1.ac.lk", "ns3.ac.lk"}, n, n, n, "whois.nic.lk", e, n, t}
	z[1662] = Zone{"xn--42c2d9a" /* คอม */, r, x, 0x40, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--42c2d9a", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--42c2d9a", e, w{"https://tld-rdap.verisign.com/xn--42c2d9a/v1/"}, t}
	z[1663] = Zone{"xn--o3cw4h" /* ไทย */, r, z[5091:5092], 0xa0, e, e, w{"a.thains.co.th", "b.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, "whois.thnic.co.th", e, n, t}
}

//go:noinline
func i1663() {
	z[1664] = Zone{"xn--q7ce6a" /* ລາວ */, r, x, 0xa0, e, e, w{"a.xn--q7ce6a.centralnic-dns.com", "b.xn--q7ce6a.centralnic-dns.com", "c.xn--q7ce6a.centralnic-dns.com", "d.xn--q7ce6a.centralnic-dns.com"}, n, n, n, "whois.nic.la", e, n, t}
	z[1665] = Zone{"xn--node" /* გე */, r, x, 0xa0, e, "http://nic.xn--node/", w{"a.xn--node.globalanycastcloud.freenom.net", "b.xn--node.globalanycastcloud.freenom.net", "c.xn--node.globalanycastcloud.freenom.net", "d.xn--node.globalanycastcloud.freenom.net", "xn--node.ns.anycast.pch.net"}, w{"188.93.95.11"}, n, n, "whois.itdc.ge", e, n, t}
	z[1666] = Zone{"xn--q9jyb4c" /* みんな */, r, x, 0x40, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1667] = Zone{"xn--cckwcxetd" /* アマゾン */, r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.xn--cckwcxetd/", w{"dns1.nic.xn--cckwcxetd", "dns2.nic.xn--cckwcxetd", "dns3.nic.xn--cckwcxetd", "dns4.nic.xn--cckwcxetd", "dnsa.nic.xn--cckwcxetd", "dnsb.nic.xn--cckwcxetd", "dnsc.nic.xn--cckwcxetd", "dnsd.nic.xn--cckwcxetd"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--cckwcxetd", e, w{"https://rdap.nominet.uk/xn--cckwcxetd/"}, t}
	z[1668] = Zone{"xn--gckr3f0f" /* クラウド */, r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.xn--gckr3f0f/", w{"a.nic.xn--gckr3f0f", "b.nic.xn--gckr3f0f", "c.nic.xn--gckr3f0f", "ns1.dns.nic.xn--gckr3f0f", "ns2.dns.nic.xn--gckr3f0f", "ns3.dns.nic.xn--gckr3f0f"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--gckr3f0f", e, w{"https://rdap.nic.xn--/-meu0a2h0g"}, t}
	z[1669] = Zone{"xn--qcka1pmc" /* グーグル */, r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1670] = Zone{"xn--tckwe" /* コム */, r, x, 0, "VeriSign Sarl", e, w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--tckwe", e, w{"https://tld-rdap.verisign.com/xn--tckwe/v1/"}, t}
	z[1671] = Zone{"xn--cck2b3b" /* ストア */, r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.xn--cck2b3b/", w{"a.nic.xn--cck2b3b", "b.nic.xn--cck2b3b", "c.nic.xn--cck2b3b", "ns1.dns.nic.xn--cck2b3b", "ns2.dns.nic.xn--cck2b3b", "ns3.dns.nic.xn--cck2b3b"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--cck2b3b", e, w{"https://rdap.nic.xn--/-eeu2cwc"}, t}
}

//go:noinline
func i1671() {
	z[1672] = Zone{"xn--1ck2e1b" /* セール */, r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.xn--1ck2e1b/", w{"a.nic.xn--1ck2e1b", "b.nic.xn--1ck2e1b", "c.nic.xn--1ck2e1b", "ns1.dns.nic.xn--1ck2e1b", "ns2.dns.nic.xn--1ck2e1b", "ns3.dns.nic.xn--1ck2e1b"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--1ck2e1b", e, w{"https://rdap.nic.xn--/-tfuygrc"}, t}
	z[1673] = Zone{"xn--zckzah" /* テスト */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1674] = Zone{"xn--bck1b9a5dre4c" /* ファッション */, r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.xn--bck1b9a5dre4c/", w{"a.nic.xn--bck1b9a5dre4c", "b.nic.xn--bck1b9a5dre4c", "c.nic.xn--bck1b9a5dre4c", "ns1.dns.nic.xn--bck1b9a5dre4c", "ns2.dns.nic.xn--bck1b9a5dre4c", "ns3.dns.nic.xn--bck1b9a5dre4c"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--bck1b9a5dre4c", e, w{"https://rdap.nic.xn--/-ceu0c2b2e5esd"}, t}
	z[1675] = Zone{"xn--eckvdtc9d" /* ポイント */, r, x, 0x40, "Amazon Registry Services, Inc.", "https://nic.xn--eckvdtc9d/", w{"a.nic.xn--eckvdtc9d", "b.nic.xn--eckvdtc9d", "c.nic.xn--eckvdtc9d", "ns1.dns.nic.xn--eckvdtc9d", "ns2.dns.nic.xn--eckvdtc9d", "ns3.dns.nic.xn--eckvdtc9d"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--eckvdtc9d", e, w{"https://rdap.nic.xn--/-ieu2end5e"}, t}
	z[1676] = Zone{"xn--4gq48lf9j" /* 一号店 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1677] = Zone{"xn--rhqv96g" /* 世界 */, r, x, 0, "Stable Tone Limited", e, w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh-Hans", "zh-Hant"}, "whois.nic.xn--rhqv96g", e, w{"https://rdap.teleinfo.cn/xn--rhqv96g/"}, t}
	z[1678] = Zone{"xn--fiq64b" /* 中信 */, r, x, 0x42, "CITIC Group Corporation", "https://www.icann.org/en/registry-agreements/details/xn--fiq64b", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--fiq64b/"}, t}
	z[1679] = Zone{"xn--fiqs8s" /* 中国 */, r, x, 0xa0, e, "https://wanwang.aliyun.com/hosting/cn_domain", w{"h.dns.cn", "i.dns.cn", "j.dns.cn", "k.dns.cn", "l.dns.cn"}, w{"218.241.105.10", "wildcard.cnnic.cn"}, n, w{"zh-Hans-CN"}, "cwhois.cnnic.cn", e, n, t}
}

//go:noinline
func i1679() {
	z[1680] = Zone{"xn--fiqz9s" /* 中國 */, r, x, 0xa0, e, "https://wanwang.aliyun.com/hosting/cn_domain", w{"h.dns.cn", "i.dns.cn", "j.dns.cn", "k.dns.cn", "l.dns.cn"}, w{"218.241.105.10", "wildcard.cnnic.cn"}, n, w{"zh-Hant-CN"}, "cwhois.cnnic.cn", e, n, t}
	z[1681] = Zone{"xn--fiq228c5hs" /* 中文网 */, r, x, 0x40, "TLD REGISTRY LIMITED OY", e, w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"zh-CN", "zh-Hans", "zh-Hant"}, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/xn--fiq228c5hs/"}, t}
	z[1682] = Zone{"xn--jlq480n2rg" /* 亚马逊 */, r, x, 0x42, "Amazon Registry Services, Inc.", "https://nic.xn--jlq480n2rg/", w{"dns1.nic.xn--jlq480n2rg", "dns2.nic.xn--jlq480n2rg", "dns3.nic.xn--jlq480n2rg", "dns4.nic.xn--jlq480n2rg", "dnsa.nic.xn--jlq480n2rg", "dnsb.nic.xn--jlq480n2rg", "dnsc.nic.xn--jlq480n2rg", "dnsd.nic.xn--jlq480n2rg"}, n, n, w{"da", "de", "es", "fi", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--jlq480n2rg", e, w{"https://rdap.nominet.uk/xn--jlq480n2rg/"}, t}
	z[1683] = Zone{"xn--vhquv" /* 企业 */, r, x, 0, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.xn--vhquv", "v0n1.nic.xn--vhquv", "v0n2.nic.xn--vhquv", "v0n3.nic.xn--vhquv", "v2n0.nic.xn--vhquv", "v2n1.nic.xn--vhquv"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.xn--vhquv", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1684] = Zone{"xn--1qqw23a" /* 佛山 */, r, x, 0xc4, "Guangzhou YU Wei Information Technology Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xn--1qqw23a", w{"ta.ngtld.cn", "tb.ngtld.cn", "tc.ngtld.cn", "td.ngtld.cn", "te.ngtld.cn"}, n, w{"Foshan"}, w{"zh-Hans"}, "whois.ngtld.cn", e, w{"https://restwhois.ngtld.cn/"}, t}
	z[1685] = Zone{"xn--vuq861b" /* 信息 */, r, x, 0, "Beijing Tele-info Technology Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xn--vuq861b", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/"}, t}
	z[1686] = Zone{"xn--nyqy26a" /* 健康 */, r, x, 0, "Stable Tone Limited", "https://www.icann.org/en/registry-agreements/details/xn--nyqy26a", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "it", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh-Hans", "zh-Hant"}, "whois.nic.xn--nyqy26a", e, w{"https://rdap.teleinfo.cn/xn--nyqy26a/"}, t}
	z[1687] = Zone{"xn--45q11c" /* 八卦 */, r, x, 0, "Zodiac Gemini Ltd", "https://www.icann.org/en/registry-agreements/details/xn--45q11c", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/XN--45Q11C/"}, t}
}

//go:noinline
func i1687() {
	z[1688] = Zone{"xn--55qx5d" /* 公司 */, r, x, 0, "China Internet Network Information Center (CNNIC)", e, w{"a.ngtld.cn", "b.ngtld.cn", "c.ngtld.cn", "d.ngtld.cn", "e.ngtld.cn"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.ngtld.cn", e, w{"https://restwhois.ngtld.cn/"}, t}
	z[1689] = Zone{"xn--55qw42g" /* 公益 */, r, x, 0, "China Organizational Name Administration Center", "https://www.icann.org/en/registry-agreements/details/xn--55qw42g", w{"ns1.conac.cn", "ns2.conac.cn", "ns3.conac.cn", "ns4.conac.cn", "ns5.conac.cn"}, n, n, w{"zh"}, "whois.conac.cn", e, w{"https://rdap.conac.cn/"}, t}
	z[1690] = Zone{"xn--kprw13d" /* 台湾 */, r, x, 0xa0, e, e, w{"anytld.apnic.net", "d.dns.tw", "e.dns.tw", "f.dns.tw", "g.dns.tw", "h.dns.tw"}, n, n, w{"zh-Hans-TW"}, "whois.twnic.net.tw", e, n, t}
	z[1691] = Zone{"xn--kpry57d" /* 台灣 */, r, x, 0xa0, e, e, w{"anytld.apnic.net", "d.dns.tw", "e.dns.tw", "f.dns.tw", "g.dns.tw", "h.dns.tw"}, n, n, w{"zh-Hant-TW"}, "whois.twnic.net.tw", e, n, t}
	z[1692] = Zone{"xn--czru2d" /* 商城 */, r, x, 0, "Zodiac Aquarius Limited", e, w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--czru2d/"}, t}
	z[1693] = Zone{"xn--czrs0t" /* 商店 */, r, x, 0, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.xn--czrs0t", "v0n1.nic.xn--czrs0t", "v0n2.nic.xn--czrs0t", "v0n3.nic.xn--czrs0t", "v2n0.nic.xn--czrs0t", "v2n1.nic.xn--czrs0t"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.xn--czrs0t", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1694] = Zone{"xn--czr694b" /* 商标 */, r, x, 0, "Internet DotTrademark Organisation Limited", "https://www.icann.org/en/registry-agreements/details/xn--czr694b", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, e, e, w{"https://rdap.zdnsgtld.com/xn--czr694b/"}, t}
	z[1695] = Zone{"xn--w4rs40l" /* 嘉里 */, r, x, 0x2, "Kerry Trading Co. Limited", "https://www.icann.org/en/registry-agreements/details/xn--w4rs40l", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.xn--w4rs40l", e, w{"https://tld-rdap.verisign.com/xn--w4rs40l/v1/"}, t}
}

//go:noinline
func i1695() {
	z[1696] = Zone{"xn--w4r85el8fhu5dnra" /* 嘉里大酒店 */, r, x, 0x40, "Kerry Trading Co. Limited", "https://www.icann.org/en/registry-agreements/details/xn--w4r85el8fhu5dnra", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.xn--w4r85el8fhu5dnra", e, w{"https://tld-rdap.verisign.com/xn--w4r85el8fhu5dnra/v1/"}, t}
	z[1697] = Zone{"xn--3ds443g" /* 在线 */, r, x, 0x40, "TLD REGISTRY LIMITED OY", e, w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"zh-CN", "zh-Hans", "zh-Hant"}, "whois.teleinfo.cn", e, w{"https://rdap.teleinfo.cn/xn--3ds443g/"}, t}
	z[1698] = Zone{"xn--3oq18vl8pn36a" /* 大众汽车 */, r, x, 0x1042, "Volkswagen (China) Investment Co., Ltd.", e, n, n, n, w{"be", "bg", "bs", "da", "de", "es", "hu", "is", "ko", "lt", "lv", "mk", "pl", "ru", "sr", "sr-ME", "sv", "uk", "zh-Hans", "zh-Hant"}, "whois.nic.xn--3oq18vl8pn36a", e, w{"https://rdap.afilias-srs.net/rdap/xn--3oq18vl8pn36a/"}, t}
	z[1699] = Zone{"xn--pssy2u" /* 大拿 */, r, x, 0, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--pssy2u", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--pssy2u", e, w{"https://tld-rdap.verisign.com/xn--pssy2u/v1/"}, t}
	z[1700] = Zone{"xn--tiq49xqyj" /* 天主教 */, r, x, 0x40, "Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)", "http://nic.xn--tiq49xqyj/", w{"a.nic.xn--tiq49xqyj", "b.nic.xn--tiq49xqyj", "c.nic.xn--tiq49xqyj", "x.nic.xn--tiq49xqyj", "y.nic.xn--tiq49xqyj", "z.nic.xn--tiq49xqyj"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.xn--tiq49xqyj", e, w{"https://rdap.nic.xn--/-dr6ar36avim"}, t}
	z[1701] = Zone{"xn--fjq720a" /* 娱乐 */, r, x, 0, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.xn--fjq720a", "v0n1.nic.xn--fjq720a", "v0n2.nic.xn--fjq720a", "v0n3.nic.xn--fjq720a", "v2n0.nic.xn--fjq720a", "v2n1.nic.xn--fjq720a"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.xn--fjq720a", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1702] = Zone{"xn--fct429k" /* 家電 */, r, x, 0, "Amazon Registry Services, Inc.", "https://nic.xn--fct429k/", w{"a.nic.xn--fct429k", "b.nic.xn--fct429k", "c.nic.xn--fct429k", "ns1.dns.nic.xn--fct429k", "ns2.dns.nic.xn--fct429k", "ns3.dns.nic.xn--fct429k"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--fct429k", e, w{"https://rdap.nic.xn--/-ke2bt46o"}, t}
	z[1703] = Zone{"xn--estv75g" /* 工行 */, r, x, 0x1000, e, e, n, n, n, n, "whois.nic.xn--estv75g", e, n, t}
}

//go:noinline
func i1703() {
	z[1704] = Zone{"xn--xhq521b" /* 广东 */, r, x, 0, "Guangzhou YU Wei Information Technology Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xn--xhq521b", w{"ta.ngtld.cn", "tb.ngtld.cn", "tc.ngtld.cn", "td.ngtld.cn", "te.ngtld.cn"}, n, n, w{"zh-Hans"}, "whois.ngtld.cn", e, w{"https://restwhois.ngtld.cn/"}, t}
	z[1705] = Zone{"xn--6rtwn" /* 广州 */, r, x, 0x10040, "Guangzhou YU Wei Information Technology Co., Ltd.", "https://gtldresult.icann.org/applicationstatus/applicationdetails/1379", n, n, n, n, e, e, n, t}
	z[1706] = Zone{"xn--9krt00a" /* 微博 */, r, x, 0x42, "Sina Corporation", "https://www.icann.org/en/registry-agreements/details/xn--9krt00a", w{"a0.nic.xn--9krt00a", "a2.nic.xn--9krt00a", "b0.nic.xn--9krt00a", "c0.nic.xn--9krt00a"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--9krt00a", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1707] = Zone{"xn--30rr7y" /* 慈善 */, r, x, 0, "Excellent First Limited", "https://www.icann.org/en/registry-agreements/details/xn--30rr7y", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--30rr7y/"}, t}
	z[1708] = Zone{"xn--6qq986b3xl" /* 我爱你 */, r, x, 0x40, "Tycoon Treasure Limited", "http://520.newgtld.cn/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--6qq986b3xl/"}, t}
	z[1709] = Zone{"xn--kput3i" /* 手机 */, r, x, 0x40, "Beijing RITT-Net Technology Development Co., Ltd", "https://www.icann.org/en/registry-agreements/details/xn--kput3i", w{"ns1.teleinfo.cn", "ns2.teleinfoo.com", "ns3.teleinfo.cn", "ns4.teleinfoo.com"}, n, n, w{"zh-CN", "zh-Hans", "zh-Hant", "zh-TW"}, "whois.nic.xn--kput3i", e, w{"https://rdap.teleinfo.cn/xn--kput3i/"}, t}
	z[1710] = Zone{"xn--kpu716f" /* 手表 */, r, x, 0x1000, e, e, n, n, n, n, "whois.nic.xn--kpu716f", e, n, t}
	z[1711] = Zone{"xn--otu796d" /* 招聘 */, r, x, 0, "Jiang Yu Liang Cai Technology Company Limited", "https://www.icann.org/en/registry-agreements/details/xn--otu796d", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, e, e, w{"https://rdap.zdnsgtld.com/xn--otu796d/"}, t}
}

//go:noinline
func i1711() {
	z[1712] = Zone{"xn--zfr164b" /* 政务 */, r, x, 0, "China Organizational Name Administration Center", "https://www.icann.org/en/registry-agreements/details/xn--zfr164b", w{"ns1.conac.cn", "ns2.conac.cn", "ns3.conac.cn", "ns4.conac.cn", "ns5.conac.cn"}, n, n, w{"zh"}, "whois.conac.cn", e, w{"https://rdap.conac.cn/"}, t}
	z[1713] = Zone{"xn--mxtq1m" /* 政府 */, r, x, 0, "Net-Chinese Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xn--mxtq1m", w{"a.nic.xn--mxtq1m", "b.nic.xn--mxtq1m", "c.nic.xn--mxtq1m", "d.nic.xn--mxtq1m"}, n, n, n, "whois.nic.xn--mxtq1m", e, w{"https://rdap.twnic.tw/rdap/"}, t}
	z[1714] = Zone{"xn--yfro4i67o" /* 新加坡 */, r, x, 0xa0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "ns4.apnic.net", "pch.sgzones.sg"}, n, n, n, "whois.sgnic.sg", e, n, t}
	z[1715] = Zone{"xn--efvy88h" /* 新闻 */, r, x, 0, "Guangzhou YU Wei Information Technology Co., Ltd.", "https://www.icann.org/en/registry-agreements/details/xn--efvy88h", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.xn--efvy88h", e, w{"https://rdap.zdnsgtld.com/XN--EFVY88H/"}, t}
	z[1716] = Zone{"xn--9et52u" /* 时尚 */, r, x, 0, "RISE VICTORY LIMITED", "https://www.icann.org/en/registry-agreements/details/xn--9et52u", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--9et52u/"}, t}
	z[1717] = Zone{"xn--kcrx7bb75ajk3b" /* 普利司通 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1718] = Zone{"xn--rovu88b" /* 書籍 */, r, x, 0, "Amazon Registry Services, Inc.", "https://nic.xn--rovu88b/", w{"a.nic.xn--rovu88b", "b.nic.xn--rovu88b", "c.nic.xn--rovu88b", "ns1.dns.nic.xn--rovu88b", "ns2.dns.nic.xn--rovu88b", "ns3.dns.nic.xn--rovu88b"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--rovu88b", e, w{"https://rdap.nic.xn--/-826b334d"}, t}
	z[1719] = Zone{"xn--nqv7f" /* 机构 */, r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/domain-products/%E6%9C%BA%E6%9E%84/", w{"a0.nic.xn--nqv7f", "a2.nic.xn--nqv7f", "b0.nic.xn--nqv7f", "c0.nic.xn--nqv7f"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--nqv7f", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
}

//go:noinline
func i1719() {
	z[1720] = Zone{"xn--tqq33ed31aqia" /* 机构体制 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1721] = Zone{"xn--dkwm73cwpn" /* 欧莱雅 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1722] = Zone{"xn--0zwm56d" /* 测试 */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1723] = Zone{"xn--b4w605ferd" /* 淡马锡 */, r, x, 0x42, "Temasek Holdings (Private) Limited", "https://www.temasek.com.sg/zh/site-services/nictemasek", w{"a0.nic.xn--b4w605ferd", "a2.nic.xn--b4w605ferd", "b0.nic.xn--b4w605ferd", "c0.nic.xn--b4w605ferd"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--b4w605ferd", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1724] = Zone{"xn--fes124c" /* 深圳 */, r, x, 0x10040, "Guangzhou YU Wei Information Technology Co., Ltd.", "https://gtldresult.icann.org/applicationstatus/applicationdetails/1380", n, n, n, n, e, e, n, t}
	z[1725] = Zone{"xn--g6w251d" /* 測試 */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1726] = Zone{"xn--unup4y" /* 游戏 */, r, x, 0, "Binky Moon, LLC", "https://identity.digital/", w{"v0n0.nic.xn--unup4y", "v0n1.nic.xn--unup4y", "v0n2.nic.xn--unup4y", "v0n3.nic.xn--unup4y", "v2n0.nic.xn--unup4y", "v2n1.nic.xn--unup4y"}, n, n, w{"de", "es", "fr", "ja", "ko", "mul-Arab", "mul-Cyrl", "mul-Deva", "mul-Grek", "mul-Hebr", "mul-Latn", "mul-Taml", "mul-Thai", "zh"}, "whois.nic.xn--unup4y", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1727] = Zone{"xn--mix891f" /* 澳門 */, r, x, 0xa0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, "whois.monic.mo", e, n, t}
}

//go:noinline
func i1727() {
	z[1728] = Zone{"xn--3pxu8k" /* 点看 */, r, x, 0, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--3pxu8k", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--3pxu8k", e, w{"https://tld-rdap.verisign.com/xn--3pxu8k/v1/"}, t}
	z[1729] = Zone{"xn--pbt977c" /* 珠宝 */, r, x, 0x1000, e, e, n, n, n, n, "whois.nic.xn--pbt977c", e, n, t}
	z[1730] = Zone{"xn--hxt035cmppuel" /* 盛貿飯店 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1731] = Zone{"xn--6frz82g" /* 移动 */, r, x, 0, "Identity Digital Limited", e, w{"a0.nic.xn--6frz82g", "a2.nic.xn--6frz82g", "b0.nic.xn--6frz82g", "c0.nic.xn--6frz82g"}, n, n, w{"ar", "be", "bg", "bs", "da", "de", "es", "fi", "fr", "hi", "hu", "is", "it", "ko", "lt", "lv", "mk", "pl", "pt", "ru", "sr", "sr-ME", "sv", "uk", "zh-CN", "zh-TW"}, "whois.nic.xn--6frz82g", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1732] = Zone{"xn--nqv7fs00ema" /* 组织机构 */, r, x, 0x40, "Public Interest Registry", "https://thenew.org/org-people/work-with-us/contact/", w{"a0.nic.xn--nqv7fs00ema", "a2.nic.xn--nqv7fs00ema", "b0.nic.xn--nqv7fs00ema", "c0.nic.xn--nqv7fs00ema"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--nqv7fs00ema", e, w{"https://rdap.publicinterestregistry.org/rdap/"}, t}
	z[1733] = Zone{"xn--ses554g" /* 网址 */, r, x, 0, "KNET Co., Ltd.", "http://nic.xn--ses554g/", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"de", "es", "fr", "ja", "ru", "zh-Hans", "zh-Hant"}, "whois.nic.xn--ses554g", e, w{"https://rdap.zdnsgtld.com/xn--ses554g/"}, t}
	z[1734] = Zone{"xn--hxt814e" /* 网店 */, r, x, 0, "Zodiac Taurus Limited", "https://www.icann.org/en/registry-agreements/details/xn--hxt814e", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--hxt814e/"}, t}
	z[1735] = Zone{"xn--5tzm5g" /* 网站 */, r, x, 0x40, "Global Website TLD Asia Limited", "https://identity.digital/", w{"a0.nic.xn--5tzm5g", "a2.nic.xn--5tzm5g", "b0.nic.xn--5tzm5g", "c0.nic.xn--5tzm5g"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--5tzm5g", e, w{"https://rdap.donuts.co/rdap/"}, t}
}

//go:noinline
func i1735() {
	z[1736] = Zone{"xn--io0a7i" /* 网络 */, r, x, 0, "China Internet Network Information Center (CNNIC)", e, w{"a.ngtld.cn", "b.ngtld.cn", "c.ngtld.cn", "d.ngtld.cn", "e.ngtld.cn"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.ngtld.cn", e, w{"https://restwhois.ngtld.cn/"}, t}
	z[1737] = Zone{"xn--8y0a063a" /* 联通 */, r, x, 0x2, "China United Network Communications Corporation Limited", "https://www.icann.org/en/registry-agreements/details/xn--8y0a063a", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.nic.xn--8y0a063a", e, w{"https://rdap.zdnsgtld.com/xn--8y0a063a/"}, t}
	z[1738] = Zone{"xn--jlq61u9w7b" /* 诺基亚 */, r, x, 0x1042, "Nokia Corporation", "http://nic.xn--jlq61u9w7b/", n, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--jlq61u9w7b", e, w{"https://rdap.afilias-srs.net/rdap/xn--jlq61u9w7b/"}, t}
	z[1739] = Zone{"xn--flw351e" /* 谷歌 */, r, x, 0x42, "Charleston Road Registry Inc.", "https://www.registry.google/", w{"ns-tld1.charlestonroadregistry.com", "ns-tld2.charlestonroadregistry.com", "ns-tld3.charlestonroadregistry.com", "ns-tld4.charlestonroadregistry.com", "ns-tld5.charlestonroadregistry.com"}, n, n, w{"mul-Arab", "mul-Armn", "mul-Beng", "mul-Cyrl", "mul-Deva", "mul-Ethi", "mul-Geor", "mul-Grek", "mul-Guru", "mul-Hebr", "mul-Jpan", "mul-Khmr", "mul-Knda", "mul-Kore", "mul-Latn", "mul-Mlym", "mul-Mymr", "mul-Orya", "mul-Sinh", "mul-Taml", "mul-Telu", "mul-Thai", "mul-Tibt", "zh-Hans", "zh-Hant"}, "whois.nic.google", e, w{"https://www.registry.google/rdap/"}, t}
	z[1740] = Zone{"xn--g2xx48c" /* 购物 */, r, x, 0, "Nawang Heli(Xiamen) Network Service Co., LTD.", "http://nic.xn--g2xx48c/", w{"a.nic.xn--g2xx48c", "b.nic.xn--g2xx48c", "c.nic.xn--g2xx48c", "ns1.dns.nic.xn--g2xx48c", "ns2.dns.nic.xn--g2xx48c", "ns3.dns.nic.xn--g2xx48c"}, n, n, w{"zh"}, "whois.nic.xn--g2xx48c", e, w{"https://rdap.nic.xn--/-mv1c947e"}, t}
	z[1741] = Zone{"xn--55qx5d8y0buji4b870u" /* 通用电气公司 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1742] = Zone{"xn--gk3at1e" /* 通販 */, r, x, 0, "Amazon Registry Services, Inc.", "https://nic.xn--gk3at1e/", w{"a.nic.xn--gk3at1e", "b.nic.xn--gk3at1e", "c.nic.xn--gk3at1e", "ns1.dns.nic.xn--gk3at1e", "ns2.dns.nic.xn--gk3at1e", "ns3.dns.nic.xn--gk3at1e"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--gk3at1e", e, w{"https://rdap.nic.xn--/-mu2dk8g"}, t}
	z[1743] = Zone{"xn--3bst00m" /* 集团 */, r, x, 0, "Eagle Horizon Limited", "https://www.icann.org/en/registry-agreements/details/xn--3bst00m", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, "whois.gtld.knet.cn", e, w{"https://rdap.zdnsgtld.com/xn--3bst00m/"}, t}
}

//go:noinline
func i1743() {
	z[1744] = Zone{"xn--fzys8d69uvgm" /* 電訊盈科 */, r, x, 0x42, "PCCW Enterprises Limited", "https://www.icann.org/en/registry-agreements/details/xn--fzys8d69uvgm", w{"a0.nic.xn--fzys8d69uvgm", "a2.nic.xn--fzys8d69uvgm", "b0.nic.xn--fzys8d69uvgm", "c0.nic.xn--fzys8d69uvgm"}, n, n, w{"zh-CN", "zh-TW"}, "whois.nic.xn--fzys8d69uvgm", e, w{"https://rdap.identitydigital.services/rdap/"}, t}
	z[1745] = Zone{"xn--kcrx77d1x4a" /* 飞利浦 */, r, x, 0x42, "Koninklijke Philips N.V.", "http://nic.xn--kcrx77d1x4a/", w{"a.nic.xn--kcrx77d1x4a", "b.nic.xn--kcrx77d1x4a", "c.nic.xn--kcrx77d1x4a", "x.nic.xn--kcrx77d1x4a", "y.nic.xn--kcrx77d1x4a", "z.nic.xn--kcrx77d1x4a"}, n, n, w{"zh", "zh-Hans", "zh-Hant"}, "whois.nic.xn--kcrx77d1x4a", e, w{"https://rdap.nic.xn--/-ue8at05fzx3b"}, t}
	z[1746] = Zone{"xn--jvr189m" /* 食品 */, r, x, 0, "Amazon Registry Services, Inc.", "https://nic.xn--jvr189m/", w{"a.nic.xn--jvr189m", "b.nic.xn--jvr189m", "c.nic.xn--jvr189m", "ns1.dns.nic.xn--jvr189m", "ns2.dns.nic.xn--jvr189m", "ns3.dns.nic.xn--jvr189m"}, n, n, w{"ar", "da", "de", "es", "fi", "fr", "hu", "is", "ja", "ko", "lt", "lv", "no", "pl", "pt", "ru", "sv", "zh"}, "whois.nic.xn--jvr189m", e, w{"https://rdap.nic.xn--/-sh9an60s"}, t}
	z[1747] = Zone{"xn--imr513n" /* 餐厅 */, r, x, 0, "Internet DotTrademark Organisation Limited", "https://www.icann.org/en/registry-agreements/details/xn--imr513n", w{"a.zdnscloud.com", "b.zdnscloud.com", "c.zdnscloud.com", "d.zdnscloud.com", "f.zdnscloud.com", "g.zdnscloud.com", "i.zdnscloud.com", "j.zdnscloud.com"}, n, n, w{"zh-Hans", "zh-Hant"}, e, e, w{"https://rdap.zdnsgtld.com/xn--imr513n/"}, t}
	z[1748] = Zone{"xn--5su34j936bgsg" /* 香格里拉 */, r, x, 0x42, "Shangri‐La International Hotel Management Limited", "https://www.icann.org/en/registry-agreements/details/xn--5su34j936bgsg", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "de", "el", "es", "fr", "hr", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "sr", "sv", "uk", "zh"}, "whois.nic.xn--5su34j936bgsg", e, w{"https://tld-rdap.verisign.com/xn--5su34j936bgsg/v1/"}, t}
	z[1749] = Zone{"xn--j6w193g" /* 香港 */, r, z[5092:5098], 0xa0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, w{"zh-Hans-HK"}, "whois.hkirc.hk", e, n, t}
	z[1750] = Zone{"xn--j6w470d71issc" /* 香港電訊 */, r, x, 0x10040, e, e, n, n, n, n, e, e, n, t}
	z[1751] = Zone{"xn--c1yn36f" /* 點看 */, r, x, 0x10040, "VeriSign Sarl", "https://gtldresult.icann.org/applicationstatus/applicationdetails/1142", n, n, n, n, e, e, n, t}
}

//go:noinline
func i1751() {
	z[1752] = Zone{"xn--t60b56a" /* 닷넷 */, r, x, 0, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--t60b56a", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--t60b56a", e, w{"https://tld-rdap.verisign.com/xn--t60b56a/v1/"}, t}
	z[1753] = Zone{"xn--mk1bu44c" /* 닷컴 */, r, x, 0, "VeriSign Sarl", "https://www.icann.org/en/registry-agreements/details/xn--mk1bu44c", w{"ac1.nstld.com", "ac2.nstld.com", "ac3.nstld.com", "ac4.nstld.com"}, n, n, w{"az", "be", "bg", "el", "ja", "ko", "ku", "mk", "mul-Arab", "mul-Armi", "mul-Armn", "mul-Avst", "mul-Bali", "mul-Bamu", "mul-Batk", "mul-Beng", "mul-Bopo", "mul-Brah", "mul-Bugi", "mul-Buhd", "mul-Cans", "mul-Cari", "mul-Cham", "mul-Cher", "mul-Copt", "mul-Cyrl", "mul-Deva", "mul-Egyp", "mul-Ethi", "mul-Geor", "mul-Glag", "mul-Grek", "mul-Gujr", "mul-Guru", "mul-Hang", "mul-Hani", "mul-Hano", "mul-Hebr", "mul-Hira", "mul-Java", "mul-Kali", "mul-Kana", "mul-Khar", "mul-Khmr", "mul-Knda", "mul-Kthi", "mul-Lana", "mul-Laoo", "mul-Latn", "mul-Lepc", "mul-Limb", "mul-Lisu", "mul-Lyci", "mul-Lydi", "mul-Mand", "mul-Mlym", "mul-Mong", "mul-Mtei", "mul-Mymr", "mul-Nkoo", "mul-Ogam", "mul-Olck", "mul-Orkh", "mul-Orya", "mul-Phag", "mul-Phli", "mul-Phnx", "mul-Prti", "mul-Rjng", "mul-Runr", "mul-Samr", "mul-Sarb", "mul-Saur", "mul-Sinh", "mul-Sund", "mul-Sylo", "mul-Syrc", "mul-Tagb", "mul-Tale", "mul-Talu", "mul-Taml", "mul-Tavt", "mul-Telu", "mul-Tfng", "mul-Tglg", "mul-Thaa", "mul-Thai", "mul-Tibt", "mul-Vaii", "mul-Xpeo", "mul-Xsux", "mul-Yiii", "pl", "ro", "ru", "uk", "zh"}, "whois.nic.xn--mk1bu44c", e, w{"https://tld-rdap.verisign.com/xn--mk1bu44c/v1/"}, t}
	z[1754] = Zone{"xn--cg4bki" /* 삼성 */, r, x, 0x42, "SAMSUNG SDS CO., LTD", "http://nic.xn--cg4bki/", w{"ns1.xn--cg4bki.centralnic-dns.com", "ns2.xn--cg4bki.centralnic-dns.com", "ns3.xn--cg4bki.centralnic-dns.com", "ns4.xn--cg4bki.centralnic-dns.com", "tld01.akam.net", "tld02.akam.net", "tld03.akam.net", "tld04.akam.net"}, n, n, w{"ko"}, "whois.kr", e, w{"https://nic.samsung:8443/rdap/"}, t}
	z[1755] = Zone{"xn--9t4b11yi5a" /* 테스트 */, r, x, 0x9000, "IANA", "https://www.iana.org/domains/reserved", n, n, n, n, e, e, n, t}
	z[1756] = Zone{"xn--3e0b707e" /* 한국 */, r, x, 0xa0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, "whois.kr", e, n, t}
	z[1757] = Zone{"com.ac", &z[10], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1758] = Zone{"gov.ac", &z[10], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1759] = Zone{"mil.ac", &z[10], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1759() {
	z[1760] = Zone{"net.ac", &z[10], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1761] = Zone{"org.ac", &z[10], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1762] = Zone{"nom.ad", &z[19], x, 0, e, e, w{"ns1.andorratelecom.ad", "ns2.andorratelecom.ad"}, n, n, n, e, e, n, f}
	z[1763] = Zone{"ac.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1764] = Zone{"co.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, t}
	z[1765] = Zone{"gov.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1766] = Zone{"mil.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1767] = Zone{"name.ae", &z[23], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1767() {
	z[1768] = Zone{"net.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1769] = Zone{"org.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1770] = Zone{"pro.ae", &z[23], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1771] = Zone{"sch.ae", &z[23], x, 0, e, e, w{"ns1.aedns.ae", "ns2.aedns.ae", "ns4.apnic.net", "nsext-pch.aedns.ae"}, n, n, n, e, e, n, f}
	z[1772] = Zone{"airport.aero", &z[25], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1773] = Zone{"cargo.aero", &z[25], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1774] = Zone{"charter.aero", &z[25], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1775] = Zone{"bank.af", &z[27], x, 0, e, e, w{"ns-us.1and1-dns.com", "ns-us.1and1-dns.us"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1775() {
	z[1776] = Zone{"com.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1777] = Zone{"edu.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1778] = Zone{"gov.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1779] = Zone{"hotel.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1780] = Zone{"law.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1781] = Zone{"music.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1782] = Zone{"net.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1783] = Zone{"org.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1783() {
	z[1784] = Zone{"tv.af", &z[27], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1785] = Zone{"co.ag", &z[32], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[1786] = Zone{"com.ag", &z[32], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[1787] = Zone{"net.ag", &z[32], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[1788] = Zone{"nom.ag", &z[32], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[1789] = Zone{"ok.ag", &z[32], x, 0, e, e, w{"ns1.webagentur.at", "ns2.webagentur.at"}, n, n, n, e, e, n, t}
	z[1790] = Zone{"org.ag", &z[32], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[1791] = Zone{"com.ai", &z[35], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1791() {
	z[1792] = Zone{"net.ai", &z[35], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1793] = Zone{"off.ai", &z[35], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1794] = Zone{"org.ai", &z[35], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1795] = Zone{"com.al", &z[42], x, 0, e, e, w{"al.cctld.authdns.ripe.net", "ns-al.isti.cnr.it", "ns1.nic.al", "ns2.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[1796] = Zone{"edu.al", &z[42], x, 0, e, e, w{"al.cctld.authdns.ripe.net", "ns-al.isti.cnr.it", "ns1.nic.al", "ns2.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[1797] = Zone{"gov.al", &z[42], x, 0, e, e, w{"al.cctld.authdns.ripe.net", "ns-al.isti.cnr.it", "ns1.nic.al", "ns2.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[1798] = Zone{"inima.al", &z[42], x, 0, e, e, w{"ns1.he.net", "ns2.he.net", "ns3.he.net", "ns4.he.net", "ns5.he.net"}, n, n, n, e, e, n, f}
	z[1799] = Zone{"net.al", &z[42], x, 0, e, e, w{"al.cctld.authdns.ripe.net", "ns-al.isti.cnr.it", "ns1.nic.al", "ns2.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1799() {
	z[1800] = Zone{"org.al", &z[42], x, 0, e, e, w{"al.cctld.authdns.ripe.net", "ns-al.isti.cnr.it", "ns1.nic.al", "ns2.nic.al", "nsx.nic.al", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[1801] = Zone{"soros.al", &z[42], x, 0, e, e, w{"ns1.siteground.net", "ns2.siteground.net"}, n, n, n, e, e, n, f}
	z[1802] = Zone{"tirana.al", &z[42], x, 0, e, e, w{"ivan.ns.cloudflare.com", "lady.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[1803] = Zone{"uniti.al", &z[42], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1804] = Zone{"upt.al", &z[42], x, 0, e, e, w{"ns1.keminet.net", "ns1.keminetdns.net", "ns2.keminet.net", "ns2.keminetdns.net"}, n, n, n, e, e, n, f}
	z[1805] = Zone{"co.am", &z[55], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1806] = Zone{"com.am", &z[55], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1807] = Zone{"net.am", &z[55], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1807() {
	z[1808] = Zone{"north.am", &z[55], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com"}, n, n, n, e, e, n, f}
	z[1809] = Zone{"org.am", &z[55], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1810] = Zone{"radio.am", &z[55], x, 0, e, e, w{"a.nic.fm", "b.nic.fm", "c.nic.fm", "d.nic.fm"}, n, n, n, e, e, w{"https://rdap.centralnic.com/radio.am/"}, t}
	z[1811] = Zone{"south.am", &z[55], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com"}, n, n, n, e, e, n, f}
	z[1812] = Zone{"com.an", &z[64], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1813] = Zone{"edu.an", &z[64], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1814] = Zone{"net.an", &z[64], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[1815] = Zone{"org.an", &z[64], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1815() {
	z[1816] = Zone{"co.ao", &z[72], z[5098:5100], 0, e, e, w{"fred.nic.ao", "ns-ext.nic.cz", "ns0.nic.ao", "ns02.fccn.pt", "ns2.reg.it.ao"}, n, n, n, e, e, n, f}
	z[1817] = Zone{"ed.ao", &z[72], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1818] = Zone{"gv.ao", &z[72], x, 0, e, e, w{"ns1.gov.ao", "ns2.gov.ao", "ns3.gov.ao", "ns4.gov.ao"}, n, n, n, e, e, n, f}
	z[1819] = Zone{"it.ao", &z[72], x, 0, e, e, w{"fred.nic.ao", "ns-ext.nic.cz", "ns0.nic.ao", "ns02.fccn.pt", "ns2.reg.it.ao"}, n, n, n, e, e, n, f}
	z[1820] = Zone{"og.ao", &z[72], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1821] = Zone{"pb.ao", &z[72], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1822] = Zone{"com.aq", &z[77], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1823] = Zone{"com.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
}

//go:noinline
func i1823() {
	z[1824] = Zone{"edu.ar", &z[80], z[5100:5101], 0, e, e, w{"a.riu.edu.ar", "b.riu.edu.ar", "noc.uncu.edu.ar", "ns1.uba.ar", "ns2.switch.ch", "unlp.unlp.edu.ar"}, n, n, n, e, e, n, f}
	z[1825] = Zone{"gob.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, f}
	z[1826] = Zone{"gov.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, f}
	z[1827] = Zone{"int.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
	z[1828] = Zone{"mil.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
	z[1829] = Zone{"net.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
	z[1830] = Zone{"org.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
	z[1831] = Zone{"tur.ar", &z[80], x, 0, e, e, w{"a.dns.ar", "ar.cctld.authdns.ripe.net", "b.dns.ar", "c.dns.ar", "e.dns.ar", "f.dns.ar"}, n, n, n, e, e, n, t}
}

//go:noinline
func i1831() {
	z[1832] = Zone{"6tisch.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc9031", n, n, n, n, e, e, n, t}
	z[1833] = Zone{"as112.arpa", &z[86], z[5101:5102], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc7535", w{"a.iana-servers.net", "b.iana-servers.net", "c.iana-servers.net"}, n, n, n, e, e, n, t}
	z[1834] = Zone{"e164.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc2916", w{"ns3.afrinic.net", "ns3.lacnic.net", "ns4.apnic.net", "pri.authdns.ripe.net", "rirns.arin.net"}, n, n, n, "whois.ripe.net", e, n, f}
	z[1835] = Zone{"eap-noob.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc9140", n, n, n, n, e, e, n, t}
	z[1836] = Zone{"home.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8375", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[1837] = Zone{"in-addr.arpa", &z[86], z[5102:5106], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc1034", w{"a.in-addr-servers.arpa", "b.in-addr-servers.arpa", "c.in-addr-servers.arpa", "d.in-addr-servers.arpa", "e.in-addr-servers.arpa", "f.in-addr-servers.arpa"}, n, n, n, e, e, n, f}
	z[1838] = Zone{"in-addr-servers.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc5855", w{"a.in-addr-servers.arpa", "b.in-addr-servers.arpa", "c.in-addr-servers.arpa", "d.in-addr-servers.arpa", "e.in-addr-servers.arpa", "f.in-addr-servers.arpa"}, n, n, n, e, e, n, t}
	z[1839] = Zone{"ip6.arpa", &z[86], z[5106:5107], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc3152", w{"a.ip6-servers.arpa", "b.ip6-servers.arpa", "c.ip6-servers.arpa", "d.ip6-servers.arpa", "e.ip6-servers.arpa", "f.ip6-servers.arpa"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1839() {
	z[1840] = Zone{"ip6-servers.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc5855", w{"a.ip6-servers.arpa", "b.ip6-servers.arpa", "c.ip6-servers.arpa", "d.ip6-servers.arpa", "e.ip6-servers.arpa", "f.ip6-servers.arpa"}, n, n, n, e, e, n, t}
	z[1841] = Zone{"ipv4only.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8880", w{"a.iana-servers.net", "b.iana-servers.net", "c.iana-servers.net", "ns.icann.org"}, n, n, n, e, e, n, t}
	z[1842] = Zone{"iris.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc4698", w{"a.iana-servers.net", "b.iana-servers.net", "c.iana-servers.net", "ns3.lacnic.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[1843] = Zone{"ns.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc9120", w{"a.ns.arpa", "b.ns.arpa", "c.ns.arpa", "d.ns.arpa", "e.ns.arpa", "f.ns.arpa", "g.ns.arpa", "h.ns.arpa", "i.ns.arpa", "k.ns.arpa", "l.ns.arpa", "m.ns.arpa"}, n, n, n, e, e, n, t}
	z[1844] = Zone{"resolver.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc9462", n, n, n, n, e, e, n, t}
	z[1845] = Zone{"uri.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc3405", w{"a.iana-servers.net", "b.iana-servers.net", "c.iana-servers.net", "ns3.lacnic.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[1846] = Zone{"urn.arpa", &z[86], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc3405", w{"a.iana-servers.net", "b.iana-servers.net", "c.iana-servers.net", "ns3.lacnic.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[1847] = Zone{"ac.at", &z[94], x, 0, e, e, w{"d.ns.at", "n.ns.at", "ns1.univie.ac.at", "ns2.univie.ac.at"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1847() {
	z[1848] = Zone{"co.at", &z[94], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1849] = Zone{"gv.at", &z[94], x, 0, e, e, w{"dns4.statistik.gv.at", "ns12.govix.at", "ns3.wien.at", "ns4.wien.gv.at", "ns7.univie.ac.at", "ns8.univie.ac.at"}, n, n, n, e, e, n, f}
	z[1850] = Zone{"or.at", &z[94], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1851] = Zone{"priv.at", &z[94], x, 0, e, e, w{"adns3.utanet.at", "mosthamster.nic.priv.at", "ns2.nextlayer.at", "ns5.univie.ac.at", "priv-ns.sil.at", "sec1.rcode0.net"}, n, n, n, "whois.nic.priv.at", e, n, t}
	z[1852] = Zone{"act.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1853] = Zone{"asn.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[1854] = Zone{"com.au", &z[97], x, 0, e, e, w{"q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[1855] = Zone{"conf.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1855() {
	z[1856] = Zone{"csiro.au", &z[97], x, 0, e, e, w{"ns1.aarnet.net.au", "ns1.csiro.au", "ns2.aarnet.net.au", "ns2.csiro.au", "ns3.aarnet.net.au", "ns3.csiro.au"}, n, n, n, e, e, n, f}
	z[1857] = Zone{"edu.au", &z[97], z[5107:5115], 0, e, "https://www.domainname.edu.au/", w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1858] = Zone{"gov.au", &z[97], x, 0, e, e, w{"q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1859] = Zone{"id.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1860] = Zone{"info.au", &z[97], x, 0x1000, e, "http://www.aunic.net/", n, n, n, n, e, e, n, f}
	z[1861] = Zone{"net.au", &z[97], x, 0, e, e, w{"q.au", "r.au", "s.au", "t.au"}, n, n, n, "eleanba.connect.com.au", e, n, t}
	z[1862] = Zone{"nsw.au", &z[97], x, 0x80, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, w{"AU-NSW"}, n, e, e, n, f}
	z[1863] = Zone{"nt.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1863() {
	z[1864] = Zone{"org.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[1865] = Zone{"oz.au", &z[97], x, 0, e, e, w{"munnari.oz.au", "ns.sw.oz.au", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[1866] = Zone{"qld.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1867] = Zone{"sa.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1868] = Zone{"tas.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1869] = Zone{"telememo.au", &z[97], x, 0, e, e, w{"ns-1046.awsdns-02.org", "ns-2022.awsdns-60.co.uk", "ns-467.awsdns-58.com", "ns-554.awsdns-05.net"}, n, n, n, e, e, n, f}
	z[1870] = Zone{"vic.au", &z[97], x, 0, e, "http://www.aucd.org.au/", w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
	z[1871] = Zone{"wa.au", &z[97], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1871() {
	z[1872] = Zone{"com.aw", &z[109], x, 0, e, e, w{"ns1.setarnet.aw", "ns2.setarnet.aw"}, n, n, n, e, e, n, f}
	z[1873] = Zone{"bilesuvar.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1874] = Zone{"biz.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1875] = Zone{"co.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1876] = Zone{"com.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1877] = Zone{"edu.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1878] = Zone{"ganja.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1879] = Zone{"gov.az", &z[114], x, 0, e, e, w{"ns1.gov.az", "ns2.gov.az", "ns3.gov.az", "ns4.gov.az"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1879() {
	z[1880] = Zone{"imishli.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1881] = Zone{"info.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1882] = Zone{"int.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1883] = Zone{"mil.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1884] = Zone{"name.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1885] = Zone{"net.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1886] = Zone{"org.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1887] = Zone{"pp.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1887() {
	z[1888] = Zone{"pro.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1889] = Zone{"samux.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1890] = Zone{"shamaxi.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1891] = Zone{"shusha.az", &z[114], x, 0, e, e, w{"ns1.gov.az", "ns2.gov.az"}, n, n, n, e, e, n, t}
	z[1892] = Zone{"sumgait.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1893] = Zone{"zaqatala.az", &z[114], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1894] = Zone{"co.ba", &z[116], x, 0, e, e, w{"ns1.co.ba", "ns2.co.ba", "ns3.co.ba"}, n, n, n, e, e, n, t}
	z[1895] = Zone{"com.ba", &z[116], x, 0, e, e, w{"dns01.hosting.bhtelecom.ba", "dns02.hosting.bhtelecom.ba"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1895() {
	z[1896] = Zone{"edu.ba", &z[116], x, 0, e, e, w{"lim.utic.net.ba", "ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1897] = Zone{"gov.ba", &z[116], x, 0, e, e, w{"lim.utic.net.ba", "ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1898] = Zone{"mil.ba", &z[116], x, 0, e, e, w{"bosna.utic.net.ba", "ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1899] = Zone{"net.ba", &z[116], x, 0, e, e, w{"ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1900] = Zone{"org.ba", &z[116], x, 0, e, e, w{"ns.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1901] = Zone{"rs.ba", &z[116], x, 0, e, e, w{"ns1.rs.ba", "ns2.rs.ba", "ns3.rs.ba"}, n, n, n, e, e, n, f}
	z[1902] = Zone{"unbi.ba", &z[116], x, 0, e, e, w{"linhost04.utic.net.ba", "sava.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1903] = Zone{"unmo.ba", &z[116], x, 0, e, e, w{"dns2019.unmo.ba", "pdc2019.unmo.ba", "web2019.unmo.ba"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1903() {
	z[1904] = Zone{"unsa.ba", &z[116], x, 0, e, e, w{"bosna.utic.net.ba", "sava.utic.net.ba", "una.utic.net.ba"}, n, n, n, e, e, n, f}
	z[1905] = Zone{"untz.ba", &z[116], x, 0, e, e, w{"gradina.untz.ba", "slatina.untz.ba"}, n, n, n, e, e, n, f}
	z[1906] = Zone{"unze.ba", &z[116], x, 0, e, e, w{"ns1.first-ns.de", "pi.unze.ba", "robotns2.second-ns.de", "robotns3.second-ns.com"}, n, n, n, e, e, n, f}
	z[1907] = Zone{"biz.bb", &z[134], x, 0, e, e, w{"bds-tlcm-ns-001.gov.bb", "bds-tlcm-ns-002.gov.bb", "ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1908] = Zone{"co.bb", &z[134], x, 0, e, e, w{"bds-tlcm-ns-001.gov.bb", "bds-tlcm-ns-002.gov.bb", "ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1909] = Zone{"com.bb", &z[134], x, 0, e, e, w{"ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1910] = Zone{"gov.bb", &z[134], x, 0, e, e, w{"ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1911] = Zone{"info.bb", &z[134], x, 0, e, e, w{"bds-tlcm-ns-001.gov.bb", "bds-tlcm-ns-002.gov.bb", "ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1911() {
	z[1912] = Zone{"net.bb", &z[134], x, 0, e, e, w{"ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1913] = Zone{"org.bb", &z[134], x, 0, e, e, w{"ns1.nic.bb", "ns2.nic.bb", "ns3.nic.bb", "ns4.nic.bb", "ns5.nic.bb", "ns6.nic.bb"}, n, n, n, e, e, n, f}
	z[1914] = Zone{"store.bb", &z[134], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1915] = Zone{"tv.bb", &z[134], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1916] = Zone{"ac.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1917] = Zone{"com.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1918] = Zone{"edu.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1919] = Zone{"gov.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
}

//go:noinline
func i1919() {
	z[1920] = Zone{"mil.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1921] = Zone{"net.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1922] = Zone{"org.bd", &z[141], x, 0, e, e, w{"bd-ns.anycast.pch.net", "dns.bd", "jamuna.btcl.net.bd", "surma.btcl.net.bd"}, n, n, n, e, e, n, f}
	z[1923] = Zone{"gov.bf", &z[152], x, 0, e, e, w{"ntoo.gov.bf", "oubri.gov.bf"}, n, n, n, e, e, n, f}
	z[1924] = Zone{"0.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1925] = Zone{"1.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1926] = Zone{"2.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1927] = Zone{"3.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1927() {
	z[1928] = Zone{"4.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1929] = Zone{"5.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1930] = Zone{"6.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1931] = Zone{"7.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1932] = Zone{"8.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1933] = Zone{"9.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1934] = Zone{"a.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1935] = Zone{"b.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i1935() {
	z[1936] = Zone{"c.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1937] = Zone{"d.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1938] = Zone{"e.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1939] = Zone{"f.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1940] = Zone{"g.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1941] = Zone{"h.bg", &z[153], x, 0, e, "https://www.register.bg/user/", n, n, n, n, e, e, n, f}
	z[1942] = Zone{"i.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1943] = Zone{"j.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1943() {
	z[1944] = Zone{"k.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1945] = Zone{"l.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1946] = Zone{"m.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1947] = Zone{"n.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1948] = Zone{"o.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1949] = Zone{"p.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1950] = Zone{"q.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1951] = Zone{"r.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1951() {
	z[1952] = Zone{"s.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1953] = Zone{"t.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1954] = Zone{"u.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1955] = Zone{"v.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1956] = Zone{"w.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1957] = Zone{"x.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1958] = Zone{"y.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1959] = Zone{"z.bg", &z[153], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1959() {
	z[1960] = Zone{"biz.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1961] = Zone{"cc.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1962] = Zone{"com.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1963] = Zone{"edu.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1964] = Zone{"gov.bh", &z[154], x, 0, e, e, w{"ns1.gov.bh", "ns2.gov.bh"}, n, n, n, e, e, n, f}
	z[1965] = Zone{"info.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1966] = Zone{"me.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, t}
	z[1967] = Zone{"med.bh", &z[154], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i1967() {
	z[1968] = Zone{"name.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, t}
	z[1969] = Zone{"net.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1970] = Zone{"org.bh", &z[154], x, 0, e, e, w{"a.bh.centralnic-dns.com", "b.bh.centralnic-dns.com", "c.bh.centralnic-dns.com", "d.bh.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[1971] = Zone{"co.bi", &z[156], x, 0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, f}
	z[1972] = Zone{"com.bi", &z[156], x, 0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, f}
	z[1973] = Zone{"edu.bi", &z[156], x, 0, e, e, w{"bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, f}
	z[1974] = Zone{"gov.bi", &z[156], x, 0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, t}
	z[1975] = Zone{"info.bi", &z[156], x, 0, e, e, w{"bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns-bi.ripe.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, t}
}

//go:noinline
func i1975() {
	z[1976] = Zone{"mo.bi", &z[156], x, 0, e, e, w{"ns1.hostinger.com", "ns2.hostinger.com", "ns3.hostinger.com", "ns4.hostinger.com"}, n, n, n, e, e, n, t}
	z[1977] = Zone{"net.bi", &z[156], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1978] = Zone{"or.bi", &z[156], x, 0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, f}
	z[1979] = Zone{"org.bi", &z[156], x, 0, e, e, w{"anyns.nic.bi", "bi.cctld.authdns.ripe.net", "ns-bi.afrinic.net", "ns.nic.bi", "ns1.nic.bi"}, n, n, n, e, e, n, f}
	z[1980] = Zone{"ote.bi", &z[156], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[1981] = Zone{"1x.biz", &z[163], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"1x.biz", "95.217.58.108"}, n, n, e, e, n, t}
	z[1982] = Zone{"auz.biz", &z[163], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com", "ns3.instradns.com"}, n, n, n, e, e, n, f}
	z[1983] = Zone{"asso.bj", &z[164], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i1983() {
	z[1984] = Zone{"barreau.bj", &z[164], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1985] = Zone{"com.bj", &z[164], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1986] = Zone{"edu.bj", &z[164], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[1987] = Zone{"gouv.bj", &z[164], x, 0, e, e, w{"beau.ns.cloudflare.com", "emma.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[1988] = Zone{"gov.bj", &z[164], x, 0, e, e, w{"ns1.intnet.bj"}, n, n, n, e, e, n, f}
	z[1989] = Zone{"mil.bj", &z[164], x, 0, e, e, w{"beau.ns.cloudflare.com", "emma.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[1990] = Zone{"art.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "art.blog"}, n, n, e, e, n, t}
	z[1991] = Zone{"business.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "business.blog"}, n, n, e, e, n, t}
}

//go:noinline
func i1991() {
	z[1992] = Zone{"car.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "car.blog"}, n, n, e, e, n, t}
	z[1993] = Zone{"code.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "code.blog"}, n, n, e, e, n, t}
	z[1994] = Zone{"family.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "family.blog"}, n, n, e, e, n, t}
	z[1995] = Zone{"fashion.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "fashion.blog"}, n, n, e, e, n, t}
	z[1996] = Zone{"finance.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "finance.blog"}, n, n, e, e, n, t}
	z[1997] = Zone{"fitness.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "fitness.blog"}, n, n, e, e, n, t}
	z[1998] = Zone{"food.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "food.blog"}, n, n, e, e, n, t}
	z[1999] = Zone{"game.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "game.blog"}, n, n, e, e, n, t}
}

//go:noinline
func i1999() {
	z[2000] = Zone{"health.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "health.blog"}, n, n, e, e, n, t}
	z[2001] = Zone{"home.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "home.blog"}, n, n, e, e, n, t}
	z[2002] = Zone{"law.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "law.blog"}, n, n, e, e, n, t}
	z[2003] = Zone{"movie.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "movie.blog"}, n, n, e, e, n, t}
	z[2004] = Zone{"music.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "music.blog"}, n, n, e, e, n, t}
	z[2005] = Zone{"news.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "news.blog"}, n, n, e, e, n, t}
	z[2006] = Zone{"photo.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "photo.blog"}, n, n, e, e, n, t}
	z[2007] = Zone{"poetry.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "poetry.blog"}, n, n, e, e, n, t}
}

//go:noinline
func i2007() {
	z[2008] = Zone{"politics.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "politics.blog"}, n, n, e, e, n, t}
	z[2009] = Zone{"school.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "school.blog"}, n, n, e, e, n, t}
	z[2010] = Zone{"science.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "science.blog"}, n, n, e, e, n, t}
	z[2011] = Zone{"sport.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "sport.blog"}, n, n, e, e, n, t}
	z[2012] = Zone{"tech.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "tech.blog"}, n, n, e, e, n, t}
	z[2013] = Zone{"travel.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "travel.blog"}, n, n, e, e, n, t}
	z[2014] = Zone{"video.blog", &z[169], x, 0, e, e, w{"ns1.automattic.com", "ns2.automattic.com", "ns3.automattic.com", "ns4.automattic.com"}, w{"192.0.78.30", "192.0.78.31", "video.blog"}, n, n, e, e, n, t}
	z[2015] = Zone{"water.blog", &z[169], x, 0, e, e, w{"ns1.wordpress.com", "ns2.wordpress.com", "ns3.wordpress.com"}, w{"192.0.78.24", "192.0.78.25", "water.blog"}, n, n, e, e, n, t}
}

//go:noinline
func i2015() {
	z[2016] = Zone{"com.bm", &z[173], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2017] = Zone{"edu.bm", &z[173], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2018] = Zone{"gov.bm", &z[173], x, 0, e, e, w{"dns1.gov.bm", "dns2.gov.bm", "dns3.gov.bm"}, n, n, n, e, e, n, f}
	z[2019] = Zone{"net.bm", &z[173], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2020] = Zone{"org.bm", &z[173], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2021] = Zone{"com.bn", &z[176], x, 0, e, e, w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2022] = Zone{"edu.bn", &z[176], x, 0, e, e, w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2023] = Zone{"gov.bn", &z[176], x, 0, e, e, w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2023() {
	z[2024] = Zone{"net.bn", &z[176], x, 0, e, e, w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2025] = Zone{"org.bn", &z[176], x, 0, e, e, w{"bn-ns.anycast.pch.net", "ns1.bnnic.bn", "ns2.bnnic.bn", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2026] = Zone{"com.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2027] = Zone{"edu.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2028] = Zone{"gob.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2029] = Zone{"gov.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2030] = Zone{"int.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2031] = Zone{"mil.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2031() {
	z[2032] = Zone{"net.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2033] = Zone{"org.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2034] = Zone{"tv.bo", &z[179], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2035] = Zone{"abc.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2036] = Zone{"adm.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2037] = Zone{"adv.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2038] = Zone{"agr.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2039] = Zone{"am.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2039() {
	z[2040] = Zone{"aparecida.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2041] = Zone{"app.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2042] = Zone{"arq.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2043] = Zone{"art.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2044] = Zone{"ato.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2045] = Zone{"b.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2046] = Zone{"belem.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2047] = Zone{"bhz.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2047() {
	z[2048] = Zone{"bib.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2049] = Zone{"bio.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2050] = Zone{"blog.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2051] = Zone{"bmd.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2052] = Zone{"boavista.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2053] = Zone{"bsb.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2054] = Zone{"campinas.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2055] = Zone{"caxias.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2055() {
	z[2056] = Zone{"cim.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2057] = Zone{"cng.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2058] = Zone{"cnt.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2059] = Zone{"com.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2060] = Zone{"coop.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2061] = Zone{"coz.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2062] = Zone{"curitiba.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2063] = Zone{"des.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2063() {
	z[2064] = Zone{"det.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2065] = Zone{"dev.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2066] = Zone{"ecn.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2067] = Zone{"eco.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2068] = Zone{"edu.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2069] = Zone{"emp.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2070] = Zone{"enf.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2071] = Zone{"eng.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2071() {
	z[2072] = Zone{"esp.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2073] = Zone{"etc.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2074] = Zone{"eti.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2075] = Zone{"far.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2076] = Zone{"flog.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2077] = Zone{"floripa.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2078] = Zone{"fm.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2079] = Zone{"fnd.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2079() {
	z[2080] = Zone{"fortal.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2081] = Zone{"fot.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2082] = Zone{"foz.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2083] = Zone{"fst.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2084] = Zone{"g12.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2085] = Zone{"geo.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2086] = Zone{"ggf.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2087] = Zone{"gov.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2087() {
	z[2088] = Zone{"gru.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2089] = Zone{"imb.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2090] = Zone{"ind.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2091] = Zone{"inf.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2092] = Zone{"jampa.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2093] = Zone{"jor.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2094] = Zone{"jus.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2095] = Zone{"lel.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2095() {
	z[2096] = Zone{"log.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2097] = Zone{"macapa.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2098] = Zone{"maceio.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2099] = Zone{"manaus.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2100] = Zone{"mat.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2101] = Zone{"med.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2102] = Zone{"mil.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2103] = Zone{"mus.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2103() {
	z[2104] = Zone{"natal.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2105] = Zone{"net.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2106] = Zone{"nom.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2107] = Zone{"not.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2108] = Zone{"ntr.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2109] = Zone{"odo.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2110] = Zone{"org.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2111] = Zone{"palmas.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2111() {
	z[2112] = Zone{"poa.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2113] = Zone{"ppg.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2114] = Zone{"pro.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2115] = Zone{"psc.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2116] = Zone{"psi.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2117] = Zone{"qsl.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2118] = Zone{"radio.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2119] = Zone{"rec.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2119() {
	z[2120] = Zone{"recife.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2121] = Zone{"rep.br", &z[195], z[5115:5116], 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2122] = Zone{"rio.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2123] = Zone{"salvador.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2124] = Zone{"sec3.br", &z[195], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2125] = Zone{"seg.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2126] = Zone{"sjc.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2127] = Zone{"slg.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2127() {
	z[2128] = Zone{"srv.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2129] = Zone{"taxi.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2130] = Zone{"tec.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2131] = Zone{"teo.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2132] = Zone{"tmp.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2133] = Zone{"trd.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2134] = Zone{"tur.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2135] = Zone{"tv.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2135() {
	z[2136] = Zone{"vet.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2137] = Zone{"vix.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2138] = Zone{"vlog.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2139] = Zone{"wiki.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2140] = Zone{"zlg.br", &z[195], x, 0, e, e, w{"a.dns.br", "b.dns.br", "c.dns.br", "d.dns.br", "e.dns.br", "f.dns.br"}, n, n, n, e, e, n, t}
	z[2141] = Zone{"com.bs", &z[202], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2142] = Zone{"edu.bs", &z[202], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2143] = Zone{"gov.bs", &z[202], x, 0, e, e, w{"ib-ext1.gov.bs", "ib-ext2.gov.bs", "ibgb-ext1.gov.bs", "ibgb-ext2.gov.bs"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2143() {
	z[2144] = Zone{"net.bs", &z[202], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2145] = Zone{"org.bs", &z[202], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2146] = Zone{"we.bs", &z[202], x, 0, e, e, w{"ns1.nic.we.bs", "ns2.nic.we.bs"}, n, n, n, e, e, n, t}
	z[2147] = Zone{"com.bt", &z[203], x, 0, e, e, w{"ns1.druknet.bt", "ns2.druknet.bt", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2148] = Zone{"edu.bt", &z[203], x, 0, e, e, w{"ns1.druknet.bt", "ns2.druknet.bt", "phloem.uoregon.edu"}, n, n, n, e, e, n, t}
	z[2149] = Zone{"gov.bt", &z[203], x, 0, e, e, w{"ns1.druknet.bt", "ns2.druknet.bt"}, n, n, n, e, e, n, t}
	z[2150] = Zone{"net.bt", &z[203], x, 0, e, e, w{"ns1.druknet.bt", "ns2.druknet.bt"}, n, n, n, e, e, n, t}
	z[2151] = Zone{"org.bt", &z[203], x, 0, e, e, w{"ns1.druknet.bt", "ns2.druknet.bt", "phloem.uoregon.edu"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2151() {
	z[2152] = Zone{"ac.bw", &z[213], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2153] = Zone{"co.bw", &z[213], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2154] = Zone{"net.bw", &z[213], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2155] = Zone{"org.bw", &z[213], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2156] = Zone{"com.by", &z[215], x, 0, e, e, w{"a1.domain.by", "a2.domain.by"}, n, n, n, e, e, n, f}
	z[2157] = Zone{"gov.by", &z[215], x, 0, e, e, w{"ns1.g-cloud.by", "ns2.g-cloud.by", "ns3.g-cloud.by", "ns4.becloudby.tech"}, n, n, n, e, e, n, f}
	z[2158] = Zone{"mil.by", &z[215], x, 0, e, e, w{"ns1.g-cloud.by", "ns2.g-cloud.by", "ns3.g-cloud.by", "ns4.becloudby.tech"}, n, n, n, e, e, n, f}
	z[2159] = Zone{"minsk.by", &z[215], x, 0, e, e, w{"a1.domain.by", "a2.domain.by"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2159() {
	z[2160] = Zone{"net.by", &z[215], x, 0, e, e, w{"a1.domain.by", "a2.domain.by"}, n, n, n, e, e, n, f}
	z[2161] = Zone{"co.bz", &z[216], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2162] = Zone{"com.bz", &z[216], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2163] = Zone{"edu.bz", &z[216], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2164] = Zone{"gov.bz", &z[216], x, 0, e, e, w{"ns1-03.azure-dns.com", "ns2-03.azure-dns.net", "ns3-03.azure-dns.org", "ns4-03.azure-dns.info"}, n, n, n, e, e, n, f}
	z[2165] = Zone{"net.bz", &z[216], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2166] = Zone{"org.bz", &z[216], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2167] = Zone{"za.bz", &z[216], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net", "ns5.centralnic.net"}, n, n, n, "whois.centralnic.com", e, n, t}
}

//go:noinline
func i2167() {
	z[2168] = Zone{"ab.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2169] = Zone{"bc.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2170] = Zone{"co.ca", &z[218], x, 0x200, e, e, w{"primary.relay.co.ca", "quaternary.relay.co.ca", "quinary.relay.co.ca", "secondary.relay.co.ca", "senary.relay.co.ca", "tertiary.relay.co.ca"}, n, n, n, "whois.co.ca", e, n, t}
	z[2171] = Zone{"gc.ca", &z[218], x, 0, e, e, w{"dns2.nrc.ca", "ns1.d-zone.ca", "ns2.d-zone.ca"}, n, n, n, e, e, n, f}
	z[2172] = Zone{"mb.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2173] = Zone{"nb.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2174] = Zone{"nf.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2175] = Zone{"nl.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2175() {
	z[2176] = Zone{"ns.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2177] = Zone{"nt.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2178] = Zone{"nu.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2179] = Zone{"on.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2180] = Zone{"pe.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2181] = Zone{"qc.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2182] = Zone{"sk.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2183] = Zone{"yk.ca", &z[218], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2183() {
	z[2184] = Zone{"com.cc", &z[256], x, 0, e, e, w{"ns-1287.awsdns-32.org", "ns-1613.awsdns-09.co.uk", "ns-472.awsdns-59.com", "ns-680.awsdns-21.net"}, n, n, n, e, e, n, t}
	z[2185] = Zone{"edu.cc", &z[256], x, 0, e, e, w{"ns1.netdns.co.nz", "ns2.netdns.co.nz"}, n, n, n, e, e, n, t}
	z[2186] = Zone{"net.cc", &z[256], x, 0, e, e, w{"ns-1028.awsdns-00.org", "ns-1980.awsdns-55.co.uk", "ns-218.awsdns-27.com", "ns-779.awsdns-33.net"}, n, n, n, e, e, n, t}
	z[2187] = Zone{"org.cc", &z[256], x, 0, e, e, w{"ns.cocca.fr", "ns.coccaregistry.org"}, n, n, n, e, e, n, t}
	z[2188] = Zone{"ac.cd", &z[257], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2189] = Zone{"com.cd", &z[257], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2190] = Zone{"edu.cd", &z[257], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2191] = Zone{"gouv.cd", &z[257], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2191() {
	z[2192] = Zone{"net.cd", &z[257], x, 0, e, e, w{"ns-root-21.scpt-network.net", "ns-root-22.scpt-network.net"}, n, n, n, e, e, n, f}
	z[2193] = Zone{"org.cd", &z[257], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2194] = Zone{"ac.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2195] = Zone{"aeroport.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2196] = Zone{"asso.ci", &z[285], x, 0, e, "https://www.nic.ci/", w{"a.dnspod.com", "b.dnspod.com", "c.dnspod.com"}, n, n, n, e, e, n, t}
	z[2197] = Zone{"assoc.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2198] = Zone{"co.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2199] = Zone{"com.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2199() {
	z[2200] = Zone{"ed.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2201] = Zone{"edu.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2202] = Zone{"go.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2203] = Zone{"gov.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2204] = Zone{"in.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2205] = Zone{"int.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2206] = Zone{"net.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2207] = Zone{"nom.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2207() {
	z[2208] = Zone{"or.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2209] = Zone{"org.ci", &z[285], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2210] = Zone{"presse.ci", &z[285], x, 0, e, e, w{"a.dnspod.com", "b.dnspod.com", "c.dnspod.com"}, n, n, n, e, e, n, f}
	z[2211] = Zone{"biz.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2212] = Zone{"co.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2213] = Zone{"edu.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2214] = Zone{"gen.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2215] = Zone{"gov.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2215() {
	z[2216] = Zone{"info.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2217] = Zone{"net.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2218] = Zone{"org.ck", &z[295], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2219] = Zone{"co.cm", &z[306], x, 0, e, e, w{"a.cm.dyntld.net", "b.cm.dyntld.net", "ns1.nic.cm", "ns2.nic.cm"}, n, n, n, e, e, n, f}
	z[2220] = Zone{"com.cm", &z[306], x, 0, e, e, w{"a.cm.dyntld.net", "b.cm.dyntld.net", "ns1.nic.cm", "ns2.nic.cm"}, n, n, n, e, e, n, f}
	z[2221] = Zone{"edu.cm", &z[306], x, 0, e, e, w{"ns1.nic.cm", "ns2.nic.cm"}, w{"85.25.140.105"}, n, n, e, e, n, t}
	z[2222] = Zone{"net.cm", &z[306], x, 0, e, e, w{"a.cm.dyntld.net", "b.cm.dyntld.net", "ns1.nic.cm", "ns2.nic.cm"}, n, n, n, e, e, n, f}
	z[2223] = Zone{"ac.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2223() {
	z[2224] = Zone{"ah.cn", &z[307], x, 0, e, "https://www.belizenic.bz/", w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2225] = Zone{"bj.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2226] = Zone{"com.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2227] = Zone{"cq.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2228] = Zone{"edu.cn", &z[307], x, 0, e, e, w{"deneb.dfn.de", "dns.edu.cn", "dns2.edu.cn", "ns2.cernet.net", "ns2.cuhk.hk", "ns3.cernet.net", "ns4.cernet.net", "ns5.cernet.net"}, n, n, n, "whois.edu.cn", e, n, f}
	z[2229] = Zone{"fj.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2230] = Zone{"gd.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2231] = Zone{"gov.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2231() {
	z[2232] = Zone{"gs.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2233] = Zone{"gx.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2234] = Zone{"gz.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2235] = Zone{"ha.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2236] = Zone{"hb.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2237] = Zone{"he.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2238] = Zone{"hi.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2239] = Zone{"hk.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2239() {
	z[2240] = Zone{"hl.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2241] = Zone{"hn.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2242] = Zone{"jl.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2243] = Zone{"js.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2244] = Zone{"jx.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2245] = Zone{"keyword.cn", &z[307], z[5116:5117], 0, e, e, w{"ns11.xincache.com", "ns12.xincache.com"}, n, n, n, e, e, n, t}
	z[2246] = Zone{"ln.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2247] = Zone{"mil.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "d.dns.cn", "e.dns.cn", "ns3.mil.cn", "ns4.mil.cn"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2247() {
	z[2248] = Zone{"mo.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2249] = Zone{"net.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2250] = Zone{"nm.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2251] = Zone{"nx.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2252] = Zone{"org.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2253] = Zone{"qh.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2254] = Zone{"sc.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2255] = Zone{"sd.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2255() {
	z[2256] = Zone{"sh.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2257] = Zone{"sn.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2258] = Zone{"sx.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2259] = Zone{"tj.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2260] = Zone{"tw.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2261] = Zone{"xj.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2262] = Zone{"xz.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2263] = Zone{"yn.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2263() {
	z[2264] = Zone{"zj.cn", &z[307], x, 0, e, e, w{"a.dns.cn", "b.dns.cn", "c.dns.cn", "cns.cernet.net", "d.dns.cn", "e.dns.cn"}, n, n, n, e, e, n, f}
	z[2265] = Zone{"com.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2266] = Zone{"edu.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2267] = Zone{"gov.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2268] = Zone{"mil.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2269] = Zone{"net.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2270] = Zone{"nom.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2271] = Zone{"org.co", &z[308], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2271() {
	z[2272] = Zone{"4u.com", &z[314], x, 0x200, e, e, w{"ns59.worldnic.com", "ns60.worldnic.com"}, w{"208.91.197.27"}, n, n, e, e, n, t}
	z[2273] = Zone{"ae.com", &z[314], x, 0x200, e, e, w{"asia3.akam.net", "eur2.akam.net", "eur5.akam.net", "ns1-161.akam.net", "ns1-163.akam.net", "usc3.akam.net", "use4.akam.net", "usw1.akam.net"}, n, n, n, e, e, n, t}
	z[2274] = Zone{"africa.com", &z[314], x, 0x200, e, e, w{"ns2.dns.business", "ns3.dns.business", "ns4.dns.business"}, n, n, n, "whois.centralnic.com", e, n, f}
	z[2275] = Zone{"ar.com", &z[314], x, 0x200, e, e, w{"pdns07.domaincontrol.com", "pdns08.domaincontrol.com"}, n, n, n, "whois.centralnic.com", e, n, t}
	z[2276] = Zone{"asia.com", &z[314], x, 0x200, e, e, w{"pdns1.ultradns.net", "pdns2.ultradns.net", "pdns3.ultradns.org", "pdns4.ultradns.org", "pdns5.ultradns.info", "pdns6.ultradns.co.uk"}, w{"72.55.150.59"}, n, n, e, e, n, f}
	z[2277] = Zone{"au.com", &z[314], x, 0x200, e, e, w{"dns103.kddi.ne.jp", "dns104.kddi.ne.jp", "dnsa01.kddi.ne.jp", "dnsa02.kddi.ne.jp"}, n, n, n, e, e, n, f}
	z[2278] = Zone{"br.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/br.com/"}, t}
	z[2279] = Zone{"cn.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/cn.com/"}, t}
}

//go:noinline
func i2279() {
	z[2280] = Zone{"co.com", &z[314], x, 0x200, e, "https://registry.co.com/", w{"ns1.nic.co.com", "ns2.nic.co.com", "ns3.nic.co.com", "ns4.nic.co.com"}, w{"173.192.115.17"}, n, n, "whois.centralnic.net", e, w{"https://rdap.centralnic.com/co.com/"}, t}
	z[2281] = Zone{"de.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/de.com/"}, t}
	z[2282] = Zone{"eu.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/eu.com/"}, t}
	z[2283] = Zone{"example.com", &z[314], x, 0x2108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"a.iana-servers.net", "b.iana-servers.net"}, n, n, n, e, e, n, t}
	z[2284] = Zone{"gb.com", &z[314], x, 0x200, e, e, w{"ns-cloud-c1.googledomains.com", "ns-cloud-c2.googledomains.com", "ns-cloud-c3.googledomains.com", "ns-cloud-c4.googledomains.com"}, n, n, n, "whois.centralnic.com", e, n, t}
	z[2285] = Zone{"gr.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"72.34.38.64"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/gr.com/"}, f}
	z[2286] = Zone{"hk.com", &z[314], x, 0x200, e, e, w{"a.udrtld.net", "b.udrtld.net", "c.udrtld.net", "d.udrtld.net", "e.udrtld.net", "f.udrtld.net", "g.udrtld.net", "h.udrtld.net", "j.udrtld.net", "s.udrtld.net"}, n, n, n, "whois.registry.hk.com", e, n, f}
	z[2287] = Zone{"hu.com", &z[314], x, 0x200, e, e, w{"ns3.dns.com", "ns4.dns.com"}, n, n, n, "whois.centralnic.com", e, n, t}
}

//go:noinline
func i2287() {
	z[2288] = Zone{"it.com", &z[314], x, 0, e, e, w{"ns1.it.com", "ns2.it.com", "ns3.it.com"}, n, n, n, e, e, n, t}
	z[2289] = Zone{"jpn.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/jpn.com/"}, t}
	z[2290] = Zone{"kr.com", &z[314], x, 0x200, e, e, w{"sl1.sedo.com", "sl2.sedo.com"}, w{"91.195.241.232"}, n, n, "whois.centralnic.com", e, n, t}
	z[2291] = Zone{"mex.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183", "mex.com"}, n, n, e, e, w{"https://rdap.centralnic.com/mex.com/"}, f}
	z[2292] = Zone{"no.com", &z[314], x, 0x200, e, e, w{"pdns07.domaincontrol.com", "pdns08.domaincontrol.com"}, n, n, n, "whois.centralnic.com", e, n, t}
	z[2293] = Zone{"nv.com", &z[314], x, 0x200, e, e, w{"a.root-servers.nv.com", "b.root-servers.nv.com", "c.root-servers.nv.com", "d.root-servers.nv.com"}, w{"205.159.223.111", "hosting.gi.net"}, n, n, e, e, n, f}
	z[2294] = Zone{"pty-ltd.com", &z[314], x, 0x200, e, e, w{"ns1.instradns.com", "ns2.instradns.com", "ns3.instradns.com"}, n, n, n, e, e, n, f}
	z[2295] = Zone{"qb.com", &z[314], x, 0x200, e, e, w{"ns1.sedoparking.com", "ns2.sedoparking.com"}, w{"64.190.63.111"}, n, n, e, e, n, f}
}

//go:noinline
func i2295() {
	z[2296] = Zone{"qc.com", &z[314], x, 0x200, e, e, w{"pdns07.domaincontrol.com", "pdns08.domaincontrol.com"}, n, n, n, "whois.centralnic.com", e, n, t}
	z[2297] = Zone{"ru.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/ru.com/"}, t}
	z[2298] = Zone{"sa.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/sa.com/"}, t}
	z[2299] = Zone{"uk.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/uk.com/"}, t}
	z[2300] = Zone{"ukots.com", &z[314], x, 0, e, e, w{"blerghmaheuikcnho.hydrapiglephant.com", "blerghoqannsvwfvm.hydrapiglephant.com"}, n, n, n, e, e, n, t}
	z[2301] = Zone{"us.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"45.56.79.23"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/us.com/"}, t}
	z[2302] = Zone{"uy.com", &z[314], x, 0x200, e, e, w{"ns49.domaincontrol.com", "ns50.domaincontrol.com"}, n, n, n, "whois.centralnic.com", e, n, t}
	z[2303] = Zone{"za.com", &z[314], x, 0x200, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/za.com/"}, t}
}

//go:noinline
func i2303() {
	z[2304] = Zone{"de.cool", &z[330], x, 0, e, e, w{"jake.ns.cloudflare.com", "jill.ns.cloudflare.com"}, w{"2a00:f48:2000:affe::50", "91.216.248.20", "91.216.248.21", "91.216.248.22"}, n, n, e, e, n, t}
	z[2305] = Zone{"ac.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
	z[2306] = Zone{"co.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
	z[2307] = Zone{"ed.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, t}
	z[2308] = Zone{"eu.cr", &z[339], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "eu.cr"}, n, n, e, e, n, t}
	z[2309] = Zone{"fi.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
	z[2310] = Zone{"go.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
	z[2311] = Zone{"or.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2311() {
	z[2312] = Zone{"sa.cr", &z[339], x, 0, e, e, w{"a.lactld.org", "ca1.nic.cr", "ca2.nic.cr", "de.nic.cr", "dns.nic.cr", "p.nic.cr"}, n, n, n, e, e, n, f}
	z[2313] = Zone{"co.cu", &z[349], x, 0, e, e, w{"ns3.etecsa.net", "ns3.etecsa.net.cu", "ns4.etecsa.net", "ns4.etecsa.net.cu", "ns5.etecsa.net"}, n, n, n, e, e, n, f}
	z[2314] = Zone{"com.cu", &z[349], x, 0, e, e, w{"ns.canada.cuba.cu", "ns.ceniai.net.cu", "ns.citmatel.com.cu"}, n, n, n, e, e, n, f}
	z[2315] = Zone{"cuba.cu", &z[349], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2316] = Zone{"cyt.cu", &z[349], x, 0, e, e, w{"mail1.tur.cu", "mail2.tur.cu", "ns1.tur.cu"}, n, n, n, e, e, n, f}
	z[2317] = Zone{"edu.cu", &z[349], x, 0, e, e, w{"cu.cctld.authdns.ripe.net", "ns.ceniai.net.cu", "ns.citmatel.com.cu", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2318] = Zone{"get.cu", &z[349], x, 0, e, e, w{"mail1.tur.cu", "mail2.tur.cu", "ns1.tur.cu"}, n, n, n, e, e, n, f}
	z[2319] = Zone{"gov.cu", &z[349], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2319() {
	z[2320] = Zone{"inf.cu", &z[349], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2321] = Zone{"net.cu", &z[349], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2322] = Zone{"org.cu", &z[349], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2323] = Zone{"tur.cu", &z[349], x, 0, e, e, w{"ns.ceniai.net.cu", "ns1.g4its.com", "ns1.tur.cu", "ns2.g4its.com", "ns2.tur.cu", "ns3.tur.cu", "ns4.tur.cu", "ns5.tur.cu"}, n, n, n, e, e, n, f}
	z[2324] = Zone{"com.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2325] = Zone{"edu.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2326] = Zone{"gov.cv", &z[351], x, 0, e, e, w{"ns1.gov.cv", "ns2.gov.cv", "ns3.gov.cv", "ns4.gov.cv"}, n, n, n, e, e, n, f}
	z[2327] = Zone{"int.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2327() {
	z[2328] = Zone{"net.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2329] = Zone{"nome.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2330] = Zone{"org.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2331] = Zone{"publ.cv", &z[351], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2332] = Zone{"com.cw", &z[352], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2333] = Zone{"net.cw", &z[352], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2334] = Zone{"com.cx", &z[353], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2335] = Zone{"edu.cx", &z[353], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2335() {
	z[2336] = Zone{"gov.cx", &z[353], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2337] = Zone{"net.cx", &z[353], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2338] = Zone{"org.cx", &z[353], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2339] = Zone{"ac.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2340] = Zone{"biz.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2341] = Zone{"com.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2342] = Zone{"ekloges.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2343] = Zone{"gov.cy", &z[354], x, 0, e, e, w{"dns.cit.cornell.edu", "ns01.gov.cy", "ns02.gov.cy", "sns0.grnet.gr", "sns1.grnet.gr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2343() {
	z[2344] = Zone{"ltd.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2345] = Zone{"name.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2346] = Zone{"net.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2347] = Zone{"org.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2348] = Zone{"parliament.cy", &z[354], x, 0, e, e, w{"armfazh.ns.cloudflare.com", "mona.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[2349] = Zone{"press.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2350] = Zone{"pro.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[2351] = Zone{"tm.cy", &z[354], x, 0, e, e, w{"cy-ns.anycast.pch.net", "cynic4.dns.cy", "cynic6.dns.cy", "estia.ics.forth.gr", "ns31.rcode0.net", "ns4.apnic.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2351() {
	z[2352] = Zone{"co.cz", &z[357], x, 0, e, e, w{"ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2353] = Zone{"1x.de", &z[368], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, n, n, n, e, e, n, t}
	z[2354] = Zone{"co.de", &z[368], x, 0, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "co.de"}, n, n, "whois.co.de", e, n, t}
	z[2355] = Zone{"com.de", &z[368], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183", "com.de"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/com.de/"}, f}
	z[2356] = Zone{"biz.dk", &z[397], x, 0, e, e, w{"dns2000.euro-isp.net", "dns2001.euro-isp.net", "mail.digitalmarketing.net", "ns.digitalmarketing.net", "ns1.dk.org", "post.digitalmarketing.net", "root-dns.euro-isp.net"}, n, n, n, e, e, n, f}
	z[2357] = Zone{"co.dk", &z[397], x, 0, e, e, w{"dns2000.euro-isp.net", "dns2001.euro-isp.net", "mail.digitalmarketing.net", "ns.digitalmarketing.net", "ns1.dk.org", "post.digitalmarketing.net", "root-dns.euro-isp.net"}, n, n, n, e, e, n, f}
	z[2358] = Zone{"co.dm", &z[398], x, 0, e, e, w{"ns.blacknightsolutions.com", "ns1.uniregistry.net", "ns2.blacknightsolutions.com", "ns2.nic.dm", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns34.cdns.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[2359] = Zone{"com.dm", &z[398], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2359() {
	z[2360] = Zone{"edu.dm", &z[398], x, 0, e, e, w{"ns.blacknightsolutions.com", "ns1.uniregistry.net", "ns2.blacknightsolutions.com", "ns2.nic.dm", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns34.cdns.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[2361] = Zone{"gov.dm", &z[398], x, 0, e, e, w{"ns.blacknightsolutions.com", "ns1.uniregistry.net", "ns2.blacknightsolutions.com", "ns2.nic.dm", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns34.cdns.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[2362] = Zone{"net.dm", &z[398], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2363] = Zone{"org.dm", &z[398], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2364] = Zone{"art.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2365] = Zone{"com.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2366] = Zone{"edu.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2367] = Zone{"gob.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2367() {
	z[2368] = Zone{"gov.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2369] = Zone{"mil.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2370] = Zone{"net.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2371] = Zone{"org.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2372] = Zone{"sid.do", &z[401], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2373] = Zone{"sld.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2374] = Zone{"web.do", &z[401], x, 0, e, e, w{"a.lactld.org", "ns.nic.do", "ns1.nic.do", "ns2.nic.do", "ns4.nic.do", "ns5.nic.do", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2375] = Zone{"art.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2375() {
	z[2376] = Zone{"asso.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2377] = Zone{"com.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2378] = Zone{"edu.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2379] = Zone{"gov.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2380] = Zone{"net.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2381] = Zone{"org.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2382] = Zone{"pol.dz", &z[425], x, 0, e, e, w{"ns1.nic.dz", "ns2.nic.dz", "ns3.nic.dz"}, n, n, n, e, e, n, f}
	z[2383] = Zone{"com.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2383() {
	z[2384] = Zone{"edu.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2385] = Zone{"fin.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec"}, n, n, n, e, e, n, f}
	z[2386] = Zone{"gob.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, t}
	z[2387] = Zone{"gov.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2388] = Zone{"info.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2389] = Zone{"med.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec"}, n, n, n, e, e, n, f}
	z[2390] = Zone{"mil.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec"}, n, n, n, e, e, n, f}
	z[2391] = Zone{"net.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2391() {
	z[2392] = Zone{"org.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2393] = Zone{"pro.ec", &z[428], x, 0, e, e, w{"a.lactld.org", "n1.nic.ec", "n2.nic.ec", "n3.dns.ec", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2394] = Zone{"co.ee", &z[434], x, 0, e, e, w{"ns.elkdata.ee", "ns2.elkdata.ee"}, n, n, n, e, e, n, f}
	z[2395] = Zone{"com.ee", &z[434], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2396] = Zone{"edu.ee", &z[434], x, 0, e, e, w{"ns.eenet.ee", "ns2.eenet.ee"}, w{"127.0.0.1"}, n, n, e, e, n, f}
	z[2397] = Zone{"fie.ee", &z[434], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2398] = Zone{"gov.ee", &z[434], x, 0, e, e, w{"ans2.aso.ee", "ns.gov.ee", "sec1.rcode0.net", "sec2.rcode0.net"}, n, n, n, e, e, n, f}
	z[2399] = Zone{"hiiumaa.ee", &z[434], x, 0, e, e, w{"ns.zone.eu", "ns2.zone.ee", "ns3.zonedata.net"}, w{"217.146.69.46"}, n, n, e, e, n, f}
}

//go:noinline
func i2399() {
	z[2400] = Zone{"kul.ee", &z[434], x, 0, e, e, w{"ans2.aso.ee", "dns.rmit.ee"}, n, n, n, e, e, n, f}
	z[2401] = Zone{"med.ee", &z[434], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2402] = Zone{"org.ee", &z[434], x, 0, e, e, w{"ns.eenet.ee", "ns2.eenet.ee"}, w{"127.0.0.1"}, n, n, e, e, n, f}
	z[2403] = Zone{"parnu.ee", &z[434], x, 0, e, e, w{"ns.eenet.ee", "ns2.eenet.ee"}, n, n, n, e, e, n, f}
	z[2404] = Zone{"parnumaa.ee", &z[434], x, 0, e, e, w{"ns.elkdata.ee", "ns2.elkdata.ee", "ns3.elkdata.net"}, n, n, n, e, e, n, f}
	z[2405] = Zone{"polvamaa.ee", &z[434], x, 0, e, e, w{"ns.zone.eu", "ns2.zone.ee", "ns3.zonedata.net"}, n, n, n, e, e, n, f}
	z[2406] = Zone{"pri.ee", &z[434], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2407] = Zone{"tartu.ee", &z[434], x, 0, e, e, w{"ns.eenet.ee", "ns2.eenet.ee"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2407() {
	z[2408] = Zone{"com.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2409] = Zone{"edu.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2410] = Zone{"eun.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "ns.idsc.gov.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2411] = Zone{"gov.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "ns.idsc.gov.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2412] = Zone{"info.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2413] = Zone{"mil.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2414] = Zone{"name.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2415] = Zone{"net.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2415() {
	z[2416] = Zone{"org.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com", "www.frcu.eun.eg"}, n, n, n, e, e, n, f}
	z[2417] = Zone{"sci.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[2418] = Zone{"sport.eg", &z[435], x, 0, e, e, w{"frcu.eun.eg", "rip.psg.com"}, n, n, n, e, e, n, t}
	z[2419] = Zone{"tv.eg", &z[435], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2420] = Zone{"com.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, t}
	z[2421] = Zone{"edu.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[2422] = Zone{"gov.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[2423] = Zone{"ind.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2423() {
	z[2424] = Zone{"mil.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[2425] = Zone{"net.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[2426] = Zone{"org.er", &z[446], x, 0, e, e, w{"er.cctld.authdns.ripe.net", "sawanew.noc.net.er", "zaranew.noc.net.er"}, n, n, n, e, e, n, f}
	z[2427] = Zone{"com.es", &z[449], x, 0, e, e, w{"c.nic.es", "g.nic.es", "h.nic.es", "n3ns.nic.es"}, n, n, n, e, e, n, f}
	z[2428] = Zone{"edu.es", &z[449], x, 0, e, e, w{"c.nic.es", "g.nic.es", "h.nic.es", "n3ns.nic.es"}, n, n, n, e, e, n, f}
	z[2429] = Zone{"gob.es", &z[449], x, 0, e, e, w{"c.nic.es", "g.nic.es", "h.nic.es", "n3ns.nic.es"}, n, n, n, e, e, n, f}
	z[2430] = Zone{"nom.es", &z[449], x, 0, e, e, w{"c.nic.es", "g.nic.es", "h.nic.es", "n3ns.nic.es"}, n, n, n, e, e, n, f}
	z[2431] = Zone{"org.es", &z[449], x, 0, e, e, w{"c.nic.es", "g.nic.es", "h.nic.es", "n3ns.nic.es"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2431() {
	z[2432] = Zone{"biz.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2433] = Zone{"com.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2434] = Zone{"edu.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2435] = Zone{"gov.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2436] = Zone{"info.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2437] = Zone{"name.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2438] = Zone{"net.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
	z[2439] = Zone{"org.et", &z[453], x, 0, e, e, w{"a.nic.et", "b.nic.et"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2439() {
	z[2440] = Zone{"24.eu", &z[455], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.eu", "95.217.58.108"}, n, n, e, e, n, t}
	z[2441] = Zone{"ecb.eu", &z[455], x, 0, e, e, w{"ns1.coltnet.at", "ns8.colt.net", "pdns109.ultradns.biz", "pdns109.ultradns.com", "pdns109.ultradns.net", "pdns109.ultradns.org"}, n, n, n, e, e, n, f}
	z[2442] = Zone{"europa.eu", &z[455], x, 0, e, e, w{"ans1.cw.net", "ans2.cw.net", "ns1.bt.net", "ns1bru.europa.eu", "ns1lux.europa.eu", "ns2bru.europa.eu", "ns2eu.bt.net", "ns2lux.europa.eu", "ns3bru.europa.eu", "ns3lux.europa.eu", "ns4az1.europa.eu"}, n, n, n, e, e, n, f}
	z[2443] = Zone{"ac.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2444] = Zone{"biz.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2445] = Zone{"com.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2446] = Zone{"gov.fj", &z[498], x, 0, e, e, w{"dnssecondary.auth-servers.net", "itcns1.gov.fj", "itcns2.gov.fj", "ns-ext.vix.com", "tabua.tfl.com.fj"}, n, n, n, e, e, n, t}
	z[2447] = Zone{"info.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2447() {
	z[2448] = Zone{"mil.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2449] = Zone{"name.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2450] = Zone{"net.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2451] = Zone{"org.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2452] = Zone{"pro.fj", &z[498], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2453] = Zone{"school.fj", &z[498], x, 0, e, e, w{"drua.tfl.com.fj", "ns-ext.isc.org", "tabua.tfl.com.fj"}, n, n, n, e, e, n, t}
	z[2454] = Zone{"ac.fk", &z[499], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2455] = Zone{"co.fk", &z[499], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2455() {
	z[2456] = Zone{"gov.fk", &z[499], x, 0, e, e, w{"ns1.horizon.net.fk", "ns2.horizon.net.fk"}, n, n, n, e, e, n, f}
	z[2457] = Zone{"net.fk", &z[499], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2458] = Zone{"nom.fk", &z[499], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2459] = Zone{"org.fk", &z[499], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2460] = Zone{"0.fm", &z[508], x, 0, e, e, w{"ns1.bb-online.net", "ns2.bb-online.org"}, w{"0.fm", "139.162.244.72"}, n, n, e, e, n, t}
	z[2461] = Zone{"2.fm", &z[508], x, 0, e, e, w{"ns1.idnz.net", "ns2.idnz.net", "ns3.idnz.net"}, w{"109.234.111.119"}, n, n, e, e, n, t}
	z[2462] = Zone{"4.fm", &z[508], x, 0, e, e, w{"ns1.bb-online.net", "ns2.bb-online.org"}, w{"139.162.244.72", "4.fm"}, n, n, e, e, n, t}
	z[2463] = Zone{"6.fm", &z[508], x, 0, e, e, w{"ns1.bb-online.net", "ns2.bb-online.org"}, w{"139.162.244.72", "6.fm"}, n, n, e, e, n, t}
}

//go:noinline
func i2463() {
	z[2464] = Zone{"8.fm", &z[508], x, 0, e, e, w{"ns1.bb-online.net", "ns2.bb-online.org"}, w{"139.162.244.72", "8.fm"}, n, n, e, e, n, t}
	z[2465] = Zone{"radio.fm", &z[508], x, 0, e, e, w{"a.nic.fm", "b.nic.fm", "c.nic.fm", "d.nic.fm"}, n, n, n, e, e, w{"https://rdap.centralnic.com/radio.fm/"}, t}
	z[2466] = Zone{"aero.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2467] = Zone{"biz.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2468] = Zone{"co.fo", &z[509], x, 0, e, e, w{"ns1.cctld.com", "ns2.cctld.com"}, n, n, n, e, e, n, f}
	z[2469] = Zone{"com.fo", &z[509], x, 0, e, e, w{"ns1.cctld.com", "ns2.cctld.com"}, n, n, n, e, e, n, f}
	z[2470] = Zone{"coop.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2471] = Zone{"edu.fo", &z[509], x, 0, e, e, w{"ns1.nic.fo", "ns2.nic.fo", "ns3.nic.fo"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2471() {
	z[2472] = Zone{"flp.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2473] = Zone{"ftp.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2474] = Zone{"gov.fo", &z[509], x, 0, e, e, w{"ns1.elektron.fo", "ns2.elektron.fo"}, n, n, n, e, e, n, f}
	z[2475] = Zone{"info.fo", &z[509], x, 0, e, e, w{"ns1.bodis.com", "ns2.bodis.com"}, w{"199.59.243.225"}, n, n, e, e, n, f}
	z[2476] = Zone{"int.fo", &z[509], x, 0, e, e, w{"1-you.njalla.no", "2-can.njalla.in", "3-get.njalla.fo"}, n, n, n, e, e, n, f}
	z[2477] = Zone{"internet.fo", &z[509], x, 0, e, e, w{"ns1.olivant.fo", "ns2.olivant.fo"}, n, n, n, e, e, n, f}
	z[2478] = Zone{"irc.fo", &z[509], x, 0, e, e, w{"art.ns.cloudflare.com", "sandy.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[2479] = Zone{"mail.fo", &z[509], x, 0, e, e, w{"ns1.nic.fo", "ns2.nic.fo", "ns3.nic.fo"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2479() {
	z[2480] = Zone{"mil.fo", &z[509], x, 0, e, e, w{"ns1.afriregister.com", "ns3.afriregister.com"}, n, n, n, e, e, n, f}
	z[2481] = Zone{"museum.fo", &z[509], x, 0, e, e, w{"ns1.xn--kisnavn-p1a.fo", "ns2.xn--kisnavn-p1a.fo", "ns3.xn--kisnavn-p1a.fo"}, n, n, n, e, e, n, f}
	z[2482] = Zone{"name.fo", &z[509], x, 0, e, e, w{"ns1.xn--kisnavn-p1a.fo", "ns2.xn--kisnavn-p1a.fo", "ns3.xn--kisnavn-p1a.fo"}, n, n, n, e, e, n, f}
	z[2483] = Zone{"org.fo", &z[509], x, 0, e, e, w{"ns1.idvdns.com", "ns2.idvdns.com", "ns3.idvdns.com", "ns4.idvdns.com"}, n, n, n, e, e, n, f}
	z[2484] = Zone{"telenet.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2485] = Zone{"telnet.fo", &z[509], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2486] = Zone{"web.fo", &z[509], x, 0, e, e, w{"ns1.xn--kisnavn-p1a.fo", "ns2.xn--kisnavn-p1a.fo", "ns3.xn--kisnavn-p1a.fo"}, n, n, n, e, e, n, f}
	z[2487] = Zone{"www.fo", &z[509], x, 0, e, e, w{"ns1.xn--kisnavn-p1a.fo", "ns2.xn--kisnavn-p1a.fo", "ns3.xn--kisnavn-p1a.fo"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2487() {
	z[2488] = Zone{"aeroport.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2489] = Zone{"asso.fr", &z[520], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2490] = Zone{"avocat.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2491] = Zone{"chambagri.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2492] = Zone{"chirurgiens-dentistes.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2493] = Zone{"com.fr", &z[520], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2494] = Zone{"experts-comptables.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2495] = Zone{"geometre-expert.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
}

//go:noinline
func i2495() {
	z[2496] = Zone{"gouv.fr", &z[520], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2497] = Zone{"medecin.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2498] = Zone{"nom.fr", &z[520], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2499] = Zone{"notaires.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2500] = Zone{"pharmacien.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2501] = Zone{"port.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2502] = Zone{"prd.fr", &z[520], x, 0, e, e, w{"ns1.renater.fr", "ns2.renater.fr", "ns4.cnrs.fr"}, n, n, n, e, e, n, f}
	z[2503] = Zone{"presse.fr", &z[520], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2503() {
	z[2504] = Zone{"tm.fr", &z[520], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2505] = Zone{"veterinaire.fr", &z[520], x, 0, e, e, w{"ns1.smallregistry.net", "ns2.smallregistry.net", "ns3.smallregistry.net", "ns4.smallregistry.net"}, n, n, n, "whois.smallregistry.net", e, n, t}
	z[2506] = Zone{"ac.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2507] = Zone{"aeroport.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2508] = Zone{"asso.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga", "ns3.aninf.ga"}, n, n, n, e, e, n, f}
	z[2509] = Zone{"co.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga", "ns3.aninf.ga"}, n, n, n, e, e, n, f}
	z[2510] = Zone{"com.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2511] = Zone{"ed.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2511() {
	z[2512] = Zone{"edu.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga", "ns3.aninf.ga"}, n, n, n, e, e, n, f}
	z[2513] = Zone{"go.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2514] = Zone{"int.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2515] = Zone{"net.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga", "ns3.aninf.ga"}, n, n, n, e, e, n, f}
	z[2516] = Zone{"or.ga", &z[535], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2517] = Zone{"org.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga"}, n, n, n, e, e, n, f}
	z[2518] = Zone{"presse.ga", &z[535], x, 0, e, e, w{"ns1.aninf.ga", "ns2.aninf.ga", "ns3.aninf.ga"}, n, n, n, e, e, n, f}
	z[2519] = Zone{"hmg.gb", &z[546], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2519() {
	z[2520] = Zone{"com.gd", &z[549], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2521] = Zone{"edu.gd", &z[549], x, 0, e, e, w{"a.nic.gd", "b.nic.gd", "c.nic.gd", "d.nic.gd"}, n, n, n, e, e, n, f}
	z[2522] = Zone{"gov.gd", &z[549], x, 0, e, e, w{"pam.ns.cloudflare.com", "reese.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[2523] = Zone{"mlt.gd", &z[549], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2524] = Zone{"net.gd", &z[549], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2525] = Zone{"org.gd", &z[549], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2526] = Zone{"sci.gd", &z[549], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2527] = Zone{"com.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2527() {
	z[2528] = Zone{"edu.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, f}
	z[2529] = Zone{"gov.ge", &z[551], x, 0, e, e, w{"ns01.gov.ge", "ns02.gov.ge"}, n, n, n, e, e, n, f}
	z[2530] = Zone{"mil.ge", &z[551], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2531] = Zone{"net.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, f}
	z[2532] = Zone{"org.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, f}
	z[2533] = Zone{"pvt.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, f}
	z[2534] = Zone{"school.ge", &z[551], x, 0, e, e, w{"ns.nic.ge", "ns1.nic-ge.com", "ns2.nic-ge.com", "ns2.nic.ge"}, n, n, n, e, e, n, t}
	z[2535] = Zone{"24.gg", &z[559], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.gg", "95.217.58.108"}, n, n, e, e, n, t}
}

//go:noinline
func i2535() {
	z[2536] = Zone{"ac.gg", &z[559], x, 0, e, e, w{"dns0.mtgsy.com", "dns1.name-s.net", "dns2.name-s.net", "dns3.mtgsy.com", "dns4.mtgsy.com"}, n, n, n, e, e, n, f}
	z[2537] = Zone{"co.gg", &z[559], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[2538] = Zone{"eu.gg", &z[559], x, 0x200, e, e, w{"ns1.dan.com", "ns2.dan.com"}, w{"3.64.163.50"}, n, n, e, e, n, t}
	z[2539] = Zone{"gov.gg", &z[559], x, 0, e, e, w{"dns1.name-s.net", "dns2.name-s.net", "dns3.mtgsy.com", "dns4.mtgsy.com", "gsy1.ns.mtgsy.com"}, n, n, n, e, e, n, f}
	z[2540] = Zone{"net.gg", &z[559], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[2541] = Zone{"org.gg", &z[559], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[2542] = Zone{"sch.gg", &z[559], x, 0, e, e, w{"dns1.name-s.net", "dns2.name-s.net", "dns3.mtgsy.com", "dns4.mtgsy.com", "gsy1.ns.mtgsy.com"}, n, n, n, e, e, n, f}
	z[2543] = Zone{"tr.gg", &z[559], x, 0, e, e, w{"ns01.webme.com", "ns02.webme.com"}, w{"193.238.27.36"}, n, n, e, e, n, f}
}

//go:noinline
func i2543() {
	z[2544] = Zone{"tv.gg", &z[559], x, 0x200, e, e, n, n, n, n, e, e, n, t}
	z[2545] = Zone{"web.gg", &z[559], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "web.gg"}, n, n, e, e, n, t}
	z[2546] = Zone{"com.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, f}
	z[2547] = Zone{"edu.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, f}
	z[2548] = Zone{"gov.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, f}
	z[2549] = Zone{"mil.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, f}
	z[2550] = Zone{"net.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, t}
	z[2551] = Zone{"org.gh", &z[561], x, 0, e, e, w{"ns.dns.br", "ns1.nic.gh", "ns2.nic.gh"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2551() {
	z[2552] = Zone{"com.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2553] = Zone{"edu.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2554] = Zone{"gov.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2555] = Zone{"ltd.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2556] = Zone{"mod.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2557] = Zone{"org.gi", &z[562], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[2558] = Zone{"co.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2559] = Zone{"com.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2559() {
	z[2560] = Zone{"edu.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2561] = Zone{"gov.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, t}
	z[2562] = Zone{"net.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2563] = Zone{"org.gl", &z[567], x, 0, e, e, w{"d.nic.gl", "ns1.anycastdns.cz", "ns2.anycastdns.cz"}, n, n, n, e, e, n, f}
	z[2564] = Zone{"ac.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2565] = Zone{"co.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2566] = Zone{"com.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2567] = Zone{"gov.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2567() {
	z[2568] = Zone{"net.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2569] = Zone{"org.gn", &z[580], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2570] = Zone{"24.gp", &z[594], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.gp", "95.217.58.108"}, n, n, e, e, n, t}
	z[2571] = Zone{"asia.gp", &z[594], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "asia.gp"}, n, n, e, e, n, t}
	z[2572] = Zone{"asso.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, f}
	z[2573] = Zone{"co.gp", &z[594], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "co.gp"}, n, n, e, e, n, t}
	z[2574] = Zone{"com.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, f}
	z[2575] = Zone{"edu.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2575() {
	z[2576] = Zone{"eu.gp", &z[594], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "eu.gp"}, n, n, e, e, n, t}
	z[2577] = Zone{"mobi.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, f}
	z[2578] = Zone{"net.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, f}
	z[2579] = Zone{"online.gp", &z[594], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "online.gp"}, n, n, e, e, n, t}
	z[2580] = Zone{"org.gp", &z[594], x, 0, e, e, w{"ns-gp.nic.fr", "ns1.nic.gp", "ns2.nic.gp"}, n, n, n, e, e, n, f}
	z[2581] = Zone{"com.gr", &z[596], x, 0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, e, n, f}
	z[2582] = Zone{"edu.gr", &z[596], x, 0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, e, n, f}
	z[2583] = Zone{"gov.gr", &z[596], x, 0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2583() {
	z[2584] = Zone{"net.gr", &z[596], x, 0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, e, n, f}
	z[2585] = Zone{"org.gr", &z[596], x, 0, e, e, w{"estia.ics.forth.gr", "gr-at.ics.forth.gr", "gr-c.ics.forth.gr", "gr-d.ics.forth.gr", "grdns-m.ics.forth.gr", "grdns.ics.forth.gr"}, n, n, n, e, e, n, f}
	z[2586] = Zone{"com.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2587] = Zone{"edu.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2588] = Zone{"gob.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2589] = Zone{"ind.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2590] = Zone{"mil.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2591] = Zone{"net.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2591() {
	z[2592] = Zone{"org.gt", &z[606], x, 0, e, e, w{"a.lactld.org", "gt-e.ns.nic.cz", "gt.anycastdns.cz", "ns.dns.br", "pch.gt", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[2593] = Zone{"com.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2594] = Zone{"edu.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2595] = Zone{"gov.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2596] = Zone{"guam.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, t}
	z[2597] = Zone{"net.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2598] = Zone{"org.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[2599] = Zone{"web.gu", &z[607], x, 0, e, e, w{"gold.uog.edu", "green.uog.edu", "gu.cctld.authdns.ripe.net", "phloem.uoregon.edu"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2599() {
	z[2600] = Zone{"co.gy", &z[617], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2601] = Zone{"com.gy", &z[617], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2602] = Zone{"net.gy", &z[617], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2603] = Zone{"org.gy", &z[617], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2604] = Zone{"com.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
	z[2605] = Zone{"edu.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
	z[2606] = Zone{"gov.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
	z[2607] = Zone{"idv.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2607() {
	z[2608] = Zone{"inc.hk", &z[640], x, 0, e, e, w{"a.udrtld.net", "b.udrtld.net", "c.udrtld.net", "dns7.cloudns.net", "dns8.cloudns.net"}, n, n, n, "whois.registry.hk.com", e, n, f}
	z[2609] = Zone{"ltd.hk", &z[640], x, 0, e, e, w{"a.udrtld.net", "b.udrtld.net", "c.udrtld.net", "dns7.cloudns.net", "dns8.cloudns.net"}, n, n, n, "whois.registry.hk.com", e, n, f}
	z[2610] = Zone{"net.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
	z[2611] = Zone{"org.hk", &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, f}
	z[2612] = Zone{"xn--55qx5d.hk" /* 公司.hk */, &z[640], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[2613] = Zone{"com.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2614] = Zone{"edu.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2615] = Zone{"gob.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2615() {
	z[2616] = Zone{"mil.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2617] = Zone{"net.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2618] = Zone{"org.hn", &z[643], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2619] = Zone{"com.hr", &z[666], x, 0, e, e, w{"dns-ez-1.carnet.hr", "dns1.com.hr", "dns2.com.hr"}, n, n, n, e, e, n, f}
	z[2620] = Zone{"from.hr", &z[666], x, 0, e, e, w{"dns-ez-1.carnet.hr", "dns1.from.hr", "dns2.from.hr"}, n, n, n, e, e, n, f}
	z[2621] = Zone{"iz.hr", &z[666], x, 0, e, e, w{"dns-ez-1.carnet.hr", "dns1.from.hr", "dns1.iz.hr", "dns2.from.hr", "dns2.iz.hr"}, n, n, n, e, e, n, f}
	z[2622] = Zone{"name.hr", &z[666], x, 0, e, e, w{"dns-ez-1.carnet.hr", "dns1.from.hr", "dns1.name.hr", "dns2.from.hr", "dns2.name.hr"}, n, n, n, e, e, n, f}
	z[2623] = Zone{"adult.ht", &z[668], x, 0x1, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2623() {
	z[2624] = Zone{"art.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2625] = Zone{"asso.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2626] = Zone{"com.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2627] = Zone{"coop.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2628] = Zone{"edu.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2629] = Zone{"firm.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2630] = Zone{"gouv.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2631] = Zone{"info.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2631() {
	z[2632] = Zone{"med.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2633] = Zone{"net.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2634] = Zone{"org.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2635] = Zone{"perso.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2636] = Zone{"pol.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2637] = Zone{"pro.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2638] = Zone{"rel.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2639] = Zone{"shop.ht", &z[668], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2639() {
	z[2640] = Zone{"2000.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2641] = Zone{"ac.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2642] = Zone{"agrar.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2643] = Zone{"bolt.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2644] = Zone{"casino.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2645] = Zone{"city.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2646] = Zone{"co.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2647] = Zone{"edu.hu", &z[670], x, 0, e, e, w{"kubiac.iif.hu", "ns2.iif.hu", "ns2.sztaki.hbone.hu"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2647() {
	z[2648] = Zone{"erotica.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2649] = Zone{"erotika.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2650] = Zone{"film.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2651] = Zone{"forum.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2652] = Zone{"games.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2653] = Zone{"gov.hu", &z[670], x, 0, e, e, w{"adns0.gov.hu", "adns1.gov.hu", "adns2.gov.hu"}, n, n, n, e, e, n, f}
	z[2654] = Zone{"hotel.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2655] = Zone{"info.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2655() {
	z[2656] = Zone{"ingatlan.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2657] = Zone{"jogasz.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2658] = Zone{"konyvelo.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2659] = Zone{"lakas.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2660] = Zone{"media.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2661] = Zone{"news.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2662] = Zone{"org.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2663] = Zone{"priv.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2663() {
	z[2664] = Zone{"reklam.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2665] = Zone{"sex.hu", &z[670], x, 0x1, e, e, n, n, n, n, e, e, n, f}
	z[2666] = Zone{"shop.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2667] = Zone{"sport.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2668] = Zone{"suli.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2669] = Zone{"szex.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2670] = Zone{"tm.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2671] = Zone{"tozsde.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2671() {
	z[2672] = Zone{"utazas.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2673] = Zone{"video.hu", &z[670], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2674] = Zone{"ac.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2675] = Zone{"biz.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2676] = Zone{"co.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2677] = Zone{"desa.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, t}
	z[2678] = Zone{"go.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2679] = Zone{"mi.id", &z[678], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2679() {
	z[2680] = Zone{"mil.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2681] = Zone{"my.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2682] = Zone{"net.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2683] = Zone{"or.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2684] = Zone{"sch.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2685] = Zone{"web.id", &z[678], x, 0, e, e, w{"b.dns.id", "c.dns.id", "d.dns.id", "e.dns.id"}, n, n, n, e, e, n, f}
	z[2686] = Zone{"com.ie", &z[680], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2687] = Zone{"gov.ie", &z[680], x, 0, e, e, w{"cwext.gn.gov.ie", "ext3.gn.gov.ie", "gbext.gn.gov.ie"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2687() {
	z[2688] = Zone{"net.ie", &z[680], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2689] = Zone{"nli.ie", &z[680], x, 0, e, e, w{"auth-ns2.heanet.ie", "auth-ns3.heanet.ie", "ns.heanet.ie"}, n, n, n, e, e, n, f}
	z[2690] = Zone{"org.ie", &z[680], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[2691] = Zone{"ac.il", &z[685], x, 0, e, e, w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2692] = Zone{"co.il", &z[685], x, 0, e, "https://www.isoc.org.il/domain-name-registry", w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2693] = Zone{"gov.il", &z[685], x, 0, e, e, w{"dns3.gov.il", "ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2694] = Zone{"idf.il", &z[685], x, 0, e, e, w{"ns-1425.awsdns-50.org", "ns-1637.awsdns-12.co.uk", "ns-458.awsdns-57.com", "ns-817.awsdns-38.net"}, n, n, n, e, e, n, f}
	z[2695] = Zone{"k12.il", &z[685], x, 0, e, e, w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2695() {
	z[2696] = Zone{"muni.il", &z[685], x, 0, e, e, w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2697] = Zone{"net.il", &z[685], x, 0, e, "https://www.isoc.org.il/domain-name-registry", w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2698] = Zone{"org.il", &z[685], x, 0, e, "https://www.isoc.org.il/domain-name-registry", w{"ilns.ilan.net.il", "lookup.iucc.ac.il", "ns1.ns.il", "ns3.ns.il", "ns4.ns.il", "nsa.ns.il", "nsb.ns.il", "nse.ns.il"}, n, n, n, e, e, n, f}
	z[2699] = Zone{"ac.im", &z[686], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2700] = Zone{"co.im", &z[686], z[5117:5119], 0, e, e, n, n, n, n, e, e, n, f}
	z[2701] = Zone{"com.im", &z[686], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2702] = Zone{"gov.im", &z[686], x, 0, e, e, w{"barney.advsys.co.uk", "hoppy.iom.com", "pebbles.iom.com"}, n, n, n, e, e, n, f}
	z[2703] = Zone{"net.im", &z[686], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2703() {
	z[2704] = Zone{"org.im", &z[686], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2705] = Zone{"5g.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2706] = Zone{"6g.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2707] = Zone{"ac.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.ncst.ernet.in", e, w{"https://rdap.registry.in/"}, f}
	z[2708] = Zone{"ahmdabad.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2709] = Zone{"ai.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2710] = Zone{"am.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2711] = Zone{"bihar.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2711() {
	z[2712] = Zone{"biz.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2713] = Zone{"business.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2714] = Zone{"ca.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2715] = Zone{"cn.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2716] = Zone{"co.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.inregistry.net", e, w{"https://rdap.registry.in/"}, f}
	z[2717] = Zone{"com.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2718] = Zone{"coop.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2719] = Zone{"cs.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2719() {
	z[2720] = Zone{"delhi.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2721] = Zone{"dr.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2722] = Zone{"edu.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.ncst.ernet.in", e, w{"https://rdap.registry.in/"}, f}
	z[2723] = Zone{"er.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2724] = Zone{"ernet.in", &z[691], x, 0x200, e, e, w{"dns.ernet.in", "dns1.ernet.in", "dns2.ernet.in", "dns3.ernet.in", "dns4.ernet.in", "e-eihq01.eis.ernet.in", "e-eihq02.eis.ernet.in", "ns1.ernet.in"}, n, n, n, e, e, n, t}
	z[2725] = Zone{"firm.in", &z[691], x, 0x200, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.inregistry.net", e, n, f}
	z[2726] = Zone{"gen.in", &z[691], x, 0x200, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.inregistry.net", e, n, f}
	z[2727] = Zone{"gov.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.ncst.ernet.in", e, w{"https://rdap.registry.in/"}, f}
}

//go:noinline
func i2727() {
	z[2728] = Zone{"gujarat.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2729] = Zone{"ind.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, e, e, w{"https://rdap.registry.in/"}, f}
	z[2730] = Zone{"info.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2731] = Zone{"int.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2732] = Zone{"internet.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2733] = Zone{"io.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2734] = Zone{"me.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2735] = Zone{"mil.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.ncst.ernet.in", e, n, f}
}

//go:noinline
func i2735() {
	z[2736] = Zone{"net.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.inregistry.net", e, w{"https://rdap.registry.in/"}, f}
	z[2737] = Zone{"org.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.inregistry.net", e, w{"https://rdap.registry.in/"}, f}
	z[2738] = Zone{"pg.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2739] = Zone{"post.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2740] = Zone{"pro.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2741] = Zone{"res.in", &z[691], x, 0, e, e, w{"ns1.registry.in", "ns2.registry.in", "ns3.registry.in", "ns4.registry.in", "ns5.registry.in", "ns6.registry.in"}, n, n, n, "whois.ncst.ernet.in", e, w{"https://rdap.registry.in/"}, f}
	z[2742] = Zone{"travel.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2743] = Zone{"tv.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2743() {
	z[2744] = Zone{"uk.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2745] = Zone{"up.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2746] = Zone{"us.in", &z[691], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2747] = Zone{"auz.info", &z[696], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com", "ns3.instradns.com"}, n, n, n, e, e, n, f}
	z[2748] = Zone{"eu.int", &z[704], x, 0, e, e, w{"ns1.be.colt.net", "ns1.bt.net", "ns1bru.europa.eu", "ns1lux.europa.eu", "ns2bru.europa.eu", "ns2eu.bt.net", "ns2lux.europa.eu"}, n, n, n, e, e, n, f}
	z[2749] = Zone{"com.io", &z[710], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[2750] = Zone{"org.io", &z[710], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[2751] = Zone{"biz.iq", &z[712], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2751() {
	z[2752] = Zone{"com.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, f}
	z[2753] = Zone{"edu.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, f}
	z[2754] = Zone{"gov.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, f}
	z[2755] = Zone{"info.iq", &z[712], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2756] = Zone{"mil.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, f}
	z[2757] = Zone{"name.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, t}
	z[2758] = Zone{"net.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, t}
	z[2759] = Zone{"org.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2759() {
	z[2760] = Zone{"tv.iq", &z[712], x, 0, e, e, w{"ns.cmc.iq", "ns01.trs-dns.com", "ns01.trs-dns.net", "nsp-anycast.cmc.iq"}, n, n, n, e, e, n, t}
	z[2761] = Zone{"ac.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2762] = Zone{"co.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2763] = Zone{"gov.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2764] = Zone{"id.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2765] = Zone{"net.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2766] = Zone{"org.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
	z[2767] = Zone{"sch.ir", &z[713], x, 0, e, e, w{"a.nic.ir", "b.nic.ir", "c.nic.ir", "d.nic.ir"}, n, n, n, e, e, n, f}
}

//go:noinline
func i2767() {
	z[2768] = Zone{"abr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2769] = Zone{"abruzzo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2770] = Zone{"ag.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2771] = Zone{"agrigento.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2772] = Zone{"al.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2773] = Zone{"alessandria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2774] = Zone{"alto-adige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2775] = Zone{"altoadige.it", &z[722], x, 0, e, e, w{"guy.ns.cloudflare.com", "rita.ns.cloudflare.com"}, n, n, n, e, e, n, t}
}

//go:noinline
func i2775() {
	z[2776] = Zone{"an.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2777] = Zone{"ancona.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2778] = Zone{"andria-barletta-trani.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2779] = Zone{"andria-trani-barletta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2780] = Zone{"andriabarlettatrani.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2781] = Zone{"andriatranibarletta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2782] = Zone{"ao.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2783] = Zone{"aosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2783() {
	z[2784] = Zone{"aoste.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2785] = Zone{"ap.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2786] = Zone{"aq.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2787] = Zone{"aquila.it", &z[722], x, 0, e, e, w{"ns1.parkingcrew.net", "ns2.parkingcrew.net"}, w{"127.0.0.10"}, n, n, e, e, n, t}
	z[2788] = Zone{"ar.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2789] = Zone{"arezzo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2790] = Zone{"ascoli-piceno.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2791] = Zone{"ascolipiceno.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2791() {
	z[2792] = Zone{"asti.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2793] = Zone{"at.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2794] = Zone{"av.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2795] = Zone{"avellino.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2796] = Zone{"ba.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2797] = Zone{"balsan.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2798] = Zone{"bari.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2799] = Zone{"barletta-trani-andria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2799() {
	z[2800] = Zone{"barlettatraniandria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2801] = Zone{"bas.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2802] = Zone{"basilicata.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2803] = Zone{"belluno.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2804] = Zone{"benevento.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2805] = Zone{"bergamo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2806] = Zone{"bg.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2807] = Zone{"bi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2807() {
	z[2808] = Zone{"biella.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2809] = Zone{"bl.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2810] = Zone{"bn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2811] = Zone{"bo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2812] = Zone{"bologna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2813] = Zone{"bolzano.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2814] = Zone{"bozen.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2815] = Zone{"br.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2815() {
	z[2816] = Zone{"brescia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2817] = Zone{"brindisi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2818] = Zone{"bs.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2819] = Zone{"bt.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2820] = Zone{"bz.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[2821] = Zone{"ca.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2822] = Zone{"cagliari.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2823] = Zone{"cal.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2823() {
	z[2824] = Zone{"calabria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2825] = Zone{"caltanissetta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2826] = Zone{"cam.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2827] = Zone{"campania.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2828] = Zone{"campidano-medio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2829] = Zone{"campidanomedio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2830] = Zone{"campobasso.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2831] = Zone{"carbonia-iglesias.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2831() {
	z[2832] = Zone{"carboniaiglesias.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2833] = Zone{"carrara-massa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2834] = Zone{"carraramassa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2835] = Zone{"caserta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2836] = Zone{"catania.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2837] = Zone{"catanzaro.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2838] = Zone{"cb.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2839] = Zone{"ce.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2839() {
	z[2840] = Zone{"cesena-forli.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2841] = Zone{"cesenaforli.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2842] = Zone{"ch.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2843] = Zone{"chieti.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2844] = Zone{"ci.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2845] = Zone{"cl.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2846] = Zone{"cn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2847] = Zone{"co.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i2847() {
	z[2848] = Zone{"como.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2849] = Zone{"cosenza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2850] = Zone{"cr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2851] = Zone{"cremona.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2852] = Zone{"crotone.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2853] = Zone{"cs.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2854] = Zone{"ct.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2855] = Zone{"cuneo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2855() {
	z[2856] = Zone{"cz.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2857] = Zone{"dell-ogliastra.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2858] = Zone{"dellogliastra.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2859] = Zone{"edu.it", &z[722], x, 0, e, e, w{"a.dns.it", "dns.edu.it", "m.dns.it", "r.dns.it"}, n, n, n, e, e, n, f}
	z[2860] = Zone{"emilia-romagna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2861] = Zone{"emiliaromagna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2862] = Zone{"emr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2863] = Zone{"en.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2863() {
	z[2864] = Zone{"enna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2865] = Zone{"fc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2866] = Zone{"fe.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2867] = Zone{"fermo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2868] = Zone{"ferrara.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2869] = Zone{"fg.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2870] = Zone{"fi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2871] = Zone{"firenze.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2871() {
	z[2872] = Zone{"florence.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2873] = Zone{"fm.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2874] = Zone{"foggia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2875] = Zone{"forli-cesena.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2876] = Zone{"forlicesena.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2877] = Zone{"fr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2878] = Zone{"friuli-v-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2879] = Zone{"friuli-ve-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2879() {
	z[2880] = Zone{"friuli-vegiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2881] = Zone{"friuli-venezia-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2882] = Zone{"friuli-veneziagiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2883] = Zone{"friuli-vgiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2884] = Zone{"friuliv-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2885] = Zone{"friulive-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2886] = Zone{"friulivegiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2887] = Zone{"friulivenezia-giulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2887() {
	z[2888] = Zone{"friuliveneziagiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2889] = Zone{"friulivgiulia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2890] = Zone{"frosinone.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2891] = Zone{"fvg.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2892] = Zone{"ge.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2893] = Zone{"genoa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2894] = Zone{"genova.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2895] = Zone{"go.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2895() {
	z[2896] = Zone{"gorizia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2897] = Zone{"gov.it", &z[722], x, 0, e, e, w{"a.dns.it", "dns.gov.it", "m.dns.it", "r.dns.it"}, n, n, n, e, e, n, f}
	z[2898] = Zone{"gr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2899] = Zone{"grosseto.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2900] = Zone{"iglesias-carbonia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2901] = Zone{"iglesiascarbonia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2902] = Zone{"im.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2903] = Zone{"imperia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2903() {
	z[2904] = Zone{"is.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2905] = Zone{"isernia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2906] = Zone{"kr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2907] = Zone{"la-spezia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2908] = Zone{"laquila.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2909] = Zone{"laspezia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2910] = Zone{"latina.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2911] = Zone{"laz.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2911() {
	z[2912] = Zone{"lazio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2913] = Zone{"lc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2914] = Zone{"le.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2915] = Zone{"lecce.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2916] = Zone{"lecco.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2917] = Zone{"li.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2918] = Zone{"lig.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2919] = Zone{"liguria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2919() {
	z[2920] = Zone{"livorno.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2921] = Zone{"lo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2922] = Zone{"lodi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2923] = Zone{"lom.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2924] = Zone{"lombardia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2925] = Zone{"lombardy.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2926] = Zone{"lt.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2927] = Zone{"lu.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2927() {
	z[2928] = Zone{"lucania.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2929] = Zone{"lucca.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2930] = Zone{"macerata.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2931] = Zone{"mantova.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2932] = Zone{"mar.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2933] = Zone{"marche.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2934] = Zone{"massa-carrara.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2935] = Zone{"massacarrara.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2935() {
	z[2936] = Zone{"matera.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2937] = Zone{"mb.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2938] = Zone{"mc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2939] = Zone{"me.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2940] = Zone{"medio-campidano.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2941] = Zone{"mediocampidano.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2942] = Zone{"messina.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2943] = Zone{"mi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2943() {
	z[2944] = Zone{"milan.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2945] = Zone{"milano.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2946] = Zone{"mn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2947] = Zone{"mo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2948] = Zone{"modena.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2949] = Zone{"mol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2950] = Zone{"molise.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2951] = Zone{"monza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2951() {
	z[2952] = Zone{"monza-brianza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2953] = Zone{"monza-e-della-brianza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2954] = Zone{"monzabrianza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2955] = Zone{"monzaebrianza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2956] = Zone{"monzaedellabrianza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2957] = Zone{"ms.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2958] = Zone{"mt.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2959] = Zone{"na.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2959() {
	z[2960] = Zone{"naples.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2961] = Zone{"napoli.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2962] = Zone{"no.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2963] = Zone{"novara.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2964] = Zone{"nu.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2965] = Zone{"nuoro.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2966] = Zone{"og.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2967] = Zone{"ogliastra.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2967() {
	z[2968] = Zone{"olbia-tempio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2969] = Zone{"olbiatempio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2970] = Zone{"or.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2971] = Zone{"oristano.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2972] = Zone{"ot.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2973] = Zone{"pa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2974] = Zone{"padova.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2975] = Zone{"padua.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2975() {
	z[2976] = Zone{"palermo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2977] = Zone{"parma.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2978] = Zone{"pavia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2979] = Zone{"pc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2980] = Zone{"pd.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2981] = Zone{"pe.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2982] = Zone{"perugia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2983] = Zone{"pesaro-urbino.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2983() {
	z[2984] = Zone{"pesarourbino.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2985] = Zone{"pescara.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2986] = Zone{"pg.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2987] = Zone{"pi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2988] = Zone{"piacenza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2989] = Zone{"piedmont.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2990] = Zone{"piemonte.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2991] = Zone{"pisa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2991() {
	z[2992] = Zone{"pistoia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2993] = Zone{"pmn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2994] = Zone{"pn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2995] = Zone{"po.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2996] = Zone{"pordenone.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2997] = Zone{"potenza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2998] = Zone{"pr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[2999] = Zone{"prato.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i2999() {
	z[3000] = Zone{"pt.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3001] = Zone{"pu.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3002] = Zone{"pug.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3003] = Zone{"puglia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3004] = Zone{"pv.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3005] = Zone{"pz.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3006] = Zone{"ra.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3007] = Zone{"ragusa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3007() {
	z[3008] = Zone{"ravenna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3009] = Zone{"rc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3010] = Zone{"re.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3011] = Zone{"reggio-calabria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3012] = Zone{"reggio-emilia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3013] = Zone{"reggiocalabria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3014] = Zone{"reggioemilia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3015] = Zone{"rg.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3015() {
	z[3016] = Zone{"ri.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3017] = Zone{"rieti.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3018] = Zone{"rimini.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3019] = Zone{"rm.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3020] = Zone{"rn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3021] = Zone{"ro.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3022] = Zone{"roma.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3023] = Zone{"rome.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3023() {
	z[3024] = Zone{"rovigo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3025] = Zone{"sa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3026] = Zone{"salerno.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3027] = Zone{"sar.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3028] = Zone{"sardegna.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3029] = Zone{"sardinia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3030] = Zone{"sassari.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3031] = Zone{"savona.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3031() {
	z[3032] = Zone{"si.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3033] = Zone{"sic.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3034] = Zone{"sicilia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3035] = Zone{"sicily.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3036] = Zone{"siena.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3037] = Zone{"siracusa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3038] = Zone{"so.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3039] = Zone{"sondrio.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3039() {
	z[3040] = Zone{"sp.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3041] = Zone{"sr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3042] = Zone{"ss.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3043] = Zone{"suedtirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3044] = Zone{"sv.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3045] = Zone{"ta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3046] = Zone{"taa.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3047] = Zone{"taranto.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3047() {
	z[3048] = Zone{"te.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3049] = Zone{"tempio-olbia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3050] = Zone{"tempioolbia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3051] = Zone{"teramo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3052] = Zone{"terni.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3053] = Zone{"tn.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3054] = Zone{"to.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3055] = Zone{"torino.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3055() {
	z[3056] = Zone{"tos.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3057] = Zone{"toscana.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3058] = Zone{"tp.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3059] = Zone{"tr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3060] = Zone{"trani-andria-barletta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3061] = Zone{"trani-barletta-andria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3062] = Zone{"traniandriabarletta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3063] = Zone{"tranibarlettaandria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3063() {
	z[3064] = Zone{"trapani.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3065] = Zone{"trentino.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3066] = Zone{"trentino-a-adige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3067] = Zone{"trentino-aadige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3068] = Zone{"trentino-alto-adige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3069] = Zone{"trentino-altoadige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3070] = Zone{"trentino-s-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3071] = Zone{"trentino-stirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3071() {
	z[3072] = Zone{"trentino-sud-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3073] = Zone{"trentino-sudtirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3074] = Zone{"trentino-sued-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3075] = Zone{"trentino-suedtirol.it", &z[722], x, 0, e, e, w{"dns3.aknet.it", "dns4.aknet.it"}, n, n, n, e, e, n, t}
	z[3076] = Zone{"trentinoa-adige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3077] = Zone{"trentinoaadige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3078] = Zone{"trentinoalto-adige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3079] = Zone{"trentinoaltoadige.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3079() {
	z[3080] = Zone{"trentinos-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3081] = Zone{"trentinosud-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3082] = Zone{"trentinosudtirol.it", &z[722], x, 0, e, e, w{"dns.widhost.net", "dns2.widhost.net"}, n, n, n, e, e, n, t}
	z[3083] = Zone{"trentinosued-tirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3084] = Zone{"trentinosuedtirol.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3085] = Zone{"trento.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3086] = Zone{"treviso.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3087] = Zone{"trieste.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3087() {
	z[3088] = Zone{"ts.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3089] = Zone{"turin.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3090] = Zone{"tuscany.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3091] = Zone{"tv.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3092] = Zone{"ud.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3093] = Zone{"udine.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3094] = Zone{"umb.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3095] = Zone{"umbria.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3095() {
	z[3096] = Zone{"urbino-pesaro.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3097] = Zone{"urbinopesaro.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3098] = Zone{"va.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3099] = Zone{"val-d-aosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3100] = Zone{"val-daosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3101] = Zone{"vald-aosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3102] = Zone{"valdaosta.it", &z[722], x, 0, e, e, w{"dns2.alsitech.it", "dns4.alsitech.it"}, n, n, n, e, e, n, t}
	z[3103] = Zone{"valle-d-aosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3103() {
	z[3104] = Zone{"valle-daosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3105] = Zone{"valled-aosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3106] = Zone{"valledaosta.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3107] = Zone{"vao.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3108] = Zone{"varese.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3109] = Zone{"vb.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3110] = Zone{"vc.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3111] = Zone{"vda.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3111() {
	z[3112] = Zone{"ve.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3113] = Zone{"ven.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3114] = Zone{"veneto.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3115] = Zone{"venezia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3116] = Zone{"venice.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3117] = Zone{"verbania.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3118] = Zone{"vercelli.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3119] = Zone{"verona.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3119() {
	z[3120] = Zone{"vi.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3121] = Zone{"vibo-valentia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3122] = Zone{"vibovalentia.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3123] = Zone{"vicenza.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3124] = Zone{"viterbo.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3125] = Zone{"vr.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3126] = Zone{"vs.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3127] = Zone{"vt.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3127() {
	z[3128] = Zone{"vv.it", &z[722], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3129] = Zone{"24.je", &z[731], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.je", "95.217.58.108"}, n, n, e, e, n, t}
	z[3130] = Zone{"co.je", &z[731], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[3131] = Zone{"gov.je", &z[731], x, 0, e, e, w{"ns1-38.azure-dns.com", "ns2-38.azure-dns.net", "ns3-38.azure-dns.org", "ns4-38.azure-dns.info"}, n, n, n, e, e, n, f}
	z[3132] = Zone{"net.je", &z[731], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[3133] = Zone{"org.je", &z[731], x, 0, e, e, w{"c.ci-servers.org", "dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk"}, n, n, n, e, e, n, f}
	z[3134] = Zone{"sch.je", &z[731], x, 0, e, e, w{"ns1-35.azure-dns.com", "ns2-35.azure-dns.net", "ns3-35.azure-dns.org", "ns4-35.azure-dns.info"}, n, n, n, e, e, n, f}
	z[3135] = Zone{"com.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3135() {
	z[3136] = Zone{"edu.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3137] = Zone{"gov.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3138] = Zone{"mil.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3139] = Zone{"net.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3140] = Zone{"org.jm", &z[738], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3141] = Zone{"com.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3142] = Zone{"edu.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, t}
	z[3143] = Zone{"gov.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3143() {
	z[3144] = Zone{"mil.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3145] = Zone{"name.jo", &z[741], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3146] = Zone{"net.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3147] = Zone{"org.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3148] = Zone{"sch.jo", &z[741], x, 0, e, e, w{"amra.nic.gov.jo", "cloudns.nic.net.jo", "jo.cctld.authdns.ripe.net", "jordan1st.nic.gov.jo", "petra.nic.gov.jo", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3149] = Zone{"ac.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3150] = Zone{"ad.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3151] = Zone{"aichi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3151() {
	z[3152] = Zone{"akita.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3153] = Zone{"aomori.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3154] = Zone{"chiba.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3155] = Zone{"co.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3156] = Zone{"ed.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3157] = Zone{"ehime.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3158] = Zone{"fukui.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3159] = Zone{"fukuoka.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3159() {
	z[3160] = Zone{"fukushima.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3161] = Zone{"gifu.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3162] = Zone{"go.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3163] = Zone{"gr.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3164] = Zone{"gunma.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3165] = Zone{"hiroshima.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3166] = Zone{"hokkaido.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3167] = Zone{"hyogo.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3167() {
	z[3168] = Zone{"ibaraki.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3169] = Zone{"ishikawa.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3170] = Zone{"iwate.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3171] = Zone{"kagawa.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3172] = Zone{"kagoshima.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3173] = Zone{"kanagawa.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3174] = Zone{"kochi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3175] = Zone{"kumamoto.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3175() {
	z[3176] = Zone{"kyoto.jp", &z[746], x, 0x84, e, e, n, n, w{"Kyoto"}, n, e, e, n, f}
	z[3177] = Zone{"lg.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3178] = Zone{"mie.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3179] = Zone{"miyagi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3180] = Zone{"miyazaki.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3181] = Zone{"nagano.jp", &z[746], x, 0x84, e, e, n, n, w{"Nagano", "JP-20"}, n, e, e, n, f}
	z[3182] = Zone{"nagasaki.jp", &z[746], x, 0x84, e, e, n, n, w{"Nagasaki"}, n, e, e, n, f}
	z[3183] = Zone{"nara.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3183() {
	z[3184] = Zone{"ne.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3185] = Zone{"niigata.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3186] = Zone{"oita.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3187] = Zone{"okayama.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3188] = Zone{"okinawa.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3189] = Zone{"or.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3190] = Zone{"osaka.jp", &z[746], x, 0x84, e, e, n, n, w{"Osaka"}, n, e, e, n, f}
	z[3191] = Zone{"saga.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3191() {
	z[3192] = Zone{"saitama.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3193] = Zone{"shiga.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3194] = Zone{"shimane.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3195] = Zone{"shizuoka.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3196] = Zone{"tochigi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3197] = Zone{"tokushima.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3198] = Zone{"tokyo.jp", &z[746], x, 0x84, e, e, n, n, w{"Tokyo"}, n, e, e, n, f}
	z[3199] = Zone{"tottori.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3199() {
	z[3200] = Zone{"toyama.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3201] = Zone{"wakayama.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3202] = Zone{"yamagata.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3203] = Zone{"yamaguchi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3204] = Zone{"yamanashi.jp", &z[746], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3205] = Zone{"xn--ehqz56n.jp" /* 三重.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3206] = Zone{"xn--1lqs03n.jp" /* 京都.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3207] = Zone{"xn--qqqt11m.jp" /* 佐賀.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3207() {
	z[3208] = Zone{"xn--f6qx53a.jp" /* 兵庫.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3209] = Zone{"xn--djrs72d6uy.jp" /* 北海道.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3210] = Zone{"xn--mkru45i.jp" /* 千葉.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3211] = Zone{"xn--0trq7p7nn.jp" /* 和歌山.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3212] = Zone{"xn--5js045d.jp" /* 埼玉.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3213] = Zone{"xn--kbrq7o.jp" /* 大分.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3214] = Zone{"xn--pssu33l.jp" /* 大阪.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3215] = Zone{"xn--ntsq17g.jp" /* 奈良.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3215() {
	z[3216] = Zone{"xn--uisz3g.jp" /* 宮城.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3217] = Zone{"xn--6btw5a.jp" /* 宮崎.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3218] = Zone{"xn--1ctwo.jp" /* 富山.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3219] = Zone{"xn--6orx2r.jp" /* 山口.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3220] = Zone{"xn--rht61e.jp" /* 山形.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3221] = Zone{"xn--rht27z.jp" /* 山梨.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3222] = Zone{"xn--nit225k.jp" /* 岐阜.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3223] = Zone{"xn--rht3d.jp" /* 岡山.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3223() {
	z[3224] = Zone{"xn--djty4k.jp" /* 岩手.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3225] = Zone{"xn--klty5x.jp" /* 島根.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3226] = Zone{"xn--kltx9a.jp" /* 広島.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3227] = Zone{"xn--kltp7d.jp" /* 徳島.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3228] = Zone{"xn--c3s14m.jp" /* 愛媛.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3229] = Zone{"xn--vgu402c.jp" /* 愛知.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3230] = Zone{"xn--efvn9s.jp" /* 新潟.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3231] = Zone{"xn--1lqs71d.jp" /* 東京.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3231() {
	z[3232] = Zone{"xn--4pvxs.jp" /* 栃木.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3233] = Zone{"xn--uuwu58a.jp" /* 沖縄.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3234] = Zone{"xn--zbx025d.jp" /* 滋賀.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3235] = Zone{"xn--8pvr4u.jp" /* 熊本.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3236] = Zone{"xn--5rtp49c.jp" /* 石川.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3237] = Zone{"xn--ntso0iqx3a.jp" /* 神奈川.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3238] = Zone{"xn--elqq16h.jp" /* 福井.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3239] = Zone{"xn--4it168d.jp" /* 福岡.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3239() {
	z[3240] = Zone{"xn--klt787d.jp" /* 福島.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3241] = Zone{"xn--rny31h.jp" /* 秋田.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3242] = Zone{"xn--7t0a264c.jp" /* 群馬.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3243] = Zone{"xn--uist22h.jp" /* 茨城.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3244] = Zone{"xn--8ltr62k.jp" /* 長崎.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3245] = Zone{"xn--2m4a15e.jp" /* 長野.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3246] = Zone{"xn--32vp30h.jp" /* 青森.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3247] = Zone{"xn--4it797k.jp" /* 静岡.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3247() {
	z[3248] = Zone{"xn--5rtq34k.jp" /* 香川.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3249] = Zone{"xn--k7yn95e.jp" /* 高知.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3250] = Zone{"xn--tor131o.jp" /* 鳥取.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3251] = Zone{"xn--d5qv7z876c.jp" /* 鹿児島.jp */, &z[746], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3252] = Zone{"ac.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3253] = Zone{"co.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3254] = Zone{"go.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3255] = Zone{"info.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3255() {
	z[3256] = Zone{"me.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3257] = Zone{"mobi.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3258] = Zone{"ne.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3259] = Zone{"or.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3260] = Zone{"sc.ke", &z[755], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3261] = Zone{"com.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "ns.kg"}, n, n, n, e, e, n, f}
	z[3262] = Zone{"edu.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "ns.kg"}, n, n, n, e, e, n, t}
	z[3263] = Zone{"gov.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "kg.cctld.authdns.ripe.net", "ns.kg"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3263() {
	z[3264] = Zone{"mil.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "kg.cctld.authdns.ripe.net", "ns.kg"}, n, n, n, e, e, n, f}
	z[3265] = Zone{"net.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "kg.cctld.authdns.ripe.net", "ns.kg"}, n, n, n, e, e, n, f}
	z[3266] = Zone{"org.kg", &z[763], x, 0, e, e, w{"as.asiainfo.kg", "kg.cctld.authdns.ripe.net", "ns.kg"}, n, n, n, e, e, n, f}
	z[3267] = Zone{"com.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3268] = Zone{"edu.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3269] = Zone{"gov.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3270] = Zone{"mil.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3271] = Zone{"net.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3271() {
	z[3272] = Zone{"org.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3273] = Zone{"per.kh", &z[764], x, 0, e, e, w{"dns1.online.com.kh", "ns.camnet.com.kh", "ns1.dns.net.kh", "ns4.apnic.net"}, n, n, n, e, e, n, f}
	z[3274] = Zone{"biz.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3275] = Zone{"com.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3276] = Zone{"edu.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3277] = Zone{"eu.ki", &z[765], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "eu.ki"}, n, n, e, e, n, t}
	z[3278] = Zone{"gov.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3279] = Zone{"info.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3279() {
	z[3280] = Zone{"mob.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3281] = Zone{"mobi.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3282] = Zone{"net.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3283] = Zone{"org.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3284] = Zone{"phone.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3285] = Zone{"tel.ki", &z[765], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3286] = Zone{"asso.km", &z[775], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[3287] = Zone{"com.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3287() {
	z[3288] = Zone{"coop.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3289] = Zone{"edu.km", &z[775], x, 0, e, e, w{"dns1.nic.km", "dns2.nic.km"}, n, n, n, e, e, n, f}
	z[3290] = Zone{"gouv.km", &z[775], x, 0, e, e, w{"dns1.nic.km", "dns2.nic.km"}, n, n, n, e, e, n, f}
	z[3291] = Zone{"medecin.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3292] = Zone{"mil.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3293] = Zone{"nom.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3294] = Zone{"notaires.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3295] = Zone{"org.km", &z[775], x, 0, e, e, w{"dns1.nic.km", "dns2.nic.km"}, n, n, n, e, e, n, t}
}

//go:noinline
func i3295() {
	z[3296] = Zone{"pharmaciens.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3297] = Zone{"presse.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3298] = Zone{"tm.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3299] = Zone{"veterinaire.km", &z[775], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3300] = Zone{"co.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3301] = Zone{"com.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3302] = Zone{"edu.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3303] = Zone{"gov.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3303() {
	z[3304] = Zone{"net.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3305] = Zone{"org.kn", &z[776], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3306] = Zone{"com.kp", &z[782], x, 0, e, e, w{"ns1.com.kp", "ns2.com.kp"}, n, n, n, e, e, n, f}
	z[3307] = Zone{"net.kp", &z[782], x, 0, e, e, w{"ns1.net.kp", "ns2.net.kp"}, n, n, n, e, e, n, f}
	z[3308] = Zone{"org.kp", &z[782], x, 0, e, e, w{"ns1.org.kp", "ns2.org.kp"}, n, n, n, e, e, n, f}
	z[3309] = Zone{"rep.kp", &z[782], x, 0, e, e, w{"ns1.rep.kp", "ns2.rep.kp"}, n, n, n, e, e, n, f}
	z[3310] = Zone{"ac.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3311] = Zone{"busan.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3311() {
	z[3312] = Zone{"chungbuk.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3313] = Zone{"chungnam.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3314] = Zone{"co.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3315] = Zone{"daegu.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3316] = Zone{"daejeon.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3317] = Zone{"es.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3318] = Zone{"gangwon.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3319] = Zone{"go.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3319() {
	z[3320] = Zone{"gwangju.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3321] = Zone{"gyeongbuk.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3322] = Zone{"gyeonggi.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3323] = Zone{"gyeongnam.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3324] = Zone{"hs.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3325] = Zone{"incheon.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3326] = Zone{"jeju.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3327] = Zone{"jeonbuk.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3327() {
	z[3328] = Zone{"jeonnam.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3329] = Zone{"kg.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3330] = Zone{"mil.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3331] = Zone{"ms.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3332] = Zone{"ne.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3333] = Zone{"or.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3334] = Zone{"pe.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3335] = Zone{"re.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3335() {
	z[3336] = Zone{"sc.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3337] = Zone{"seoul.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3338] = Zone{"ulsan.kr", &z[785], x, 0, e, e, w{"b.dns.kr", "c.dns.kr", "d.dns.kr", "e.dns.kr", "f.dns.kr", "g.dns.kr"}, n, n, n, e, e, n, f}
	z[3339] = Zone{"com.kw", &z[789], x, 0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, e, n, f}
	z[3340] = Zone{"edu.kw", &z[789], x, 0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, e, n, f}
	z[3341] = Zone{"gov.kw", &z[789], x, 0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, e, n, f}
	z[3342] = Zone{"net.kw", &z[789], x, 0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, e, n, f}
	z[3343] = Zone{"org.kw", &z[789], x, 0, e, e, w{"a.nic.kw", "b.nic.kw", "c.nic.kw", "d.nic.kw", "pch.nic.kw"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3343() {
	z[3344] = Zone{"com.ky", &z[790], x, 0, e, e, w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[3345] = Zone{"edu.ky", &z[790], x, 0, e, e, w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[3346] = Zone{"gov.ky", &z[790], x, 0, e, e, w{"dns1.p04.nsone.net", "dns2.p04.nsone.net", "dns3.p04.nsone.net", "dns4.p04.nsone.net"}, n, n, n, e, e, n, f}
	z[3347] = Zone{"net.ky", &z[790], x, 0, e, e, w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[3348] = Zone{"org.ky", &z[790], x, 0, e, e, w{"ns1.uniregistry.net", "ns2.uniregistry.info", "ns3.uniregistry.net", "ns4.uniregistry.info"}, n, n, n, e, e, n, f}
	z[3349] = Zone{"com.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
	z[3350] = Zone{"edu.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
	z[3351] = Zone{"gov.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3351() {
	z[3352] = Zone{"mil.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
	z[3353] = Zone{"net.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
	z[3354] = Zone{"org.kz", &z[793], x, 0, e, e, w{"ns.nic.kz", "ns1.nic.kz", "ns2.nic.kz"}, n, n, n, e, e, n, f}
	z[3355] = Zone{"com.lb", &z[811], x, 0, e, e, w{"nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3356] = Zone{"edu.lb", &z[811], x, 0, e, e, w{"nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3357] = Zone{"gov.lb", &z[811], x, 0, e, e, w{"nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3358] = Zone{"net.lb", &z[811], x, 0, e, e, w{"nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3359] = Zone{"org.lb", &z[811], x, 0, e, e, w{"nabil.beirutix.net", "nn.uninett.no", "ns-jp.lbdr.org.lb", "ns3.seacomnet.com", "ns4.seacomnet.com", "rip.psg.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3359() {
	z[3360] = Zone{"24.lc", &z[812], x, 0x200, e, e, w{"ns1.website.org", "ns2.website.org"}, w{"138.201.129.184", "24.lc"}, n, n, e, e, n, t}
	z[3361] = Zone{"co.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[3362] = Zone{"com.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[3363] = Zone{"edu.lc", &z[812], x, 0, e, e, w{"dns2.server.ag", "ns2.dnsserver.ag"}, n, n, n, e, e, n, f}
	z[3364] = Zone{"gov.lc", &z[812], x, 0, e, e, w{"dns2.server.ag", "ns2.dnsserver.ag"}, n, n, n, e, e, n, f}
	z[3365] = Zone{"l.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[3366] = Zone{"net.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[3367] = Zone{"org.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3367() {
	z[3368] = Zone{"p.lc", &z[812], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org", "dns.server.ag"}, n, n, n, e, e, n, f}
	z[3369] = Zone{"assn.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3370] = Zone{"com.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3371] = Zone{"edu.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3372] = Zone{"gov.lk", &z[840], x, 0, e, e, w{"c.nic.lk", "d.nic.lk", "m.nic.lk", "ns1.gov.lk", "ns2.gov.lk"}, n, n, n, e, e, n, f}
	z[3373] = Zone{"grp.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3374] = Zone{"hotel.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3375] = Zone{"int.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3375() {
	z[3376] = Zone{"ltd.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3377] = Zone{"net.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3378] = Zone{"ngo.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3379] = Zone{"org.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3380] = Zone{"sch.lk", &z[840], x, 0, e, e, w{"ns1.sch.lk"}, n, n, n, e, e, n, f}
	z[3381] = Zone{"soc.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3382] = Zone{"web.lk", &z[840], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3383] = Zone{"com.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3383() {
	z[3384] = Zone{"edu.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3385] = Zone{"gov.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3386] = Zone{"net.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3387] = Zone{"org.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3388] = Zone{"vcom.lr", &z[858], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3389] = Zone{"ac.ls", &z[859], x, 0, e, e, w{"ls-ns.anycast.pch.net", "ns-ls.afrinic.net", "ns1.nic.ls", "ns2.nic.ls", "thaba.nul.ls"}, n, n, n, e, e, n, t}
	z[3390] = Zone{"co.ls", &z[859], x, 0, e, e, w{"ls-ns.anycast.pch.net", "ns-ls.afrinic.net", "ns1.nic.ls", "ns2.nic.ls"}, n, n, n, e, e, n, f}
	z[3391] = Zone{"gov.ls", &z[859], x, 0, e, e, w{"acst23ns01.gov.ls", "afdp23ns01.gov.ls", "secdns1.posix.co.za", "thaba.nul.ls"}, n, n, n, e, e, n, t}
}

//go:noinline
func i3391() {
	z[3392] = Zone{"net.ls", &z[859], x, 0, e, e, w{"ls-ns.anycast.pch.net", "ns1.nic.ls", "ns2.nic.ls"}, n, n, n, e, e, n, t}
	z[3393] = Zone{"org.ls", &z[859], x, 0, e, e, w{"ls-ns.anycast.pch.net", "ns-ls.afrinic.net", "ns1.nic.ls", "ns2.nic.ls"}, n, n, n, e, e, n, f}
	z[3394] = Zone{"sc.ls", &z[859], x, 0, e, e, w{"ns1.nic.ls", "ns1.telewebls.com", "ns2.nic.ls"}, n, n, n, e, e, n, f}
	z[3395] = Zone{"gov.lt", &z[860], x, 0, e, e, w{"ns2.domreg.lt", "ns3.is.lt", "ns4.is.lt", "nsa.domreg.lt"}, n, n, n, e, e, n, f}
	z[3396] = Zone{"asn.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3397] = Zone{"com.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3398] = Zone{"conf.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3399] = Zone{"edu.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3399() {
	z[3400] = Zone{"gov.lv", &z[868], x, 0, e, e, w{"a.gov.lv", "b.gov.lv", "c.gov.lv"}, n, n, n, e, e, n, f}
	z[3401] = Zone{"id.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3402] = Zone{"mil.lv", &z[868], x, 0, e, e, w{"ns1.mil.lv", "ns2.mil.lv"}, n, n, n, e, e, n, f}
	z[3403] = Zone{"net.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3404] = Zone{"org.lv", &z[868], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3405] = Zone{"com.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3406] = Zone{"edu.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3407] = Zone{"gov.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3407() {
	z[3408] = Zone{"id.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3409] = Zone{"med.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3410] = Zone{"net.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3411] = Zone{"org.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3412] = Zone{"plc.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3413] = Zone{"sch.ly", &z[869], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3414] = Zone{"ac.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
	z[3415] = Zone{"co.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3415() {
	z[3416] = Zone{"gov.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
	z[3417] = Zone{"net.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
	z[3418] = Zone{"org.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
	z[3419] = Zone{"press.ma", &z[870], x, 0, e, e, w{"a.tld.ma", "b.tld.ma", "c.tld.ma", "d.tld.ma", "dns.inria.fr", "e.tld.ma", "f.tld.ma", "ns-ma.nic.fr"}, n, n, n, e, e, n, f}
	z[3420] = Zone{"asso.mc", &z[891], x, 0, e, e, w{"mc.cctld.authdns.ripe.net", "ns1.nic.mc", "ns2.nic.mc", "ns3.nic.mc"}, n, n, n, e, e, n, f}
	z[3421] = Zone{"tm.mc", &z[891], x, 0, e, e, w{"mc.cctld.authdns.ripe.net", "ns1.nic.mc", "ns2.nic.mc", "ns3.nic.mc"}, n, n, n, e, e, n, f}
	z[3422] = Zone{"ac.me", &z[896], x, 0, e, e, w{"ns.ac.me", "ns2.ac.me"}, n, n, n, e, e, n, f}
	z[3423] = Zone{"co.me", &z[896], x, 0, e, e, w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3423() {
	z[3424] = Zone{"edu.me", &z[896], x, 0, e, e, w{"ns.edu.me", "ns2.edu.me"}, n, n, n, e, e, n, f}
	z[3425] = Zone{"gov.me", &z[896], x, 0, e, e, w{"derek.ns.cloudflare.com", "rosalie.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[3426] = Zone{"its.me", &z[896], x, 0, e, e, w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, e, e, n, f}
	z[3427] = Zone{"net.me", &z[896], x, 0, e, e, w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, e, e, n, f}
	z[3428] = Zone{"org.me", &z[896], x, 0, e, e, w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, e, e, n, f}
	z[3429] = Zone{"priv.me", &z[896], x, 0, e, e, w{"a0.nic.me", "a2.nic.me", "b0.nic.me", "b2.nic.me", "c0.nic.me"}, n, n, n, e, e, n, f}
	z[3430] = Zone{"co.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3431] = Zone{"com.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3431() {
	z[3432] = Zone{"edu.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3433] = Zone{"gov.mg", &z[911], x, 0, e, e, w{"ns.gov.mg", "ns1.gov.mg", "ns2.gov.mg"}, n, n, n, e, e, n, f}
	z[3434] = Zone{"in.mg", &z[911], x, 0, e, e, w{"ns1.in.mg", "ns2.in.mg"}, n, n, n, e, e, n, t}
	z[3435] = Zone{"mil.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3436] = Zone{"net.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3437] = Zone{"nom.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3438] = Zone{"org.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
	z[3439] = Zone{"prd.mg", &z[911], x, 0, e, e, w{"censvrns0001.ird.fr", "dns-tld.ird.fr", "ns-mg.malagasy.com", "ns.dts.mg", "ns.nic.mg", "pch.nic.mg"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3439() {
	z[3440] = Zone{"tm.mg", &z[911], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3441] = Zone{"com.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3442] = Zone{"edu.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3443] = Zone{"gov.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3444] = Zone{"inf.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3445] = Zone{"name.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3446] = Zone{"net.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
	z[3447] = Zone{"org.mk", &z[923], x, 0, e, e, w{"d.ext.nic.cz", "dns-mk.univie.ac.at", "ns2.arnes.si", "tld1.marnet.mk"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3447() {
	z[3448] = Zone{"pro.mk", &z[923], x, 0, e, e, w{"ns1.he.net", "ns2.he.net", "ns3.he.net", "ns4.he.net"}, w{"162.218.232.95"}, n, n, e, e, n, t}
	z[3449] = Zone{"com.ml", &z[924], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3450] = Zone{"edu.ml", &z[924], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3451] = Zone{"gov.ml", &z[924], x, 0, e, e, w{"dns01.gouv.ml", "dns02.gouv.ml"}, n, n, n, e, e, n, f}
	z[3452] = Zone{"info.ml", &z[924], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3453] = Zone{"net.ml", &z[924], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3454] = Zone{"org.ml", &z[924], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3455] = Zone{"presse.ml", &z[924], x, 0, e, e, w{"dns1.namecheaphosting.com", "dns2.namecheaphosting.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3455() {
	z[3456] = Zone{"biz.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3457] = Zone{"com.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3458] = Zone{"edu.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3459] = Zone{"gov.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3460] = Zone{"mil.mm", &z[927], x, 0, e, e, w{"master.mil.mm", "ns1.mil.mm", "ns2.mil.mm"}, n, n, n, e, e, n, t}
	z[3461] = Zone{"net.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3462] = Zone{"org.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
	z[3463] = Zone{"per.mm", &z[927], x, 0, e, e, w{"a.nic.net.mm", "b.nic.net.mm", "c.nic.net.mm", "d.nic.net.mm", "ptd.net.mm"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3463() {
	z[3464] = Zone{"edu.mn", &z[929], x, 0, e, e, w{"ns.erdemnet.mn", "ns1.erdemnet.mn", "ns2.erdemnet.mn"}, n, n, n, e, e, n, f}
	z[3465] = Zone{"gov.mn", &z[929], x, 0, e, e, w{"ns.gov.mn", "ns1.gov.mn", "ns3.gov.mn", "ns4.gov.mn"}, n, n, n, e, e, n, f}
	z[3466] = Zone{"org.mn", &z[929], x, 0, e, e, w{"ns1.magic.mn", "ns2.magic.mn", "ns3.magic.mn"}, n, n, n, e, e, n, f}
	z[3467] = Zone{"co.mo", &z[931], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3468] = Zone{"com.mo", &z[931], x, 0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, e, e, n, f}
	z[3469] = Zone{"edu.mo", &z[931], x, 0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, e, e, n, f}
	z[3470] = Zone{"gov.mo", &z[931], x, 0, e, e, w{"ns1.gov.mo", "ns2.gov.mo", "ns3.gov.mo", "ns4.gov.mo", "ns5.gov.mo"}, n, n, n, e, e, n, f}
	z[3471] = Zone{"net.mo", &z[931], x, 0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3471() {
	z[3472] = Zone{"org.mo", &z[931], x, 0, e, e, w{"a.monic.mo", "b.monic.mo", "c.monic.mo", "d.monic.mo", "e.monic.mo", "ns17.cdns.net", "ns2.cuhk.edu.hk"}, n, n, n, e, e, n, f}
	z[3473] = Zone{"co.mp", &z[953], x, 0, e, e, w{"ns1.nic.net.mp", "ns2.nic.net.mp"}, n, n, n, e, e, n, t}
	z[3474] = Zone{"com.mp", &z[953], x, 0, e, e, w{"a.dnspod.com", "b.dnspod.com", "c.dnspod.com"}, n, n, n, e, e, n, f}
	z[3475] = Zone{"gov.mp", &z[953], x, 0, e, e, w{"ns1.nic.net.mp", "ns2.nic.net.mp"}, n, n, n, e, e, n, f}
	z[3476] = Zone{"org.mp", &z[953], x, 0, e, e, w{"ns1.nic.net.mp", "ns2.nic.net.mp"}, n, n, n, e, e, n, f}
	z[3477] = Zone{"edu.mr", &z[955], x, 0, e, e, w{"ns-mr.afrinic.net", "ns-mr.nic.fr", "ns-mr.nic.tn", "ns1.nic.mr", "ns2.nic.mr", "ns3.nic.mr"}, n, n, n, e, e, n, t}
	z[3478] = Zone{"gov.mr", &z[955], x, 0, e, e, w{"dns.mauritania.mr", "dns2.mauritania.mr"}, n, n, n, e, e, n, f}
	z[3479] = Zone{"org.mr", &z[955], x, 0, e, e, w{"ns-mr.afrinic.net", "ns-mr.nic.fr", "ns-mr.nic.tn", "ns1.nic.mr", "ns2.nic.mr", "ns3.nic.mr"}, n, n, n, e, e, n, t}
}

//go:noinline
func i3479() {
	z[3480] = Zone{"perso.mr", &z[955], x, 0, e, e, w{"ns-mr.afrinic.net", "ns-mr.nic.fr", "ns-mr.nic.tn", "ns1.nic.mr", "ns2.nic.mr", "ns3.nic.mr"}, n, n, n, e, e, n, t}
	z[3481] = Zone{"co.ms", &z[958], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3482] = Zone{"com.ms", &z[958], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3483] = Zone{"net.ms", &z[958], x, 0, e, e, w{"ns1.idnscan.net", "ns6.idnscan.net"}, n, n, n, e, e, n, t}
	z[3484] = Zone{"org.ms", &z[958], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3485] = Zone{"com.mt", &z[960], x, 0, e, e, w{"a.ns.mt", "b.ns.mt", "f.ns.mt", "p.ns.mt"}, n, n, n, e, e, n, f}
	z[3486] = Zone{"edu.mt", &z[960], x, 0, e, e, w{"a.ns.mt", "b.ns.mt", "f.ns.mt", "p.ns.mt"}, n, n, n, e, e, n, f}
	z[3487] = Zone{"gov.mt", &z[960], x, 0, e, e, w{"ns1-09.azure-dns.com", "ns2-09.azure-dns.net", "ns3-09.azure-dns.org", "ns4-09.azure-dns.info"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3487() {
	z[3488] = Zone{"net.mt", &z[960], x, 0, e, e, w{"a.ns.mt", "b.ns.mt", "f.ns.mt", "p.ns.mt"}, n, n, n, e, e, n, f}
	z[3489] = Zone{"org.mt", &z[960], x, 0, e, e, w{"a.ns.mt", "b.ns.mt", "f.ns.mt", "p.ns.mt"}, n, n, n, e, e, n, f}
	z[3490] = Zone{"24.mu", &z[964], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.mu", "95.217.58.108"}, n, n, e, e, n, t}
	z[3491] = Zone{"ac.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, f}
	z[3492] = Zone{"co.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, t}
	z[3493] = Zone{"com.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, f}
	z[3494] = Zone{"gov.mu", &z[964], x, 0, e, e, w{"ns0.gov.mu", "ns1.gov.mu", "ns2.gov.mu", "ns3.gov.mu", "ns4.gov.mu"}, w{"::1"}, n, n, e, e, n, f}
	z[3495] = Zone{"net.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3495() {
	z[3496] = Zone{"nom.mu", &z[964], x, 0, e, e, w{"ns1.mu-dns.net", "ns2.mu-dns.net"}, w{"142.54.186.122", "172.96.172.79", "185.150.190.161", "66.187.75.46"}, n, n, e, e, n, f}
	z[3497] = Zone{"or.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, f}
	z[3498] = Zone{"org.mu", &z[964], x, 0, e, e, w{"anycast1.irondns.net", "fork.sth.dnsnode.net", "mu1.dyntld.net", "mu2.dyntld.net", "mu3.dyntld.net", "mu4.dyntld.net", "udns1.tld.mu", "udns2.tld.mu"}, n, n, n, e, e, n, f}
	z[3499] = Zone{"aero.mv", &z[971], x, 0, e, "http://www.brunet.bn/products_webrelated_domain_main.htm", w{"mv-ns.anycast.pch.net", "ns-cache01.dhivehinet.net.mv", "ns.aero.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[3500] = Zone{"biz.mv", &z[971], x, 0, e, e, w{"webns.mv"}, n, n, n, e, e, n, f}
	z[3501] = Zone{"com.mv", &z[971], x, 0, e, e, w{"mv-ns.anycast.pch.net", "ns.com.mv", "ns.dhivehinet.net.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[3502] = Zone{"coop.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3503] = Zone{"edu.mv", &z[971], x, 0, e, e, w{"mv-ns.anycast.pch.net", "ns.edu.mv", "ns1.edu.mv", "ns2.edu.mv"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3503() {
	z[3504] = Zone{"gov.mv", &z[971], x, 0, e, e, w{"mv-ns.anycast.pch.net", "ns.dhivehinet.net.mv", "ns.gov.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[3505] = Zone{"info.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3506] = Zone{"int.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3507] = Zone{"mil.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3508] = Zone{"museum.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3509] = Zone{"name.mv", &z[971], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3510] = Zone{"net.mv", &z[971], x, 0, e, e, w{"mv-ns.anycast.pch.net", "ns.net.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[3511] = Zone{"org.mv", &z[971], x, 0, e, e, w{"mv-ns.anycast.pch.net", "ns.org.mv", "ns2.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3511() {
	z[3512] = Zone{"pro.mv", &z[971], x, 0, e, e, w{"host13.dhivehinet.net.mv"}, n, n, n, e, e, n, f}
	z[3513] = Zone{"ac.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "domwe.sdn.mw", "mw-e.ns.nic.cz", "mw.cctld.authdns.ripe.net", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3514] = Zone{"co.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "d.ext.nic.cz", "domwe.sdn.mw", "mw-e.ns.nic.cz", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3515] = Zone{"com.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "d.ext.nic.cz", "domwe.sdn.mw", "mw-e.ns.nic.cz", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3516] = Zone{"coop.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "d.ext.nic.cz", "domwe.sdn.mw", "mw-e.ns.nic.cz", "pch1.nic.mw", "rip.psg.com", "rip.psg.mw"}, n, n, n, e, e, n, f}
	z[3517] = Zone{"edu.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "domwe.sdn.mw", "mw-e.ns.nic.cz", "mw.cctld.authdns.ripe.net", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3518] = Zone{"gov.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "domwe.sdn.mw", "mw-e.ns.nic.cz", "mw.cctld.authdns.ripe.net", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3519] = Zone{"int.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "d.ext.nic.cz", "domwe.sdn.mw", "mw-e.ns.nic.cz", "pch1.nic.mw", "rip.psg.com", "rip.psg.mw"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3519() {
	z[3520] = Zone{"museum.mw", &z[972], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3521] = Zone{"net.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "d.ext.nic.cz", "domwe.sdn.mw", "mw-e.ns.nic.cz", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3522] = Zone{"org.mw", &z[972], x, 0, e, e, w{"chambo.sdnp.org.mw", "domwe.sdn.mw", "mw-e.ns.nic.cz", "mw.cctld.authdns.ripe.net", "ns4.apnic.net", "pch1.nic.mw", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[3523] = Zone{"com.mx", &z[973], x, 0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, e, e, n, f}
	z[3524] = Zone{"edu.mx", &z[973], x, 0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, e, e, n, f}
	z[3525] = Zone{"gob.mx", &z[973], x, 0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, e, e, n, f}
	z[3526] = Zone{"net.mx", &z[973], x, 0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, e, e, n, f}
	z[3527] = Zone{"org.mx", &z[973], x, 0, e, e, w{"c.mx-ns.mx", "e.mx-ns.mx", "i.mx-ns.mx", "m.mx-ns.mx", "o.mx-ns.mx", "x.mx-ns.mx"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3527() {
	z[3528] = Zone{"com.my", &z[974], x, 0, e, "https://mynic.my/", w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, t}
	z[3529] = Zone{"edi.my", &z[974], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[3530] = Zone{"edu.my", &z[974], x, 0, e, e, w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[3531] = Zone{"gov.my", &z[974], x, 0, e, e, w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[3532] = Zone{"mil.my", &z[974], x, 0, e, e, w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[3533] = Zone{"name.my", &z[974], x, 0, e, "https://mynic.my/", w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[3534] = Zone{"net.my", &z[974], x, 0, e, "https://mynic.my/", w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
	z[3535] = Zone{"org.my", &z[974], x, 0, e, e, w{"a.mynic.centralnic-dns.com", "b.mynic.centralnic-dns.com", "b.nic.my", "b1.nic.my", "c.mynic.centralnic-dns.com", "d.mynic.centralnic-dns.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3535() {
	z[3536] = Zone{"ac.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3537] = Zone{"adv.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3538] = Zone{"co.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3539] = Zone{"edu.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3540] = Zone{"gov.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3541] = Zone{"net.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, t}
	z[3542] = Zone{"org.mz", &z[975], x, 0, e, e, w{"anyns.uem.mz", "dzowo.uem.mz", "ns-mz.afrinic.net", "oceano.uem.mz", "phloem.uoregon.edu", "zebra.uem.mz"}, n, n, n, e, e, n, f}
	z[3543] = Zone{"alt.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3543() {
	z[3544] = Zone{"cc.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3545] = Zone{"co.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
	z[3546] = Zone{"com.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
	z[3547] = Zone{"dr.na", &z[977], x, 0, e, e, w{"ns1.click.na", "ns2.click.na"}, n, n, n, e, e, n, f}
	z[3548] = Zone{"edu.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
	z[3549] = Zone{"in.na", &z[977], x, 0, e, e, w{"ns21.cloudns.net", "ns22.cloudns.net", "ns23.cloudns.net", "ns24.cloudns.net"}, n, n, n, e, e, n, f}
	z[3550] = Zone{"info.na", &z[977], x, 0, e, e, w{"clint.ns.cloudflare.com", "jocelyn.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[3551] = Zone{"mobi.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3551() {
	z[3552] = Zone{"name.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3553] = Zone{"net.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
	z[3554] = Zone{"of.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3555] = Zone{"or.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3556] = Zone{"org.na", &z[977], x, 0, e, e, w{"anyc2.irondns.net", "katima.omadhina.net", "na-ns.anycast.pch.net", "ns-na.afrinic.net"}, n, n, n, e, e, n, f}
	z[3557] = Zone{"pro.na", &z[977], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[3558] = Zone{"school.na", &z[977], x, 0, e, e, w{"beau.ns.cloudflare.com", "reza.ns.cloudflare.com"}, n, n, n, e, e, n, f}
	z[3559] = Zone{"tv.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3559() {
	z[3560] = Zone{"ws.na", &z[977], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3561] = Zone{"kyrylkov.name", &z[981], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3562] = Zone{"puig.name", &z[981], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3563] = Zone{"santos.name", &z[981], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3564] = Zone{"asso.nc", &z[987], x, 0, e, e, w{"any-ns1.nc", "ns1.nc", "ns2.nc"}, n, n, n, e, e, n, f}
	z[3565] = Zone{"com.nc", &z[987], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3566] = Zone{"nom.nc", &z[987], x, 0, e, e, w{"any-ns1.nc", "ns1.nc", "ns2.nc"}, n, n, n, e, e, n, f}
	z[3567] = Zone{"com.ne", &z[988], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3567() {
	z[3568] = Zone{"info.ne", &z[988], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3569] = Zone{"int.ne", &z[988], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3570] = Zone{"org.ne", &z[988], x, 0, e, e, w{"nafarko.netlink-niger.biz", "server1.netlink-niger.biz"}, n, n, n, e, e, n, f}
	z[3571] = Zone{"perso.ne", &z[988], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3572] = Zone{"1x.net", &z[990], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"1x.net", "95.217.58.108"}, n, n, e, e, n, t}
	z[3573] = Zone{"auz.net", &z[990], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com", "ns3.instradns.com"}, n, n, n, e, e, n, f}
	z[3574] = Zone{"example.net", &z[990], x, 0x2108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"a.iana-servers.net", "b.iana-servers.net"}, n, n, n, e, e, n, t}
	z[3575] = Zone{"gb.net", &z[990], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/gb.net/"}, t}
}

//go:noinline
func i3575() {
	z[3576] = Zone{"hu.net", &z[990], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183", "hu.net"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/hu.net/"}, f}
	z[3577] = Zone{"in.net", &z[990], x, 0, e, e, w{"ns1.nic.in.net", "ns2.nic.in.net", "ns3.nic.in.net", "ns4.nic.in.net"}, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/in.net/"}, f}
	z[3578] = Zone{"jp.net", &z[990], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"205.164.14.88"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/jp.net/"}, f}
	z[3579] = Zone{"ru.net", &z[990], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[3580] = Zone{"se.net", &z[990], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/se.net/"}, t}
	z[3581] = Zone{"uk.net", &z[990], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/uk.net/"}, t}
	z[3582] = Zone{"za.net", &z[990], x, 0, e, e, w{"igubu.saix.net", "ns0.is.co.za", "ns0.plig.net", "ns1.plig.net"}, n, n, n, "whois.za.net", e, n, t}
	z[3583] = Zone{"arts.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3583() {
	z[3584] = Zone{"com.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3585] = Zone{"firm.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3586] = Zone{"info.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3587] = Zone{"net.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3588] = Zone{"org.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3589] = Zone{"other.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3590] = Zone{"per.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3591] = Zone{"rec.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3591() {
	z[3592] = Zone{"store.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3593] = Zone{"us.nf", &z[1002], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "us.nf"}, n, n, e, e, n, t}
	z[3594] = Zone{"web.nf", &z[1002], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3595] = Zone{"com.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3596] = Zone{"edu.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3597] = Zone{"gov.ng", &z[1004], z[5119:5120], 0, e, e, n, n, n, n, e, e, n, f}
	z[3598] = Zone{"i.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3599] = Zone{"mil.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3599() {
	z[3600] = Zone{"mobi.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3601] = Zone{"name.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3602] = Zone{"net.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3603] = Zone{"org.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3604] = Zone{"sch.ng", &z[1004], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3605] = Zone{"ac.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3606] = Zone{"biz.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3607] = Zone{"co.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3607() {
	z[3608] = Zone{"com.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3609] = Zone{"edu.ni", &z[1007], x, 0, e, e, w{"auth01.ns.uu.net", "dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3610] = Zone{"gob.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3611] = Zone{"gov.ni", &z[1007], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3612] = Zone{"in.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3613] = Zone{"info.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3614] = Zone{"int.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, w{"72.5.65.111", "sinkhole.paloaltonetworks.com"}, n, n, e, e, n, f}
	z[3615] = Zone{"mil.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3615() {
	z[3616] = Zone{"net.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3617] = Zone{"nom.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3618] = Zone{"org.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3619] = Zone{"web.ni", &z[1007], x, 0, e, e, w{"dns-ext.nic.cr", "ns.ideay.net.ni", "ns.ni", "ns.uu.net", "ns2.ni", "ns3.ni"}, n, n, n, e, e, n, f}
	z[3620] = Zone{"24.nl", &z[1014], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.nl", "95.217.58.108"}, n, n, e, e, n, t}
	z[3621] = Zone{"co.nl", &z[1014], x, 0x200, e, e, w{"a.nic.co.nl", "b.nic.co.nl", "c.nic.co.nl", "d.nic.co.nl"}, n, n, n, e, e, w{"https://rdap.centralnic.com/co.nl/"}, f}
	z[3622] = Zone{"com.nl", &z[1014], x, 0x200, e, e, w{"ns1.com.nl", "ns2.com.nl"}, w{"185.85.211.136"}, n, n, e, e, n, f}
	z[3623] = Zone{"net.nl", &z[1014], x, 0x200, e, e, w{"ns1.com.nl", "ns2.com.nl"}, w{"185.85.211.136"}, n, n, e, e, n, f}
}

//go:noinline
func i3623() {
	z[3624] = Zone{"nom.nl", &z[1014], x, 0, e, e, w{"ns01.hostnet.nl", "ns02.hostnet.nl"}, n, n, n, e, e, n, f}
	z[3625] = Zone{"punt.nl", &z[1014], x, 0x200, e, "http://www.punt.nl/", w{"nsn1.mijndomein.nl", "nsn2.mijndomein.nl"}, w{"2a05:d018:964:c0a:a58f:6b32:e401:3d8a", "34.240.216.169", "hosting-1a.mijndomein-ws.nl"}, n, n, e, e, n, f}
	z[3626] = Zone{"co.no", &z[1015], x, 0, e, "https://www.norid.no/index.en.html", w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, e, e, w{"https://rdap.centralnic.com/co.no/"}, t}
	z[3627] = Zone{"fhs.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3628] = Zone{"folkebibl.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3629] = Zone{"fylkesbibl.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3630] = Zone{"gs.no", &z[1015], x, 0, e, e, w{"ns1.hyp.net", "ns2.hyp.net", "ns3.hyp.net"}, n, n, n, e, e, n, f}
	z[3631] = Zone{"idrett.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3631() {
	z[3632] = Zone{"museum.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3633] = Zone{"priv.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3634] = Zone{"uenorge.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3635] = Zone{"vgs.no", &z[1015], x, 0, e, e, w{"x.nic.no", "y.nic.no", "z.nic.no"}, n, n, n, e, e, n, f}
	z[3636] = Zone{"academy.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3637] = Zone{"accountants.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3638] = Zone{"actor.np", &z[1023], x, 0, e, "https://register.com.np/", w{"manaslu.mos.com.np", "np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np", "sungava.mos.com.np"}, n, n, n, e, e, n, f}
	z[3639] = Zone{"aero.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3639() {
	z[3640] = Zone{"agency.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3641] = Zone{"asia.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3642] = Zone{"associates.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3643] = Zone{"audio.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3644] = Zone{"bar.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3645] = Zone{"bargains.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3646] = Zone{"beer.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3647] = Zone{"bid.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3647() {
	z[3648] = Zone{"bike.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3649] = Zone{"bio.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3650] = Zone{"biz.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3651] = Zone{"black.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3652] = Zone{"blue.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "pch.nnic.np", "sec2.apnic.net"}, n, n, n, e, e, n, f}
	z[3653] = Zone{"boutique.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3654] = Zone{"build.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3655] = Zone{"builders.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3655() {
	z[3656] = Zone{"buzz.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3657] = Zone{"cab.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3658] = Zone{"camera.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3659] = Zone{"camp.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3660] = Zone{"capital.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3661] = Zone{"cards.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3662] = Zone{"care.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3663] = Zone{"careers.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3663() {
	z[3664] = Zone{"cash.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3665] = Zone{"catering.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3666] = Zone{"center.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3667] = Zone{"ceo.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3668] = Zone{"christmas.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3669] = Zone{"clinic.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3670] = Zone{"clothing.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3671] = Zone{"club.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3671() {
	z[3672] = Zone{"codes.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3673] = Zone{"coffee.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3674] = Zone{"college.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3675] = Zone{"com.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3676] = Zone{"community.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3677] = Zone{"company.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3678] = Zone{"computer.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3679] = Zone{"cool.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3679() {
	z[3680] = Zone{"coop.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3681] = Zone{"country.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3682] = Zone{"credit.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3683] = Zone{"creditcard.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3684] = Zone{"dental.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3685] = Zone{"diamonds.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3686] = Zone{"edu.np", &z[1023], x, 0, e, "https://register.com.np/", w{"lana.ns.cloudflare.com", "np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3687] = Zone{"email.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3687() {
	z[3688] = Zone{"engineering.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3689] = Zone{"estate.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3690] = Zone{"events.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3691] = Zone{"expert.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3692] = Zone{"finance.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3693] = Zone{"financial.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3694] = Zone{"fish.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3695] = Zone{"fishing.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3695() {
	z[3696] = Zone{"fitness.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3697] = Zone{"flights.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3698] = Zone{"florist.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3699] = Zone{"fund.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3700] = Zone{"furniture.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3701] = Zone{"futbol.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3702] = Zone{"gallery.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3703] = Zone{"gov.np", &z[1023], x, 0, e, "https://nitc.gov.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns-np.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3703() {
	z[3704] = Zone{"guitars.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3705] = Zone{"guru.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3706] = Zone{"hiphop.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3707] = Zone{"hiv.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3708] = Zone{"house.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3709] = Zone{"industries.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3710] = Zone{"info.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3711] = Zone{"ink.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3711() {
	z[3712] = Zone{"jobs.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3713] = Zone{"limited.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3714] = Zone{"link.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3715] = Zone{"management.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3716] = Zone{"marketing.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3717] = Zone{"media.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3718] = Zone{"menu.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3719] = Zone{"mil.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np.cctld.authdns.ripe.net", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3719() {
	z[3720] = Zone{"mobi.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3721] = Zone{"museum.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3722] = Zone{"name.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3723] = Zone{"net.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3724] = Zone{"ninja.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3725] = Zone{"onl.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3726] = Zone{"org.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3727] = Zone{"partners.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3727() {
	z[3728] = Zone{"parts.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3729] = Zone{"photo.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3730] = Zone{"photos.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3731] = Zone{"pics.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3732] = Zone{"pink.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3733] = Zone{"pro.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "np.cctld.authdns.ripe.net", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3734] = Zone{"productions.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3735] = Zone{"products.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3735() {
	z[3736] = Zone{"properties.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3737] = Zone{"pub.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3738] = Zone{"red.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3739] = Zone{"rentals.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3740] = Zone{"repair.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3741] = Zone{"rest.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3742] = Zone{"rocks.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3743] = Zone{"services.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3743() {
	z[3744] = Zone{"shiksha.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3745] = Zone{"shoes.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3746] = Zone{"social.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3747] = Zone{"solar.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3748] = Zone{"solutions.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3749] = Zone{"space.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3750] = Zone{"supplies.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3751] = Zone{"supply.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3751() {
	z[3752] = Zone{"support.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3753] = Zone{"surf.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3754] = Zone{"surgery.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3755] = Zone{"systems.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3756] = Zone{"tattoo.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3757] = Zone{"tax.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3758] = Zone{"technology.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3759] = Zone{"tel.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3759() {
	z[3760] = Zone{"tips.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3761] = Zone{"today.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net"}, n, n, n, e, e, n, f}
	z[3762] = Zone{"tools.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3763] = Zone{"town.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3764] = Zone{"trade.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3765] = Zone{"training.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3766] = Zone{"travel.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns-ext.vix.com", "ns4.apnic.net", "pch.nnic.np", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3767] = Zone{"university.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3767() {
	z[3768] = Zone{"vacations.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3769] = Zone{"ventures.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3770] = Zone{"villas.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3771] = Zone{"vision.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3772] = Zone{"vodka.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3773] = Zone{"voting.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3774] = Zone{"voyage.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3775] = Zone{"watch.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
}

//go:noinline
func i3775() {
	z[3776] = Zone{"webcam.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3777] = Zone{"wiki.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3778] = Zone{"works.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3779] = Zone{"wtf.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3780] = Zone{"xyz.np", &z[1023], x, 0, e, "https://register.com.np/", w{"np-ns.npix.net.np", "ns4.apnic.net", "pch.nnic.np", "sec2.apnic.net", "shikhar.mos.com.np"}, n, n, n, e, e, n, f}
	z[3781] = Zone{"zone.np", &z[1023], x, 0, e, "https://register.com.np/", n, n, n, n, e, e, n, f}
	z[3782] = Zone{"biz.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr"}, n, n, n, e, e, n, f}
	z[3783] = Zone{"com.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3783() {
	z[3784] = Zone{"edu.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr"}, n, n, n, e, e, n, f}
	z[3785] = Zone{"gov.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[3786] = Zone{"info.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr"}, n, n, n, e, e, n, f}
	z[3787] = Zone{"net.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[3788] = Zone{"org.nr", &z[1024], x, 0, e, e, w{"ns0.cenpac.net.nr", "ns1.cenpac.net.nr", "ns2.cenpac.net.nr", "phloem.uoregon.edu"}, n, n, n, e, e, n, f}
	z[3789] = Zone{"co.nu", &z[1028], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "co.nu"}, n, n, e, e, n, t}
	z[3790] = Zone{"com.nu", &z[1028], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "com.nu"}, n, n, e, e, n, t}
	z[3791] = Zone{"eu.nu", &z[1028], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "eu.nu"}, n, n, e, e, n, t}
}

//go:noinline
func i3791() {
	z[3792] = Zone{"ie.nu", &z[1028], x, 0, e, e, w{"brain.bb-online.co.uk", "ns1.bb-online.net", "ns2.bb-online.net", "ns5.bb-online.net"}, n, n, n, e, e, n, t}
	z[3793] = Zone{"info.nu", &z[1028], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "info.nu"}, n, n, e, e, n, t}
	z[3794] = Zone{"ac.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3795] = Zone{"co.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3796] = Zone{"cri.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3797] = Zone{"geek.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3798] = Zone{"gen.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3799] = Zone{"govt.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3799() {
	z[3800] = Zone{"health.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, t}
	z[3801] = Zone{"iwi.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3802] = Zone{"kiwi.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3803] = Zone{"maori.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3804] = Zone{"mil.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3805] = Zone{"net.nz", &z[1030], z[5120:5121], 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3806] = Zone{"org.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3807] = Zone{"parliament.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3807() {
	z[3808] = Zone{"school.nz", &z[1030], x, 0, e, e, w{"ns1.dns.net.nz", "ns2.dns.net.nz", "ns3.dns.net.nz", "ns4.dns.net.nz", "ns5.dns.net.nz", "ns6.dns.net.nz", "ns7.dns.net.nz"}, n, n, n, e, e, n, f}
	z[3809] = Zone{"ac.om", &z[1041], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3810] = Zone{"biz.om", &z[1041], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3811] = Zone{"co.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3812] = Zone{"com.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3813] = Zone{"edu.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3814] = Zone{"gov.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3815] = Zone{"med.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3815() {
	z[3816] = Zone{"mil.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3817] = Zone{"museum.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3818] = Zone{"net.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3819] = Zone{"org.om", &z[1041], x, 0, e, e, w{"cctld.alpha.aridns.net.au", "cctld.beta.aridns.net.au", "cctld.delta.aridns.net.au", "cctld.gamma.aridns.net.au", "ns1.registry.om", "ns2.registry.om"}, n, n, n, e, e, n, f}
	z[3820] = Zone{"pro.om", &z[1041], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3821] = Zone{"sch.om", &z[1041], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3822] = Zone{"ae.org", &z[1053], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183", "ae.org"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/ae.org/"}, t}
	z[3823] = Zone{"eu.org", &z[1053], z[5121:5173], 0, e, e, w{"dns.elm.net", "dns4.gandi.net", "gra.wolfhugel.eu", "hobbes.bsd-dk.dk", "ns.bortzmeyer.eu.org", "ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net", "wombat.billaud.eu.org"}, n, n, n, "whois.eu.org", e, n, t}
}

//go:noinline
func i3823() {
	z[3824] = Zone{"example.org", &z[1053], x, 0x2108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"a.iana-servers.net", "b.iana-servers.net"}, n, n, n, e, e, n, t}
	z[3825] = Zone{"hk.org", &z[1053], x, 0, e, e, w{"a.udrtld.net", "b.udrtld.net", "c.udrtld.net", "dns7.cloudns.net", "dns8.cloudns.net", "ns3.he.net"}, n, n, n, "whois.registry.hk.com", e, n, f}
	z[3826] = Zone{"js.org", &z[1053], x, 0, e, e, w{"miles.ns.cloudflare.com", "pam.ns.cloudflare.com"}, w{"162.0.217.23"}, n, n, e, e, n, t}
	z[3827] = Zone{"us.org", &z[1053], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, w{"54.153.56.183", "us.org"}, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/us.org/"}, f}
	z[3828] = Zone{"za.org", &z[1053], x, 0, e, e, w{"igubu.saix.net", "ns0.is.co.za", "ns0.plig.net", "ns1.plig.net"}, n, n, n, "whois.za.org", e, n, t}
	z[3829] = Zone{"abo.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3830] = Zone{"ac.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3831] = Zone{"com.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3831() {
	z[3832] = Zone{"edu.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3833] = Zone{"gob.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3834] = Zone{"ing.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3835] = Zone{"med.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3836] = Zone{"net.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3837] = Zone{"nom.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3838] = Zone{"org.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[3839] = Zone{"sld.pa", &z[1062], x, 0, e, e, w{"a.lactld.org", "dns-ext.nic.cr", "ns.dns.br", "ns.nic.pa", "ns.pa", "ns2.pa", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3839() {
	z[3840] = Zone{"com.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3841] = Zone{"edu.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3842] = Zone{"gob.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3843] = Zone{"gov.pe", &z[1078], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3844] = Zone{"mil.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3845] = Zone{"net.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3846] = Zone{"nom.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
	z[3847] = Zone{"org.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3847() {
	z[3848] = Zone{"sld.pe", &z[1078], x, 0, e, e, w{"a.lactld.org", "pch.rcp.pe", "pe1.dnsnode.net", "quipu.rcp.net.pe"}, n, n, n, e, e, n, t}
	z[3849] = Zone{"asso.pf", &z[1082], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3850] = Zone{"com.pf", &z[1082], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3851] = Zone{"edu.pf", &z[1082], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3852] = Zone{"gov.pf", &z[1082], x, 0, e, e, w{"ns1.gov.pf", "ns1.mana.pf", "ns2.gov.pf", "ns2.mana.pf", "reverse.gov.pf"}, n, n, n, e, e, n, f}
	z[3853] = Zone{"org.pf", &z[1082], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3854] = Zone{"ac.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3855] = Zone{"com.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3855() {
	z[3856] = Zone{"gov.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3857] = Zone{"mil.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3858] = Zone{"net.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3859] = Zone{"org.pg", &z[1084], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3860] = Zone{"com.ph", &z[1085], x, 0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, e, n, f}
	z[3861] = Zone{"edi.ph", &z[1085], x, 0, e, e, w{"ns1.dp47.parklogic.com", "ns2.dp47.parklogic.com", "ns3.dp47.parklogic.com", "ns4.dp47.parklogic.com", "ns5.dp47.parklogic.com"}, w{"2607:fad0:3801:4::1", "69.16.231.56"}, n, n, e, e, n, t}
	z[3862] = Zone{"edu.ph", &z[1085], x, 0, e, e, w{"evans.gov.ph", "gabriela.ph.net", "gomez.ph.net"}, n, n, n, e, e, n, f}
	z[3863] = Zone{"gov.ph", &z[1085], x, 0, e, e, w{"gabriela.ph.net", "gomez.ph.net", "ns1.gov.ph", "ns2.gov.ph"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3863() {
	z[3864] = Zone{"i.ph", &z[1085], x, 0, e, e, w{"nsi1.domains.ph", "nsi2.domains.ph"}, n, n, n, e, e, n, f}
	z[3865] = Zone{"mil.ph", &z[1085], x, 0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, e, n, f}
	z[3866] = Zone{"net.ph", &z[1085], x, 0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, e, n, f}
	z[3867] = Zone{"ngo.ph", &z[1085], x, 0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, e, n, f}
	z[3868] = Zone{"org.ph", &z[1085], x, 0, e, e, w{"1.ns.ph", "ns2.cuhk.edu.hk", "ns4.apnic.net", "ph.communitydns.net"}, w{"45.79.222.138"}, n, n, e, e, n, f}
	z[3869] = Zone{"biz.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3870] = Zone{"com.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3871] = Zone{"edu.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i3871() {
	z[3872] = Zone{"fam.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3873] = Zone{"gkp.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3874] = Zone{"gob.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3875] = Zone{"gog.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3876] = Zone{"gok.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3877] = Zone{"gon.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3878] = Zone{"gop.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3879] = Zone{"gos.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i3879() {
	z[3880] = Zone{"gov.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3881] = Zone{"net.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3882] = Zone{"org.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3883] = Zone{"web.pk", &z[1106], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[3884] = Zone{"agro.pl", &z[1107], x, 0, e, "https://domain.by/", w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3885] = Zone{"aid.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3886] = Zone{"art.pl", &z[1107], x, 0, e, e, w{"dns14.ovh.net", "ns14.ovh.net"}, w{"146.59.33.164"}, n, n, e, e, n, f}
	z[3887] = Zone{"atm.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3887() {
	z[3888] = Zone{"augustow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3889] = Zone{"auto.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3890] = Zone{"babia-gora.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3891] = Zone{"bedzin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3892] = Zone{"beskidy.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3893] = Zone{"bialowieza.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, t}
	z[3894] = Zone{"bialystok.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3895] = Zone{"bielawa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3895() {
	z[3896] = Zone{"bieszczady.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3897] = Zone{"biz.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3898] = Zone{"boleslawiec.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3899] = Zone{"bydgoszcz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3900] = Zone{"bytom.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3901] = Zone{"ciesyn.pl", &z[1107], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3902] = Zone{"cieszyn.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3903] = Zone{"co.pl", &z[1107], x, 0, e, e, w{"ns1.co.pl", "ns2.co.pl"}, n, n, n, "whois.co.pl", e, n, t}
}

//go:noinline
func i3903() {
	z[3904] = Zone{"com.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3905] = Zone{"czeladz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3906] = Zone{"czest.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3907] = Zone{"dlugoleka.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3908] = Zone{"edu.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3909] = Zone{"elblag.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3910] = Zone{"elk.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3911] = Zone{"gda.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "ns.tpnet.pl", "ns1.task.gda.pl", "ns2.task.gda.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3911() {
	z[3912] = Zone{"gdansk.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "ns.tpnet.pl", "ns1.task.gda.pl", "ns2.task.gda.pl"}, n, n, n, e, e, n, f}
	z[3913] = Zone{"glogow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3914] = Zone{"gmina.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3915] = Zone{"gniezno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3916] = Zone{"gniezon.pl", &z[1107], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3917] = Zone{"gorlice.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3918] = Zone{"gov.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3919] = Zone{"grajewo.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3919() {
	z[3920] = Zone{"gsm.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3921] = Zone{"ilawa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3922] = Zone{"info.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3923] = Zone{"jaworzno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3924] = Zone{"jaworzono.pl", &z[1107], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3925] = Zone{"jelenia-gora.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3926] = Zone{"jgora.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3927] = Zone{"kalisz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3927() {
	z[3928] = Zone{"karpacz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3929] = Zone{"kartuzy.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3930] = Zone{"kaszuby.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3931] = Zone{"katowice.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3932] = Zone{"kazimierz-dolny.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3933] = Zone{"kepno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3934] = Zone{"ketrzyn.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3935] = Zone{"klodzko.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3935() {
	z[3936] = Zone{"kobierzyce.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3937] = Zone{"kolobrzeg.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3938] = Zone{"konin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3939] = Zone{"konskowola.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3940] = Zone{"krakow.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "dns.fuw.edu.pl", "info.cyf-kr.edu.pl", "nms.cyf-kr.edu.pl"}, n, n, n, e, e, n, f}
	z[3941] = Zone{"ksazuby.pl", &z[1107], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[3942] = Zone{"kutno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3943] = Zone{"lapy.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3943() {
	z[3944] = Zone{"lebork.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3945] = Zone{"legnica.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3946] = Zone{"lezajsk.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3947] = Zone{"limanowa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3948] = Zone{"lodz.pl", &z[1107], x, 0x480, e, e, w{"dns.man.lodz.pl", "dns2.man.lodz.pl", "ns1.task.gda.pl", "ns1.tpnet.pl"}, n, n, n, e, e, n, f}
	z[3949] = Zone{"lomza.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3950] = Zone{"lowicz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3951] = Zone{"lubin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3951() {
	z[3952] = Zone{"lublin.pl", &z[1107], x, 0x480, e, e, w{"bilbo.nask.org.pl", "ns1.lublin.pl", "ns2.lublin.pl"}, n, n, n, e, e, n, f}
	z[3953] = Zone{"lukow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3954] = Zone{"mail.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3955] = Zone{"malbork.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3956] = Zone{"malopolska.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3957] = Zone{"mazowsze.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3958] = Zone{"mazury.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3959] = Zone{"media.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3959() {
	z[3960] = Zone{"miasta.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3961] = Zone{"mielec.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3962] = Zone{"mielno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3963] = Zone{"mil.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3964] = Zone{"mragowo.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3965] = Zone{"naklo.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3966] = Zone{"net.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3967] = Zone{"ngo.pl", &z[1107], x, 0, e, e, w{"ns1.ngo.pl", "ns2.ngo.pl", "ns3.ngo.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3967() {
	z[3968] = Zone{"nieruchomosci.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3969] = Zone{"nom.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3970] = Zone{"nowaruda.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3971] = Zone{"nysa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3972] = Zone{"olawa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3973] = Zone{"olecko.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3974] = Zone{"olkusz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3975] = Zone{"olsztyn.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3975() {
	z[3976] = Zone{"opoczno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3977] = Zone{"opole.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3978] = Zone{"org.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3979] = Zone{"ostroda.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3980] = Zone{"ostroleka.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3981] = Zone{"ostrowiec.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3982] = Zone{"ostrowwlkp.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3983] = Zone{"pc.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3983() {
	z[3984] = Zone{"pila.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3985] = Zone{"pisz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3986] = Zone{"podhale.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3987] = Zone{"podlasie.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3988] = Zone{"polkowice.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3989] = Zone{"pomorskie.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3990] = Zone{"pomorze.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3991] = Zone{"powiat.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3991() {
	z[3992] = Zone{"poznan.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "dns.amu.edu.pl", "dns3.amu.edu.pl", "kirdan.nask.net.pl"}, n, n, n, e, e, n, f}
	z[3993] = Zone{"priv.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3994] = Zone{"prochowice.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3995] = Zone{"pruszkow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3996] = Zone{"przeworsk.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3997] = Zone{"pulawy.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3998] = Zone{"radom.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[3999] = Zone{"rawa-maz.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i3999() {
	z[4000] = Zone{"realestate.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4001] = Zone{"rel.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4002] = Zone{"rybnik.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4003] = Zone{"rzeszow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4004] = Zone{"sanok.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4005] = Zone{"sejny.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4006] = Zone{"sex.pl", &z[1107], x, 0x1, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4007] = Zone{"shop.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4007() {
	z[4008] = Zone{"sklep.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4009] = Zone{"skoczow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4010] = Zone{"slask.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4011] = Zone{"slupsk.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4012] = Zone{"sos.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4013] = Zone{"sosnowiec.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4014] = Zone{"stalowa-wola.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4015] = Zone{"starachowice.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4015() {
	z[4016] = Zone{"stargard.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4017] = Zone{"suwalki.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4018] = Zone{"swidnica.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4019] = Zone{"swiebodzin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4020] = Zone{"swinoujscie.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4021] = Zone{"szczecin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4022] = Zone{"szczytno.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4023] = Zone{"szkola.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4023() {
	z[4024] = Zone{"targi.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4025] = Zone{"tarnobrzeg.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4026] = Zone{"tgory.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4027] = Zone{"tm.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4028] = Zone{"torun.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "blackbox.uci.uni.torun.pl", "flis.man.torun.pl", "koala.uci.uni.torun.pl"}, n, n, n, e, e, n, f}
	z[4029] = Zone{"tourism.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4030] = Zone{"travel.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4031] = Zone{"turek.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4031() {
	z[4032] = Zone{"turystyka.pl", &z[1107], x, 0, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4033] = Zone{"tychy.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4034] = Zone{"ustka.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4035] = Zone{"walbrzych.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4036] = Zone{"warmia.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4037] = Zone{"warszawa.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4038] = Zone{"waw.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4039] = Zone{"wegrow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4039() {
	z[4040] = Zone{"wielun.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4041] = Zone{"wlocl.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4042] = Zone{"wloclawek.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4043] = Zone{"wodzislaw.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4044] = Zone{"wolomin.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4045] = Zone{"wroc.pl", &z[1107], x, 0, e, e, w{"bilbo.nask.org.pl", "kirdan.nask.net.pl", "ldhpux.immt.pwr.wroc.pl", "ns1.net.icm.edu.pl", "sun2.pwr.wroc.pl", "wask.wask.wroc.pl"}, n, n, n, e, e, n, f}
	z[4046] = Zone{"wroclaw.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4047] = Zone{"zachpomor.pl", &z[1107], x, 0x480, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4047() {
	z[4048] = Zone{"zagan.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4049] = Zone{"zarow.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4050] = Zone{"zgora.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4051] = Zone{"zgorzelec.pl", &z[1107], x, 0x84, e, e, w{"a-dns.pl", "b-dns.pl", "d-dns.pl", "f-dns.pl", "h-dns.pl", "j-dns.pl"}, n, n, n, e, e, n, f}
	z[4052] = Zone{"co.pn", &z[1114], x, 0, e, "https://nic.pn/", w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, e, n, t}
	z[4053] = Zone{"edu.pn", &z[1114], x, 0x200, e, e, w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, e, n, t}
	z[4054] = Zone{"eu.pn", &z[1114], x, 0x200, e, e, w{"ns1.he.net", "ns2.he.net", "ns3.he.net", "ns4.he.net", "ns5.he.net"}, n, n, n, e, e, n, t}
	z[4055] = Zone{"gov.pn", &z[1114], x, 0, e, e, w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4055() {
	z[4056] = Zone{"in.pn", &z[1114], x, 0x200, e, e, w{"ns-1425.awsdns-50.org", "ns-1677.awsdns-17.co.uk", "ns-441.awsdns-55.com", "ns-906.awsdns-49.net"}, n, n, n, e, e, n, t}
	z[4057] = Zone{"it.pn", &z[1114], x, 0x200, e, e, w{"ns1.he.net", "ns2.he.net", "ns3.he.net", "ns4.he.net", "ns5.he.net"}, n, n, n, e, e, n, t}
	z[4058] = Zone{"net.pn", &z[1114], x, 0, e, "https://nic.pn/", w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, e, n, t}
	z[4059] = Zone{"org.pn", &z[1114], x, 0, e, "https://nic.pn/", w{"dns1.nominetdns.uk", "dns2.nominetdns.uk", "dns3.nominetdns.uk", "dns4.nominetdns.uk", "dnsa.nominetdns.uk", "dnsb.nominetdns.uk", "dnsc.nominetdns.uk", "dnsd.nominetdns.uk"}, n, n, n, e, e, n, t}
	z[4060] = Zone{"ac.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4061] = Zone{"ag.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4062] = Zone{"ai.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4063] = Zone{"alumni.pr", &z[1122], x, 0, e, e, w{"descartes.nic.pr", "golomb.nic.pr", "pascal.nic.pr", "pr-dns.denic.de", "pr-ns.anycast.pch.net"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4063() {
	z[4064] = Zone{"at.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4065] = Zone{"au.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4066] = Zone{"aw.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4067] = Zone{"bb.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4068] = Zone{"biz.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4069] = Zone{"bm.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4070] = Zone{"br.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4071] = Zone{"bs.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4071() {
	z[4072] = Zone{"ch.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4073] = Zone{"cl.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4074] = Zone{"cn.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4075] = Zone{"com.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4076] = Zone{"de.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4077] = Zone{"dm.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4078] = Zone{"do.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4079] = Zone{"edu.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4079() {
	z[4080] = Zone{"en.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4081] = Zone{"es.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4082] = Zone{"est.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4083] = Zone{"eu.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4084] = Zone{"exalumnos.pr", &z[1122], x, 0, e, e, w{"descartes.nic.pr", "golomb.nic.pr", "pascal.nic.pr", "pr-dns.denic.de", "pr-ns.anycast.pch.net"}, n, n, n, e, e, n, t}
	z[4085] = Zone{"fr.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4086] = Zone{"gd.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4087] = Zone{"gov.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4087() {
	z[4088] = Zone{"gp.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4089] = Zone{"ht.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4090] = Zone{"info.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4091] = Zone{"isla.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4092] = Zone{"it.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4093] = Zone{"jm.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4094] = Zone{"jp.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4095] = Zone{"ky.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4095() {
	z[4096] = Zone{"lc.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4097] = Zone{"ms.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4098] = Zone{"mx.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4099] = Zone{"name.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4100] = Zone{"net.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4101] = Zone{"nl.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4102] = Zone{"nom.pr", &z[1122], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4103] = Zone{"org.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4103() {
	z[4104] = Zone{"pro.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4105] = Zone{"prof.pr", &z[1122], x, 0, e, e, w{"a0.pr.afilias-nst.info", "a2.pr.afilias-nst.info", "b0.pr.afilias-nst.org", "b2.pr.afilias-nst.org", "c0.pr.afilias-nst.info", "d0.pr.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4106] = Zone{"tc.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4107] = Zone{"tt.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4108] = Zone{"uk.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4109] = Zone{"us.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4110] = Zone{"vg.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4111] = Zone{"vi.pr", &z[1122], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4111() {
	z[4112] = Zone{"aaa.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4113] = Zone{"aca.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4114] = Zone{"acct.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4115] = Zone{"arc.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4116] = Zone{"avocat.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4117] = Zone{"bar.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4118] = Zone{"bus.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4119] = Zone{"cfp.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4119() {
	z[4120] = Zone{"chi.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4121] = Zone{"chiro.pro", &z[1127], x, 0, e, e, w{"dns111.ovh.net", "ns111.ovh.net"}, n, n, n, e, e, n, f}
	z[4122] = Zone{"cpa.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4123] = Zone{"dds.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4124] = Zone{"den.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4125] = Zone{"dent.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4126] = Zone{"ed.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4127] = Zone{"eng.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4127() {
	z[4128] = Zone{"jur.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4129] = Zone{"law.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4130] = Zone{"med.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4131] = Zone{"min.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4132] = Zone{"nitr.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4133] = Zone{"nur.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4134] = Zone{"nurse.pro", &z[1127], x, 0, e, e, w{"park1.encirca.net", "park2.encirca.net"}, n, n, n, e, e, n, f}
	z[4135] = Zone{"opt.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4135() {
	z[4136] = Zone{"pa.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4137] = Zone{"pha.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4138] = Zone{"pharma.pro", &z[1127], x, 0, e, e, w{"ns1.sedoparking.com", "ns2.sedoparking.com"}, w{"64.190.63.111"}, n, n, e, e, n, f}
	z[4139] = Zone{"pod.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4140] = Zone{"pr.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4141] = Zone{"prof.pro", &z[1127], x, 0, e, e, w{"standard3.encirca.net", "standard4.encirca.net"}, n, n, n, e, e, n, f}
	z[4142] = Zone{"prx.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4143] = Zone{"pt.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4143() {
	z[4144] = Zone{"recht.pro", &z[1127], x, 0, e, e, w{"a0.pro.afilias-nst.info", "a2.pro.afilias-nst.info", "b0.pro.afilias-nst.org", "b2.pro.afilias-nst.org", "c0.pro.afilias-nst.info", "d0.pro.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4145] = Zone{"rel.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4146] = Zone{"teach.pro", &z[1127], x, 0, e, e, w{"curitiba.ns.porkbun.com", "fortaleza.ns.porkbun.com", "maceio.ns.porkbun.com", "salvador.ns.porkbun.com"}, w{"44.227.65.245", "44.227.76.166", "pixie.porkbun.com"}, n, n, e, e, n, f}
	z[4147] = Zone{"vet.pro", &z[1127], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4148] = Zone{"com.ps", &z[1138], x, 0, e, "https://www.pnina.ps/registration-policy/", w{"ns1.pnina.ps", "ns2.pnina.ps", "ote.pnina.ps", "ps-ns.anycast.pch.net", "ps.cctld.authdns.ripe.net", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[4149] = Zone{"edu.ps", &z[1138], x, 0, e, "https://www.tonic.to/", w{"dns1.gov.ps", "ns1.pnina.ps", "ns2.pnina.ps"}, n, n, n, e, e, n, f}
	z[4150] = Zone{"gov.ps", &z[1138], x, 0, e, e, w{"dns1.gov.ps", "dns3.gov.ps", "ns1.gov.ps", "ns1.palgov.net", "ns2.palgov.net"}, n, n, n, e, e, n, f}
	z[4151] = Zone{"mobi.ps", &z[1138], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "mobi.ps"}, n, n, e, e, n, t}
}

//go:noinline
func i4151() {
	z[4152] = Zone{"mun.ps", &z[1138], x, 0, e, e, w{"dns1.gov.ps", "ns1.gov.ps", "ns2.palgov.net", "ns2.pnina.ps"}, n, n, n, e, e, n, f}
	z[4153] = Zone{"net.ps", &z[1138], x, 0, e, e, w{"ns1.pnina.ps", "ns2.pnina.ps", "ote.pnina.ps", "ps-ns.anycast.pch.net", "ps.cctld.authdns.ripe.net", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[4154] = Zone{"org.ps", &z[1138], x, 0, e, e, w{"ns1.pnina.ps", "ns2.pnina.ps", "ote.pnina.ps", "ps-ns.anycast.pch.net", "ps.cctld.authdns.ripe.net", "rip.psg.com"}, n, n, n, e, e, n, f}
	z[4155] = Zone{"plo.ps", &z[1138], x, 0, e, e, w{"dns3.gov.ps", "dns4.pna.ps", "dns5.gov.ps"}, n, n, n, e, e, n, t}
	z[4156] = Zone{"sch.ps", &z[1138], x, 0, e, e, w{"dns1.gov.ps", "ns1.pnina.ps", "ns2.pnina.ps"}, n, n, n, e, e, n, f}
	z[4157] = Zone{"co.pt", &z[1139], x, 0, e, e, w{"ns61.plako.net", "ns62.plako.net"}, w{"185.98.250.164"}, n, n, e, e, n, t}
	z[4158] = Zone{"com.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4159] = Zone{"edu.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4159() {
	z[4160] = Zone{"gov.pt", &z[1139], x, 0, e, e, w{"a.dns.pt", "dns1.gov.pt", "europe1.dnsnode.net", "ns02.fccn.pt", "nsp.dnsnode.net"}, n, n, n, e, e, n, f}
	z[4161] = Zone{"int.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4162] = Zone{"net.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4163] = Zone{"nome.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4164] = Zone{"org.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4165] = Zone{"publ.pt", &z[1139], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4166] = Zone{"belau.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4167] = Zone{"co.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4167() {
	z[4168] = Zone{"com.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4169] = Zone{"ed.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4170] = Zone{"edu.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4171] = Zone{"go.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4172] = Zone{"gov.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4173] = Zone{"ne.pw", &z[1141], x, 0, e, e, w{"a.dnspod.com", "b.dnspod.com", "c.dnspod.com"}, n, n, n, e, e, n, t}
	z[4174] = Zone{"net.pw", &z[1141], x, 0, e, e, w{"ns1390.ztomy.com", "ns2390.ztomy.com"}, w{"127.0.0.1"}, n, n, e, e, n, f}
	z[4175] = Zone{"or.pw", &z[1141], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4175() {
	z[4176] = Zone{"org.pw", &z[1141], x, 0, e, e, w{"ns1390.ztomy.com", "ns2390.ztomy.com"}, w{"127.0.0.1"}, n, n, e, e, n, f}
	z[4177] = Zone{"com.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4178] = Zone{"coop.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4179] = Zone{"edu.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4180] = Zone{"gov.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4181] = Zone{"mil.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4182] = Zone{"net.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4183] = Zone{"org.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "l.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4183() {
	z[4184] = Zone{"una.py", &z[1143], x, 0, e, e, w{"b.dns.py", "c.dns.py", "p.dns.py", "u.dns.py"}, n, n, n, e, e, n, f}
	z[4185] = Zone{"com.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, f}
	z[4186] = Zone{"edu.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, t}
	z[4187] = Zone{"gov.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, t}
	z[4188] = Zone{"mil.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, t}
	z[4189] = Zone{"name.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, f}
	z[4190] = Zone{"net.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, f}
	z[4191] = Zone{"org.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4191() {
	z[4192] = Zone{"sch.qa", &z[1144], x, 0, e, e, w{"a.registry.qa", "b.registry.qa", "c.registry.qa", "d.registry.qa", "e.registry.qa", "f.registry.qa", "g.registry.qa", "h.registry.qa", "i.registry.qa"}, n, n, n, e, e, n, t}
	z[4193] = Zone{"asso.re", &z[1154], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4194] = Zone{"com.re", &z[1154], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4195] = Zone{"nom.re", &z[1154], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4196] = Zone{"arts.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4197] = Zone{"co.ro", &z[1189], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, e, e, n, f}
	z[4198] = Zone{"com.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4199] = Zone{"firm.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4199() {
	z[4200] = Zone{"info.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4201] = Zone{"ne.ro", &z[1189], x, 0, e, e, w{"ns1.eunoc.net", "ns2.eunoc.net", "ns3.eunoc.net"}, n, n, n, e, e, n, f}
	z[4202] = Zone{"nom.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4203] = Zone{"nt.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4204] = Zone{"or.ro", &z[1189], x, 0, e, e, w{"ns1.eunoc.net", "ns2.eunoc.net", "ns3.eunoc.net"}, n, n, n, e, e, n, f}
	z[4205] = Zone{"org.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4206] = Zone{"rec.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4207] = Zone{"ro.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4207() {
	z[4208] = Zone{"sa.ro", &z[1189], x, 0, e, e, w{"ns1.eunoc.net", "ns2.eunoc.net", "ns3.eunoc.net"}, n, n, n, e, e, n, f}
	z[4209] = Zone{"srl.ro", &z[1189], x, 0, e, e, w{"ns1.eunoc.net", "ns2.eunoc.net", "ns3.eunoc.net"}, n, n, n, e, e, n, f}
	z[4210] = Zone{"store.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4211] = Zone{"tm.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4212] = Zone{"www.ro", &z[1189], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4213] = Zone{"ac.rs", &z[1198], x, 0, e, e, w{"ban.junis.ni.ac.rs", "gaea.rcub.bg.ac.rs", "ns.etf.bg.ac.rs", "ns.rcub.bg.ac.rs", "ns.unic.kg.ac.rs", "ns1.uns.ac.rs", "ns2.iif.hu", "odisej.telekom.rs"}, n, n, n, e, e, n, t}
	z[4214] = Zone{"co.rs", &z[1198], x, 0, e, e, w{"a.nic.rs", "b.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, n, e, e, n, f}
	z[4215] = Zone{"edu.rs", &z[1198], x, 0, e, e, w{"a.nic.rs", "b.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4215() {
	z[4216] = Zone{"gov.rs", &z[1198], x, 0, e, e, w{"dagda1.isp.telekom.rs", "gov1.nic.rs", "gov2.nic.rs", "ns1.gov.rs", "ns2.gov.rs", "ns3.gov.rs", "ns4.gov.rs"}, n, n, n, e, e, n, f}
	z[4217] = Zone{"in.rs", &z[1198], x, 0, e, e, w{"a.nic.rs", "b.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, n, e, e, n, f}
	z[4218] = Zone{"org.rs", &z[1198], x, 0, e, e, w{"a.nic.rs", "b.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, n, e, e, n, f}
	z[4219] = Zone{"ac.ru", &z[1200], x, 0, e, e, w{"ns-soa.darenet.dk", "ns1.free.net", "ns2.free.net"}, n, n, n, e, e, n, f}
	z[4220] = Zone{"adygeya.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4221] = Zone{"bashkiria.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4222] = Zone{"bir.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4223] = Zone{"cbg.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4223() {
	z[4224] = Zone{"com.ru", &z[1200], x, 0, e, e, w{"ns3-com.nic.ru", "ns4-com.nic.ru", "ns8-com.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, f}
	z[4225] = Zone{"dagestan.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4226] = Zone{"edu.ru", &z[1200], x, 0, e, e, w{"ns.informika.ru", "ns.msu.ru", "ns.runnet.ru", "ns2.free.net", "ns2.informika.ru"}, n, n, n, "whois.informika.ru", e, n, f}
	z[4227] = Zone{"gov.ru", &z[1200], x, 0, e, e, w{"acl.dns.ripn.net", "bcl.dns.ripn.net", "ccl.dns.ripn.net", "ns.gov.ru", "ns2.gov.ru"}, n, n, n, e, e, n, f}
	z[4228] = Zone{"grozny.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4229] = Zone{"int.ru", &z[1200], x, 0, e, e, w{"ns.macomnet.ru", "ns2.macomnet.ru", "ns3.macomnet.ru"}, n, n, n, e, e, n, f}
	z[4230] = Zone{"kalmykia.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4231] = Zone{"kustanai.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4231() {
	z[4232] = Zone{"marine.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4233] = Zone{"mil.ru", &z[1200], x, 0, e, e, w{"ns1.mil.ru", "ns2.mil.ru", "ns4.mil.ru"}, n, n, n, e, e, n, f}
	z[4234] = Zone{"mordovia.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4235] = Zone{"msk.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"185.76.145.93"}, n, n, e, e, n, f}
	z[4236] = Zone{"mytis.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4237] = Zone{"nalchik.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4238] = Zone{"net.ru", &z[1200], x, 0, e, e, w{"a.dns-servers.net.ru", "b.dns-servers.net.ru", "d.dns-servers.net.ru", "e.dns-servers.net.ru", "f.dns-servers.net.ru"}, n, n, n, e, e, n, f}
	z[4239] = Zone{"nov.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4239() {
	z[4240] = Zone{"org.ru", &z[1200], x, 0, e, e, w{"a.dns-servers.net.ru", "b.dns-servers.net.ru", "d.dns-servers.net.ru", "e.dns-servers.net.ru", "f.dns-servers.net.ru"}, n, n, n, e, e, n, f}
	z[4241] = Zone{"pp.ru", &z[1200], x, 0, e, e, w{"a.dns-servers.net.ru", "b.dns-servers.net.ru", "d.dns-servers.net.ru", "e.dns-servers.net.ru", "f.dns-servers.net.ru"}, n, n, n, e, e, n, f}
	z[4242] = Zone{"pyatigorsk.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4243] = Zone{"relarn.ru", &z[1200], x, 0, e, e, w{"ns1.timeweb.ru", "ns2.timeweb.ru", "ns3.timeweb.org", "ns4.timeweb.org"}, n, n, n, e, e, n, f}
	z[4244] = Zone{"spb.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, f}
	z[4245] = Zone{"vladikavkaz.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4246] = Zone{"vladimir.ru", &z[1200], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4247] = Zone{"ac.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4247() {
	z[4248] = Zone{"co.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4249] = Zone{"com.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4250] = Zone{"edu.rw", &z[1204], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns-rw.afrinic.net", "ns1.ricta.org.rw", "pch.ricta.org.rw"}, n, n, n, e, e, n, f}
	z[4251] = Zone{"gouv.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4252] = Zone{"gov.rw", &z[1204], x, 0, e, e, w{"ns1.rdb.rw", "ns2.aos.rw", "ns3.aos.rw"}, n, n, n, e, e, n, f}
	z[4253] = Zone{"int.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4254] = Zone{"mil.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4255] = Zone{"net.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4255() {
	z[4256] = Zone{"org.rw", &z[1204], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4257] = Zone{"com.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4258] = Zone{"edu.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4259] = Zone{"gov.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4260] = Zone{"med.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4261] = Zone{"net.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4262] = Zone{"org.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4263] = Zone{"pub.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4263() {
	z[4264] = Zone{"sch.sa", &z[1207], x, 0, e, e, w{"c1.dns.sa", "c2.dns.sa", "i1.dns.sa", "m1.dns.sa", "m2.dns.sa", "n1.dns.sa", "p1.dns.sa", "s1.dns.sa", "s2.dns.sa", "sh1.dns.sa"}, n, n, n, e, e, n, f}
	z[4265] = Zone{"com.sb", &z[1227], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4266] = Zone{"edu.sb", &z[1227], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4267] = Zone{"gov.sb", &z[1227], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4268] = Zone{"net.sb", &z[1227], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4269] = Zone{"org.sb", &z[1227], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4270] = Zone{"com.sc", &z[1230], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4271] = Zone{"edu.sc", &z[1230], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4271() {
	z[4272] = Zone{"gov.sc", &z[1230], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4273] = Zone{"net.sc", &z[1230], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4274] = Zone{"org.sc", &z[1230], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4275] = Zone{"com.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4276] = Zone{"edu.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4277] = Zone{"gov.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4278] = Zone{"info.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4279] = Zone{"med.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4279() {
	z[4280] = Zone{"net.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4281] = Zone{"org.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4282] = Zone{"tv.sd", &z[1244], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4283] = Zone{"a.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4284] = Zone{"ac.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4285] = Zone{"b.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4286] = Zone{"bd.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4287] = Zone{"c.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4287() {
	z[4288] = Zone{"com.se", &z[1245], x, 0, e, e, w{"ns1.centralnic.net", "ns2.centralnic.net", "ns3.centralnic.net", "ns4.centralnic.net"}, n, n, n, "whois.centralnic.com", e, w{"https://rdap.centralnic.com/com.se/"}, f}
	z[4289] = Zone{"d.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4290] = Zone{"e.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4291] = Zone{"f.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4292] = Zone{"g.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4293] = Zone{"h.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4294] = Zone{"i.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4295] = Zone{"k.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4295() {
	z[4296] = Zone{"l.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4297] = Zone{"m.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4298] = Zone{"mil.se", &z[1245], x, 0, e, e, w{"dns5.telia.com", "dns6.telia.com", "ns.mil.se", "ns2.mil.se", "ns3.mil.se", "pitea.dns.swip.net"}, n, n, n, e, e, n, f}
	z[4299] = Zone{"n.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4300] = Zone{"o.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4301] = Zone{"org.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4302] = Zone{"p.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4303] = Zone{"parti.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4303() {
	z[4304] = Zone{"pp.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4305] = Zone{"press.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4306] = Zone{"r.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4307] = Zone{"s.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4308] = Zone{"t.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4309] = Zone{"tm.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4310] = Zone{"u.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4311] = Zone{"w.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4311() {
	z[4312] = Zone{"x.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4313] = Zone{"y.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4314] = Zone{"z.se", &z[1245], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4315] = Zone{"com.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, t}
	z[4316] = Zone{"edu.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, f}
	z[4317] = Zone{"gov.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, f}
	z[4318] = Zone{"net.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, f}
	z[4319] = Zone{"org.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4319() {
	z[4320] = Zone{"per.sg", &z[1260], x, 0, e, e, w{"dsany.sgnic.sg", "dsany2.sgnic.sg", "dsany3.sgnic.sg", "pch.sgzones.sg"}, n, n, n, e, e, n, f}
	z[4321] = Zone{"co.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4322] = Zone{"com.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4323] = Zone{"edu.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4324] = Zone{"gov.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4325] = Zone{"net.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4326] = Zone{"nom.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4327] = Zone{"org.sh", &z[1261], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4327() {
	z[4328] = Zone{"ae.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4329] = Zone{"at.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4330] = Zone{"cn.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4331] = Zone{"co.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4332] = Zone{"de.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4333] = Zone{"uk.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4334] = Zone{"us.si", &z[1276], x, 0, e, e, w{"b.ns14.net", "c.ns14.net", "d.ns14.net", "ns.2020.si"}, n, n, n, e, e, n, f}
	z[4335] = Zone{"org.sk", &z[1282], x, 0, e, e, w{"a.tld.sk", "e.tld.sk", "f.tld.sk", "g.tld.sk", "h.tld.sk"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4335() {
	z[4336] = Zone{"com.sl", &z[1289], x, 0, e, e, w{"ns1.neoip.com", "ns2.neoip.com"}, n, n, n, e, e, n, f}
	z[4337] = Zone{"edu.sl", &z[1289], x, 0, e, e, w{"ns1.neoip.com"}, n, n, n, e, e, n, t}
	z[4338] = Zone{"gov.sl", &z[1289], x, 0, e, e, w{"ns1.neoip.com", "ns2.neoip.com"}, n, n, n, e, e, n, f}
	z[4339] = Zone{"net.sl", &z[1289], x, 0, e, e, w{"ns1.neoip.com", "ns2.neoip.com"}, n, n, n, e, e, n, f}
	z[4340] = Zone{"org.sl", &z[1289], x, 0, e, e, w{"ns1.neoip.com", "ns2.neoip.com"}, n, n, n, e, e, n, f}
	z[4341] = Zone{"art.sn", &z[1294], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4342] = Zone{"com.sn", &z[1294], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4343] = Zone{"edu.sn", &z[1294], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4343() {
	z[4344] = Zone{"gouv.sn", &z[1294], x, 0, e, e, w{"ns.gouv.sn", "ns.ucad.sn", "ns2.gouv.sn", "ns3.gouv.sn"}, n, n, n, e, e, n, t}
	z[4345] = Zone{"org.sn", &z[1294], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4346] = Zone{"perso.sn", &z[1294], x, 0, e, e, w{"ns1.hellomouse.net"}, n, n, n, e, e, n, f}
	z[4347] = Zone{"univ.sn", &z[1294], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4348] = Zone{"com.so", &z[1296], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4349] = Zone{"edu.so", &z[1296], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4350] = Zone{"net.so", &z[1296], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4351] = Zone{"org.so", &z[1296], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4351() {
	z[4352] = Zone{"uber.space", &z[1308], x, 0, e, e, w{"ns1.jonaspasche.com", "ns2.jonaspasche.com", "ns3.jonaspasche.com"}, n, n, n, e, e, n, t}
	z[4353] = Zone{"biz.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4354] = Zone{"com.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4355] = Zone{"edu.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4356] = Zone{"me.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4357] = Zone{"net.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4358] = Zone{"org.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4359] = Zone{"sch.ss", &z[1317], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4359() {
	z[4360] = Zone{"co.st", &z[1318], x, 0, e, e, w{"ns1-se.ns.gov.st", "ns1-us.ns.gov.st", "ns2-se.ns.gov.st"}, n, n, n, e, e, n, f}
	z[4361] = Zone{"com.st", &z[1318], x, 0, e, e, w{"ns1-se.ns.gov.st", "ns1-us.ns.gov.st", "ns2-se.ns.gov.st"}, n, n, n, e, e, n, f}
	z[4362] = Zone{"consulado.st", &z[1318], x, 0, e, e, w{"ns1.bahnhof.net", "ns2.bahnhof.net"}, n, n, n, e, e, n, f}
	z[4363] = Zone{"edu.st", &z[1318], x, 0, e, e, w{"ns1.nic.st", "ns2.nic.st", "ns3.nic.st"}, n, n, n, e, e, n, f}
	z[4364] = Zone{"embaixada.st", &z[1318], x, 0, e, e, w{"ns1.bahnhof.net", "ns2.bahnhof.net"}, n, n, n, e, e, n, t}
	z[4365] = Zone{"embaizada.st", &z[1318], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4366] = Zone{"gov.st", &z[1318], x, 0, e, e, w{"ns1.gov.st", "ns2.gov.st"}, n, n, n, e, e, n, f}
	z[4367] = Zone{"mil.st", &z[1318], x, 0, e, e, w{"ns.bahnhof.se", "ns2.bahnhof.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4367() {
	z[4368] = Zone{"net.st", &z[1318], x, 0, e, e, w{"ns1.nic.st", "ns2.nic.st", "ns3.nic.st"}, n, n, n, e, e, n, f}
	z[4369] = Zone{"org.st", &z[1318], x, 0, e, e, w{"ns1.webhosting.st", "ns2.webhosting.st", "ns3.webhosting.st"}, n, n, n, e, e, n, f}
	z[4370] = Zone{"principe.st", &z[1318], x, 0, e, e, w{"ns1.govhost.st", "ns2.govhost.st", "ns3.govhost.st"}, n, n, n, e, e, n, f}
	z[4371] = Zone{"saotome.st", &z[1318], x, 0, e, e, w{"ns1.govhost.st", "ns2.govhost.st"}, n, n, n, e, e, n, f}
	z[4372] = Zone{"store.st", &z[1318], x, 0, e, e, w{"ns.bahnhof.se", "ns2.bahnhof.net"}, n, n, n, e, e, n, f}
	z[4373] = Zone{"abkhazia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4374] = Zone{"adygeya.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4375] = Zone{"aktyubinsk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4375() {
	z[4376] = Zone{"arkhangelsk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4377] = Zone{"armenia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4378] = Zone{"ashgabad.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4379] = Zone{"azerbaijan.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4380] = Zone{"balashov.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4381] = Zone{"bashkiria.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4382] = Zone{"bryansk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4383] = Zone{"bukhara.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4383() {
	z[4384] = Zone{"chimkent.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4385] = Zone{"dagestan.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4386] = Zone{"east-kazakhstan.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4387] = Zone{"exnet.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4388] = Zone{"georgia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4389] = Zone{"grozny.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4390] = Zone{"ivanovo.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4391] = Zone{"jambyl.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4391() {
	z[4392] = Zone{"kalmykia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4393] = Zone{"kaluga.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4394] = Zone{"karacol.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4395] = Zone{"karaganda.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4396] = Zone{"karelia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4397] = Zone{"khakassia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4398] = Zone{"krasnodar.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4399] = Zone{"kurgan.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4399() {
	z[4400] = Zone{"kustanai.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4401] = Zone{"lenug.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4402] = Zone{"mangyshlak.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4403] = Zone{"mordovia.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4404] = Zone{"msk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4405] = Zone{"murmansk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4406] = Zone{"nalchik.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4407] = Zone{"navoi.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4407() {
	z[4408] = Zone{"north-kazakhstan.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4409] = Zone{"nov.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4410] = Zone{"obninsk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4411] = Zone{"penza.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4412] = Zone{"pokrovsk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4413] = Zone{"sochi.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4414] = Zone{"spb.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4415] = Zone{"tashkent.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4415() {
	z[4416] = Zone{"termez.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4417] = Zone{"togliatti.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4418] = Zone{"troitsk.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4419] = Zone{"tselinograd.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4420] = Zone{"tula.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4421] = Zone{"tuva.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4422] = Zone{"vladikavkaz.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4423] = Zone{"vladimir.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
}

//go:noinline
func i4423() {
	z[4424] = Zone{"vologda.su", &z[1336], x, 0, e, e, w{"ns3-geo.nic.ru", "ns4-geo.nic.ru", "ns8-geo.nic.ru"}, w{"178.210.89.119"}, n, n, e, e, n, t}
	z[4425] = Zone{"com.sv", &z[1345], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4426] = Zone{"edu.sv", &z[1345], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4427] = Zone{"gob.sv", &z[1345], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4428] = Zone{"org.sv", &z[1345], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4429] = Zone{"red.sv", &z[1345], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4430] = Zone{"com.sy", &z[1351], x, 0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
	z[4431] = Zone{"edu.sy", &z[1351], x, 0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4431() {
	z[4432] = Zone{"gov.sy", &z[1351], x, 0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
	z[4433] = Zone{"mil.sy", &z[1351], x, 0, e, e, w{"freedns1.registrar-servers.com", "freedns2.registrar-servers.com", "freedns3.registrar-servers.com", "freedns4.registrar-servers.com", "freedns5.registrar-servers.com"}, n, n, n, e, e, n, f}
	z[4434] = Zone{"net.sy", &z[1351], x, 0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
	z[4435] = Zone{"org.sy", &z[1351], x, 0, e, e, w{"ns1.tld.sy", "pch.anycast.tld.sy", "sy.cctld.authdns.ripe.net"}, n, n, n, e, e, n, f}
	z[4436] = Zone{"ac.sz", &z[1355], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4437] = Zone{"co.sz", &z[1355], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4438] = Zone{"org.sz", &z[1355], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4439] = Zone{"biz.tc", &z[1367], x, 0, e, "https://www.biz.tc/", w{"ns1.nudrv.com", "ns2.nudrv.com"}, w{"64.136.20.39"}, n, n, e, e, n, f}
}

//go:noinline
func i4439() {
	z[4440] = Zone{"com.tc", &z[1367], x, 0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc", "root8.zone.tc"}, n, n, n, e, e, n, t}
	z[4441] = Zone{"edu.tc", &z[1367], x, 0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc", "root8.zone.tc"}, n, n, n, e, e, n, f}
	z[4442] = Zone{"net.tc", &z[1367], x, 0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc"}, n, n, n, e, e, n, f}
	z[4443] = Zone{"org.tc", &z[1367], x, 0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc", "root8.zone.tc"}, n, n, n, e, e, n, t}
	z[4444] = Zone{"pro.tc", &z[1367], x, 0, e, e, w{"root1.zone.tc", "root2.zone.tc", "root3.zone.tc", "root4.zone.tc", "root5.zone.tc", "root6.zone.tc", "root7.zone.tc", "root8.zone.tc"}, n, n, n, e, e, n, f}
	z[4445] = Zone{"com.td", &z[1369], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4446] = Zone{"net.td", &z[1369], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4447] = Zone{"org.td", &z[1369], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4447() {
	z[4448] = Zone{"tourism.td", &z[1369], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4449] = Zone{"at.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4450] = Zone{"bg.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4451] = Zone{"ca.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4452] = Zone{"ch.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4453] = Zone{"cz.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4454] = Zone{"de.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4455] = Zone{"edu.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4455() {
	z[4456] = Zone{"eu.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4457] = Zone{"int.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4458] = Zone{"net.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4459] = Zone{"pl.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4460] = Zone{"ru.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4461] = Zone{"sg.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4462] = Zone{"us.tf", &z[1382], x, 0, e, e, w{"ns1.idnscan.net", "ns2.idnscan.net"}, n, n, n, e, e, n, f}
	z[4463] = Zone{"ac.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4463() {
	z[4464] = Zone{"co.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4465] = Zone{"go.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4466] = Zone{"in.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4467] = Zone{"mi.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4468] = Zone{"net.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4469] = Zone{"or.th", &z[1384], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "c.thains.co.th", "ns.thnic.net", "p.thains.co.th"}, n, n, n, e, e, n, f}
	z[4470] = Zone{"ac.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4471] = Zone{"aero.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4471() {
	z[4472] = Zone{"biz.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4473] = Zone{"co.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4474] = Zone{"com.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4475] = Zone{"coop.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4476] = Zone{"dyn.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4477] = Zone{"edu.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4478] = Zone{"go.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4479] = Zone{"gov.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4479() {
	z[4480] = Zone{"info.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4481] = Zone{"int.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4482] = Zone{"mil.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4483] = Zone{"museum.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4484] = Zone{"my.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4485] = Zone{"name.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4486] = Zone{"net.tj", &z[1399], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4487] = Zone{"org.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4487() {
	z[4488] = Zone{"per.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4489] = Zone{"pro.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4490] = Zone{"web.tj", &z[1399], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4491] = Zone{"24.tl", &z[1404], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"24.tl", "95.217.58.108"}, n, n, e, e, n, t}
	z[4492] = Zone{"com.tl", &z[1404], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4493] = Zone{"gov.tl", &z[1404], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4494] = Zone{"in.tl", &z[1404], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4495] = Zone{"na.tl", &z[1404], x, 0, e, e, w{"ns1.afraid.org", "ns2.afraid.org", "ns3.afraid.org", "ns4.afraid.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4495() {
	z[4496] = Zone{"net.tl", &z[1404], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4497] = Zone{"org.tl", &z[1404], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4498] = Zone{"co.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, t}
	z[4499] = Zone{"com.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, f}
	z[4500] = Zone{"edu.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, f}
	z[4501] = Zone{"gov.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, f}
	z[4502] = Zone{"mil.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, t}
	z[4503] = Zone{"net.tm", &z[1405], x, 0, e, e, w{"webns1.telecom.tm", "webns2.telecom.tm"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4503() {
	z[4504] = Zone{"nom.tm", &z[1405], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4505] = Zone{"org.tm", &z[1405], x, 0, e, e, w{"dns01.nic.tm", "dns1.nic.tm", "dns2.nic.tm", "dns3.nic.tm", "dns4.nic.tm", "dns5.nic.tm", "dns6.nic.tm", "dns7.nic.tm"}, n, n, n, e, e, n, f}
	z[4506] = Zone{"agrinet.tn", &z[1407], x, 0, e, e, w{"ns4.ati.tn", "ns6.ati.tn"}, n, n, n, e, e, n, f}
	z[4507] = Zone{"com.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4508] = Zone{"defense.tn", &z[1407], x, 0, e, e, w{"ns1.defense.tn", "ns2.defense.tn"}, n, n, n, e, e, n, f}
	z[4509] = Zone{"edunet.tn", &z[1407], x, 0, e, e, w{"ns1.cnte.tn", "ns2.cnte.tn"}, n, n, n, e, e, n, f}
	z[4510] = Zone{"ens.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4511] = Zone{"fin.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4511() {
	z[4512] = Zone{"gov.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4513] = Zone{"ind.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4514] = Zone{"info.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4515] = Zone{"intl.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4516] = Zone{"mincom.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4517] = Zone{"nat.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4518] = Zone{"net.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4519] = Zone{"org.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4519() {
	z[4520] = Zone{"perso.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4521] = Zone{"rnrt.tn", &z[1407], x, 0, e, e, w{"ns1.cck.rnu.tn", "ns2.cck.rnu.tn", "ns3.cck.rnu.tn"}, n, n, n, e, e, n, f}
	z[4522] = Zone{"rns.tn", &z[1407], x, 0, e, e, w{"ns4.ati.tn", "ns6.ati.tn"}, n, n, n, e, e, n, f}
	z[4523] = Zone{"rnu.tn", &z[1407], x, 0, e, e, w{"ns1.cck.rnu.tn", "ns2.cck.rnu.tn", "ns3.cck.rnu.tn"}, n, n, n, e, e, n, f}
	z[4524] = Zone{"tourism.tn", &z[1407], x, 0, e, e, w{"ns1.ati.tn", "ns2.ati.tn"}, n, n, n, e, e, n, f}
	z[4525] = Zone{"com.tp", &z[1421], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4526] = Zone{"or.tp", &z[1421], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4527] = Zone{"org.tp", &z[1421], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4527() {
	z[4528] = Zone{"av.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4529] = Zone{"bbs.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4530] = Zone{"bel.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4531] = Zone{"biz.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4532] = Zone{"com.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4533] = Zone{"dr.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4534] = Zone{"edu.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4535] = Zone{"gen.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4535() {
	z[4536] = Zone{"gov.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4537] = Zone{"info.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4538] = Zone{"k12.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4539] = Zone{"mil.tr", &z[1422], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4540] = Zone{"name.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4541] = Zone{"nc.tr", &z[1422], x, 0, e, e, w{"ns1.nc.tr", "ns2.nc.tr"}, n, n, n, e, e, n, t}
	z[4542] = Zone{"net.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4543] = Zone{"org.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4543() {
	z[4544] = Zone{"pol.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4545] = Zone{"tel.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4546] = Zone{"tsk.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4547] = Zone{"tv.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4548] = Zone{"web.tr", &z[1422], x, 0, e, e, w{"lns41.nic.tr", "lns42.nic.tr", "lns61.nic.tr", "lns71.nic.tr", "lns72.nic.tr", "lns73.nic.tr", "lns74.nic.tr"}, n, n, n, e, e, n, t}
	z[4549] = Zone{"aero.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4550] = Zone{"at.tt", &z[1437], z[5173:5174], 0, e, e, w{"ns1.viennaweb.at", "ns2.viennaweb.at"}, w{"195.225.236.206"}, n, n, e, e, n, f}
	z[4551] = Zone{"au.tt", &z[1437], z[5174:5175], 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4551() {
	z[4552] = Zone{"be.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4553] = Zone{"biz.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4554] = Zone{"ca.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4555] = Zone{"cat.tt", &z[1437], x, 0, e, e, w{"ns1.cctld.com", "ns2.cctld.com"}, n, n, n, e, e, n, t}
	z[4556] = Zone{"ch.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4557] = Zone{"co.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4558] = Zone{"com.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4559] = Zone{"coop.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4559() {
	z[4560] = Zone{"de.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4561] = Zone{"edu.tt", &z[1437], x, 0, e, e, w{"ns1.edu.tt", "ns2.edu.tt"}, n, n, n, e, e, n, f}
	z[4562] = Zone{"es.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4563] = Zone{"eu.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4564] = Zone{"fr.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4565] = Zone{"gov.tt", &z[1437], x, 0, e, e, w{"dns5.gov.tt", "dns6.gov.tt", "dns7.gov.tt", "dns8.gov.tt"}, n, n, n, e, e, n, f}
	z[4566] = Zone{"info.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4567] = Zone{"int.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i4567() {
	z[4568] = Zone{"it.tt", &z[1437], x, 0, e, e, w{"pns41.cloudns.net", "pns42.cloudns.net", "pns43.cloudns.net", "pns44.cloudns.net"}, n, n, n, e, e, n, f}
	z[4569] = Zone{"jobs.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4570] = Zone{"mobi.tt", &z[1437], x, 0, e, e, w{"pns101.cloudns.net", "pns102.cloudns.net", "pns103.cloudns.net", "pns104.cloudns.net"}, w{"127.0.0.1", "::1"}, n, n, e, e, n, t}
	z[4571] = Zone{"museum.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4572] = Zone{"name.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4573] = Zone{"net.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4574] = Zone{"nl.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4575] = Zone{"org.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4575() {
	z[4576] = Zone{"pro.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4577] = Zone{"tel.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4578] = Zone{"travel.tt", &z[1437], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4579] = Zone{"uk.tt", &z[1437], z[5175:5178], 0, e, e, n, n, n, n, e, e, n, f}
	z[4580] = Zone{"us.tt", &z[1437], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4581] = Zone{"club.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4582] = Zone{"com.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "anytld.apnic.net", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4583] = Zone{"ebiz.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4583() {
	z[4584] = Zone{"edu.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "edudns-a1.edu.tw", "edudns-a2.edu.tw", "edudns-a3.edu.tw", "moemoon.edu.tw", "moestar.edu.tw"}, n, n, n, e, e, n, t}
	z[4585] = Zone{"game.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4586] = Zone{"gov.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4587] = Zone{"idv.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "anytld.apnic.net", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4588] = Zone{"mil.tw", &z[1444], x, 0, e, e, w{"dns1.mil.tw", "dns2.mil.tw"}, n, n, n, e, e, n, t}
	z[4589] = Zone{"net.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "anytld.apnic.net", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4590] = Zone{"org.tw", &z[1444], x, 0, e, e, w{"a.twnic.net.tw", "anytld.apnic.net", "b.twnic.net.tw", "c.twnic.net.tw", "d.twnic.net.tw", "e.twnic.net.tw", "f.twnic.net.tw", "g.twnic.net.tw", "h.dns.tw", "h.twnic.net.tw"}, n, n, n, e, e, n, t}
	z[4591] = Zone{"xn--czrw28b.tw" /* 商業.tw */, &z[1444], x, 0, e, "https://www.twnic.tw/", n, n, n, n, e, e, n, t}
}

//go:noinline
func i4591() {
	z[4592] = Zone{"xn--uc0atv.tw" /* 組織.tw */, &z[1444], x, 0, e, "https://www.twnic.tw/", n, n, n, n, e, e, n, t}
	z[4593] = Zone{"xn--zf0ao64a.tw" /* 網路.tw */, &z[1444], x, 0, e, "https://www.twnic.tw/", n, n, n, n, e, e, n, t}
	z[4594] = Zone{"ac.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4595] = Zone{"co.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4596] = Zone{"go.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4597] = Zone{"hotel.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4598] = Zone{"info.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4599] = Zone{"me.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4599() {
	z[4600] = Zone{"mil.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz"}, n, n, n, e, e, n, f}
	z[4601] = Zone{"mobi.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4602] = Zone{"ne.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4603] = Zone{"or.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "rip.psg.com", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4604] = Zone{"sc.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4605] = Zone{"tv.tz", &z[1445], x, 0, e, e, w{"fork.sth.dnsnode.net", "ns.anycast.co.tz", "ns2.tznic.or.tz", "tz-e.ns.nic.cz"}, n, n, n, e, e, n, f}
	z[4606] = Zone{"biz.ua", &z[1446], x, 0, e, e, w{"ns1.uadns.com", "ns2.uadns.com", "ns3.uadns.com"}, n, n, n, "whois.biz.ua", e, n, f}
	z[4607] = Zone{"cherkassy.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nn.ns.ua", "ns2.km.ua", "zt.ns.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4607() {
	z[4608] = Zone{"cherkasy.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "nn.ns.ua", "ns2.km.ua", "zt.ns.ua"}, n, n, n, e, e, n, t}
	z[4609] = Zone{"chernigov.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"nix.ns.ua", "ns.dn.ua", "ns3.cn.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4610] = Zone{"chernivtsi.ua", &z[1446], x, 0, e, e, w{"nix.ns.ua", "ns.cv.ua", "ns.kharkov.ua", "ns2.km.ua"}, n, n, n, e, e, n, t}
	z[4611] = Zone{"chernovtsy.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nix.ns.ua", "ns.cv.ua", "ns2.km.ua"}, n, n, n, e, e, n, f}
	z[4612] = Zone{"ck.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nn.ns.ua", "ns2.km.ua", "zt.ns.ua"}, n, n, n, e, e, n, f}
	z[4613] = Zone{"cn.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"nix.ns.ua", "ns.dn.ua", "ns3.cn.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4614] = Zone{"co.ua", &z[1446], x, 0, e, e, w{"ns1.uadns.com", "ns2.uadns.com", "ns3.uadns.com"}, n, n, n, "whois.co.ua", e, n, f}
	z[4615] = Zone{"com.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"cfdf.com.ns.ua", "ho1.com.ns.ua", "nn.ns.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4615() {
	z[4616] = Zone{"crimea.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.crimea.ua", "nn.ns.ua", "ns.crimea.ua"}, n, n, n, "whois.crimea.ua", e, n, f}
	z[4617] = Zone{"cv.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nix.ns.ua", "ns.cv.ua", "ns2.km.ua"}, n, n, n, e, e, n, f}
	z[4618] = Zone{"dn.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4619] = Zone{"dnepropetrovsk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ns.dp.ua", "s1.ns.dp.ua", "s2.ns.dp.ua", "s3.ns.dp.ua", "s4.ns.dp.ua"}, n, n, n, e, e, n, f}
	z[4620] = Zone{"dnipropetrovsk.ua", &z[1446], x, 0, e, e, w{"ns.dp.ua", "s1.ns.dp.ua", "s2.ns.dp.ua", "s3.ns.dp.ua", "s4.ns.dp.ua"}, n, n, n, e, e, n, t}
	z[4621] = Zone{"donetsk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4622] = Zone{"dp.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ns.dp.ua", "s1.ns.dp.ua", "s2.ns.dp.ua", "s3.ns.dp.ua", "s4.ns.dp.ua"}, n, n, n, e, e, n, f}
	z[4623] = Zone{"edu.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "bg.ns.ua", "ho1.edu.ns.ua", "nn.ns.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4623() {
	z[4624] = Zone{"gov.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "bg.ns.ua", "ho1.gov.ns.ua", "nn.ns.ua"}, n, n, n, e, e, n, f}
	z[4625] = Zone{"if.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.if.ua", "nn.ns.ua", "ns.dn.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4626] = Zone{"in.ua", &z[1446], x, 0, e, e, w{"ho1.ns.od.ua", "nn.ns.ua", "ns.in.ua", "sx.ns.ua"}, n, n, n, "whois.in.ua", e, n, f}
	z[4627] = Zone{"ivano-frankivsk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.if.ua", "nix.ns.ua", "ns.dn.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4628] = Zone{"kh.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4629] = Zone{"kharkiv.ua", &z[1446], x, 0, e, e, w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, t}
	z[4630] = Zone{"kharkov.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4631] = Zone{"kherson.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.ns.ks.ua", "nix.ns.ua", "ns.dn.ua", "ns.kharkov.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4631() {
	z[4632] = Zone{"khmelnitskiy.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "nn.ns.ua", "ns.km.ua", "ns2.km.ua", "ns3.km.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4633] = Zone{"kiev.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"cfdf.kiev.ns.ua", "ho1.kv.ns.ua", "nn.ns.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4634] = Zone{"kirovograd.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nix.ns.ua", "ns.host.kr.ua", "ns.infocom.kr.ua"}, n, n, n, e, e, n, f}
	z[4635] = Zone{"km.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nn.ns.ua", "ns.km.ua", "ns2.km.ua", "ns3.km.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4636] = Zone{"kr.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "nix.ns.ua", "ns.host.kr.ua", "ns.infocom.kr.ua"}, n, n, n, e, e, n, f}
	z[4637] = Zone{"ks.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.ns.ks.ua", "nix.ns.ua", "ns.dn.ua", "ns.kharkov.ua"}, n, n, n, e, e, n, f}
	z[4638] = Zone{"kyiv.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "bg.ns.ua", "cz.ns.ua", "ho1.kv.ns.ua", "nn.ns.ua"}, n, n, n, e, e, n, t}
	z[4639] = Zone{"lg.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4639() {
	z[4640] = Zone{"lt.ua", &z[1446], x, 0, e, e, w{"ho1.ns.lutsk.ua", "nix.ns.ua", "ns.cv.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4641] = Zone{"lugansk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4642] = Zone{"lutsk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.ns.lutsk.ua", "nix.ns.ua", "ns.zp.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4643] = Zone{"lviv.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"eta.uar.net", "iota.uar.net", "kappa.uar.net", "ns.lviv.ua", "theta.uar.net", "zeta.uar.net"}, n, n, n, "whois.lviv.ua", e, n, f}
	z[4644] = Zone{"mk.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.od.ua", "nix.ns.ua", "ns.mk.ua"}, n, n, n, e, e, n, f}
	z[4645] = Zone{"net.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "ho1.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua", "sx.ns.ua"}, n, n, n, "whois.net.ua", e, n, f}
	z[4646] = Zone{"nikolaev.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.od.ua", "nix.ns.ua", "ns.nikolaev.ua"}, n, n, n, e, e, n, f}
	z[4647] = Zone{"od.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.ns.od.ua", "nix.ns.ua", "ns.od.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4647() {
	z[4648] = Zone{"odesa.ua", &z[1446], x, 0, e, e, w{"ho1.ns.od.ua", "nix.ns.ua", "ns.od.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4649] = Zone{"odessa.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.ns.od.ua", "nix.ns.ua", "ns.od.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4650] = Zone{"org.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "cz.ns.ua", "ho1.ns.org.ua", "nn.ns.ua", "ns.dn.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4651] = Zone{"pl.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.pl.ua", "nn.ns.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4652] = Zone{"poltava.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.pl.ua", "nn.ns.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4653] = Zone{"pp.ua", &z[1446], x, 0, e, e, w{"ns1.uadns.com", "ns2.uadns.com", "ns3.uadns.com"}, n, n, n, "whois.pp.ua", e, n, f}
	z[4654] = Zone{"rivne.ua", &z[1446], x, 0, e, e, w{"ba1.ns.ua", "dns1.easydns.com", "dns2.easydns.net", "dns3.easydns.org", "pns41.cloudns.net", "pns42.cloudns.net"}, n, n, n, e, e, n, t}
	z[4655] = Zone{"rovno.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "dns1.easydns.com", "dns2.easydns.net", "dns3.easydns.org", "pns41.cloudns.net", "pns42.cloudns.net"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4655() {
	z[4656] = Zone{"rv.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "dns1.easydns.com", "dns2.easydns.net", "dns3.easydns.org", "pns41.cloudns.net", "pns42.cloudns.net"}, n, n, n, e, e, n, f}
	z[4657] = Zone{"sebastopol.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.crimea.ua", "nn.ns.ua", "ns.crimea.ua"}, n, n, n, "whois.sebastopol.ua", e, n, f}
	z[4658] = Zone{"sm.ua", &z[1446], x, 0, e, e, w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4659] = Zone{"sumy.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4660] = Zone{"te.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"nix.ns.ua", "ns.dn.ua", "ns.te.ua", "ns2.km.ua"}, n, n, n, e, e, n, f}
	z[4661] = Zone{"ternopil.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"nix.ns.ua", "ns.dn.ua", "ns.te.ua", "ns2.km.ua"}, n, n, n, e, e, n, f}
	z[4662] = Zone{"uz.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.uz.ua", "nn.ns.ua", "ns.dn.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4663] = Zone{"uzhgorod.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "ho1.ns.uz.ua", "nix.ns.ua", "ns.dn.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4663() {
	z[4664] = Zone{"vinnica.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "bat.enet.vn.ua", "ho1.ks.ns.ua", "nix.ns.ua", "ns1.vn.ua", "ns4.nest.vn.ua"}, n, n, n, e, e, n, f}
	z[4665] = Zone{"vn.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "bat.enet.vn.ua", "ho1.ks.ns.ua", "nix.ns.ua", "ns1.vn.ua", "ns4.nest.vn.ua"}, n, n, n, e, e, n, f}
	z[4666] = Zone{"volyn.ua", &z[1446], x, 0, e, e, w{"ho1.ns.lutsk.ua", "nix.ns.ua", "ns.zp.ua", "sx.ns.ua"}, n, n, n, e, e, n, f}
	z[4667] = Zone{"yalta.ua", &z[1446], x, 0, e, e, w{"ho1.ns.crimea.ua", "nn.ns.ua", "ns.crimea.ua"}, n, n, n, "whois.yalta.ua", e, n, f}
	z[4668] = Zone{"zaporizhzhe.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua", "ns.zp.ua"}, n, n, n, e, e, n, f}
	z[4669] = Zone{"zhitomir.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "impuls.zhitomir.ua", "nix.ns.ua", "zt.ns.ua"}, n, n, n, e, e, n, f}
	z[4670] = Zone{"zp.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ho1.eunic.net.ua", "nn.ns.ua", "ns.dn.ua", "ns.kh.ua", "ns.net.ua", "ns.sm.ua"}, n, n, n, e, e, n, f}
	z[4671] = Zone{"zt.ua", &z[1446], x, 0, e, "http://www.nic.net.ua/", w{"ba1.ns.ua", "impuls.zhitomir.ua", "nix.ns.ua", "zt.ns.ua"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4671() {
	z[4672] = Zone{"ac.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4673] = Zone{"co.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4674] = Zone{"com.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4675] = Zone{"go.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4676] = Zone{"mil.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4677] = Zone{"ne.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4678] = Zone{"or.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4679] = Zone{"org.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4679() {
	z[4680] = Zone{"sc.ug", &z[1450], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4681] = Zone{"ac.uk", &z[1451], x, 0, e, e, w{"auth03.ns.uu.net", "dns-3.dfn.de", "ns0.ja.net", "ns1.surfnet.nl", "ns2.ja.net", "ns3.ja.net", "ns4.ja.net"}, n, n, n, "whois.ja.net", e, n, f}
	z[4682] = Zone{"co.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, "whois.nic.uk", e, n, f}
	z[4683] = Zone{"gov.uk", &z[1451], x, 0, e, e, w{"auth00.ns.de.uu.net", "auth50.ns.de.uu.net", "ns0.ja.net", "ns1.surfnet.nl", "ns2.ja.net", "ns3.ja.net", "ns4.ja.net"}, n, n, n, "whois.ja.net", e, n, f}
	z[4684] = Zone{"ltd.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
	z[4685] = Zone{"me.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
	z[4686] = Zone{"mod.uk", &z[1451], x, 0, e, e, w{"ns0047.secondary.cloudflare.com", "ns0170.secondary.cloudflare.com"}, n, n, n, e, e, n, f}
	z[4687] = Zone{"net.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4687() {
	z[4688] = Zone{"org.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
	z[4689] = Zone{"plc.uk", &z[1451], x, 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
	z[4690] = Zone{"sch.uk", &z[1451], z[5178:5193], 0, e, e, w{"dns1.nic.uk", "dns2.nic.uk", "dns3.nic.uk", "dns4.nic.uk", "nsa.nic.uk", "nsb.nic.uk", "nsc.nic.uk", "nsd.nic.uk"}, n, n, n, e, e, n, f}
	z[4691] = Zone{"ak.us", &z[1461], x, 0x480, e, "https://doa.alaska.gov/ets/eash/DNSrequestform.html", n, n, n, n, e, e, n, t}
	z[4692] = Zone{"al.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4693] = Zone{"ar.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, t}
	z[4694] = Zone{"as.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4695] = Zone{"az.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4695() {
	z[4696] = Zone{"ca.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4697] = Zone{"co.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4698] = Zone{"ct.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4699] = Zone{"dc.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4700] = Zone{"de.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4701] = Zone{"dni.us", &z[1461], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4702] = Zone{"fed.us", &z[1461], x, 0, e, e, n, n, n, n, "whois.dotgov.gov", e, n, f}
	z[4703] = Zone{"fl.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4703() {
	z[4704] = Zone{"ga.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4705] = Zone{"gu.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4706] = Zone{"hi.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4707] = Zone{"ia.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4708] = Zone{"id.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4709] = Zone{"il.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4710] = Zone{"in.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4711] = Zone{"isa.us", &z[1461], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4711() {
	z[4712] = Zone{"kids.us", &z[1461], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4713] = Zone{"ks.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4714] = Zone{"ky.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4715] = Zone{"la.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4716] = Zone{"ma.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4717] = Zone{"md.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4718] = Zone{"me.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4719] = Zone{"mi.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4719() {
	z[4720] = Zone{"mn.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4721] = Zone{"mo.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4722] = Zone{"ms.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4723] = Zone{"mt.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4724] = Zone{"nc.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4725] = Zone{"nd.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4726] = Zone{"ne.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4727] = Zone{"nh.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4727() {
	z[4728] = Zone{"nj.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4729] = Zone{"nm.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4730] = Zone{"nsn.us", &z[1461], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4731] = Zone{"nv.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4732] = Zone{"ny.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4733] = Zone{"oh.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4734] = Zone{"ok.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4735] = Zone{"or.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4735() {
	z[4736] = Zone{"pa.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4737] = Zone{"pr.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4738] = Zone{"ri.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4739] = Zone{"sc.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4740] = Zone{"sd.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4741] = Zone{"tn.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4742] = Zone{"tx.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4743] = Zone{"ut.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4743() {
	z[4744] = Zone{"va.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4745] = Zone{"vi.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4746] = Zone{"vt.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4747] = Zone{"wa.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4748] = Zone{"wi.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4749] = Zone{"wv.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4750] = Zone{"wy.us", &z[1461], x, 0x480, e, e, n, n, n, n, e, e, n, f}
	z[4751] = Zone{"com.uy", &z[1462], x, 0, e, e, w{"ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4751() {
	z[4752] = Zone{"edu.uy", &z[1462], x, 0, e, e, w{"a.lactld.org", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy", "ultra.edu.uy"}, n, n, n, e, e, n, f}
	z[4753] = Zone{"gub.uy", &z[1462], x, 0, e, e, w{"a.lactld.org", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy", "ultra.edu.uy"}, n, n, n, e, e, n, f}
	z[4754] = Zone{"mil.uy", &z[1462], x, 0, e, e, w{"a.lactld.org", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy", "ultra.edu.uy"}, n, n, n, e, e, n, f}
	z[4755] = Zone{"net.uy", &z[1462], x, 0, e, e, w{"a.lactld.org", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy", "ultra.edu.uy"}, n, n, n, e, e, n, f}
	z[4756] = Zone{"org.uy", &z[1462], x, 0, e, e, w{"a.lactld.org", "ns.dns.br", "ns1.anteldata.com.uy", "ns2.anteldata.com.uy", "seciu.edu.uy", "ultra.edu.uy"}, n, n, n, e, e, n, f}
	z[4757] = Zone{"ac.uz", &z[1463], x, 0, e, e, w{"p3.dc.uz", "p4.dc.uz"}, n, n, n, e, e, n, f}
	z[4758] = Zone{"aero.uz", &z[1463], x, 0, e, e, w{"ns3.uzinfocom.uz", "rdns3.uzinfocom.uz"}, n, n, n, e, e, n, f}
	z[4759] = Zone{"biz.uz", &z[1463], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "biz.uz"}, n, n, e, e, n, f}
}

//go:noinline
func i4759() {
	z[4760] = Zone{"co.uz", &z[1463], x, 0, e, e, w{"ns1.reg.uz", "ns2.reg.uz"}, n, n, n, e, e, n, f}
	z[4761] = Zone{"com.uz", &z[1463], x, 0, e, e, w{"ns1.reg.uz", "ns2.reg.uz"}, n, n, n, e, e, n, f}
	z[4762] = Zone{"coop.uz", &z[1463], x, 0, e, e, w{"ns3.uzinfocom.uz", "rdns3.uzinfocom.uz"}, n, n, n, e, e, n, f}
	z[4763] = Zone{"edu.uz", &z[1463], x, 0, e, e, w{"dns1.webspace.uz", "dns2.webspace.uz", "dns3.webspace.uz", "dns4.webspace.uz"}, n, n, n, e, e, n, f}
	z[4764] = Zone{"info.uz", &z[1463], x, 0, e, e, w{"ns.ol.uz", "ns.tshtt.uz"}, n, n, n, e, e, n, f}
	z[4765] = Zone{"int.uz", &z[1463], x, 0, e, e, w{"ns3.uzinfocom.uz", "rdns3.uzinfocom.uz"}, n, n, n, e, e, n, f}
	z[4766] = Zone{"museum.uz", &z[1463], x, 0, e, e, w{"dns1.webspace.uz", "dns2.webspace.uz", "dns3.webspace.uz", "dns4.webspace.uz"}, n, n, n, e, e, n, f}
	z[4767] = Zone{"name.uz", &z[1463], x, 0, e, e, w{"ns1.billur.com", "ns2.billur.com"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4767() {
	z[4768] = Zone{"net.uz", &z[1463], x, 0, e, e, w{"ns1.reg.uz", "ns2.reg.uz"}, n, n, n, e, e, n, f}
	z[4769] = Zone{"org.uz", &z[1463], x, 0, e, e, w{"ns1.reg.uz", "ns2.reg.uz"}, n, n, n, e, e, n, f}
	z[4770] = Zone{"pro.uz", &z[1463], x, 0x8, e, e, w{"ns3.uzinfocom.uz", "rdns3.uzinfocom.uz"}, n, n, n, e, e, n, f}
	z[4771] = Zone{"com.vc", &z[1469], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4772] = Zone{"net.vc", &z[1469], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4773] = Zone{"org.vc", &z[1469], x, 0, e, e, w{"a0.cctld.afilias-nst.info", "a2.cctld.afilias-nst.info", "b0.cctld.afilias-nst.org", "b2.cctld.afilias-nst.org", "c0.cctld.afilias-nst.info", "d0.cctld.afilias-nst.org"}, n, n, n, e, e, n, f}
	z[4774] = Zone{"co.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4775] = Zone{"com.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4775() {
	z[4776] = Zone{"edu.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4777] = Zone{"gob.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4778] = Zone{"info.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4779] = Zone{"int.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, t}
	z[4780] = Zone{"mil.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4781] = Zone{"net.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4782] = Zone{"org.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4783] = Zone{"tec.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4783() {
	z[4784] = Zone{"web.ve", &z[1470], x, 0, e, e, w{"a.lactld.org", "ns3.nic.ve", "ns4.nic.ve", "ns5.nic.ve", "ns6.nic.ve", "ssdns-tld.nic.cl"}, n, n, n, e, e, n, f}
	z[4785] = Zone{"pro.vg", &z[1478], x, 0x200, e, e, w{"ns1.subdomain.net", "ns2.subdomain.net"}, w{"95.217.58.108", "pro.vg"}, n, n, e, e, n, t}
	z[4786] = Zone{"biz.vi", &z[1479], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4787] = Zone{"co.vi", &z[1479], x, 0, e, e, w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, e, n, f}
	z[4788] = Zone{"com.vi", &z[1479], x, 0, e, e, w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, e, n, t}
	z[4789] = Zone{"k12.vi", &z[1479], x, 0, e, e, w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, e, n, f}
	z[4790] = Zone{"net.vi", &z[1479], x, 0, e, e, w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, e, n, f}
	z[4791] = Zone{"org.vi", &z[1479], x, 0, e, e, w{"ns3.nic.vi", "pch.nic.vi"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4791() {
	z[4792] = Zone{"ac.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4793] = Zone{"biz.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4794] = Zone{"com.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4795] = Zone{"edu.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4796] = Zone{"gov.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4797] = Zone{"health.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4798] = Zone{"info.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4799] = Zone{"int.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4799() {
	z[4800] = Zone{"name.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4801] = Zone{"net.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4802] = Zone{"org.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4803] = Zone{"pro.vn", &z[1495], x, 0, e, e, w{"a.dns-servers.vn", "b.dns-servers.vn", "c.dns-servers.vn", "d.dns-servers.vn", "e.dns-servers.vn", "f.dns-servers.vn", "g.dns-servers.vn", "h.dns-servers.vn"}, n, n, n, e, e, n, f}
	z[4804] = Zone{"co.vu", &z[1504], x, 0, e, "https://codotvu.co/", w{"ns1.codotvu.com", "ns2.codotvu.com", "ns3.codotvu.com", "ns4.codotvu.com"}, n, n, n, e, e, n, t}
	z[4805] = Zone{"com.vu", &z[1504], x, 0, e, e, w{"ns1.tldns.vu", "ns2.tldns.vu", "ns3.tldns.vu", "ns4.tldns.vu"}, n, n, n, e, e, n, f}
	z[4806] = Zone{"name.vu", &z[1504], x, 0x200, e, e, n, n, n, n, e, e, n, t}
	z[4807] = Zone{"net.vu", &z[1504], x, 0, e, e, w{"ns1.tldns.vu", "ns2.tldns.vu", "ns3.tldns.vu", "ns4.tldns.vu"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4807() {
	z[4808] = Zone{"org.vu", &z[1504], x, 0, e, e, w{"ns1.tldns.vu", "ns2.tldns.vu", "ns3.tldns.vu", "ns4.tldns.vu"}, n, n, n, e, e, n, f}
	z[4809] = Zone{"com.ws", &z[1543], x, 0, e, e, w{"europe1.dnsnode.net", "ns2.dns.ws", "ns5.dns.ws", "nsp.dnsnode.net", "nsu.dnsnode.net", "us3.dns.ws", "us4.dns.ws"}, w{"202.4.48.211"}, n, n, e, e, n, f}
	z[4810] = Zone{"edu.ws", &z[1543], x, 0, e, e, w{"europe1.dnsnode.net", "ns2.dns.ws", "ns5.dns.ws", "nsp.dnsnode.net", "nsu.dnsnode.net", "us3.dns.ws", "us4.dns.ws"}, w{"202.4.48.211"}, n, n, e, e, n, f}
	z[4811] = Zone{"gov.ws", &z[1543], x, 0, e, e, w{"europe1.dnsnode.net", "ns2.dns.ws", "ns5.dns.ws", "nsp.dnsnode.net", "nsu.dnsnode.net", "us3.dns.ws", "us4.dns.ws"}, n, n, n, e, e, n, f}
	z[4812] = Zone{"net.ws", &z[1543], x, 0, e, e, w{"europe1.dnsnode.net", "ns2.dns.ws", "ns5.dns.ws", "nsp.dnsnode.net", "nsu.dnsnode.net", "us3.dns.ws", "us4.dns.ws"}, w{"202.4.48.211"}, n, n, e, e, n, f}
	z[4813] = Zone{"org.ws", &z[1543], x, 0, e, e, w{"europe1.dnsnode.net", "ns2.dns.ws", "ns5.dns.ws", "nsp.dnsnode.net", "nsu.dnsnode.net", "us3.dns.ws", "us4.dns.ws"}, w{"202.4.48.211"}, n, n, e, e, n, f}
	z[4814] = Zone{"co.ye", &z[1558], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4815] = Zone{"com.ye", &z[1558], x, 0, e, e, w{"tld1.ye", "tld2.ye"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4815() {
	z[4816] = Zone{"gov.ye", &z[1558], x, 0, e, e, w{"tld1.ye", "tld2.ye"}, n, n, n, e, e, n, t}
	z[4817] = Zone{"ltd.ye", &z[1558], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4818] = Zone{"me.ye", &z[1558], x, 0, e, e, w{"tld1.ye", "tld2.ye"}, n, n, n, e, e, n, f}
	z[4819] = Zone{"net.ye", &z[1558], x, 0, e, e, w{"tld1.ye", "tld2.ye"}, n, n, n, e, e, n, f}
	z[4820] = Zone{"org.ye", &z[1558], x, 0, e, e, w{"ns1.yemen.net.ye", "ns2.yemen.net.ye", "sah1.ye", "sah2.ye", "tld1.ye", "tld2.ye"}, n, n, n, e, e, n, f}
	z[4821] = Zone{"plc.ye", &z[1558], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4822] = Zone{"ac.yu", &z[1565], z[5193:5194], 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4823] = Zone{"cg.yu", &z[1565], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4823() {
	z[4824] = Zone{"co.yu", &z[1565], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[4825] = Zone{"edu.yu", &z[1565], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4826] = Zone{"gov.yu", &z[1565], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4827] = Zone{"org.yu", &z[1565], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4828] = Zone{"ac.za", &z[1567], x, 0, e, e, w{"disa.tenet.ac.za", "ns2us.dns.business", "ns3.dns.business", "ns4.dns.business", "za-ns.anycast.pch.net"}, n, n, n, "whois.ac.za", e, n, f}
	z[4829] = Zone{"agric.za", &z[1567], x, 0, e, e, w{"demeter.is.co.za", "jupiter.is.co.za", "titan.is.co.za"}, n, n, n, e, e, n, f}
	z[4830] = Zone{"alt.za", &z[1567], x, 0, e, e, w{"ans.dnsstudy.africa", "ns-za.afrinic.net", "ns1.iafrica.com", "ns2.iafrica.com", "psg.com", "za-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4831] = Zone{"bourse.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4831() {
	z[4832] = Zone{"city.za", &z[1567], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4833] = Zone{"co.za", &z[1567], x, 0, e, e, w{"coza1.dnsnode.net", "ns.coza.net.za", "ns0.is.co.za", "ns2us.dns.business"}, n, n, n, "coza-whois.registry.net.za", e, n, f}
	z[4834] = Zone{"cybernet.za", &z[1567], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4835] = Zone{"db.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4836] = Zone{"edu.za", &z[1567], x, 0, e, e, w{"ans.dnsstudy.africa", "ns-za.afrinic.net", "ns1.iafrica.com", "ns2.iafrica.com", "za-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4837] = Zone{"gov.za", &z[1567], x, 0, e, e, w{"ns1.gov.za", "ns1.inx.net.za", "ns2.gov.za", "ns2.inx.net.za", "ns3.gov.za", "ns3.inx.net.za"}, n, n, n, "whois.gov.za", e, n, f}
	z[4838] = Zone{"grondar.za", &z[1567], x, 0, e, e, w{"ns.rail.eu.org", "ns0.keltia.net", "ns2.bortzmeyer.org"}, n, n, n, e, e, n, f}
	z[4839] = Zone{"iaccess.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4839() {
	z[4840] = Zone{"imt.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4841] = Zone{"inca.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4842] = Zone{"landesign.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4843] = Zone{"law.za", &z[1567], x, 0, e, e, w{"ns0.is.co.za", "ns1.dns.net.za", "ns2us.dns.business", "za-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4844] = Zone{"mil.za", &z[1567], x, 0, e, e, w{"ns1.gov.za", "ns2.gov.za", "ns3.gov.za"}, n, n, n, e, e, n, f}
	z[4845] = Zone{"net.za", &z[1567], x, 0, e, e, w{"ns0.is.co.za", "za-ns.anycast.pch.net"}, n, n, n, "net-whois.registry.net.za", e, n, f}
	z[4846] = Zone{"ngo.za", &z[1567], x, 0, e, e, w{"antares.wisenet.co.za", "ns1.sn.apc.org", "ns2.sn.apc.org"}, w{"127.0.0.1", "::1"}, n, n, e, e, n, f}
	z[4847] = Zone{"nis.za", &z[1567], x, 0, e, e, w{"ns1.mywebserver.co.za", "ns2.mywebserver.co.za", "ns3.mywebserver.co.za"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4847() {
	z[4848] = Zone{"nom.za", &z[1567], x, 0, e, e, w{"ins1.c6dns.com", "nomza.cryger.com", "ns0.plig.net"}, w{"127.0.0.1"}, n, n, e, e, n, f}
	z[4849] = Zone{"olivetti.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4850] = Zone{"org.za", &z[1567], x, 0, e, e, w{"ns0.is.co.za", "ns0.plig.net", "za-ns.anycast.pch.net"}, n, n, n, "org-whois.registry.net.za", e, n, f}
	z[4851] = Zone{"pix.za", &z[1567], x, 0x1000, e, e, n, n, n, n, e, e, n, f}
	z[4852] = Zone{"school.za", &z[1567], z[5194:5203], 0, e, e, w{"ns10.dns.business", "ns11.dns.business", "za-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4853] = Zone{"tm.za", &z[1567], x, 0, e, e, w{"ns0.is.co.za", "ns1.dnsza.co.za", "ns1.iafrica.com", "ns1.lolo.co.za", "ns2.iafrica.com", "ns2.lolo.co.za", "ns3.lolo.co.za", "reaper.org", "sticky.spider.web.za", "za-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4854] = Zone{"web.za", &z[1567], x, 0, e, e, w{"disa.tenet.ac.za", "ns0.is.co.za", "za-ns.anycast.pch.net"}, n, n, n, "web-whois.registry.net.za", e, n, f}
	z[4855] = Zone{"ac.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, f}
}

//go:noinline
func i4855() {
	z[4856] = Zone{"co.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4857] = Zone{"com.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4858] = Zone{"edu.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4859] = Zone{"gov.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4860] = Zone{"net.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4861] = Zone{"org.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, f}
	z[4862] = Zone{"sch.zm", &z[1573], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4863] = Zone{"ac.zw", &z[1577], x, 0, e, e, w{"ns1.liquidtelecom.net", "ns2.liquidtelecom.net", "uofzns01.uz.ac.zw", "uofzns02.uz.ac.zw"}, n, n, n, e, e, n, f}
}

//go:noinline
func i4863() {
	z[4864] = Zone{"co.zw", &z[1577], x, 0, e, e, w{"ans.dnsstudy.africa", "ns1.liquidtelecom.net", "ns1.zispa.org.zw", "ns2.liquidtelecom.net", "ns2.zispa.org.zw"}, n, n, n, e, e, n, t}
	z[4865] = Zone{"com.zw", &z[1577], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[4866] = Zone{"org.zw", &z[1577], x, 0, e, e, w{"ns1.liquidtelecom.net", "ns1zim.telone.co.zw", "ns2.liquidtelecom.net", "ns2zim.telone.co.zw", "zw-ns.anycast.pch.net"}, n, n, n, e, e, n, f}
	z[4867] = Zone{"001.xn--p1acf" /* 001.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4868] = Zone{"002.xn--p1acf" /* 002.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4869] = Zone{"003.xn--p1acf" /* 003.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4870] = Zone{"004.xn--p1acf" /* 004.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4871] = Zone{"005.xn--p1acf" /* 005.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4871() {
	z[4872] = Zone{"006.xn--p1acf" /* 006.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4873] = Zone{"007.xn--p1acf" /* 007.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4874] = Zone{"008.xn--p1acf" /* 008.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4875] = Zone{"009.xn--p1acf" /* 009.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4876] = Zone{"010.xn--p1acf" /* 010.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4877] = Zone{"011.xn--p1acf" /* 011.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4878] = Zone{"012.xn--p1acf" /* 012.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4879] = Zone{"013.xn--p1acf" /* 013.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4879() {
	z[4880] = Zone{"014.xn--p1acf" /* 014.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4881] = Zone{"015.xn--p1acf" /* 015.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4882] = Zone{"016.xn--p1acf" /* 016.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4883] = Zone{"017.xn--p1acf" /* 017.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4884] = Zone{"018.xn--p1acf" /* 018.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4885] = Zone{"019.xn--p1acf" /* 019.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4886] = Zone{"021.xn--p1acf" /* 021.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4887] = Zone{"022.xn--p1acf" /* 022.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4887() {
	z[4888] = Zone{"023.xn--p1acf" /* 023.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4889] = Zone{"024.xn--p1acf" /* 024.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4890] = Zone{"025.xn--p1acf" /* 025.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4891] = Zone{"026.xn--p1acf" /* 026.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4892] = Zone{"027.xn--p1acf" /* 027.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4893] = Zone{"028.xn--p1acf" /* 028.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4894] = Zone{"029.xn--p1acf" /* 029.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4895] = Zone{"030.xn--p1acf" /* 030.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4895() {
	z[4896] = Zone{"031.xn--p1acf" /* 031.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4897] = Zone{"032.xn--p1acf" /* 032.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4898] = Zone{"033.xn--p1acf" /* 033.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4899] = Zone{"034.xn--p1acf" /* 034.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4900] = Zone{"035.xn--p1acf" /* 035.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4901] = Zone{"036.xn--p1acf" /* 036.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4902] = Zone{"037.xn--p1acf" /* 037.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4903] = Zone{"038.xn--p1acf" /* 038.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4903() {
	z[4904] = Zone{"039.xn--p1acf" /* 039.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4905] = Zone{"040.xn--p1acf" /* 040.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4906] = Zone{"041.xn--p1acf" /* 041.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4907] = Zone{"042.xn--p1acf" /* 042.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4908] = Zone{"043.xn--p1acf" /* 043.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4909] = Zone{"044.xn--p1acf" /* 044.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4910] = Zone{"045.xn--p1acf" /* 045.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4911] = Zone{"046.xn--p1acf" /* 046.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4911() {
	z[4912] = Zone{"047.xn--p1acf" /* 047.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4913] = Zone{"048.xn--p1acf" /* 048.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4914] = Zone{"049.xn--p1acf" /* 049.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4915] = Zone{"050.xn--p1acf" /* 050.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4916] = Zone{"051.xn--p1acf" /* 051.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4917] = Zone{"052.xn--p1acf" /* 052.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4918] = Zone{"053.xn--p1acf" /* 053.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4919] = Zone{"054.xn--p1acf" /* 054.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4919() {
	z[4920] = Zone{"055.xn--p1acf" /* 055.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4921] = Zone{"056.xn--p1acf" /* 056.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4922] = Zone{"057.xn--p1acf" /* 057.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4923] = Zone{"058.xn--p1acf" /* 058.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4924] = Zone{"059.xn--p1acf" /* 059.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4925] = Zone{"060.xn--p1acf" /* 060.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4926] = Zone{"061.xn--p1acf" /* 061.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4927] = Zone{"062.xn--p1acf" /* 062.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4927() {
	z[4928] = Zone{"063.xn--p1acf" /* 063.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4929] = Zone{"064.xn--p1acf" /* 064.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4930] = Zone{"065.xn--p1acf" /* 065.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4931] = Zone{"066.xn--p1acf" /* 066.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4932] = Zone{"067.xn--p1acf" /* 067.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4933] = Zone{"068.xn--p1acf" /* 068.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4934] = Zone{"069.xn--p1acf" /* 069.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4935] = Zone{"070.xn--p1acf" /* 070.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4935() {
	z[4936] = Zone{"071.xn--p1acf" /* 071.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4937] = Zone{"072.xn--p1acf" /* 072.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4938] = Zone{"073.xn--p1acf" /* 073.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4939] = Zone{"074.xn--p1acf" /* 074.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4940] = Zone{"075.xn--p1acf" /* 075.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4941] = Zone{"076.xn--p1acf" /* 076.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4942] = Zone{"077.xn--p1acf" /* 077.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4943] = Zone{"078.xn--p1acf" /* 078.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4943() {
	z[4944] = Zone{"079.xn--p1acf" /* 079.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4945] = Zone{"083.xn--p1acf" /* 083.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4946] = Zone{"086.xn--p1acf" /* 086.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4947] = Zone{"087.xn--p1acf" /* 087.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4948] = Zone{"089.xn--p1acf" /* 089.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4949] = Zone{"094.xn--p1acf" /* 094.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4950] = Zone{"095.xn--p1acf" /* 095.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4951] = Zone{"xn--80aaac0ct.xn--p1acf" /* абакан.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4951() {
	z[4952] = Zone{"xn--80ae0bp.xn--p1acf" /* авто.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4953] = Zone{"xn--80awd.xn--p1acf" /* алм.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4954] = Zone{"xn--80aa1ag9a.xn--p1acf" /* алтай.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4955] = Zone{"xn--80aam8ar9di.xn--p1acf" /* анадырь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4956] = Zone{"xn--80aaa5csg.xn--p1acf" /* астана.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4957] = Zone{"xn--80aaa6cmfh0a9d.xn--p1acf" /* астрахань.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4958] = Zone{"xn--80ag7c.xn--p1acf" /* ахг.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4959] = Zone{"xn--80ab2azb.xn--p1acf" /* баку.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4959() {
	z[4960] = Zone{"xn--80aab6birx.xn--p1acf" /* барнаул.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4961] = Zone{"xn--90ais.xn--p1acf" /* бел.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4962] = Zone{"xn--90aedc4atap.xn--p1acf" /* белгород.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4963] = Zone{"xn--90amc.xn--p1acf" /* биз.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, w{"77.221.159.235"}, n, n, e, e, n, t}
	z[4964] = Zone{"xn--80abamkjb7bdt.xn--p1acf" /* биробиджан.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4965] = Zone{"xn--90aiiib5f.xn--p1acf" /* бишкек.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4966] = Zone{"xn--80abedla9acxg1b7f.xn--p1acf" /* благовещенск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4967] = Zone{"xn--90aayernio.xn--p1acf" /* бобруйск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4967() {
	z[4968] = Zone{"xn--90ai6aff.xn--p1acf" /* брест.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4969] = Zone{"xn--90asilg6f.xn--p1acf" /* брянск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4970] = Zone{"xn--90aoxid7ec.xn--p1acf" /* бурятия.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4971] = Zone{"xn--b1aadecsaihi0amca3a.xn--p1acf" /* великийновгород.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4972] = Zone{"xn--b1amjiu8dp.xn--p1acf" /* вильнюс.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4973] = Zone{"xn--90abjlm5be.xn--p1acf" /* витебск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4974] = Zone{"xn--80adbhunc2aa3al.xn--p1acf" /* владивосток.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4975] = Zone{"xn--80aaafckyesce.xn--p1acf" /* владикавказ.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4975() {
	z[4976] = Zone{"xn--80adhqaok7a.xn--p1acf" /* владимир.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4977] = Zone{"xn--80add5ao.xn--p1acf" /* волга.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4978] = Zone{"xn--80addag2buct.xn--p1acf" /* волгоград.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4979] = Zone{"xn--80adde7arb.xn--p1acf" /* вологда.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4980] = Zone{"xn--b1agd0aean.xn--p1acf" /* воронеж.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4981] = Zone{"xn--c1aescj1g.xn--p1acf" /* гомель.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4982] = Zone{"xn--c1ac1acci.xn--p1acf" /* гродно.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4983] = Zone{"xn--c1aigpdl4f.xn--p1acf" /* грозный.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4983() {
	z[4984] = Zone{"xn--e1aky.xn--p1acf" /* ект.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4985] = Zone{"xn--80adia6bo.xn--p1acf" /* ереван.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4986] = Zone{"xn--80adbv1agb.xn--p1acf" /* иваново.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4987] = Zone{"xn--b1afchn5b.xn--p1acf" /* ижевск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4988] = Zone{"xn--h1aeawgfg.xn--p1acf" /* иркутск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4989] = Zone{"xn--80aauks4g.xn--p1acf" /* казань.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4990] = Zone{"xn--80aaifraljtb8a.xn--p1acf" /* калининград.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4991] = Zone{"xn--80aai0ag2c.xn--p1acf" /* калуга.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4991() {
	z[4992] = Zone{"xn--b1afaslnbn.xn--p1acf" /* кемерово.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4993] = Zone{"xn--b1afih.xn--p1acf" /* киев.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4994] = Zone{"xn--b1alfrj.xn--p1acf" /* киров.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4995] = Zone{"xn--b1afiaiu9e.xn--p1acf" /* кишинев.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4996] = Zone{"xn--j1aef.xn--p1acf" /* ком.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4997] = Zone{"xn--h1adgh.xn--p1acf" /* коми.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4998] = Zone{"xn--80atghalgi.xn--p1acf" /* кострома.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[4999] = Zone{"xn--80aalwqglfe.xn--p1acf" /* краснодар.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i4999() {
	z[5000] = Zone{"xn--80atblfjdfd2l.xn--p1acf" /* красноярск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5001] = Zone{"xn--j1ael8b.xn--p1acf" /* крым.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5002] = Zone{"xn--80afuomr.xn--p1acf" /* курган.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5003] = Zone{"xn--j1aarei.xn--p1acf" /* курск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5004] = Zone{"xn--g1afe6db.xn--p1acf" /* кызыл.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5005] = Zone{"xn--e1afhbv7b.xn--p1acf" /* липецк.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5006] = Zone{"xn--80aaakg8bm.xn--p1acf" /* магадан.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5007] = Zone{"xn--80aai8a2a.xn--p1acf" /* магас.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5007() {
	z[5008] = Zone{"xn--80asdfng.xn--p1acf" /* майкоп.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5009] = Zone{"xn--80aqdibz7g.xn--p1acf" /* марийэл.рус */, &z[1593], x, 0x200, e, e, w{"ns1.nameself.com", "ns2.nameself.com"}, n, n, n, e, e, n, t}
	z[5010] = Zone{"xn--80aqjbv3f.xn--p1acf" /* мариэл.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5011] = Zone{"xn--80aaaa2chc7eta.xn--p1acf" /* махачкала.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5012] = Zone{"xn--d1abu.xn--p1acf" /* мед.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5013] = Zone{"xn--h1aeefu.xn--p1acf" /* минск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5014] = Zone{"xn--h1ahn.xn--p1acf" /* мир.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5015] = Zone{"xn--b1abpmcm0l.xn--p1acf" /* могилёв.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5015() {
	z[5016] = Zone{"xn--j1adp.xn--p1acf" /* мск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5017] = Zone{"xn--80aueagpkl.xn--p1acf" /* мурманск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5018] = Zone{"xn--80aqgci6d8a.xn--p1acf" /* нальчик.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5019] = Zone{"xn----7sbb1ccbue7h4a.xn--p1acf" /* нарьян-мар.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5020] = Zone{"xn--b1avn.xn--p1acf" /* нвс.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5021] = Zone{"xn--m1aa.xn--p1acf" /* нн.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5022] = Zone{"xn--90absbknhbvge.xn--p1acf" /* новосибирск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5023] = Zone{"xn--j1adfn.xn--p1acf" /* омск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5023() {
	z[5024] = Zone{"xn--c1avg.xn--p1acf" /* орг.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5025] = Zone{"xn--90aee6admdx.xn--p1acf" /* оренбург.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5026] = Zone{"xn--k1afg2e.xn--p1acf" /* орёл.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5027] = Zone{"xn--80ajgxi.xn--p1acf" /* пенза.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5028] = Zone{"xn--e1aohf5d.xn--p1acf" /* пермь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5029] = Zone{"xn--80aegbkvxddlre.xn--p1acf" /* петрозаводск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5030] = Zone{"xn--b1apmck.xn--p1acf" /* псков.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5031] = Zone{"xn--c1ajhpcjhd0j.xn--p1acf" /* пятигорск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5031() {
	z[5032] = Zone{"xn--80afo5a.xn--p1acf" /* рига.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5033] = Zone{"xn-----7kcgn5cdbagnnnx.xn--p1acf" /* ростов-на-дону.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5034] = Zone{"xn--b1a1ade.xn--p1acf" /* рств.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5035] = Zone{"xn--80antj7do.xn--p1acf" /* рязань.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5036] = Zone{"xn--80aamc3a6ac9a.xn--p1acf" /* салехард.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5037] = Zone{"xn--80aaa0cvac.xn--p1acf" /* самара.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5038] = Zone{"xn--80aa4alnee.xn--p1acf" /* саранск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5039] = Zone{"xn--80aag1ciek.xn--p1acf" /* саратов.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5039() {
	z[5040] = Zone{"xn--80adi0aoagldk8i.xn--p1acf" /* севастополь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5041] = Zone{"xn--90anb6ab4e.xn--p1acf" /* сибирь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5042] = Zone{"xn--e1afkclaggf6a2g.xn--p1acf" /* симферополь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5043] = Zone{"xn--e1akbdhdtf.xn--p1acf" /* смоленск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5044] = Zone{"xn--h1aliz.xn--p1acf" /* сочи.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5045] = Zone{"xn--90a1af.xn--p1acf" /* спб.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5046] = Zone{"xn--80ae1alafffj1i.xn--p1acf" /* ставрополь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5047] = Zone{"xn--80aqialz.xn--p1acf" /* таллин.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5047() {
	z[5048] = Zone{"xn--80acd2blu.xn--p1acf" /* тамбов.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5049] = Zone{"xn--80akonve2b.xn--p1acf" /* ташкент.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5050] = Zone{"xn--b1ag7af7c.xn--p1acf" /* тверь.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5051] = Zone{"xn--j1adfnc.xn--p1acf" /* томск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5052] = Zone{"xn--80avue.xn--p1acf" /* тула.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5053] = Zone{"xn--p1add.xn--p1acf" /* тур.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5054] = Zone{"xn--e1aner7ci.xn--p1acf" /* тюмень.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5055] = Zone{"xn--b1aqbiftj7e1a.xn--p1acf" /* ульяновск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5055() {
	z[5056] = Zone{"xn--80a1bd.xn--p1acf" /* уфа.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5057] = Zone{"xn--80aacf4bwnk3a.xn--p1acf" /* хабаровск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5058] = Zone{"xn--80ablvtof7b4b.xn--p1acf" /* чебоксары.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5059] = Zone{"xn--90ahkico2a6b9d.xn--p1acf" /* челябинск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5060] = Zone{"xn--e1aaob1aia6b.xn--p1acf" /* черкесск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5061] = Zone{"xn--80ap4as.xn--p1acf" /* чита.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5062] = Zone{"xn--80apizf6d.xn--p1acf" /* элиста.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5063] = Zone{"xn--80af3b0c.xn--p1acf" /* югра.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5063() {
	z[5064] = Zone{"xn--41a.xn--p1acf" /* я.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5065] = Zone{"xn--j1aaude4e.xn--p1acf" /* якутск.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5066] = Zone{"xn--80aezclli6gta.xn--p1acf" /* ярославль.рус */, &z[1593], x, 0x200, e, e, w{"ns1.rlnic.ru", "ns2.rlnic.ru"}, n, n, n, e, e, n, t}
	z[5067] = Zone{"xn--j1aef.xn--p1ai" /* ком.рф */, &z[1594], x, 0x200, e, e, w{"ns1.nameself.com", "ns2.nameself.com"}, n, n, n, e, e, n, t}
	z[5068] = Zone{"xn--e1apq.xn--p1ai" /* нет.рф */, &z[1594], x, 0x200, e, e, w{"ns1.nameself.com", "ns2.nameself.com"}, n, n, n, e, e, n, t}
	z[5069] = Zone{"xn--c1avg.xn--p1ai" /* орг.рф */, &z[1594], x, 0x200, e, e, w{"ns1.nameself.com", "ns2.nameself.com"}, n, n, n, e, e, n, t}
	z[5070] = Zone{"xn--o-btb9b.xn--90a3ac" /* oбр.срб */, &z[1596], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5071] = Zone{"xn--o-htb.xn--90a3ac" /* oд.срб */, &z[1596], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5071() {
	z[5072] = Zone{"xn--80au.xn--90a3ac" /* ак.срб */, &z[1596], x, 0, e, e, w{"ban.junis.ni.ac.rs", "gaea.rcub.bg.ac.rs", "ns.etf.bg.ac.rs", "ns.rcub.bg.ac.rs", "ns.unic.kg.ac.rs", "ns1.uns.ac.rs", "ns2.iif.hu", "odisej.telekom.rs"}, n, n, n, e, e, n, t}
	z[5073] = Zone{"xn--o1ac.xn--90a3ac" /* пр.срб */, &z[1596], x, 0, e, e, w{"a.nic.rs", "g.nic.rs", "h.nic.rs", "l.nic.rs"}, n, n, n, e, e, n, t}
	z[5074] = Zone{"xn--o1ach.xn--90a3ac" /* упр.срб */, &z[1596], x, 0, e, e, w{"dagda1.isp.telekom.rs", "ns1.gov.rs", "ns1.nic.rs", "ns2.gov.rs", "ns3.gov.orion.rs"}, n, n, n, e, e, n, t}
	z[5075] = Zone{"xn--hhbc1oogwg.xn--mgbbh1a" /* کمپنی.بارت */, &z[1616], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5076] = Zone{"xn--fhbed7t1n.xn--mgbbh1a71e" /* كمپنی.بھارت */, &z[1619], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5077] = Zone{"xn--hhbcm0s3h.xn--mgbgu82a" /* ڪمپني.ڀارت */, &z[1640], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5078] = Zone{"xn--i1b1b4ch5i.xn--h2brj9c" /* कंपनी.भारत */, &z[1644], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5079] = Zone{"xn--i1b4e9bb5a7f.xn--h2breg3eve" /* संस्था.भारतम् */, &z[1645], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5079() {
	z[5080] = Zone{"xn--11b2bfp5fn6er.xn--h2brj9c8c" /* कोम्पानी.भारोत */, &z[1646], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5081] = Zone{"xn--p5b2bfp1g0b0b.xn--45brj9c" /* কম্পেনি.ভারত */, &z[1649], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5082] = Zone{"xn--p5b2bfp5fh3fra.xn--45brj9c" /* কোম্পানি.ভারত */, &z[1649], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5083] = Zone{"xn--p5b2bfp5fn6er.xn--45br5cyl" /* কোম্পানী.ভাৰত */, &z[1650], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5084] = Zone{"xn--d9b2bf3g1k.xn--s9brj9c" /* ਕੰਪਨੀ.ਭਾਰਤ */, &z[1651], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5085] = Zone{"xn--hdc1b4ch5i.xn--gecrj9c" /* કંપની.ભારત */, &z[1652], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5086] = Zone{"xn--ohc2bfp5fn1g.xn--3hcrj9c" /* କମ୍ପାନୀ.ଭାରତ */, &z[1653], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5087] = Zone{"xn--vlccpku2dp3h.xn--xkc2dl3a5ee0h" /* நிறுவனம்.இந்தியா */, &z[1654], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5087() {
	z[5088] = Zone{"xn--goc1b4ch5i8a.xn--fpcrj9c3d" /* కంపెనీ.భారత్ */, &z[1658], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5089] = Zone{"xn--3rc1b4ch0i.xn--2scrj9c" /* ಕಂಪನಿ.ಭಾರತ */, &z[1659], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5090] = Zone{"xn--bwc2bah0ipd.xn--rvc1e0am3e" /* കന്പനി.ഭാരതം */, &z[1660], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5091] = Zone{"xn--12co0c3b4eva.xn--o3cw4h" /* ธุรกิจ.ไทย */, &z[1663], x, 0, e, e, w{"a.thains.co.th", "b.thains.co.th", "p.thains.co.th"}, n, n, n, e, e, n, t}
	z[5092] = Zone{"xn--gmqw5a.xn--j6w193g" /* 個人.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[5093] = Zone{"xn--55qx5d.xn--j6w193g" /* 公司.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[5094] = Zone{"xn--mxtq1m.xn--j6w193g" /* 政府.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[5095] = Zone{"xn--wcvs22d.xn--j6w193g" /* 教育.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5095() {
	z[5096] = Zone{"xn--uc0atv.xn--j6w193g" /* 組織.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[5097] = Zone{"xn--od0alg.xn--j6w193g" /* 網絡.香港 */, &z[1749], x, 0, e, e, w{"c.hkirc.net.hk", "d.hkirc.net.hk", "m.hkirc.net.hk", "t.hkirc.net.hk", "u.hkirc.net.hk", "v.hkirc.net.hk", "x.hkirc.net.hk", "y.hkirc.net.hk", "z.hkirc.net.hk"}, n, n, n, e, e, n, t}
	z[5098] = Zone{"lda.co.ao", &z[1816], x, 0, e, e, w{"ns.inventadomains.com", "ns2.inventadomains.com"}, n, n, n, e, e, n, t}
	z[5099] = Zone{"sa.co.ao", &z[1816], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5100] = Zone{"esc.edu.ar", &z[1824], x, 0, e, e, w{"ns1.innova-red.net", "ns2.innova-red.net"}, n, n, n, e, e, n, t}
	z[5101] = Zone{"empty.as112.arpa", &z[1833], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc7535", w{"blackhole.as112.arpa"}, n, n, n, e, e, n, t}
	z[5102] = Zone{"10.in-addr.arpa", &z[1837], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5103] = Zone{"169.in-addr.arpa", &z[1837], z[5203:5204], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", w{"arin.authdns.ripe.net", "r.arin.net", "u.arin.net", "x.arin.net", "y.arin.net", "z.arin.net"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5103() {
	z[5104] = Zone{"172.in-addr.arpa", &z[1837], z[5204:5220], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"arin.authdns.ripe.net", "r.arin.net", "u.arin.net", "x.arin.net", "y.arin.net", "z.arin.net"}, n, n, n, e, e, n, t}
	z[5105] = Zone{"192.in-addr.arpa", &z[1837], z[5220:5222], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"arin.authdns.ripe.net", "r.arin.net", "u.arin.net", "x.arin.net", "y.arin.net", "z.arin.net"}, n, n, n, e, e, n, t}
	z[5106] = Zone{"f.ip6.arpa", &z[1839], z[5222:5223], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5107] = Zone{"act.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5108] = Zone{"nsw.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5109] = Zone{"nt.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5110] = Zone{"qld.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5111] = Zone{"sa.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5111() {
	z[5112] = Zone{"tas.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5113] = Zone{"vic.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5114] = Zone{"wa.edu.au", &z[1857], x, 0, e, e, w{"a.au", "q.au", "r.au", "s.au", "t.au"}, n, n, n, e, e, n, t}
	z[5115] = Zone{"fed.rep.br", &z[2121], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5116] = Zone{"wireless.keyword.cn", &z[2245], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5117] = Zone{"ltd.co.im", &z[2700], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5118] = Zone{"plc.co.im", &z[2700], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5119] = Zone{"lg.gov.ng", &z[3597], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5119() {
	z[5120] = Zone{"co.net.nz", &z[3805], x, 0, e, e, w{"ns1.instradns.com", "ns2.instradns.com", "ns3.instradns.com"}, n, n, n, e, e, n, t}
	z[5121] = Zone{"al.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5122] = Zone{"asso.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5123] = Zone{"at.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5124] = Zone{"au.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5125] = Zone{"be.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5126] = Zone{"bg.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5127] = Zone{"ca.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5127() {
	z[5128] = Zone{"cd.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5129] = Zone{"ch.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5130] = Zone{"cn.eu.org", &z[3823], x, 0, e, e, w{"ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5131] = Zone{"cy.eu.org", &z[3823], x, 0, e, e, w{"ns.eu.org", "ns1.absolight.net", "ns1.tee.gr", "ns2.absolight.net", "ns2.tee.gr", "ns3.absolight.net", "ns4.absolight.net"}, n, n, n, e, e, n, t}
	z[5132] = Zone{"cz.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5133] = Zone{"de.eu.org", &z[3823], x, 0, e, e, w{"ns.oberon.net", "ns1.absolight.net", "ns2.absolight.net", "ns2.lf.net", "ns3.absolight.net", "ns4.absolight.net"}, n, n, n, e, e, n, t}
	z[5134] = Zone{"dk.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5135] = Zone{"edu.eu.org", &z[3823], x, 0, e, e, w{"dns.elm.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5135() {
	z[5136] = Zone{"ee.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5137] = Zone{"es.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5138] = Zone{"fi.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5139] = Zone{"fr.eu.org", &z[3823], x, 0, e, e, w{"dns.elm.net", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5140] = Zone{"gr.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns.eu.org", "ns1.absolight.net", "ns1.tee.gr", "ns2.absolight.net", "ns2.tee.gr", "ns3.absolight.net", "ns4.absolight.net"}, n, n, n, e, e, n, t}
	z[5141] = Zone{"hr.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5142] = Zone{"hu.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5143] = Zone{"ie.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5143() {
	z[5144] = Zone{"il.eu.org", &z[3823], x, 0, e, e, w{"ns.eu.org", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5145] = Zone{"in.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5146] = Zone{"int.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5147] = Zone{"is.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5148] = Zone{"it.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5149] = Zone{"jp.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5150] = Zone{"kr.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5151] = Zone{"lt.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5151() {
	z[5152] = Zone{"lu.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5153] = Zone{"lv.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5154] = Zone{"me.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5155] = Zone{"mk.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5156] = Zone{"mt.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5157] = Zone{"my.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5158] = Zone{"net.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5159] = Zone{"ng.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5159() {
	z[5160] = Zone{"nl.eu.org", &z[3823], x, 0, e, e, w{"dns.elm.net", "ns.eu.org", "ns.lf.net", "stitch.dns.badassops.com"}, n, n, n, e, e, n, t}
	z[5161] = Zone{"no.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5162] = Zone{"nz.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5163] = Zone{"pl.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5164] = Zone{"pt.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5165] = Zone{"ro.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5166] = Zone{"ru.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5167] = Zone{"se.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5167() {
	z[5168] = Zone{"si.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5169] = Zone{"sk.eu.org", &z[3823], x, 0, e, e, w{"hobbes.bsd-dk.dk", "ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5170] = Zone{"tr.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5171] = Zone{"uk.eu.org", &z[3823], x, 0, e, e, w{"ns1.eriomem.net", "ns1.eu.org", "ns3.keltia.net"}, n, n, n, e, e, n, t}
	z[5172] = Zone{"us.eu.org", &z[3823], x, 0, e, e, w{"ns.ankh.fr.eu.org", "ns1.eriomem.net", "ns1.eu.org"}, n, n, n, e, e, n, t}
	z[5173] = Zone{"co.at.tt", &z[4550], x, 0, e, e, n, w{"195.225.236.206"}, n, n, e, e, n, t}
	z[5174] = Zone{"com.au.tt", &z[4551], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5175] = Zone{"co.uk.tt", &z[4579], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5175() {
	z[5176] = Zone{"me.uk.tt", &z[4579], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5177] = Zone{"org.uk.tt", &z[4579], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5178] = Zone{"barking-dagenham.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5179] = Zone{"barnet.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5180] = Zone{"barnsley.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5181] = Zone{"bathnes.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5182] = Zone{"beds.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5183] = Zone{"bexley.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5183() {
	z[5184] = Zone{"bham.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5185] = Zone{"blackburn.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5186] = Zone{"blackpool.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5187] = Zone{"bolton.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5188] = Zone{"bournemouth.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5189] = Zone{"bracknell-forest.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5190] = Zone{"bradford.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5191] = Zone{"brent.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
}

//go:noinline
func i5191() {
	z[5192] = Zone{"doncaster.sch.uk", &z[4690], x, 0, e, e, n, n, n, n, e, e, n, t}
	z[5193] = Zone{"bg.ac.yu", &z[4822], x, 0x1000, e, e, n, n, n, n, e, e, n, t}
	z[5194] = Zone{"ecape.school.za", &z[4852], x, 0, e, e, w{"ns10.dns.business", "ns11.dns.business"}, n, n, n, e, e, n, t}
	z[5195] = Zone{"fs.school.za", &z[4852], x, 0, e, e, w{"ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business", "nyx.marques.co.za", "ucthpx.uct.ac.za"}, n, n, n, e, e, n, t}
	z[5196] = Zone{"gp.school.za", &z[4852], x, 0, e, e, w{"ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business", "ns2.obsidian.co.za"}, n, n, n, e, e, n, t}
	z[5197] = Zone{"kzn.school.za", &z[4852], x, 0, e, e, w{"ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business", "sec.zadns.net", "ucthpx.uct.ac.za"}, n, n, n, e, e, n, t}
	z[5198] = Zone{"lp.school.za", &z[4852], x, 0, e, e, w{"lava.obsidian.co.za", "ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business"}, n, n, n, e, e, n, t}
	z[5199] = Zone{"mpm.school.za", &z[4852], x, 0, e, e, w{"lava.obsidian.co.za", "ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5199() {
	z[5200] = Zone{"ncape.school.za", &z[4852], x, 0, e, e, w{"ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business", "ucthpx.uct.ac.za"}, n, n, n, e, e, n, t}
	z[5201] = Zone{"nw.school.za", &z[4852], x, 0, e, e, w{"lava.obsidian.co.za", "ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business"}, n, n, n, e, e, n, t}
	z[5202] = Zone{"wcape.school.za", &z[4852], x, 0, e, e, w{"ns0.serendipity.org.za", "ns10.dns.business", "ns11.dns.business", "ucthpx.uct.ac.za"}, n, n, n, e, e, n, t}
	z[5203] = Zone{"254.169.in-addr.arpa", &z[5103], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5204] = Zone{"16.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5205] = Zone{"17.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5206] = Zone{"18.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5207] = Zone{"19.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5207() {
	z[5208] = Zone{"20.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5209] = Zone{"21.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5210] = Zone{"22.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5211] = Zone{"23.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5212] = Zone{"24.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5213] = Zone{"25.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5214] = Zone{"26.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5215] = Zone{"27.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
}

//go:noinline
func i5215() {
	z[5216] = Zone{"28.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5217] = Zone{"29.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5218] = Zone{"30.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5219] = Zone{"31.172.in-addr.arpa", &z[5104], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5220] = Zone{"0.192.in-addr.arpa", &z[5105], z[5223:5224], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8880", n, n, n, n, e, e, n, t}
	z[5221] = Zone{"168.192.in-addr.arpa", &z[5105], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6761", w{"blackhole-1.iana.org", "blackhole-2.iana.org"}, n, n, n, e, e, n, t}
	z[5222] = Zone{"e.f.ip6.arpa", &z[5106], z[5224:5228], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5223] = Zone{"0.0.192.in-addr.arpa", &z[5220], z[5228:5230], 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8880", n, n, n, n, e, e, n, t}
}

//go:noinline
func i5223() {
	z[5224] = Zone{"8.e.f.ip6.arpa", &z[5222], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5225] = Zone{"9.e.f.ip6.arpa", &z[5222], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5226] = Zone{"a.e.f.ip6.arpa", &z[5222], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5227] = Zone{"b.e.f.ip6.arpa", &z[5222], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc6762", n, n, n, n, e, e, n, t}
	z[5228] = Zone{"170.0.0.192.in-addr.arpa", &z[5223], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8880", n, n, n, n, e, e, n, t}
	z[5229] = Zone{"171.0.0.192.in-addr.arpa", &z[5223], x, 0x108, "IANA", "https://datatracker.ietf.org/doc/rfc8880", n, n, n, n, e, e, n, t}

}

func initZones() {
	i0()
	i7()
	i15()
	i23()
	i31()
	i39()
	i47()
	i55()
	i63()
	i71()
	i79()
	i87()
	i95()
	i103()
	i111()
	i119()
	i127()
	i135()
	i143()
	i151()
	i159()
	i167()
	i175()
	i183()
	i191()
	i199()
	i207()
	i215()
	i223()
	i231()
	i239()
	i247()
	i255()
	i263()
	i271()
	i279()
	i287()
	i295()
	i303()
	i311()
	i319()
	i327()
	i335()
	i343()
	i351()
	i359()
	i367()
	i375()
	i383()
	i391()
	i399()
	i407()
	i415()
	i423()
	i431()
	i439()
	i447()
	i455()
	i463()
	i471()
	i479()
	i487()
	i495()
	i503()
	i511()
	i519()
	i527()
	i535()
	i543()
	i551()
	i559()
	i567()
	i575()
	i583()
	i591()
	i599()
	i607()
	i615()
	i623()
	i631()
	i639()
	i647()
	i655()
	i663()
	i671()
	i679()
	i687()
	i695()
	i703()
	i711()
	i719()
	i727()
	i735()
	i743()
	i751()
	i759()
	i767()
	i775()
	i783()
	i791()
	i799()
	i807()
	i815()
	i823()
	i831()
	i839()
	i847()
	i855()
	i863()
	i871()
	i879()
	i887()
	i895()
	i903()
	i911()
	i919()
	i927()
	i935()
	i943()
	i951()
	i959()
	i967()
	i975()
	i983()
	i991()
	i999()
	i1007()
	i1015()
	i1023()
	i1031()
	i1039()
	i1047()
	i1055()
	i1063()
	i1071()
	i1079()
	i1087()
	i1095()
	i1103()
	i1111()
	i1119()
	i1127()
	i1135()
	i1143()
	i1151()
	i1159()
	i1167()
	i1175()
	i1183()
	i1191()
	i1199()
	i1207()
	i1215()
	i1223()
	i1231()
	i1239()
	i1247()
	i1255()
	i1263()
	i1271()
	i1279()
	i1287()
	i1295()
	i1303()
	i1311()
	i1319()
	i1327()
	i1335()
	i1343()
	i1351()
	i1359()
	i1367()
	i1375()
	i1383()
	i1391()
	i1399()
	i1407()
	i1415()
	i1423()
	i1431()
	i1439()
	i1447()
	i1455()
	i1463()
	i1471()
	i1479()
	i1487()
	i1495()
	i1503()
	i1511()
	i1519()
	i1527()
	i1535()
	i1543()
	i1551()
	i1559()
	i1567()
	i1575()
	i1583()
	i1591()
	i1599()
	i1607()
	i1615()
	i1623()
	i1631()
	i1639()
	i1647()
	i1655()
	i1663()
	i1671()
	i1679()
	i1687()
	i1695()
	i1703()
	i1711()
	i1719()
	i1727()
	i1735()
	i1743()
	i1751()
	i1759()
	i1767()
	i1775()
	i1783()
	i1791()
	i1799()
	i1807()
	i1815()
	i1823()
	i1831()
	i1839()
	i1847()
	i1855()
	i1863()
	i1871()
	i1879()
	i1887()
	i1895()
	i1903()
	i1911()
	i1919()
	i1927()
	i1935()
	i1943()
	i1951()
	i1959()
	i1967()
	i1975()
	i1983()
	i1991()
	i1999()
	i2007()
	i2015()
	i2023()
	i2031()
	i2039()
	i2047()
	i2055()
	i2063()
	i2071()
	i2079()
	i2087()
	i2095()
	i2103()
	i2111()
	i2119()
	i2127()
	i2135()
	i2143()
	i2151()
	i2159()
	i2167()
	i2175()
	i2183()
	i2191()
	i2199()
	i2207()
	i2215()
	i2223()
	i2231()
	i2239()
	i2247()
	i2255()
	i2263()
	i2271()
	i2279()
	i2287()
	i2295()
	i2303()
	i2311()
	i2319()
	i2327()
	i2335()
	i2343()
	i2351()
	i2359()
	i2367()
	i2375()
	i2383()
	i2391()
	i2399()
	i2407()
	i2415()
	i2423()
	i2431()
	i2439()
	i2447()
	i2455()
	i2463()
	i2471()
	i2479()
	i2487()
	i2495()
	i2503()
	i2511()
	i2519()
	i2527()
	i2535()
	i2543()
	i2551()
	i2559()
	i2567()
	i2575()
	i2583()
	i2591()
	i2599()
	i2607()
	i2615()
	i2623()
	i2631()
	i2639()
	i2647()
	i2655()
	i2663()
	i2671()
	i2679()
	i2687()
	i2695()
	i2703()
	i2711()
	i2719()
	i2727()
	i2735()
	i2743()
	i2751()
	i2759()
	i2767()
	i2775()
	i2783()
	i2791()
	i2799()
	i2807()
	i2815()
	i2823()
	i2831()
	i2839()
	i2847()
	i2855()
	i2863()
	i2871()
	i2879()
	i2887()
	i2895()
	i2903()
	i2911()
	i2919()
	i2927()
	i2935()
	i2943()
	i2951()
	i2959()
	i2967()
	i2975()
	i2983()
	i2991()
	i2999()
	i3007()
	i3015()
	i3023()
	i3031()
	i3039()
	i3047()
	i3055()
	i3063()
	i3071()
	i3079()
	i3087()
	i3095()
	i3103()
	i3111()
	i3119()
	i3127()
	i3135()
	i3143()
	i3151()
	i3159()
	i3167()
	i3175()
	i3183()
	i3191()
	i3199()
	i3207()
	i3215()
	i3223()
	i3231()
	i3239()
	i3247()
	i3255()
	i3263()
	i3271()
	i3279()
	i3287()
	i3295()
	i3303()
	i3311()
	i3319()
	i3327()
	i3335()
	i3343()
	i3351()
	i3359()
	i3367()
	i3375()
	i3383()
	i3391()
	i3399()
	i3407()
	i3415()
	i3423()
	i3431()
	i3439()
	i3447()
	i3455()
	i3463()
	i3471()
	i3479()
	i3487()
	i3495()
	i3503()
	i3511()
	i3519()
	i3527()
	i3535()
	i3543()
	i3551()
	i3559()
	i3567()
	i3575()
	i3583()
	i3591()
	i3599()
	i3607()
	i3615()
	i3623()
	i3631()
	i3639()
	i3647()
	i3655()
	i3663()
	i3671()
	i3679()
	i3687()
	i3695()
	i3703()
	i3711()
	i3719()
	i3727()
	i3735()
	i3743()
	i3751()
	i3759()
	i3767()
	i3775()
	i3783()
	i3791()
	i3799()
	i3807()
	i3815()
	i3823()
	i3831()
	i3839()
	i3847()
	i3855()
	i3863()
	i3871()
	i3879()
	i3887()
	i3895()
	i3903()
	i3911()
	i3919()
	i3927()
	i3935()
	i3943()
	i3951()
	i3959()
	i3967()
	i3975()
	i3983()
	i3991()
	i3999()
	i4007()
	i4015()
	i4023()
	i4031()
	i4039()
	i4047()
	i4055()
	i4063()
	i4071()
	i4079()
	i4087()
	i4095()
	i4103()
	i4111()
	i4119()
	i4127()
	i4135()
	i4143()
	i4151()
	i4159()
	i4167()
	i4175()
	i4183()
	i4191()
	i4199()
	i4207()
	i4215()
	i4223()
	i4231()
	i4239()
	i4247()
	i4255()
	i4263()
	i4271()
	i4279()
	i4287()
	i4295()
	i4303()
	i4311()
	i4319()
	i4327()
	i4335()
	i4343()
	i4351()
	i4359()
	i4367()
	i4375()
	i4383()
	i4391()
	i4399()
	i4407()
	i4415()
	i4423()
	i4431()
	i4439()
	i4447()
	i4455()
	i4463()
	i4471()
	i4479()
	i4487()
	i4495()
	i4503()
	i4511()
	i4519()
	i4527()
	i4535()
	i4543()
	i4551()
	i4559()
	i4567()
	i4575()
	i4583()
	i4591()
	i4599()
	i4607()
	i4615()
	i4623()
	i4631()
	i4639()
	i4647()
	i4655()
	i4663()
	i4671()
	i4679()
	i4687()
	i4695()
	i4703()
	i4711()
	i4719()
	i4727()
	i4735()
	i4743()
	i4751()
	i4759()
	i4767()
	i4775()
	i4783()
	i4791()
	i4799()
	i4807()
	i4815()
	i4823()
	i4831()
	i4839()
	i4847()
	i4855()
	i4863()
	i4871()
	i4879()
	i4887()
	i4895()
	i4903()
	i4911()
	i4919()
	i4927()
	i4935()
	i4943()
	i4951()
	i4959()
	i4967()
	i4975()
	i4983()
	i4991()
	i4999()
	i5007()
	i5015()
	i5023()
	i5031()
	i5039()
	i5047()
	i5055()
	i5063()
	i5071()
	i5079()
	i5087()
	i5095()
	i5103()
	i5111()
	i5119()
	i5127()
	i5135()
	i5143()
	i5151()
	i5159()
	i5167()
	i5175()
	i5183()
	i5191()
	i5199()
	i5207()
	i5215()
	i5223()

	ZoneMap = make(map[string]*Zone, len(z))
	for i := range z {
		ZoneMap[z[i].Domain] = &z[i]
	}
}
