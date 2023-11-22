package iface

type IServer interface {
	Serve()
	Start()
	Stop()
	AddRouter(uint32, IRouter)
	GetConnManager() IConnManager
	SetOnConnCreate(func(IConnection))
	SetOnConnDestroy(func(IConnection))
	OnConnCreate(IConnection)
	OnConnDestroy(IConnection)
}
