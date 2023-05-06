package seNet

import (
	"TCP-server-framework/seInterface"
	"log"
)

type MsgHandle struct {
	// 存放每个MsgId对应的处理方法
	Apis map[uint32]seInterface.IRouter
}

func NewMsgHandle() seInterface.IMsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]seInterface.IRouter),
	}
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
