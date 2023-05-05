package seNet

import (
	"TCP-server-framework/seInterface"
	"TCP-server-framework/utils"
	"fmt"
	"log"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	IsClose  bool
	ExitChan chan bool
	// 链接处理的方法
	Router seInterface.IRouter
}

// init Connection section
func NewConnection(conn *net.TCPConn, connId uint32, router seInterface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connId,
		IsClose:  false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID=", c.ConnID, "Reader Existed RemoteAddr is", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端输入到buf中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Reader Goroutine Read err", err)
			return
		}
		// 得到当前链接的请求数据
		req := &Request{
			conn: c,
			data: buf,
		}
		// 从路由中，找到注册绑定的Conn对应的Router调用
		go func(request *Request) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn started  ConnID=", c.ConnID)
	go c.StartReader()
	//TODO 启动从当前链接写数据的业务
}

// stop conn
func (c *Connection) Stop() {
	fmt.Println("Conn Closeed  ConnID=", c.ConnID)
	if !c.IsClose {
		return
	}
	c.IsClose = true
	c.Conn.Close()
	close(c.ExitChan)
}

// gain socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// gain conn ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// gain remote client status
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// send data to remote client
func (c *Connection) Send([]byte) error {
	return nil
}
