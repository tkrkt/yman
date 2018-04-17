package api

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"github.com/tkrkt/yman/model"
)

func Login(username string, password string) (*model.Account, error) {
	// access to server

	account := &model.Account{
		Username: username,
		Token:    "token",
	}

	// write to .ymanrc
	if err := saveAccount(account); err != nil {
		return nil, err
	}

	return account, nil
}

func Logout() error {
	// delete .ymanrc
	if err := deleteAccount(); err != nil {
		return err
	}
	return nil
}

func CurrentAccount() *model.Account {
	account, err := loadAccount()
	if err != nil {
		return nil
	}
	return account
}

func rcPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".ymanrc"), nil
}

func saveAccount(account *model.Account) error {
	// generate toml string
	const t = `username = "{{.Username}}"
token = "{{.Token}}"
`
	tmpl, parseErr := template.New("test").Parse(t)
	if parseErr != nil {
		return parseErr
	}
	var doc bytes.Buffer

	if err := tmpl.Execute(&doc, account); err != nil {
		return err
	}

	// save to file
	p, pathErr := rcPath()
	if pathErr != nil {
		return pathErr
	}

	file, ioErr := os.Create(p)
	defer file.Close()
	if ioErr != nil {
		return ioErr
	}

	if _, err := file.Write(([]byte)(doc.String())); err != nil {
		return err
	}

	return nil
}

func loadAccount() (*model.Account, error) {
	p, pathErr := rcPath()
	if pathErr != nil {
		return nil, pathErr
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	conf, err := toml.Load(string(data))
	if err != nil {
		return nil, err
	}

	username := conf.Get("username")
	token := conf.Get("token")

	if username == nil || token == nil {
		return nil, errors.New("invalid file format")
	}

	return &model.Account{
		Username: username.(string),
		Token:    token.(string),
	}, nil
}

func deleteAccount() error {
	p, pathErr := rcPath()
	if pathErr != nil {
		return pathErr
	}

	if err := os.Remove(p); err != nil {
		return err
	}

	return nil
}
