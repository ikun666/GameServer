package impl

type Message struct {
	Len  uint32 `json:"len,omitempty"`
	ID   uint32 `json:"id,omitempty"`
	Data []byte `json:"data,omitempty"`
}

func (m *Message) GetMsgLen() uint32 {
	return m.Len
}
func (m *Message) GetMsgID() uint32 {
	return m.ID
}
func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgLen(len uint32) {
	m.Len = len
}
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
