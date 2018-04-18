package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
)

// Search manuals by fetching from server or search local files
func Search(account *model.Account, query *model.Query) ([]*model.Manual, error) {
	if query == nil || query.IsEmpty() {
		return nil, errors.New("invalid query")
	}
	return loadFromLocalFile(account, query)
}

func loadFromLocalFile(account *model.Account, query *model.Query) ([]*model.Manual, error) {
	// create file
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(home, ".yman")

	// return no results if directory is not exist (it is not an error)
	if _, err := os.Stat(dir); err != nil {
		return nil, nil
	}

	// read all files
	allFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// search files matching the query
	var manuals []*model.Manual
	for _, file := range allFiles {
		if file.IsDir() {
			continue
		}

		manual := &model.Manual{}
		_, err := toml.DecodeFile(filepath.Join(dir, file.Name()), manual)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if match(manual, query) {
			manuals = append(manuals, manual)
		}
	}

	return manuals, nil
}

func match(m *model.Manual, q *model.Query) bool {
	return true
}
