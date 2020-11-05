package config

import "github.com/koding/multiconfig"

// Config our applications config
type Config struct {
	IndieAuthTokenURL string `default:""`
	IndieAuthMe       string `required:"true"`
}

type SyndicationTarget struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

// NewConfig provides our application config
func NewConfig() *Config {
	var config Config
	m := multiconfig.New()
	m.MustLoad(&config)

	return &config
}
