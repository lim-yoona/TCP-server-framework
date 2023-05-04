package seNet

import (
	"TCP-server-framework/seInterface"
	"fmt"
	"log"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	IsClose   bool
	HandleAPI seInterface.HandleFunc
	ExitChan  chan bool
}

// init Connection section
func NewConnection(conn *net.TCPConn, connId uint32, callBackFunc seInterface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connId,
		HandleAPI: callBackFunc,
		IsClose:   false,
		ExitChan:  make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID=", c.ConnID, "Reader Existed RemoteAddr is", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端输入到buf中
		buf := make([]byte, 512)
		num, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Reader Goroutine Read err", err)
			return
		}
		if err = c.HandleAPI(c.Conn, buf, num); err != nil {
			log.Println("ConnID=", c.ConnID, " HandleAPI err", err)
			return
		}

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

}
