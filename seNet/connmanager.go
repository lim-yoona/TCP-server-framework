package seNet

import (
	"TCP-server-framework/seInterface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]seInterface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() seInterface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]seInterface.IConnection),
	}
}

func (c *ConnManager) AddConn(connection seInterface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.connections[connection.GetConnID()] = connection
	fmt.Println("ConnID=", connection.GetConnID(), "has been Add to ConnManager")
	fmt.Printf("***************ConnNum=%d******************\n", len(c.connections))
}

// 删除链接
func (c *ConnManager) DeleteConn(connection seInterface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, connection.GetConnID())
	fmt.Println("ConnID=", connection.GetConnID(), "has been deleted")
}

// 根据ConnID获得链接
func (c *ConnManager) GetConn(ConnID uint32) (seInterface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	conn, err := c.connections[ConnID]
	if !err {
		fmt.Printf("GetConn err", err)
		return nil, errors.New("GetConn err")
	}
	return conn, nil
}

// 得到当前链接总数
func (c *ConnManager) GetConnNum() int {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	return len(c.connections)
}

// 清除并终止所有的链接
func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for id, conn := range c.connections {
		conn.Stop()
		delete(c.connections, id)
	}
	fmt.Println("Clear All connections succ!")
}
