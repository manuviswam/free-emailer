package config

import "code.google.com/p/gcfg"

type BasicConfig struct {
	Http struct {
		Port string
		StaticDirPath string
	}
	MailServer struct {
		UserName string
		Password string
		ServerUrl string
		MailPort string
	}
}

func LoadConfig(filename string)(config BasicConfig, err error) {
	err = gcfg.ReadFileInto(&config, filename)
	return config, err
}

