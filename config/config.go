// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package config

// Callback will use http callback,or golang callback, or lua callback.
// First supoort lua callback.

var Config configuration

func Init(filepath string) error {
	return nil
}

type configuration struct {
	Apps []app `yaml:"apps"`
}

type app struct {
	Name              string `yaml:"name"`
	RTMP              *rtmp  `yaml:"rtmp"`
	HTTP              *http  `yaml:"http"`
	PublicAddrMapping string `yaml:"addr-mapping"` // public mapping private. text,go,http,lua
	LuaPath           string `yaml:"lua-path"`
	EnbleHTTPDemo     bool   `yaml:"enble-http-demo"`
}

// Rtmp only allow publisher live streaming.
type rtmp struct {
	Enble   bool     `yaml:"enble"`
	Listen  []string `yaml:"listen"`
	AppName string   `yaml:"app-name"`
}

type http struct {
	Flv    *httpFlv `yaml:"flv"`
	Im     *im      `yaml:"im"`
	Listen []string `yaml:"listen"`
}

// HTTP-FLV only can support to play live streaming.
type httpFlv struct {
	Enble bool `yaml:"enble"`
}

// A live streaming online talk room.
// It only support websocket.
type im struct {
	Enble bool `yaml:"enble"`
}
