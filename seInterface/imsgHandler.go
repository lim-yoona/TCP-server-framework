package seInterface

type IMsgHandle interface {
	// 调度路由
	DoMsgHandle(request IRequest)
	// 添加路由
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
