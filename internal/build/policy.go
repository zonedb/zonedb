package build

import "sort"

// Policy represents a registry or technical policy applicable to subdomains of a zone.
type Policy struct {
	Type     string `json:"type"`
	Value    string `json:"value,omitempty"`
	Language string `json:"language,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

const (
	// TypeMinLength specifies the minimum number of characters required for labels beneath this zone.
	TypeMinLength = "min-length"

	// TypeIDNTable specifies a location for an IDN table a language or script.
	TypeIDNTable = "idn-table"

	// TypeIDNDisallowed specifies that IDN labels are disallowed by the operator for this zone.
	TypeIDNDisallowed = "idn-disallowed"
)

func sortPolicies(policies []Policy) {
	sort.SliceStable(policies, func(i, j int) bool {
		a := policies[i]
		b := policies[j]
		switch {
		case a.Type != b.Type:
			return a.Type < b.Type
		case a.Language != b.Language:
			return a.Language < b.Language
		case a.Value != b.Value:
			return a.Value < b.Value
		case a.Comment != b.Comment:
			return a.Comment < b.Comment
		}
		return false
	})
}
