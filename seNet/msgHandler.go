package seNet

import (
	"TCP-server-framework/seInterface"
	"TCP-server-framework/utils"
	"fmt"
	"log"
)

type MsgHandle struct {
	// 存放每个MsgId对应的处理方法
	Apis map[uint32]seInterface.IRouter
	// 负责worker取任务的消息队列
	TaskQueue []chan seInterface.IRequest
	// worker池最大数量
	WorkerPoolSize uint32
}

func NewMsgHandle() seInterface.IMsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]seInterface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan seInterface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request seInterface.IRequest) {
	// 将消息分配给不同的Worker
	// 这里用ConnId来分配
	// 达到的效果是，一个worker处理一部分链接对应的消息
	workerId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(),
		"Request MsgID=", request.GetMsgId(), "To worker", workerId)
	// 放到worker对应的channel中
	mh.TaskQueue[workerId] <- request
}

func (mh *MsgHandle) DoMsgHandle(request seInterface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		log.Println("MsgId=", request.GetMsgId(), "The router not existed!")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 添加路由
func (mh *MsgHandle) AddRouter(msgId uint32, router seInterface.IRouter) {
	// 先判断当前msgid对应的路由是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		log.Println("msgId:", msgId, " router existed!")
		return
	}
	// 未存在，添加msgid对应的路由
	mh.Apis[msgId] = router
	log.Println("msgId:", msgId, "Add succeed!")
}

// 启动一个工作池
func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 为当前worker创建一个消息队列
		mh.TaskQueue[i] = make(chan seInterface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 启动worker
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个worker
func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan seInterface.IRequest) {
	fmt.Println("WorkerId=", workerId, "is Started!")
	// 阻塞等待消息队列的消息
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandle(request)
		}
	}
}
