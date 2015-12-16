package config

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
)

var (
	isWrirteTofile = true
)

func TestDefaultConifgMarshal(t *testing.T) {
	Config.Apps = []app{app{
		Name:    "live",
		LuaPath: "/home/lua",
		RTMP: rtmp{
			Enble:  true,
			Listen: []string{"localhost:1927"},
		},
		HTTP: http{
			Flv: httpFlv{
				Enble: true,
			},
			Im: im{
				Enble: true,
			},
			Listen: []string{"localhost:9090"},
		},
		DemoServer: demoServer{
			Enble: true,
		},
	}}
	data, err := yaml.Marshal(Config)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
	if isWrirteTofile {
		// write yaml config data to file.
		path := filepath.Join(os.Getenv("GOPATH"), "/src/github.com/Alienero/IamServer/config", "config.yaml")
		file, err := os.Create(path)
		if err != nil {
			t.Error(err)
		}
		_, err = file.Write(data)
		if err != nil {
			t.Error(err)
		}
	}
}
