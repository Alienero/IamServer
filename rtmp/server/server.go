// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package server

import (
	"fmt"
	"net"
	"runtime"

	"github.com/Alienero/IamServer/callback"
	"github.com/Alienero/IamServer/im"
	"github.com/Alienero/IamServer/rtmp"
	"github.com/Alienero/IamServer/source"

	"github.com/golang/glog"
)

type SrsServer struct {
	id       uint64
	addr     string
	sources  *source.SourceManage
	cb       callback.RTMP
	enbleIM  bool
	imServer *im.IMServer
}

func NewSrsServer(addr string, cb callback.RTMP, sources *source.SourceManage,
	enbleIM bool, imServer *im.IMServer) *SrsServer {
	return &SrsServer{
		addr:     addr,
		sources:  sources,
		cb:       cb,
		id:       SrsGenerateId(),
		enbleIM:  enbleIM,
		imServer: imServer,
	}
}

func (r *SrsServer) PrintInfo() {
	fmt.Println(`Powered by
	      ___ _           _ _     _
	     / __\ | _____  _(_) |__ | | ___
	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
	   / /   | |  __/>  <| | |_) | |  __/
	   \/    |_|\___/_/\_\_|_.__/|_|\___|
	   					Team

https://github.com/FlexibleBroadband
`)
	glog.Infof("RTMP Protocol Stack:  %v", rtmp.Version)
	glog.Infof("Will use CPU Num:%v", runtime.NumCPU())
}

func (r *SrsServer) Serve() error {
	// too many open files will thows a panic.
	addr, err := net.ResolveTCPAddr("tcp", r.addr)
	if err != nil {
		glog.Errorf("resolve listen address failed, err=%v", err)
		return fmt.Errorf("resolve listen address failed, err=%v", err)
	}

	var listener *net.TCPListener
	listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		glog.Errorf("listen failed, err=%v", err)
		return fmt.Errorf("listen failed, err=%v", err)
	}
	defer listener.Close()
	for {
		glog.Info("listener ready to accept client")
		conn, err := listener.AcceptTCP()
		if err != nil {
			glog.Errorf("accept client failed, err=%v", err)
			return fmt.Errorf("accept client failed, err=%v", err)
		}
		glog.Info("TCP Connected")

		go r.serve(conn)
	}
}

func (r *SrsServer) serve(conn *net.TCPConn) {
	var (
		client *SrsClient
		err    error
	)
	if client, err = NewSrsClient(conn, r); err != nil {
		glog.Errorf("create client failed, err=%v", err)
		return
	}

	if err = client.do_cycle(); err != nil {
		glog.Errorf("do cycle err=%v", err)
	}
}
