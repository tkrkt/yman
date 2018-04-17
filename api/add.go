package api

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
)

func Add(account *model.Account, manual *model.Manual) error {
	return saveToLocalFile(account, manual)
}

func saveToLocalFile(account *model.Account, manual *model.Manual) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".yman")

	if _, err := os.Stat(dir); err != nil {
		if ioErr := os.Mkdir(dir, 0777); ioErr != nil {
			return ioErr
		}
	}

	return nil
}
