package impl

import (
	"fmt"
	"math/rand"

	"github.com/ikun666/old/v8/conf"
	"github.com/ikun666/old/v8/iface"
)

type MsgHandle struct {
	HandleMap      map[uint32]iface.IRouter
	WorkerPool     []chan iface.IRequest
	WorkerPoolSize uint32
	WorkerChanSize uint32
}

func NewMsgHandle() iface.IMsgHandle {
	return &MsgHandle{
		HandleMap:      make(map[uint32]iface.IRouter),
		WorkerPool:     make([]chan iface.IRequest, conf.GConfig.WorkerPoolSize),
		WorkerPoolSize: conf.GConfig.WorkerPoolSize,
		WorkerChanSize: conf.GConfig.WorkerChanSize,
	}
}
func (m *MsgHandle) AddRouter(id uint32, r iface.IRouter) error {
	if _, ok := m.HandleMap[id]; ok {
		fmt.Println("handle has existed")
		return fmt.Errorf("handle has existed")
	}
	m.HandleMap[id] = r
	return nil
}
func (m *MsgHandle) DoHandle(req iface.IRequest) error {
	r, ok := m.HandleMap[req.GetMessage().GetMsgID()]
	if !ok {
		fmt.Println("handle no exist")
		return fmt.Errorf("handle no exist")
	}
	r.PreHandle(req)
	r.Handle(req)
	r.PostHandle(req)
	return nil
}
func (m *MsgHandle) StartWorkerPool() {
	fmt.Println("start worker pool")
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.WorkerPool[i] = make(chan iface.IRequest, m.WorkerChanSize)
		go func(i int, worker chan iface.IRequest) {
			fmt.Println("worker ", i, " start")
			for {
				req := <-worker
				m.DoHandle(req)
			}
		}(i, m.WorkerPool[i])
	}
}
func (m *MsgHandle) Add2WorkerPool(req iface.IRequest) {
	id := rand.Intn(int(m.WorkerPoolSize))
	fmt.Println("add req ", id)
	m.WorkerPool[id] <- req
}
