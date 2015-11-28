package im

import (
	"container/list"
	"sync"

	"golang.org/x/net/websocket"
)

type IMServer struct {
	addr string
	rm   *RoomsManage
}

type RoomsManage struct {
	sync.RWMutex
	rooms map[string]*Room
}

func (rm *RoomsManage) Add(id string) {
	rm.Lock()
	room := &Room{
		id: id,
	}
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
	accessCheck func(id, psw string) bool
	sync.RWMutex
	consumersMap map[string]*list.Element
	consumers    *list.List
}

func (r *Room) Check(id, psw string) bool {
	return r.accessCheck(id, psw)
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

type Consumer struct {
	id  string
	typ byte
}
