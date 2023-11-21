package impl

import "github.com/ikun666/v9/iface"

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
