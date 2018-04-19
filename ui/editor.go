package ui

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

// Editor opens text editor to input multiline text
func Editor() (string, error) {
	// create tmp file
	tmpDir := os.TempDir()
	tmpFile, err := ioutil.TempFile(tmpDir, "yman")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	// get editor
	editorType := viper.GetString("editor")
	if editorType == "" {
		editorType = "vi"
	}
	editor, err := exec.LookPath(editorType)
	if err != nil {
		return "", err
	}

	// open editor
	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	// read from tmp file
	msg, ioErr := ioutil.ReadFile(tmpFile.Name())
	if ioErr != nil {
		return "", nil
	}

	return string(msg), nil
}
