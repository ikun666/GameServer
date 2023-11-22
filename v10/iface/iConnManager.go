package iface

type IConnManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(id uint32) IConnection
	Len() int
	ClearConn()
}
