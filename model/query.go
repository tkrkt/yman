package model

// Query is used for searching.
// Empty string or empty array does not restrict the condition.
type Query struct {
	Command string
	Tags    []string
}

// IsEmpty returns true if no conditions is specified.
// No conditions mean "all items" on list command
// and mean "invalid query (error)" on show and search command
func (q *Query) IsEmpty() bool {
	return q.Command == "" &&
		len(q.Tags) == 0
}
