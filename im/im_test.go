// Copyright © 2015 FlexibleBroadband Team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//	      ___ _           _ _     _
//	     / __\ | _____  _(_) |__ | | ___
//	    / _\ | |/ _ \ \/ / | '_ \| |/ _ \
//	   / /   | |  __/>  <| | |_) | |  __/
//	   \/    |_|\___/_/\_\_|_.__/|_|\___|

package im

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

func TestAddClientRecive(t *testing.T) {
	var server = GlobalIM
	var roomID = "master"
	server.Init()
	t.Log("add a room")
	server.Rm.Add(roomID)

	go func() {
		if err := http.ListenAndServe("localhost:9999", nil); err != nil {
			panic(err)
		}
	}()

	// let listen start.
	time.Sleep(2 * time.Second)

	ws1, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&uname=1", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}

	ws2, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&uname=2", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}
	ws3, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&uname=3", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}
	m := &msg{
		Playload: "client1",
	}

	t.Log("will send.")
	if err := websocket.JSON.Send(ws1, m); err != nil {
		t.Fatal(err)
	} else {
		t.Log("has been send.")
	}

	m = new(msg)
	if err := websocket.JSON.Receive(ws2, m); err != nil {
		t.Fatal(err)
	}
	t.Log("Get message:", string(m.Playload))
	if err := websocket.JSON.Receive(ws3, m); err != nil {
		t.Fatal(err)
	}
	t.Log("Get message:", string(m.Playload))
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Json", string(data))

}
