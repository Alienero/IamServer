// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package http

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Alienero/IamServer/im"
	"github.com/Alienero/IamServer/monitor"
	"github.com/Alienero/IamServer/source"

	"github.com/golang/glog"
)

func InitHTTP(sources *source.SourceManage) error {
	tmpl, err := template.ParseFiles("../play.tpl")
	if err != nil {
		glog.Fatal("parse template error:", err)
		return err
	}
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("http: get an request.", r.RequestURI, r.Method)
		if r.Method != "GET" {
			return
		}
		// get live source.
		key := r.FormValue("key")
		key = "/live/123" // for test.
		consumer, err := source.NewConsumer(sources, key)
		if err != nil {
			glog.Info("can not get source", err)
			return
		}
		defer consumer.Close()
		// set flv live stream http head.
		// TODO: let browser not cache sources.
		w.Header().Add("Content-Type", "video/x-flv")
		if err := consumer.Live(w); err != nil {
			glog.Info("Live get an client error:", err)
		}
	})
	index := func(w http.ResponseWriter, r *http.Request) {
		user := monitor.Monitor.GetTempInfo()
		rid := r.FormValue("room_id")
		if rid == "" {
			rid = "master"
		}
		rm := im.GlobalIM.Rm.Get(rid)
		if rm == nil {
			user.LiveCount = 0
		} else {
			user.LiveCount = rm.GetLiveCount()
		}
		if err := tmpl.Execute(w, user); err != nil {
			glog.Error(err)
		}
	}
	http.HandleFunc("/index.html", index)
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		user := monitor.Monitor.GetTempInfo()
		rid := r.FormValue("room_id")
		if rid == "" {
			rid = "master"
		}
		rm := im.GlobalIM.Rm.Get(rid)
		if rm == nil {
			user.LiveCount = 0
		} else {
			user.LiveCount = rm.GetLiveCount()
		}
		data, err := json.Marshal(user)
		if err != nil {
			glog.Errorf("marshal json error: %v", err)
			return
		}
		w.Write(data)
	})
	var fileServer = http.FileServer(http.Dir("../"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			index(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	return nil
}
