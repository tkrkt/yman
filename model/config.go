package model

// Config represents user's configurations such as username and access token,
// written in ~/.ymanrc
type Config struct {
	Username    string `toml:"username"`
	AccessToken string `toml:"access_token"`
	Repository  string `toml:"repository"`
	Editor      string `toml:"editor"`
	Lang        string `toml:"lang"`
}

// NewConfig generates new config with initial value
func NewConfig() *Config {
	c := Config{
		Repository: "https://",
		Editor:     "vi",
		Lang:       "en",
	}
	return &c
}
