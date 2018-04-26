package ui

import (
	"github.com/AlecAivazis/survey"
)

func Login() (email string, password string, err error) {
	qs := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "email"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "password"},
			Validate: survey.Required,
		},
	}

	ans := struct {
		Email    string
		Password string
	}{}

	err = survey.Ask(qs, &ans)
	if err != nil {
		return "", "", err
	}

	return ans.Email, ans.Password, nil
}
