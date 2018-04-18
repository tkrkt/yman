package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey"
)

func Text(text interface{}) {
	fmt.Println(text)
}

func Warn(text interface{}) {
	fmt.Println("WARN:", text)
}

func Error(text interface{}) {
	fmt.Println("ERROR:", text)
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
