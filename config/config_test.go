package config

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfigNullMarshal(t *testing.T) {
	data, err := yaml.Marshal(Config)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

func TestDefaultConifgMarshal(t *testing.T) {
	Config.Apps = []app{app{}}
	data, err := yaml.Marshal(Config)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
