package seNet

import "TCP-server-framework/seInterface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写就好了
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request seInterface.IRequest)  {}
func (br *BaseRouter) Handle(request seInterface.IRequest)     {}
func (br *BaseRouter) PostHandle(request seInterface.IRequest) {}

// 为什么不让用户直接重写接口，因为重写接口需要实现接口中所有的方法，但是并不是每个方法都是用户需要的，所以
// 先写一个基类，确保每个方法都有，然后让用户继承，需要哪个方法就重写哪个方法
