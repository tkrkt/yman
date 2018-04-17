package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey"
)

func Text(text interface{}) {
	fmt.Println(text)
}

func Error(text interface{}) {
	fmt.Println(text)
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
