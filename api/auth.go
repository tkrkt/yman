package api

import (
	"strings"
)

// Login to yman and create ~/.ymanrc
func Login(email string, password string) error {
	// create an account temporarily
	username := strings.Split(email, "@")[0]

	config := GetConfig()
	config.Username = username
	config.AccessToken = "token"

	// write to .ymanrc
	return SetConfig(config)
}

// Logout from yman by deleting ~/.ymanrc
func Logout() error {
	config := GetConfig()
	config.Username = ""
	config.AccessToken = ""
	return SetConfig(config)
}

// IsLogined returns account if logined
func IsLogined() bool {
	config := GetConfig()
	return config.Username != "" && config.AccessToken != ""
}
