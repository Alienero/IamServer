package im

import (
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

var server = NewIMServer()

var roomID = "123"

func TestAddClientRecive(t *testing.T) {

	server.Init()
	t.Log("add a room")
	server.rm.Add(roomID)

	go func() {
		if err := http.ListenAndServe("localhost:9999", nil); err != nil {
			panic(err)
		}
	}()

	// let listen start.
	time.Sleep(2 * time.Second)

	ws1, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&user_id=1", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}

	ws2, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&user_id=2", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}
	ws3, err := websocket.Dial("ws://localhost:9999/im?room_id="+roomID+"&user_id=3", "", "http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}
	m := &msg{
		Playload: []byte("client1"),
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

}
