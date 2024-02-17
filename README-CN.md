# TCP server framework
简体中文 | [English](README.md)

[![License](https://img.shields.io/github/license/lim-yoona/Golang-Design-Patterns)](LICENSE)

TCP server framework是一个基于TCP的轻量级服务器框架，支持高并发，用于快速构建网络通信应用。

## 架构

[![架构](https://s11.ax1x.com/2024/02/17/pFJZbS1.png)](https://imgse.com/i/pFJZbS1)  

## 特性

1. **`TLV` 消息格式**。 客⼾端服务端通讯采⽤TLV协议，实现消息包装器处理TCP字节流  
2. **多路由**。 通过消息标识符实现多路由处理消息，⽀持开发者⾃定义多路由   
3. **协程池**。 采⽤协程池处理业务，将消息负载均衡地分配给协程处理，实现协程的复⽤，避免海量协程调度的开销  
4. **顺序处理**。 将来⾃同⼀TCP连接的消息分配给协程池中的固定协程，实现消息按序处理  
5. **读写分离**。 对于每个连接，采⽤读写分离模型来处理，实现⾼内聚低耦合  

## 使用

首先， `clone` 本项目， 使用
```shell
git clone https://github.com/lim-yoona/TCP-server-framework.git
```
用于 `https` 或
```shell
git clone git@github.com:lim-yoona/TCP-server-framework.git
```
用于 `ssh`.

**demo** 文件夹下有一些用例。

具体来说, 你可以通过继承 **`BaseRouter`** 类并且重写它的方法来自定义路由。

```go
type MyRouter struct {
	seNet.BaseRouter
}

func (this *MyRouter) Handle(request seInterface.IRequest) {
	···
	err := request.GetConnection().SendMsg(200, []byte(···))
	···
}
```

然后你可以使用 **Server** 的 **`AddRouter`** 方法来注册路由, 第一个参数是 ID, 处理对应 ID 的消息, 第二个参数是你的自定义路由。

最后, 通过调用 **Server** 的 **`Serve`** 方法来启动服务器。

```go
func main() {
	s := seNet.NewServer()
	
	s.AddRouter(0, &MyRouter{})
	
	s.Serve()
}
```

你也可以通过 **Server** 的 **`SetOnConnStart`** 和 **`SetOnConnStop`** 方法来添加 **hook 函数** ， 就像 [示例](https://github.com/lim-yoona/TCP-server-framework/tree/main/demo/v1.0) 中那样.   

