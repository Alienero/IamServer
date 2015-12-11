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
	"testing"

	"github.com/yuin/gopher-lua"
)

func TestCall(t *testing.T) {
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(`
		i = 3
		function f()
			print(3)
		end

		function p(n)
			print(n)
			return "return ok!"
		end
		`); err != nil {
		t.Error(err)
	}
	fn := gl.GetCallParam("f", 0)
	if _, err := gl.Call(fn); err != nil {
		t.Error(err)
	}
	if rets, err := gl.Call(gl.GetCallParam("p", 1), "4399"); err != nil {
		t.Error(err)
	} else {
		if len(rets) != 1 {
			t.Error("ne 1")
		}
		t.Logf("len:%v, ret[0]=%v", len(rets), GetString(rets[0]))
	}
}

func TestLuaLocal(t *testing.T) {
	l1 := `
	local a = 3
	b = true
	`
	l2 := `
	print(a)
	print(b)
	`
	gl := NewGolua()
	defer gl.Close()
	gl.Load(l1)
	gl.Load(l2)
}

func TestLuaSame(t *testing.T) {
	l1 := `
	function a()
		print("first")
	end

	function a()
		print("second")
	end

	a()
	`
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(l1); err != nil {
		t.Error(err)
	}
}

func TestLuaTable(t *testing.T) {
	l := `
	function f(a) 
		temp = a["remote_addr"]
		print(temp)
		return a
	end
	`
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(l); err != nil {
		t.Error(err)
	}
	table := NewTalbe()
	table.Set("remote_addr", "ilulu.xyz")
	fn := gl.GetCallParam("f", 1)
	rets, err := gl.Call(fn, table)
	if err != nil {
		t.Error(err)
	}
	ret := GetString(GetTable(rets[0]).Get("remote_addr"))
	if ret != "ilulu.xyz" {
		t.Error("ne")
	} else {
		t.Log(ret)
	}
}

func TestLuaArray(t *testing.T) {
	arry := NewTalbe()
	i := []float64{1, 2, 3, 4, 5, 6, 7}
	for n, v := range i {
		arry.SetInt(n, v)
	}
	l := `
	function f(arry)
		for i=0,6 do
			print(i)
		end
	end
	`
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(l); err != nil {
		t.Error(err)
	}
	fn := gl.GetCallParam("f", 0)
	_, err := gl.Call(fn, arry)
	if err != nil {
		t.Error(err)
	}
}

func TestNestTable(t *testing.T) {
	arryOut := NewTalbe()
	arryIn := NewTalbe()
	i := []float64{1, 2, 3, 4, 5, 6, 7}
	for n, v := range i {
		arryIn.SetInt(n, v)
	}
	arryOut.SetInt(0, arryIn)
	arryOut.SetInt(1, arryIn)

	l := `
	function f(arry)
		for i=0,1 do
			for j = 0,6 do
				io.write(arry[i][j])
			end
			print()
		end
	end
	`
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(l); err != nil {
		t.Error(err)
	}
	fn := gl.GetCallParam("f", 0)
	_, err := gl.Call(fn, arryOut)
	if err != nil {
		t.Error(err)
	}
}

func TestAppendSlice(t *testing.T) {
	arry := NewTalbe()
	i := []float64{1, 2, 3, 4, 5, 6, 7}
	arry.Append(i)
	l := `
	function f(array)
		print("----------------Print Array--------------")
		for i=0,6 do
			io.write(array[i])
		end
		print()
		print("----------------Print Done--------------")
	end
	`
	gl := NewGolua()
	defer gl.Close()
	if err := gl.Load(l); err != nil {
		t.Error(err)
	}
	fn := gl.GetCallParam("f", 0)
	_, err := gl.Call(fn, arry)
	if err != nil {
		t.Error(err)
	}
}

func TestTwiceCall(t *testing.T) {
	l := `
	mapping = {["123"]="456",["2"]="hello"}
	function addr_mapping(public)
		return mapping[public]
	end
	`
	L := lua.NewState()
	L.DoString(l)
	L.CallByParam(lua.P{
		Fn:      L.GetGlobal("addr_mapping"),
		NRet:    1,
		Protect: true,
	}, lua.LString("123"))
	t.Log(L.Get(1).String(), L.Get(2).String())
	L.Pop(1) // must pop.
	L.CallByParam(lua.P{
		Fn:      L.GetGlobal("addr_mapping"),
		NRet:    1,
		Protect: true,
	}, lua.LString("2"))
	t.Log(L.Get(1).String(), L.GetTop())
}

func TestModuleFn(t *testing.T) {
	l := `
	test = {}
	function test.f()
		print("ok")
	end
	return test
	`
	L := lua.NewState()
	if err := L.DoString(l); err != nil {
		t.Log(err)
	}
	t.Log(L.GetTop())
	module := L.Get(-1).(*lua.LTable)

	if err := L.CallByParam(lua.P{
		Fn:      module.RawGet(lua.LString("f")),
		NRet:    0,
		Protect: true,
	}); err != nil {
		t.Error(err)
	}
}
