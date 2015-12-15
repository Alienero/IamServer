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
	"net/url"
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
	FlvAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) bool
}

// IM
type IM interface {
	IMAccessCheck(remote, url, path string, form url.Values, cookies []*http.Cookie) (uname string, access byte, ok bool)
}

type FlvCallback interface {
	FLV
}

type IMCallback interface {
	IM
}

type RTMPCallback interface {
	AppMapping
	RTMP
}
