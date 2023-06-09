package seNet

import (
	"TCP-server-framework/seInterface"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Connection struct {
	TCPServer seInterface.IServer
	Conn      *net.TCPConn
	ConnID    uint32
	IsClose   bool
	ExitChan  chan bool
	// 无缓冲channel，用于读写之间的消息通信
	msgChan chan []byte
	// 链接处理模块
	MsgHandle seInterface.IMsgHandle
	// 给链接增加属性
	property map[string]interface{}
	// 保护属性map的锁
	pLock sync.RWMutex
}

// init Connection section
func NewConnection(TCPServer seInterface.IServer, conn *net.TCPConn, connId uint32, handle seInterface.IMsgHandle) *Connection {
	c := &Connection{
		TCPServer: TCPServer,
		Conn:      conn,
		ConnID:    connId,
		IsClose:   false,
		msgChan:   make(chan []byte),
		ExitChan:  make(chan bool, 1),
		MsgHandle: handle,
		property:  make(map[string]interface{}),
	}
	TCPServer.GetConnMan().AddConn(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running...]")
	defer fmt.Println("ConnID=", c.ConnID, "Reader Existed RemoteAddr is", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端输入到msg中
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			log.Println("ReadFull err", err)
			return
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			log.Println("UnPack err", err)
			return
		}
		if msg.GetMesLen() > 0 {
			msge := msg.(*Message)
			msge.Data = make([]byte, msge.GetMesLen())
			_, err := io.ReadFull(c.Conn, msge.Data)
			if err != nil {
				log.Println("ReadFull err", err)
				return
			}
		}
		// 得到当前链接的请求数据
		req := &Request{
			conn: c,
			msg:  msg,
		}
		// 从路由中，找到注册绑定的Conn对应的Router调用
		c.MsgHandle.SendMsgToTaskQueue(req)
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running...]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!###########]")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				log.Println("Send data err", err)
				return
			}
			fmt.Println("[Writer send data succeed!]", data)
		// raeder告知writer客户端已经退出
		case <-c.ExitChan:
			fmt.Println("writer receive ExitChan")
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn started  ConnID=", c.ConnID)
	go c.StartReader()
	// 启动写数据的goroutine
	go c.StartWriter()
	// 调用OnConnStart函数
	c.TCPServer.CallOnConnStart(c)
}

func (c *Connection) SendMsg(MsgId uint32, data []byte) error {
	if c.IsClose == true {
		return errors.New("Connection Closed when send message")
	}
	dp := NewDataPack()
	msg := NewMessage(MsgId, data)
	sendmsg, err := dp.Pack(msg)
	if err != nil {
		log.Println("Pack err", err)
		return err
	}
	c.msgChan <- sendmsg
	return nil
}

// stop conn
func (c *Connection) Stop() {
	fmt.Println("Conn Closeed  ConnID=", c.ConnID)
	if c.IsClose == true {
		return
	}
	c.IsClose = true

	// 调用OnConnStop函数
	c.TCPServer.CallOnConnStop(c)
	c.Conn.Close()
	c.ExitChan <- true
	c.TCPServer.GetConnMan().DeleteConn(c)
	close(c.ExitChan)
	close(c.msgChan)
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

// 设置链接属性
func (c *Connection) SetProperty(k string, v interface{}) {
	c.pLock.Lock()
	defer c.pLock.Unlock()
	c.property[k] = v
}

// 获得链接属性
func (c *Connection) GetProperty(k string) (interface{}, error) {
	c.pLock.RLock()
	defer c.pLock.RUnlock()
	p, ok := c.property[k]
	if ok {
		return p, nil
	} else {
		fmt.Println("GetProperty k=", k, "err")
		return nil, errors.New("GetProperty err")
	}
}

// 删除链接属性
func (c *Connection) DeleteProperty(k string) {
	c.pLock.Lock()
	defer c.pLock.Unlock()
	_, ok := c.property[k]
	if ok {
		delete(c.property, k)
	} else {
		fmt.Println("DeleteProperty err No property k=", k)
	}
}
