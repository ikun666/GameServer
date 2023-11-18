package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	Read() (IMessage, error)
	Write(uint32, []byte) error
	GetConn() net.Conn
	GetID() uint32
	ConnIsClosed() bool
}
