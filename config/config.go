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
	IpWall   *ipWall  `yaml:"ip-wall"`
	RTMP     *rtmp    `yaml:"rtmp"`
	HTTP_FLV *httpFlv `yaml:"http-flv"`
	IM       *im      `yaml:"im"`
}

type ipWall struct {
	Enble     bool  `yaml:"enble"`
	ResetTime int64 `yaml:"reset-time"`
}

// Rtmp only allow publisher live streaming.
type rtmp struct {
	Enble       bool     `yaml:"enble"`
	Listen      []string `yaml:"listen"`
	AppName     string   `yaml:"app-name"`
	AccessPath  []string `yaml:"access-path"`
	AccessCheck []string `yaml:"access-check"` // callback method.
}

// HTTP-FLV only can support to play live streaming.
type httpFlv struct {
	Enble       bool     `yaml:"enble"`
	Listen      []string `yaml:"listen"`
	AccessCheck []string `yaml:"access-check"` // callback method.
	AddrMap     []string `yaml:"addr-map"`     // callback method, mapping url addr and rtmp url.
}

// A live streaming online talk room.
// It only support websocket.
type im struct {
	Enble       bool     `yaml:"enble"`
	Listen      []string `yaml:"listen"`
	AccessCheck []string `yaml:"access-check"` // callback method.
	AddrMap     []string `yaml:"addr-map"`     // callback method, mapping url addr and rtmp url.
}
