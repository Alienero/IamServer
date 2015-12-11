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
	CallBackModule  = "callback"
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
	callBackModule  *lua.Table
}

func NewLua() *Lua {
	return &Lua{
		gl: lua.NewGolua(),
	}
}

func (l *Lua) SetLuaPath(path string) {
	l.gl.SetLuaPath(path)
}

func (l *Lua) InitCallBackModule() {
	l.callBackModule = l.gl.GetModule(CallBackModule)
}

func (l *Lua) Load(source string) error {
	return l.gl.Load(source)
}

func (l *Lua) LoadFile(path string) error {
	return l.gl.LoadFile(path)
}

func (l *Lua) SetAddrMappingFn() {
	l.mappingFn = l.gl.GetCallParamWithFn(l.callBackModule.Get(AddrMappingFn), 1)
}

func (l *Lua) SetRtmpAccessCheck() {
	l.rtmpAccessCheck = l.gl.GetCallParamWithFn(l.callBackModule.Get(RTMPAccessCheck), 1)
}

func (l *Lua) SetFlvAccessCheck() {
	l.flvAccessCheck = l.gl.GetCallParamWithFn(l.callBackModule.Get(FlvAccessCheck), 1)
}

func (l *Lua) SetIMAccessCheck() {
	l.imAccessCheck = l.gl.GetCallParamWithFn(l.callBackModule.Get(IMAccessCheck), 1)
}

func (l *Lua) AddrMapping(public string) (private string) {
	rets, err := l.gl.Call(l.mappingFn, public)
	if err != nil {
		panic(err)
	}
	return lua.GetString(rets[0])
}

func (l *Lua) RtmpAccessCheck(remote, local, appname, path string) bool {
	rets, err := l.gl.Call(l.rtmpAccessCheck, remote, local, appname, path)
	if err != nil {
		panic(err)
	}
	return lua.GetBool(rets[0])
}

// remote: remote address, url: HTTP request URL
func (l *Lua) FlvAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) bool {
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
		panic(err)
	}
	return lua.GetBool(rets[0])
}

// remote: remote address, url: HTTP request URL
func (l *Lua) IMAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) bool {
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
		return false
	}
	return lua.GetBool(rets[0])
}

func (l *Lua) Close() error {
	return l.gl.Close()
}
