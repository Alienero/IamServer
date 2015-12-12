// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alienero/IamServer/callback"
	"github.com/Alienero/IamServer/config"
	serverHttp "github.com/Alienero/IamServer/http"
	"github.com/Alienero/IamServer/lua"
	rtmpServer "github.com/Alienero/IamServer/rtmp/server"
	"github.com/Alienero/IamServer/source"

	"github.com/golang/glog"
)

func InitServer() error {
	printInfo()
	if len(config.Config.Apps) == 0 {
		return errors.New("empty app list.")
	}
	for n, application := range config.Config.Apps {
		// first we should to make a source manage.
		sources := source.NewSourcerManage()
		// second init a lua.
		cb := initLua(application.LuaPath)

		if application.RTMP != nil {
			glog.Infof("Load RTMP serve:%v", n)
			// start rtmp publisher server
			for _, addr := range application.RTMP.Listen {
				s := rtmpServer.NewSrsServer(addr, cb, sources)
				go s.Serve()
			}
		} else {
			// should throws a panic.
			panic("App not has rtmp.")
		}
		// IM & HTTP use one port, by default.
		if application.HTTP != nil {
			// TODO: start  HTTP server.
			mux := http.NewServeMux()

			if application.HTTP.Flv != nil {
				glog.Infof("Load HTTP-FLV serve:%v", n)
				serverHttp.InitHTTPFlv(mux, application.Name, sources, cb)
				// start flv http server.
				for _, addr := range application.HTTP.Flv.Listen {
					go http.ListenAndServe(addr, mux)
				}
			}
			if application.HTTP.Im != nil {
				glog.Infof("Load IM serve:%v", n)
				// TODO
			}
		}
	}
	return nil
}

func initLua(luapath string) *callback.Lua {
	cl := callback.NewLua(lua.NewGolua())
	cl.SetLuaPath(luapath)
	// load lua callback module
	cl.InitCallBackModule()
	// load lua callback functions.
	cl.SetAddrMappingFn()
	cl.SetRtmpAccessCheck()
	cl.SetFlvAccessCheck()
	cl.SetIMAccessCheck()

	return cl
}

func printInfo() {
	fmt.Println(`Powered by
	      ___ _           _ _     _
	     / __\ | _____  _(_) |__ | | ___
	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
	   / /   | |  __/>  <| | |_) | |  __/
	   \/    |_|\___/_/\_\_|_.__/|_|\___|
	   					Team

https://github.com/FlexibleBroadband
`)
}
