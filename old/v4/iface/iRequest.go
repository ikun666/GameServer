package iface

type IRequest interface {
	GetConnetion() IConnection
	GetData() []byte
}
