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

func (this *PingRouter) Handle(request seInterface.IRequest) {
	fmt.Println("Router Handle Called!")
	fmt.Println("receive from client msgId:", request.GetMsgId(),
		"msgData:", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte(request.GetData()))
	if err != nil {
		log.Println("SendMsg err", err)
	}
}

// 第二个路由
type HelloRouter struct {
	seNet.BaseRouter
}

func (this *HelloRouter) Handle(request seInterface.IRequest) {
	fmt.Println("Router HelloRouter Called!")
	fmt.Println("receive from client msgId:", request.GetMsgId(),
		"msgData:", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("hello TCPServer Framework!"))
	if err != nil {
		log.Println("SendMsg err", err)
	}
}
func ConnBegin(connection seInterface.IConnection) {
	fmt.Println("ConnBegin Hook Called!")
	err := connection.SendMsg(200, []byte("OnConnStart Called!"))
	if err != nil {
		log.Println("ConnBegin err", err)
	}
	connection.SetProperty("zhangsan", "student")
	connection.SetProperty("age", 25)
}
func ConnEnd(connection seInterface.IConnection) {
	fmt.Println("ConnEnd Hook Called!")
	err := connection.SendMsg(200, []byte("OnConnEnd Called!"))
	if err != nil {
		log.Println("ConnEnd err", err)
	}
	p, err := connection.GetProperty("age")
	if err != nil {
		fmt.Println("GetProperty err", err)
	}
	fmt.Println("k=age  v=", p)
	p, err = connection.GetProperty("zhangsan")
	if err != nil {
		fmt.Println("GetProperty err", err)
	}
	fmt.Println("k=zhangsan  v=", p)
}

// 基于TCP服务器来开发的服务器端应用程序
func main() {
	// 1.使用TCPserver的API来创建一个server句柄
	s := seNet.NewServer()

	// 注册Hook函数
	s.SetOnConnStart(ConnBegin)
	s.SetOnConnStop(ConnEnd)
	// 这个版本需要添加路由了
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	// 2.启动server
	s.Serve()
}
