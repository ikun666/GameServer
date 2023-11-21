package iface

type IServer interface {
	Serve()
	Start()
	Stop()
	AddRouter(uint32, IRouter)
}
