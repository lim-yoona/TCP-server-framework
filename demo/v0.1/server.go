package main

import "TCP-server-framework/seNet"

// 基于TCP服务器来开发的服务器端应用程序
func main() {
	// 1.使用TCPserver的API来创建一个server句柄
	s := seNet.NewServer("[seNet1.0]")
	// 2.启动server
	s.Serve()
}
