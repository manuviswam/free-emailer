package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg,err := LoadConfig("test.gcfg")
	if err != nil {
		t.Errorf("Failed with error : %s",err)
	}
	if cfg.Http.Port != ":8080" {
		t.Errorf("Expected : %s \nActual : %s",":8080",cfg.Http.Port)
	}
}
