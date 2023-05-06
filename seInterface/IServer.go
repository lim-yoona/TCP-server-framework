package seInterface

// 定义一个服务器接口
type IServer interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()

	// 路由功能：給当前服务器注册一个路由方法，供客户端的链接处理使用
	AddRouter(msgId uint32, router IRouter)
	GetConnMan() IConnManager
	SetOnConnStart(func(conn IConnection))
	SetOnConnStop(func(conn IConnection))
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
