package impl

import (
	"fmt"
	"net"
	"sync"

	"github.com/ikun666/v10/conf"
	"github.com/ikun666/v10/iface"
)

type Connection struct {
	Server       iface.IServer
	Conn         net.Conn
	ID           uint32
	IsClosed     bool
	ExitChan     chan struct{}
	MsgChan      chan []byte
	MsgHandle    iface.IMsgHandle
	Property     map[string]any
	PropertyLock sync.RWMutex
}

func NewConnetion(server iface.IServer, conn net.Conn, id uint32, msgHandle iface.IMsgHandle) iface.IConnection {
	return &Connection{
		Server:    server,
		Conn:      conn,
		ID:        id,
		IsClosed:  false,
		ExitChan:  make(chan struct{}),
		MsgChan:   make(chan []byte, conf.GConfig.MaxPackageSize),
		MsgHandle: msgHandle,
		Property:  make(map[string]any),
	}
}
func (c *Connection) Read() (iface.IMessage, error) {
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
func (c *Connection) Reader() {
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
		// go func(req iface.IRequest) {
		// 	c.Router.PreHandle(req)
		// 	c.Router.Handle(req)
		// 	c.Router.PostHandle(req)
		// }(req)
		// go c.MsgHandle.DoHandle(req)
		c.MsgHandle.Add2WorkerPool(req)
	}
}
func (c *Connection) Write(id uint32, data []byte) error {
	// if c.IsClosed {
	// 	return fmt.Errorf("conn close")
	// }
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
	// _, err = c.Conn.Write(sendMsg)
	// if err != nil {
	// 	return fmt.Errorf("write msg err:%v", err)
	// }
	c.MsgChan <- sendMsg
	// fmt.Println("send pack msg ok")
	return nil
}
func (c *Connection) Writer() {
	defer c.Stop()
	for {
		select {
		case data := <-c.MsgChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Println("writer err:", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}

}
func (c *Connection) Start() {
	//add in there  remove in connection.stop()

	go c.Reader()
	go c.Writer()
	c.Server.GetConnManager().Add(c)
	c.Server.OnConnCreate(c)
}
func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}
	c.Server.OnConnDestroy(c)
	fmt.Printf("id=%v stop\n", c.ID)
	c.IsClosed = true
	c.Conn.Close()
	c.ExitChan <- struct{}{}
	close(c.ExitChan)
	close(c.MsgChan)
	c.Server.GetConnManager().Remove(c)
}
func (c *Connection) GetConn() net.Conn {
	return c.Conn
}
func (c *Connection) GetID() uint32 {
	return c.ID
}
func (c *Connection) SetProperty(key string, value any) {
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()
	c.Property[key] = value
}
func (c *Connection) GetProperty(key string) (any, error) {
	c.PropertyLock.RLock()
	defer c.PropertyLock.RUnlock()
	if v, ok := c.Property[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("no this %s property", key)
}
func (c *Connection) RemoveProperty(key string) {
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()
	delete(c.Property, key)
}
