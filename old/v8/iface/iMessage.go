package iface

type IMessage interface {
	GetMsgLen() uint32
	GetMsgID() uint32
	GetMsgData() []byte

	SetMsgLen(uint32)
	SetMsgID(uint32)
	SetMsgData([]byte)
}
