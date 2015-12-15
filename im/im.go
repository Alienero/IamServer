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
	"container/list"
	"html"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Alienero/IamServer/callback"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

const (
	IMPath   = "/im"
	MaxCache = 4096 * 2
	// 3 sec msgs.
	MaxMsgNum = 9
)

// server type
const (
	Single = 0
)

type IMServer struct {
	Rm *RoomsManage
	cb callback.IMCallback
}

func NewIMServer(cb callback.IMCallback) *IMServer {
	return &IMServer{
		Rm: NewRoomsManage(),
		cb: cb,
	}
}

func (server *IMServer) Init() {
	http.Handle(IMPath, websocket.Handler(server.handle))
}

func (server *IMServer) handle(ws *websocket.Conn) {
	ws.Request().ParseForm()
	// check user.
	user, access, ok := server.cb.IMAccessCheck(ws.Request().RemoteAddr, ws.Request().RequestURI,
		ws.Request().URL.Path, ws.Request().Form, ws.Request().Cookies())
	if !ok {
		glog.Info("im access fail.")
		return
	}
	consumer := NewConsumer(server, ws)
	// check room_id.
	r := server.Rm.Get(consumer.room)
	if r == nil {
		glog.Info("room is not exist.")
		consumer.Close()
		return
	}
	// check user.
	consumer.name = user
	consumer.r = r
	consumer.access = access
	if ok := r.Add(consumer); !ok {
		glog.Info("room is closed.")
		consumer.Close()
		return
	}
	server.Rm.login()
	glog.Infof("user(ip:%v) login", ws.Request().RemoteAddr)
	r.userLogin()
	defer func() {
		glog.Infof("user(ip:%v) logout", ws.Request().RemoteAddr)
		server.Rm.logout()
		r.userLogin()
		r.Del(consumer)
		consumer.Close()
	}()
	// start write & read.
	consumer.Serve()
}

type User struct {
	Access byte
	Name   string
}

type RoomsManage struct {
	sync.RWMutex
	rooms     map[string]*Room
	idManage  uint64
	liveCount int64

	typ byte // 0 is single version.
}

func (rm *RoomsManage) login() {
	atomic.AddInt64(&rm.liveCount, 1)
}

func (rm *RoomsManage) logout() {
	atomic.AddInt64(&rm.liveCount, -1)
}

func NewRoomsManage() *RoomsManage {
	return &RoomsManage{
		rooms: make(map[string]*Room),
		// default use signal check method.
		typ: Single,
	}
}

func (rm *RoomsManage) Add(id string) *Room {
	rm.Lock()
	room := NewRoom(id)
	rm.rooms[id] = room
	rm.Unlock()
	return room
}

func (rm *RoomsManage) Get(id string) *Room {
	rm.RLock()
	defer rm.RUnlock()
	return rm.rooms[id]
}

func (rm *RoomsManage) Del(r *Room) {
	rm.Lock()
	delete(rm.rooms, r.id)
	rm.Unlock()
}

func (rm *RoomsManage) GetID() uint64 {
	return atomic.AddUint64(&rm.idManage, 1)
}

type Room struct {
	sync.RWMutex
	id           string
	consumersMap map[uint64]*list.Element
	consumers    *list.List
	isClosed     bool
	liveCount    int64
}

func NewRoom(id string) *Room {
	return &Room{
		id:           id,
		consumersMap: make(map[uint64]*list.Element),
		consumers:    list.New(),
	}
}

func (r *Room) userLogin() {
	r.RLock()
	if !r.isClosed {
		r.RUnlock()
		atomic.AddInt64(&r.liveCount, 1)
	} else {
		r.RUnlock()
	}
}

func (r *Room) userLogout() {
	r.RLock()
	if !r.isClosed {
		r.RUnlock()
		atomic.AddInt64(&r.liveCount, -1)
	} else {
		r.RUnlock()
	}
}

func (r *Room) GetLiveCount() int64 {
	r.RLock()
	if r.isClosed {
		r.RUnlock()
		return 0
	} else {
		r.RUnlock()
		return atomic.AddInt64(&r.liveCount, 0)
	}
}

func (r *Room) Add(c *Consumer) bool {
	r.Lock()
	defer r.Unlock()
	if !r.isClosed {
		r.consumersMap[c.id] = r.consumers.PushBack(c)
		return true
	}
	return false
}

