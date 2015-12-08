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

	"github.com/Alienero/IamServer/config"
	"github.com/Alienero/IamServer/http"
	"github.com/Alienero/IamServer/source"

	"github.com/golang/glog"
)

func InitServer() error {
	if len(config.Config.Apps) == 0 {
		return errors.New("empty app list.")
	}
	for n, application := range config.Config.Apps {
		// first we should to make a source manage.
		sources := source.NewSourcerManage()
		// mapping.
		if application.RTMP != nil {
			glog.Infof("Load RTMP serve:%v", n)
			// TODO: start rtmp publisher server.s
		} else {
			// should throws a panic.
			panic("App not has rtmp.")
		}
		// IM & HTTP use one port, by default.
		if application.HTTP != nil {
			// TODO: start  HTTP server.
			if application.HTTP_FLV != nil {
				glog.Infof("Load HTTP-FLV serve:%v", n)
				// TODO
			}
			if application.IM != nil {
				glog.Infof("Load IM serve:%v", n)
				// TODO
			}
		}
	}
	return nil
}
