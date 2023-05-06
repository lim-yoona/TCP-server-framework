package seNet

import (
	"TCP-server-framework/seInterface"
	"TCP-server-framework/utils"
	"fmt"
	"log"
	"net"
)

// 接口实现，定义一个服务器模块
type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
	// 路由，服务器注册的链接对应的处理业务
	MsgHandler seInterface.IMsgHandle
}

// init server
func NewServer() seInterface.IServer {
	return &Server{
		Name:       utils.GlobalObject.Name,
		IpVersion:  "tcp4",
		Ip:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}
}

func (s *Server) AddRouter(msgId uint32, router seInterface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("AddRouter succeed!")
}

// server start
func (s *Server) Start() {
	fmt.Printf("[TCPServer] Start! ServerName:%s Listener at %s:%d is started!\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[TCPServer] Version:%s MaxConn:%d MaxPackageSize:%d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	// 放到一个goroutine中来做，这样主进程就不会阻塞，异步
	go func() {
		// 1.获取一个TCP的Addr，也就是套接字
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			log.Println("ResolveTCPAddr err", err)
			return
		}
		// 2.监听这个套接字
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			log.Println("ListenTCP err", err)
			return
		}
		fmt.Println("[Start succeed!]", s.Name, "succeed,listening...")

		// ConnID
		var ConnId uint32
		ConnId = 0
		// 3.连接到了之后Accept得到链接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("AcceptTCP err", err)
				continue
			}
			Connect := NewConnection(conn, ConnId, s.MsgHandler)
			ConnId++
			go Connect.Start()
		}
	}()

}

// server stop
func (s *Server) Stop() {
	//TODO 将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
}

// server run
func (s *Server) Serve() {
	s.Start()

	//TODO 可以做一些启动服务器之后的额外业务
	select {}
}
