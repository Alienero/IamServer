package monitor

import (
	"sync"
	"sync/atomic"
	"time"
)

var Monitor = monitor{}

func init() {
	Monitor.roomName = "测试直播间"
	Monitor.hostName = "请叫我丑的遁地"
}

type User struct {
	LiveCount int64 `json:"liveCount"` // use atomic
	RoomName  string
	HostName  string
}

type monitor struct {
	hostName       string
	roomName       string
	liveCount      int64
	trafficDown    int64
	avgTrafficDown int64
	trafficUp      int64
	avgTrafficUp   int64
	reportTime     time.Duration
	sync.RWMutex
}

func (m *monitor) SetName(host, room string) {
	m.Lock()
	m.hostName = host
	m.roomName = room
	m.Unlock()
}

func (m *monitor) GetTempInfo() *User {
	return &User{
		LiveCount: atomic.AddInt64(&m.liveCount, 0),
		RoomName:  m.roomName,
		HostName:  m.hostName,
	}
}

func (m *monitor) UserLogin() {
	atomic.AddInt64(&m.liveCount, 1)
}

func (m *monitor) UserLogout() {
	atomic.AddInt64(&m.liveCount, -1)
}

func (m *monitor) AddUpTraffic(t int64) {
	atomic.AddInt64(&m.trafficUp, t)
}

func (m *monitor) AddDownTraffic(t int64) {
	atomic.AddInt64(&m.trafficDown, t)
}
