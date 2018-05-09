package ui

import (
	"fmt"

	"github.com/mgutz/ansi"

	"github.com/AlecAivazis/survey"
)

// Text shows simple text.
func Text(text interface{}) {
	fmt.Println(text)
}

// Warn shows warning text. It might be styled.
func Warn(text interface{}) {
	fmt.Println(ansi.Yellow + "WARN: " + fmt.Sprint(text) + ansi.Reset)
}

// Error shows error text. It might be styled.
func Error(text interface{}) {
	fmt.Println(ansi.Red + "ERROR: " + fmt.Sprint(text) + ansi.Reset)
}

// Confirm displays simple confirmation to user
func Confirm(msg string) (answer bool, err error) {
	qs := &survey.Confirm{
		Message: msg,
	}
	err = survey.AskOne(qs, &answer, nil)
	if err != nil {
		return false, nil
	}
	return answer, nil
}
