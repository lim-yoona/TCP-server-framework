package seNet

import (
	"TCP-server-framework/seInterface"
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
}

// init server
func NewServer(name string) seInterface.IServer {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}
}

// server start
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at %s:%d is started!", s.Ip, s.Port)

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
		// 3.连接到了之后Accept得到链接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("AcceptTCP err", err)
				continue
			}
			// 先处理一个简单的业务
			go func() {
				for {
					buf := make([]byte, 512)
					num, err := conn.Read(buf)
					if err != nil {
						log.Println("Read err", err)
						continue
					}
					_, err = conn.Write(buf[:num])
					if err != nil {
						log.Println("Write err", err)
						continue
					}
					fmt.Println("[Return] Writer back succeed!")
				}
			}()
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
