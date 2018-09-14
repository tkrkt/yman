package api

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/tkrkt/yman/model"
)

// Add a manual by posting to server or saving to file
func Add(manual *model.Manual) error {
	if manual == nil {
		return errors.New("invalid manual")
	}
	return saveToLocalFile(manual)
}

func saveToLocalFile(manual *model.Manual) error {
	// create dir
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".yman")

	if _, err := os.Stat(dir); err != nil {
		if ioErr := os.Mkdir(dir, 0700); ioErr != nil {
			return ioErr
		}
	}

	var path string
	for {
		timestamp := strconv.FormatInt(int64(time.Now().Unix())+rand.Int63n(10), 10)
		path = filepath.Join(dir, manual.Command+"_"+timestamp+".txt")
		if _, err := os.Stat(path); err != nil {
			break
		}
	}

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	// write content as toml format
	var b bytes.Buffer
	e := toml.NewEncoder(&b)
	if e.Encode(manual); err != nil {
		return err
	}

	if _, err := file.Write(([]byte)(b.String())); err != nil {
		return err
	}

	// output for debug
	fmt.Println("Your manual is saved to file: " + path)

	return nil
}
