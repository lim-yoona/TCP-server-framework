package seInterface

// 将请求的消息封装到Message中，定义一个抽象模块
type IMessage interface {
	// 获取消息Id
	GetMesId() uint32
	// 获取消息长度
	GetMesLen() uint32
	// 获取消息内容
	GetMesData() []byte
	// 设置消息Id
	SetMesId(uint32)
	// 设置消息长度
	SetMesLen(uint32)
	// 设置消息内容
	SetMesData([]byte)
}
