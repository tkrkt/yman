package api

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
)

// Delete a manual by posting to server or deleting a file
func Delete(manual *model.Manual) error {
	if !IsLogined() {
		return errors.New("not loggined")
	}
	if manual == nil || manual.ID == nil {
		return errors.New("invalid manual")
	}

	return deleteLocalFile(manual)
}

func deleteLocalFile(manual *model.Manual) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".yman")

	path := filepath.Join(dir, *manual.ID+".txt")

	fmt.Println(path)
	if _, err := os.Stat(path); err != nil {
		return errors.New("file not found")
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
