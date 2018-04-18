package model

type Manual struct {
	Command     string // e.g. ls
	FullCommand string // e.g. ls/all
	Author      string
	Tags        []string
	Title       string // first line of the message
	Message     string
}
