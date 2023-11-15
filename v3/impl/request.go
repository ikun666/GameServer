package impl

import "github.com/ikun666/v3/iface"

type Request struct {
	conn iface.IConnection
	data []byte
}

func (r *Request) GetConnetion() iface.IConnection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.data
}
