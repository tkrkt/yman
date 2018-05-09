package api

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

var config *model.Config

// GetConfig loads configurations from ~/.ymanrc
func GetConfig() model.Config {
	if config != nil {
		return *config
	}
	config = model.NewConfig()

	path, err := getRcPath()
	if err != nil {
		ui.Error(err)
		return *config
	}

	if _, err := os.Stat(path); err == nil {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			ui.Error(err)
			return *config
		}
		if _, err := toml.Decode(string(data), &config); err != nil {
			ui.Error(err)
			return *config
		}
	}

	return *config
}

// SetConfig saves configurations into ~/.ymanrc
func SetConfig(c model.Config) error {
	path, err := getRcPath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		return err
	}

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		return err
	}

	config = &c

	return nil
}

func getRcPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".ymanrc"), nil
}
