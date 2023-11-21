package impl

import "github.com/ikun666/old/v8/iface"

type Request struct {
	conn iface.IConnection
	// data []byte
	msg iface.IMessage
}

func (r *Request) GetConnetion() iface.IConnection {
	return r.conn
}
func (r *Request) GetMessage() iface.IMessage {
	return r.msg
}
