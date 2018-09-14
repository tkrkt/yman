package model

// Manual represents a manual document created by current user or others.
type Manual struct {
	ID      *string
	Command string // e.g. ls
	Full    string // e.g. ls/all
	Tags    []string
	Title   string // first line of the message
	Message string
}
