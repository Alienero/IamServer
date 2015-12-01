package im

import (
	"container/list"
	"net/http"
	"sync"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

var GlobalIM = NewIMServer()

const (
	IMPath   = "/im"
	MaxCache = 4096 * 2
)

type IMServer struct {
	addr string
	Rm   *RoomsManage
}

func NewIMServer() *IMServer {
	return &IMServer{
		Rm: new(RoomsManage),
	}
}

func (server *IMServer) Init() {
	server.Rm = NewRoomsManage()
	http.Handle(IMPath, websocket.Handler(server.handle))
}

func (server *IMServer) handle(ws *websocket.Conn) {
	consumer := NewConsumer(ws)
	// check room_id.
	r := server.Rm.Get(consumer.room)
	if r == nil {
		glog.Info("room is not exist.")
		consumer.Close()
		return
	}
	// check user.
	// TODO: Imp me. Allow all user.
	ok, access := r.Check(consumer)
	if !ok {
		consumer.Close()
		return
	}
	consumer.r = r
	consumer.access = access
	r.Add(consumer)
	defer func() {
		r.Del(consumer)
		consumer.Close()
	}()
	// start write & read.
	consumer.Serve()
}

type RoomsManage struct {
	sync.RWMutex
	rooms map[string]*Room
}

func NewRoomsManage() *RoomsManage {
	return &RoomsManage{
		rooms: make(map[string]*Room),
	}
}

func (rm *RoomsManage) Add(id string) {
	rm.Lock()
	room := NewRoom(id)
	rm.rooms[id] = room
	rm.Unlock()
}

func (rm *RoomsManage) Get(id string) *Room {
	rm.RLock()
	defer rm.RUnlock()
	return rm.rooms[id]
}

func (rm *RoomsManage) Del(id string) {
	rm.Lock()
	delete(rm.rooms, id)
	rm.Unlock()
}

type Room struct {
	id          string
	accessCheck func(id, key, room string) (bool, byte)
	sync.RWMutex
	consumersMap map[string]*list.Element
	consumers    *list.List
	isClosed     bool
}

func NewRoom(id string) *Room {
	return &Room{
		accessCheck:  func(id, key, room string) (bool, byte) { return true, 1 },
		consumersMap: make(map[string]*list.Element),
		consumers:    list.New(),
	}
}

func (r *Room) Check(c *Consumer) (bool, byte) {
	return r.accessCheck(c.id, c.key, c.room)
}

func (r *Room) Add(c *Consumer) {
	r.Lock()
	r.consumersMap[c.id] = r.consumers.PushBack(c)
	r.Unlock()
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

func (r *Room) Broadcast(m *msg, eid string) {
	r.RLock()
	defer r.RUnlock()
	if r.isClosed {
		return
	}
	for node := r.consumers.Front(); node != nil; node = node.Next() {
		c := node.Value.(*Consumer)
		if c.id != eid {
			c.Write(m)
		}
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
	id       string
	room     string
	key      string
	typ      byte
	access   byte
	r        *Room
	isClosed bool

	sync.RWMutex

	writeChan chan *msg

	conn *websocket.Conn
}

type msg struct {
	User     string `json:"user"`
	Time     int64  `json:"time"`
	Type     byte   `json:"type"`
	Color    string `json:"color"`
	Playload []byte `json:"playload"`
}

func NewConsumer(ws *websocket.Conn) *Consumer {
	rid := ws.Request().FormValue("room_id")
	user := ws.Request().FormValue("user_id")
	key := ws.Request().FormValue("key")
	// TODO: check user from session.
	glog.Infof("im ws server got room_id:%v, user_id:%v", rid, user)
	return &Consumer{
		id:        user,
		key:       key,
		room:      rid,
		conn:      ws,
		writeChan: make(chan *msg, MaxCache),
	}
}

func (c *Consumer) Serve() {
	go c.readLoop()
	c.writeLoop()
}

func (c *Consumer) readLoop() (err error) {
	for {
		m := new(msg)
		if err = websocket.JSON.Receive(c.conn, m); err != nil {
			return
		}
		glog.Info("ws server:read a msg.")
		c.r.Broadcast(m, c.id)
	}
}

func (c *Consumer) writeLoop() {
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
	glog.Info("ws server write a msg.")
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
