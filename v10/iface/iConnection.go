package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	Read() (IMessage, error)
	Write(uint32, []byte) error
	GetConn() net.Conn
	GetID() uint32
	SetProperty(string, any)
	GetProperty(string) (any, error)
	RemoveProperty(string)
}
