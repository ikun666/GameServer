package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/old/v5/iface"
)

type Connection struct {
	Conn     net.Conn
	ID       uint32
	IsClosed bool
	ExitChan chan struct{}
	Router   iface.IRouter
}

func NewConnetion(conn net.Conn, id uint32, router iface.IRouter) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		IsClosed: false,
		ExitChan: make(chan struct{}),
		Router:   router,
	}
}
func (c *Connection) Read() (iface.IMessage, error) {
	if c.IsClosed {
		fmt.Printf("conn close\n")
		return nil, fmt.Errorf("conn close")
	}
	pack := DataPack{}
	// fmt.Println("unpack start")
	msg, err := pack.UnPack(c.Conn)

	if err != nil {
		fmt.Printf("unpack err:%v\n", err)
		return nil, err
	}
	// fmt.Println("unpack end:", msg)
	fmt.Println(string(msg.GetMsgData()))
	return msg, nil

}
func (c *Connection) Write(id uint32, data []byte) error {
	if c.IsClosed {
		return fmt.Errorf("conn close")
	}
	msg := &Message{
		Len:  uint32(len(data)),
		ID:   id,
		Data: data,
	}
	pack := DataPack{}
	// fmt.Println("pack start")
	sendMsg, err := pack.Pack(msg)
	if err != nil {
		return fmt.Errorf("pack msg err:%v", err)
	}
	// fmt.Println("pack end:", sendMsg)
	_, err = c.Conn.Write(sendMsg)
	if err != nil {
		return fmt.Errorf("write msg err:%v", err)
	}
	// fmt.Println("send pack msg ok")
	return nil
}
func (c *Connection) Start() {
	defer c.Stop()
	for {
		msg, err := c.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		req := &Request{
			conn: c,
			msg:  msg,
		}
		// fmt.Println("req start")
		go func(req iface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}

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
func (c *Connection) GetConn() net.Conn {
	return c.Conn
}
func (c *Connection) GetID() uint32 {
	return c.ID
}
func (c *Connection) ConnIsClosed() bool {
	return c.IsClosed
}
