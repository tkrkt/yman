package model

type Query struct {
	Command string
	Author  string
	Tags    []string
}

// IsEmpty returns true if no conditions is specified.
// No conditions mean "all items" on list command
// and mean "invalid query (error)" on show and search command
func (q *Query) IsEmpty() bool {
	return q.Command == "" &&
		q.Author == "" &&
		len(q.Tags) == 0
}
