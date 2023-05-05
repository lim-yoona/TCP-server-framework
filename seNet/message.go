package seNet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(msgId uint32, data []byte) *Message {
	return &Message{
		Id:      msgId,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 获取消息Id
func (m *Message) GetMesId() uint32 {
	return m.Id
}

// 获取消息长度
func (m *Message) GetMesLen() uint32 {
	return m.DataLen
}

// 获取消息内容
func (m *Message) GetMesData() []byte {
	return m.Data
}

// 设置消息Id
func (m *Message) SetMesId(id uint32) {
	m.Id = id
}

// 设置消息长度
func (m *Message) SetMesLen(len uint32) {
	m.DataLen = len
}

// 设置消息内容
func (m *Message) SetMesData(data []byte) {
	m.Data = data
}
