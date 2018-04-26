package ui

import (
	"fmt"

	"github.com/mgutz/ansi"

	"github.com/AlecAivazis/survey"
)

func Text(text interface{}) {
	fmt.Println(text)
}

func Warn(text interface{}) {
	fmt.Println(ansi.Yellow + "WARN: " + fmt.Sprint(text) + ansi.Reset)
}

func Error(text interface{}) {
	fmt.Println(ansi.Red + "ERROR: " + fmt.Sprint(text) + ansi.Reset)
}

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
