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
