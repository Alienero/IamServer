// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package lua

import (
	"io/ioutil"
	"os"
	"reflect"

	"github.com/yuin/gopher-lua"
)

type GoLua struct {
	l *lua.LState
}

func NewGolua() *GoLua {
	return &GoLua{
		l: lua.NewState(),
	}
}

func (gl *GoLua) Call(fn *Fn, args ...interface{}) (ret []interface{}, err error) {
	vs := make([]lua.LValue, len(args))
	for n, arg := range args {
		vs[n] = goToLua(arg)
	}
	if err = gl.l.CallByParam(fn.p, vs...); err != nil {
		return nil, err
	}
	ret = make([]interface{}, fn.p.NRet)
	for i := 0; i < fn.p.NRet; i++ {
		ret[i] = gl.l.Get(i + 1)
	}
	return
}

type Fn struct {
	p lua.P
}

func (gl *GoLua) GetCallParam(fn string, nret int) *Fn {
	return &Fn{
		lua.P{
			Fn:      gl.l.GetGlobal(fn),
			NRet:    nret,
			Protect: true,
		},
	}
}

func (gl *GoLua) Load(str string) error {
	return gl.l.DoString(str)
}

func (gl *GoLua) LoadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return gl.Load(string(data))
}

func (gl *GoLua) Close() error {
	gl.l.Close()
	return nil
}

func goToLua(i interface{}) (v lua.LValue) {
	kind := reflect.TypeOf(i).Kind()
	switch {
	case kind == reflect.Bool:
		v = lua.LBool(i.(bool))
	case reflect.Int <= kind && kind <= reflect.Float64:
		v = lua.LNumber(i.(float64))
	case kind == reflect.String:
		v = lua.LString(i.(string))

	}
	return
}

func GetString(lv interface{}) string {
	return string(lv.(lua.LString))
}

func GetNumber(lv interface{}) float64 {
	return float64(lv.(lua.LNumber))
}

func GetBool(lv interface{}) bool {
	return bool(lv.(lua.LBool))
}
