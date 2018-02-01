package build

import "sort"

const (
	// TypeLength specifies a length constraint for labels.
	// The Policy.Key will be "min" or "max".
	TypeLength = "length"

	// KeyMin specifies a mininum length constraint for a label.
	KeyMin = "min"

	// KeyMax specifies a maximum length constraint for a label.
	KeyMax = "max"

	// TypeIDNTable specifies a location for an IDN table a language or script.
	TypeIDNTable = "idn-table"

	// TypeIDNDisallowed specifies that IDN labels are disallowed by the operator for this zone.
	TypeIDNDisallowed = "idn-disallowed"
)

// Policy represents a registry or technical policy applicable to subdomains of a zone.
type Policy struct {
	Type    string `json:"type"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`

	// Transitional
	Language string `json:"language,omitempty"`
}

func sortPolicies(policies []Policy) {
	sort.SliceStable(policies, func(i, j int) bool {
		a := policies[i]
		b := policies[j]
		switch {
		case a.Type != b.Type:
			return a.Type < b.Type
		case a.Key != b.Key:
			return a.Key < b.Key
		case a.Value != b.Value:
			return a.Value < b.Value
		case a.Comment != b.Comment:
			return a.Comment < b.Comment
		}
		return false
	})
}
