package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/v4/iface"
)

type Connection struct {
	Conn     *net.TCPConn
	ID       uint32
	IsClosed bool
	ExitChan chan struct{}
	Router   iface.IRouter
}

func NewConnetion(conn *net.TCPConn, id uint32, router iface.IRouter) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		IsClosed: false,
		ExitChan: make(chan struct{}),
		Router:   router,
	}
}
func (c *Connection) Read() {
	defer c.Stop()
	buf := make([]byte, 512)
	for {
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("read err:%v", err)
			return
		}
		req := &Request{
			conn: c,
			data: buf[:n],
		}
		go func(req iface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}
}
func (c *Connection) Start() {
	fmt.Println("conn start")
	go c.Read()

	//TODO
	// select {}

}
func (c *Connection) Stop() {
	fmt.Printf("id=%v stop\n", c.ID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}
func (c *Connection) GetConn() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetID() uint32 {
	return c.ID
}
