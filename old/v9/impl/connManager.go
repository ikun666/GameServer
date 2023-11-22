package impl

import (
	"sync"

	"github.com/ikun666/old/v9/iface"
)

type ConnManager struct {
	ConnMap  map[uint32]iface.IConnection
	ConnLock sync.RWMutex
}

func (c *ConnManager) Add(conn iface.IConnection) {
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()
	c.ConnMap[conn.GetID()] = conn
}
func (c *ConnManager) Remove(conn iface.IConnection) {
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()
	delete(c.ConnMap, conn.GetID())
}
func (c *ConnManager) Get(id uint32) iface.IConnection {
	c.ConnLock.RLock()
	defer c.ConnLock.RUnlock()
	return c.ConnMap[id]
}
func (c *ConnManager) Len() int {
	return len(c.ConnMap)
}
func (c *ConnManager) ClearConn() {
	for _, v := range c.ConnMap {
		c.Remove(v)
		v.Stop()
	}
}
func NewConnManager() iface.IConnManager {
	return &ConnManager{
		ConnMap: make(map[uint32]iface.IConnection),
	}
}
