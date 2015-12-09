// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package callback

import (
	"net"

	"github.com/Alienero/IamServer/lua"
)

const (
	AddrMappingFn   = "addr_mapping"
	RTMPAccessCheck = "rtmp_access_check"
	FlvAccessCheck  = "flv_access_check"
	IMAccessCheck   = "im_access_check"
)

// Lua.
type Lua struct {
	gl        *lua.GoLua
	mappingFn *lua.Fn
}

func NewLua() *Lua {
	return &Lua{
		gl: lua.NewGolua(),
	}
}

func (l *Lua) Load(source string) error {
	return l.gl.Load(source)
}

func (l *Lua) SetAddrMappingFn() {
	l.mappingFn = l.gl.GetCallParam(AddrMappingFn, 1)
}

func (l *Lua) LoadFile(path string) error {
	return l.gl.LoadFile(path)
}

func (l *Lua) AddrMapping(public string) (private string, err error) {
	rets, err := l.gl.Call(l.mappingFn, public)
	if err != nil {
		return "", err
	}
	return lua.GetString(rets[0]), err
}

func (l *Lua) RtmpAccessCheck(conn net.Conn, appname, path string) (bool, error) {
}

func (l *Lua) Close() error {
	return l.gl.Close()
}