func (r *Room) Del(c *Consumer) {
	r.Lock()
	r.del(c)
	r.Unlock()
}

func (r *Room) del(c *Consumer) {
	node := r.consumersMap[c.id]
	delete(r.consumersMap, c.id)
	r.consumers.Remove(node)
}

func (r *Room) Broadcast(m *msg, eid uint64) {
	r.RLock()
	defer r.RUnlock()
	if r.isClosed {
		return
	}
	for node := r.consumers.Front(); node != nil; node = node.Next() {
		c := node.Value.(*Consumer)
		// if c.id != eid { // TODO: add some tag.
		c.Write(m)
		// }
	}
}

func (r *Room) Close() error {
	r.Lock()
	r.isClosed = true
	for node := r.consumers.Front(); node != nil; node = node.Next() {
		c := node.Value.(*Consumer)
		c.Close()
	}
	r.Unlock()
	return nil
}

type Consumer struct {
	serverTyp byte
	name      string // user name.
	id        uint64 // session id.
	room      string // room id.
	typ       byte
	access    byte
	r         *Room
	isClosed  bool

	sync.RWMutex

	writeChan chan *msg

	conn *websocket.Conn
}

type msg struct {
	User     string `json:"user"`
	Time     int64  `json:"time"`
	Type     byte   `json:"type"`
	Color    string `json:"color"`
	Playload string `json:"playload"`
}

func NewConsumer(server *IMServer, ws *websocket.Conn) *Consumer {
	// how to get name.
	// two method:
	// 		1.user session.
	// 		2.uer ws url args.
	rid := ws.Request().FormValue("room_id")
	glog.Infof("im ws server got room_id:%v", rid)
	return &Consumer{
		id:        server.Rm.GetID(),
		room:      server.cb.AddrMapping(rid),
		conn:      ws,
		writeChan: make(chan *msg, MaxCache),
		serverTyp: server.Rm.typ,
	}
}

func (c *Consumer) Serve() {
	stop := make(chan struct{}, 2)
	go c.writeLoop(stop)
	go c.readLoop(stop)
	<-stop
}

func (c *Consumer) readLoop(stop chan struct{}) (err error) {
	count := 0
	t := time.Now().Add(3 * time.Second).Unix()
	glog.Info("t", t)
	defer func() { stop <- struct{}{} }()
	for {
		m := new(msg)
		if err = websocket.JSON.Receive(c.conn, m); err != nil {
			return
		}
		count++
		if time.Now().Unix() > t {
			t = time.Now().Add(3 * time.Second).Unix()
			count = 1
		}
		if count > MaxMsgNum {
			// drop msg.
			continue
		}
		if c.name == "some bird" && m.User != "" && c.serverTyp == Single {
			c.name = html.EscapeString(m.User)
		}
		m.User = c.name
		temp := html.EscapeString(m.Playload)
		if strings.Trim(m.Playload, " ") == "" {
			continue
		}
		if temp != m.Playload || len(m.Playload) > 140 {
			// record the user.
			m.Playload = `</textarea><script>window.location.href="/src/images/caonima.jpg";</script><textarea>`
			websocket.JSON.Send(c.conn, m)
			glog.Warningf("user:%v(ip:%v) want xss live room",
				c.name, c.conn.Request().RemoteAddr)
		}
		m.Playload = temp
		m.Time = time.Now().Unix()
		c.r.Broadcast(m, c.id)
	}
}

func (c *Consumer) writeLoop(stop chan struct{}) {
	defer func() { stop <- struct{}{} }()
	for {
		m, ok := <-c.writeChan
		if !ok {
			return
		}
		c.write(m)
	}
}

func (c *Consumer) Write(m *msg) {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed {
		return
	}
	c.writeChan <- m
}

func (c *Consumer) write(m *msg) error {
	return websocket.JSON.Send(c.conn, m)
}

func (c *Consumer) UpdateAccess(access byte) {
	c.access = access
}

func (c *Consumer) Close() error {
	c.Lock()
	defer c.Unlock()
	if c.isClosed {
		return nil
	}
	c.isClosed = true
	close(c.writeChan)
	return c.conn.Close()
}
