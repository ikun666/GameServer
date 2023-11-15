package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/old/v2/iface"
)

type Connection struct {
	Conn     *net.TCPConn
	ID       uint32
	IsClosed bool
	API      iface.HandleFunc
	ExitChan chan struct{}
}

func NewConnetion(conn *net.TCPConn, id uint32, api iface.HandleFunc) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		IsClosed: false,
		API:      api,
		ExitChan: make(chan struct{}),
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
		err = c.API(c.Conn, buf, n)
		if err != nil {
			fmt.Printf("api err:%v", err)
			return
		}
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
