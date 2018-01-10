package build

// Set implements a unique set of strings.
type Set map[string]bool

// NewSet creates a new set from the supplied values.
func NewSet(vals ...string) Set {
	s := make(Set, len(vals))
	s.Add(vals...)
	return s
}

// Add inserts a string into a set. Underlying map enforces uniqueness.
func (set Set) Add(vals ...string) {
	for _, s := range vals {
		set[s] = true
	}
}

// Values returns a sorted slice of the set’s values.
func (set Set) Values() []string {
	vals := make([]string, 0, len(set))
	for s, _ := range set {
		vals = append(vals, s)
	}
	Sort(vals) // domain rank (not lexical) sort
	return vals
}
