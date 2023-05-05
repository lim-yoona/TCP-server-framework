package seInterface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsgId() uint32
}
