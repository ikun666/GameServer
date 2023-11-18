package iface

type IRequest interface {
	GetConnetion() IConnection
	GetMessage() IMessage
}
