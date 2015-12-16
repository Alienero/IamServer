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
	"path"
	"strings"

	"github.com/Alienero/IamServer/callback"
	"github.com/Alienero/IamServer/im"
	"github.com/Alienero/IamServer/source"

	"github.com/golang/glog"
)

func InitHTTP(mux *http.ServeMux, sources *source.SourceManage, imServer *im.IMServer) error {
	tmpl, err := template.ParseFiles("../player/play.tpl")
	if err != nil {
		glog.Fatal("parse template error:", err)
		return err
	}
	index := func(w http.ResponseWriter, r *http.Request) {
		rid := r.FormValue("room_id")
		if rid == "" {
			rid = "master"
		}
		rm := imServer.Rm.Get(rid)
		var liveCount int64
		if rm != nil {
			liveCount = rm.GetLiveCount()
		} else {
			liveCount = 0
		}
		type User struct {
			LiveCount int64 `json:"liveCount"` // use atomic
			RoomName  string
			HostName  string
		}
		if err := tmpl.Execute(w, User{
			LiveCount: liveCount,
		}); err != nil {
			glog.Error(err)
		}
	}
	mux.HandleFunc("/index.html", index)
	mux.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		type User struct {
			LiveCount int64 `json:"liveCount"` // use atomic
			RoomName  string
			HostName  string
		}
		user := new(User)
		rid := r.FormValue("room_id")
		if rid == "" {
			rid = "master"
		}
		rm := imServer.Rm.Get(rid)
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
	var fileServer = http.FileServer(http.Dir("../player/"))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			index(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	return nil
}

func InitHTTPFlv(mux *http.ServeMux, app string, sources *source.SourceManage, cb callback.FlvCallback) {
	prefix := path.Join("/", app) + "/"
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("http: get an request: %v", r.RequestURI)
		if r.Method != "GET" {
			return
		}
		r.ParseForm()
		// access check.
		if !cb.FlvAccessCheck(r.RemoteAddr, r.RequestURI, r.URL.Path, r.Form, r.Cookies()) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		// get path.
		_, file := path.Split(r.URL.Path)
		var key string
		key = file[:strings.Index(file, ".flv")]

		// get live source.
		consumer, err := source.NewConsumer(sources, key)
		if err != nil {
			glog.Info("can not get source", err)
			return
		}
		glog.Info("get source")
		defer consumer.Close()
		// set flv live stream http head.
		// TODO: let browser not cache sources.
		w.Header().Add("Content-Type", "video/x-flv")
		if err := consumer.Live(w); err != nil {
			glog.Info("Live get an client error:", err)
		}
	})
}
