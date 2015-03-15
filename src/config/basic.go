package config

import "code.google.com/p/gcfg"

type BasicConfig struct {
	Http struct {
		Port string
	}
}

func LoadConfig(filename string)(config BasicConfig, err error) {
	err = gcfg.ReadFileInto(&config, filename)
	return config, err
}

