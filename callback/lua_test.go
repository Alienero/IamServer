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
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var L = NewLua()

func TestLuaInit(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	t.Logf("GOPATH:%v", gopath)
	L.SetLuaPath(filepath.Join(gopath, "/src/github.com/Alienero/IamServer/test/lua/"))
	L.InitCallBackModule()

	// load callback method
	L.SetAddrMappingFn()
	L.SetRtmpAccessCheck()
	L.SetFlvAccessCheck()
	L.SetIMAccessCheck()
}

func TestMapping(t *testing.T) {
	if private := L.AddrMapping(""); private != "master" {
		t.Errorf("ne:%v", "master")
	} else {
		t.Log(private)
	}
	fmt.Println("===")
	if private := L.AddrMapping("master"); private != "master" {
		t.Errorf("ne:%v", "master")
	} else {
		t.Log(private)
	}
	fmt.Println("===")
	if private := L.AddrMapping("blank"); private != "" {
		t.Errorf("ne:%v", "blank")
	} else {
		t.Log(private)
	}
}

func TestRtmpCheck(t *testing.T) {
	if ok := L.RtmpAccessCheck("ilulu.xyz:9009", "localhost", "live", "master"); !ok {
		t.Error("ne")
	}
	fmt.Println("===")
	if ok := L.RtmpAccessCheck("ilulu.xyz:9009", "localhost", "", ""); ok {
		t.Error("ne")
	}
}

func TestFlvCheck(t *testing.T) {
	if ok := L.FlvAccessCheck("ilulu.xyz:9009", "http://ilulu.xyz:9009", "/live/master", nil, nil); !ok {
		t.Error("ne")
	}
	fmt.Println("===")
	if ok := L.FlvAccessCheck("ilulu.xyz:9009", "http://ilulu.xyz:9009", "/live", nil, nil); ok {
		t.Error("ne")
	}
}

func TestIMCheck(t *testing.T) {
	if ok := L.IMAccessCheck("ilulu.xyz:9009", "http://ilulu.xyz:9009", "/im/master", nil, nil); !ok {
		t.Error("ne")
	}
	fmt.Println("===")
	if ok := L.IMAccessCheck("ilulu.xyz:9009", "http://ilulu.xyz:9009", "/im/live", nil, nil); ok {
		t.Error("ne")
	}
}
