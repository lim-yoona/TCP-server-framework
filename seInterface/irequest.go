package seInterface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
