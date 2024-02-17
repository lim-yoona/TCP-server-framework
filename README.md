# TCP server framework
English | [简体中文](README-CN.md)

[![License](https://img.shields.io/github/license/lim-yoona/Golang-Design-Patterns)](LICENSE)

The TCP server framework is a TCP based lightweight server framework that can help developers quickly build network communication applications, supporting high concurrency.  

## Architecture  

[![Architecture](https://s11.ax1x.com/2024/02/17/pFJ95Ix.md.png)](https://imgse.com/i/pFJ95Ix)  

## Feature

1. **`TLV` message format.** `TLV` protocol is used for communication between client and server, and `TCP` byte stream is processed by a message packager  
2. **Multi-Route.** Through the message ID to implement the use of multi-route to process messages, support the developer custom multi-route to process messages  
3. **`Goroutine` Pool.** The `Goroutine` Pool is used to process messages and reuse `Goroutines`. The server distributes the message load evenly to `Goroutine` processing, avoiding the overhead of massive `Goroutine` scheduling  
4. **Sequential processing.** Messages from the same `TCP` connection are assigned to a fixed `Goroutine` in `Goroutine` Pool to implement sequential processing of messages  
5. **read-write separation.** For each connection, a read-write separation model is adopted to achieve high cohesion and low coupling  

## Usage

First, clone this project, using
```shell
git clone https://github.com/lim-yoona/TCP-server-framework.git
```
for `https` or 
```shell
git clone git@github.com:lim-yoona/TCP-server-framework.git
```
for `ssh`.  

Some use cases are in the **demo** folder.   

Specifically, you can customize the route by inheriting the **`BaseRouter`** class and rewriting its methods.  

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

You can then register the route using the **Server**'s **`AddRouter`** method, with the first parameter being the message ID, processing the message corresponding to that ID, and the second parameter being your custom route.  

Finally, start the server by calling the **`Serve`** method of **Server** .  

```go
func main() {
	s := seNet.NewServer()
	
	s.AddRouter(0, &MyRouter{})
	
	s.Serve()
}
```

You can also add **hook functions** through the **`SetOnConnStart`** and **`SetOnConnStop`** methods of the **Server** , as shown in the [example](https://github.com/lim-yoona/TCP-server-framework/tree/main/demo/v1.0).   

