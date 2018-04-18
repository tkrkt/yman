package model

type Query struct {
	Command *string
	Author  *string
	Tag     *string
}

// IsEmpty returns true if any condition is specified
func (q *Query) IsEmpty() bool {
	return (q.Command == nil || *q.Command == "") &&
		(q.Author == nil || *q.Author == "") &&
		(q.Tag == nil || *q.Tag == "")
}
