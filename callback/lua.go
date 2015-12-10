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
	"net/http"
	"net/url"

	"github.com/Alienero/IamServer/lua"
)

const (
	AddrMappingFn   = "addr_mapping"
	RTMPAccessCheck = "rtmp_access_check"
	FlvAccessCheck  = "flv_access_check"
	IMAccessCheck   = "im_access_check"
)

// Lua's callback function.
type Lua struct {
	gl              *lua.GoLua
	mappingFn       *lua.Fn
	rtmpAccessCheck *lua.Fn
	flvAccessCheck  *lua.Fn
	imAccessCheck   *lua.Fn
}

func NewLua() *Lua {
	return &Lua{
		gl: lua.NewGolua(),
	}
}

func (l *Lua) Load(source string) error {
	return l.gl.Load(source)
}

func (l *Lua) LoadFile(path string) error {
	return l.gl.LoadFile(path)
}

func (l *Lua) SetAddrMappingFn() {
	l.mappingFn = l.gl.GetCallParam(AddrMappingFn, 1)
}

func (l *Lua) SetRtmpAccessCheck() {
	l.rtmpAccessCheck = l.gl.GetCallParam(RTMPAccessCheck, 1)
}

func (l *Lua) SetFlvAccessCheck() {
	l.flvAccessCheck = l.gl.GetCallParam(FlvAccessCheck, 1)
}

func (l *Lua) SetIMAccessCheck() {
	l.imAccessCheck = l.gl.GetCallParam(IMAccessCheck, 1)
}

func (l *Lua) AddrMapping(public string) (private string, err error) {
	rets, err := l.gl.Call(l.mappingFn, public)
	if err != nil {
		return "", err
	}
	return lua.GetString(rets[0]), nil
}

func (l *Lua) RtmpAccessCheck(remote, local, appname, path string) (bool, error) {
	rets, err := l.gl.Call(l.rtmpAccessCheck, remote, local, appname, path)
	if err != nil {
		return false, err
	}
	return lua.GetBool(rets[0]), nil
}

// remote: remote address, url: HTTP request URL
func (l *Lua) FlvAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) (bool, error) {
	fms := lua.NewTalbe()
	for k, rs := range form {
		slice := lua.NewTalbe()
		slice.Append(rs)
		fms.Set(k, slice)
	}
	cs := lua.NewTalbe()
	for n, cookie := range cookies {
		c := lua.NewTalbe()
		c.Set("name", cookie.Name)
		c.Set("value", cookie.Value)
		cs.SetInt(n, c)
	}
	rets, err := l.gl.Call(l.flvAccessCheck, remote, url, path, fms, cs)
	if err != nil {
		return false, err
	}
	return lua.GetBool(rets[0]), nil
}

// remote: remote address, url: HTTP request URL
func (l *Lua) IMAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) (bool, error) {
	fms := lua.NewTalbe()
	for k, rs := range form {
		slice := lua.NewTalbe()
		slice.Append(rs)
		fms.Set(k, slice)
	}
	cs := lua.NewTalbe()
	for n, cookie := range cookies {
		c := lua.NewTalbe()
		c.Set("name", cookie.Name)
		c.Set("value", cookie.Value)
		cs.SetInt(n, c)
	}
	rets, err := l.gl.Call(l.imAccessCheck, remote, url, path, fms, cs)
	if err != nil {
		return false, err
	}
	return lua.GetBool(rets[0]), nil
}

func (l *Lua) Close() error {
	return l.gl.Close()
}
