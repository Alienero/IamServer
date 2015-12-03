package im

import (
	"container/list"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

var GlobalIM = NewIMServer()

const (
	IMPath   = "/im"
	MaxCache = 4096 * 2
)

// server type
const (
	Signal = 0
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
	user, ok := GlobalIM.Rm.Check(consumer)
	if !ok {
		consumer.Close()
		return
	}
	consumer.name = user.Name
	consumer.r = r
	consumer.access = user.Access
	r.Add(consumer)
	defer func() {
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
	rooms    map[string]*Room
	idManage uint64

	typ byte // 0 is single version.
	// callback methods.
	Check func(c *Consumer) (*User, bool) // ok,access,name
}

// TODO: Imp it.
func SessionCheck(c *Consumer) (*User, bool) {
	return nil, false
}

func SingalCheck(c *Consumer) (*User, bool) {
	user := c.conn.Request().FormValue("uname")
	if user == "" {
		user = "some bird"
	}
	return &User{
		Access: 1,
		Name:   user,
	}, true
}

func NewRoomsManage() *RoomsManage {
	return &RoomsManage{
		rooms: make(map[string]*Room),
		// default use signal check method.
		Check: SingalCheck,
		typ:   Signal,
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
}

func NewRoom(id string) *Room {
	return &Room{
		id:           id,
		consumersMap: make(map[uint64]*list.Element),
		consumers:    list.New(),
	}
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
	name     string // user name.
	id       uint64 // session id.
	room     string // room id.
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
	Playload string `json:"playload"`
}

func NewConsumer(ws *websocket.Conn) *Consumer {
	// how to get name.
	// two method:
	// 		1.user session.
	// 		2.uer ws url args.
	rid := ws.Request().FormValue("room_id")
	glog.Infof("im ws server got room_id:%v", rid)
	return &Consumer{
		id:        GlobalIM.Rm.GetID(),
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
		if c.name == "some bird" && m.User != "" && GlobalIM.Rm.typ == Signal {
			// pass.
		} else {
			m.User = c.name
		}
		m.Time = time.Now().Unix()
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
