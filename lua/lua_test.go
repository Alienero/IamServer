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

func TestLua(t *testing.T) {
	file := `
	function f(num1,num2)
		print(num1,num2)
	end
	`
	L := lua.NewState()
	L.DoString(file)
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("f"),
		NRet:    0,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(2314)); err != nil {
		panic(err)
	}
}

func TestCall(t *testing.T) {
	gl := NewGolua()
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