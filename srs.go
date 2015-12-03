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
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	myhttp "github.com/Alienero/IamServer/http"
	"github.com/Alienero/IamServer/im"
	"github.com/Alienero/IamServer/monitor"

	"github.com/golang/glog"
)

var isDebug = true

var hostName = flag.String("name", "请叫我丑的遁地啊", "")

func main() {
	if isDebug {
		f, err := os.Create("pprof")
		if err != nil {
			glog.Fatal(err)
		}
		if err = pprof.StartCPUProfile(f); err != nil {
			glog.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}
	flag.Parse()
	monitor.Monitor.SetName(*hostName, "Testing Room")
	defer glog.Flush()
	if err := flag.Set("logtostderr", "true"); err != nil {
		panic(err)
	}
	r := NewSrsServer()
	r.PrintInfo()
	// init http server
	if err := myhttp.InitHTTP(); err != nil {
		panic(err)
	}
	im.GlobalIM.Init()
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			panic(err)
		}
	}()
	go r.Serve()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	glog.Info("Signal received, initializing clean shutdown...")
}
