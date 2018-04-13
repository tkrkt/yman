package ui

import (
	"github.com/AlecAivazis/survey"
)

func Login() (username string, password string, err error) {
	qs := []*survey.Question{
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "username"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "password"},
			Validate: survey.Required,
		},
	}

	ans := struct {
		Username string
		Password string
	}{}

	err = survey.Ask(qs, &ans)
	if err != nil {
		return "", "", err
	}

	return ans.Username, ans.Password, nil
}
