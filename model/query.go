package model

type Query struct {
	Command string
	Author  string
	Tag     string
}

// IsEmpty returns true if any condition is specified
func (q *Query) IsEmpty() bool {
	return q.Command == "" &&
		q.Author == "" &&
		q.Tag == ""
}
