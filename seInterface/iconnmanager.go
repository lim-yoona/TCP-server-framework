package seInterface

type IConnManager interface {
	// 添加链接
	AddConn(connection IConnection)
	// 删除链接
	DeleteConn(connection IConnection)
	// 根据ConnID获得链接
	GetConn(ConnID uint32) (IConnection, error)
	// 得到当前链接总数
	GetConnNum() int
	// 清除并终止所有的链接
	ClearConn()
}
