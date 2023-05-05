package main

import (
	"TCP-server-framework/seInterface"
	"TCP-server-framework/seNet"
	"fmt"
	"log"
)

// 自定义路由,继承baseRouter
type PingRouter struct {
	seNet.BaseRouter
}

func (this *PingRouter) PreHandle(request seInterface.IRequest) {
	fmt.Println("Router PreHandle Called!")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before Ping..."))
	if err != nil {
		log.Println("PreHandle Write False")
		return
	}
}
func (this *PingRouter) Handle(request seInterface.IRequest) {
	fmt.Println("Router Handle Called!")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping"))
	if err != nil {
		log.Println("Handle Write False")
		return
	}
}

func (this *PingRouter) PostHandle(request seInterface.IRequest) {
	fmt.Println("Router PostHandle Called!")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After Ping..."))
	if err != nil {
		log.Println("PostHandle Write False")
		return
	}
}

// 基于TCP服务器来开发的服务器端应用程序
func main() {
	// 1.使用TCPserver的API来创建一个server句柄
	s := seNet.NewServer()

	// 这个版本需要添加路由了
	s.AddRouter(&PingRouter{})
	// 2.启动server
	s.Serve()
}
