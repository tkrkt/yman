package api

import (
	"errors"
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
)

func Login(username string, password string) (*model.User, error) {
	current := CurrentUser()
	if current != nil {
		return nil, errors.New("already logined as " + current.Username)
	}

	// access to server

	// write to .ymanrc
	home, err := homedir.Dir()
	if err != nil {
		return nil, errors.New("failed to save login status")
	}

	p := filepath.Join(home, ".ymanrc")
	fmt.Println(p)

	return nil, nil
}

func CurrentUser() *model.User {
	return nil
}
