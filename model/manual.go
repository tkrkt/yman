package model

type Manual struct {
	ID      *string
	Command string // e.g. ls
	Full    string // e.g. ls/all
	Author  string
	Tags    []string
	Title   string // first line of the message
	Message string
}
