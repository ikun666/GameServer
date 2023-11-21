package iface

type IMsgHandle interface {
	AddRouter(uint32, IRouter) error
	DoHandle(IRequest) error
	StartWorkerPool()
	Add2WorkerPool(IRequest)
}
