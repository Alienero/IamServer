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

var table = reflect.TypeOf(new(Table))

// only support number float64
func goToLua(i interface{}) (v lua.LValue) {
	t := reflect.TypeOf(i)
	kind := t.Kind()
	switch {
	case kind == reflect.Bool:
		v = lua.LBool(i.(bool))
	case reflect.Int <= kind && kind <= reflect.Float64:
		v = lua.LNumber(i.(float64))
	case kind == reflect.String:
		v = lua.LString(i.(string))
	case kind == reflect.Ptr && t == table:
		// this case only for lua table.
		v = i.(*Table).m
	default:
		panic("unknow type.")
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

func GetTable(lv interface{}) *Table {
	return newTable(lv.(*lua.LTable))
}

// Lua table.
type Table struct {
	m *lua.LTable
}

func NewTalbe() *Table {
	return newTable(new(lua.LTable))
}

func newTable(m *lua.LTable) *Table {
	return &Table{
		m: m,
	}
}

func (t *Table) Get(key interface{}) interface{} {
	return t.m.RawGetH(goToLua(key))
}

func (t *Table) Set(key, value interface{}) {
	t.m.RawSetH(goToLua(key), goToLua(value))
}

func (t *Table) Del(key interface{}) {
	t.m.RawSetH(goToLua(key), nil)
}

func (t *Table) SetInt(index int, value interface{}) {
	t.m.RawSetInt(index, goToLua(value))
}

func (t *Table) GetInt(index int) lua.LValue {
	return t.m.RawGetInt(index)
}

func (t *Table) Append(vs ...interface{}) {
	index := t.m.Len()
	if index > 0 {
		index--
	}
	for _, v := range vs {
		switch v.(type) {
		case []float64:
			for _, f := range v.([]float64) {
				t.SetInt(index, f)
				index++
			}
		case []string:
			for _, s := range v.([]string) {
				t.SetInt(index, s)
				index++
			}
		default:
			t.SetInt(index, v)
			index++
		}
	}
}

func (t *Table) Len() int {
	return t.m.Len()
}
