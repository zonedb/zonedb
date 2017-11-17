package build

// Policy represents a registry or technical policy applicable to subdomains of a zone.
type Policy struct {
	Type     string `json:"type"`
	Value    string `json:"value,omitempty"`
	Language string `json:"language,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

const (
	TypeMinLength     = "min-length"
	TypeIDNTable      = "idn-table"
	TypeIDNDisallowed = "idn-disallowed"
)
