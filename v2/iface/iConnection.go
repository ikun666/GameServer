package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	// Send([]byte)error
	// Recieve()[]byte
	GetConn() *net.TCPConn
	GetID() uint32
}
type HandleFunc func(*net.TCPConn, []byte, int) error
