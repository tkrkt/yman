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

const templateText = `[[auth]]
username = "{{ .username }}"
token = "{{ .token }}"`

func Login(username string, password string) (*model.User, error) {
	current := CurrentUser()
	if current != nil {
		return nil, errors.New("you are already logined as " + current.Username)
	}

	// access to server

	// write to .ymanrc
	if err := write(username, "token"); err != nil {
		return nil, err
	}

	return &model.User{Username: username}, nil
}

func CurrentUser() *model.User {
	username, _, err := read()
	if err != nil {
		return nil
	}

	return &model.User{Username: username}
}

func rcPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".ymanrc"), nil
}

func write(username string, token string) error {
	const t = `[[account]]
username = "{{.Username}}"
token = "{{.Token}}"
`
	tmpl, parseErr := template.New("test").Parse(t)
	if parseErr != nil {
		return parseErr
	}
	var doc bytes.Buffer

	type rc struct {
		Username string
		Token    string
	}

	if err := tmpl.Execute(&doc, rc{Username: username, Token: token}); err != nil {
		return err
	}

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

func read() (username string, token string, err error) {
	p, pathErr := rcPath()
	if pathErr != nil {
		return "", "", pathErr
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		return "", "", nil
	}
	conf, err := toml.Load(string(data))
	if err != nil {
		return "", "", nil
	}

	username = conf.Get("account.username").(string)
	token = conf.Get("account.token").(string)

	return username, token, nil
}
