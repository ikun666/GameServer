package impl

import (
	"fmt"

	"github.com/ikun666/v7/iface"
)

type MsgHandle struct {
	HandleMap map[uint32]iface.IRouter
}

func NewMsgHandle() iface.IMsgHandle {
	return &MsgHandle{
		HandleMap: make(map[uint32]iface.IRouter),
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
