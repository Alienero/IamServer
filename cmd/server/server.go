// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alienero/IamServer/app"
	"github.com/Alienero/IamServer/config"

	"github.com/golang/glog"
)

var (
	configfile string
)

func main() {
	flag.StringVar(&configfile, "conf", "config.yaml", "")

	if err := flag.Set("logtostderr", "true"); err != nil {
		panic(err)
	}
	flag.Parse()
	defer glog.Flush()

	if err := config.Init(configfile); err != nil {
		panic(err)
	}
	// print info.
	app.PrintInfo()
	// start app.
	if err := app.InitServer(); err != nil {
		panic(err)
	}
	// wait signal.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	glog.Info("Signal received, initializing clean shutdown...")
}
