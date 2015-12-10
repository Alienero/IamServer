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
	"testing"
)

func TestAddrMapping(t *testing.T) {
	lua := NewLua()
	defer lua.Close()
	if err := lua.Load(`
		-- addr-mapping
		mapping = {["123"]="456"}
		function addr_mapping(public)
			return mapping[public]
		end
		`); err != nil {
		t.Error(err)
	}
	lua.SetAddrMappingFn()

	if private, err := lua.AddrMapping("123"); private != "456" {
		if err != nil {
			t.Error(err)
		}
		t.Error("ne", private)
	} else {
		t.Log("ok")
	}

	if private, err := lua.AddrMapping("test"); private != "" {
		if err != nil {
			t.Error(err)
		}
		t.Error("ne", private)
	} else {
		t.Log("ok")
	}
}
