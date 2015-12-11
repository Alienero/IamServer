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
	"net/http"

	"golang.org/x/net/websocket"
)

// App
type AppMapping interface {
	AddrMapping(public string) (private string)
}

// RTMP
type RTMP interface {
	RtmpAccessCheck(remote, local, appname, path string) bool
}

// HTTP-FLV
type FLV interface {
	FlvAccessCheck(r *http.Request) bool
}

// IM
type IM interface {
	IMAccessCheck(ws *websocket.Conn) bool
}
