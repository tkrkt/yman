package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/fatih/color"
)

func Text(text interface{}) {
	fmt.Println(text)
}

func Warn(text interface{}) {
	color.Yellow("WARN: " + fmt.Sprint(text))
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
