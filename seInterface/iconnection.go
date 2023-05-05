package seInterface

import "net"

type IConnection interface {
	// start conn
	Start()
	// stop conn
	Stop()
	// gain socket
	GetTCPConnection() *net.TCPConn
	// gain conn ID
	GetConnID() uint32
	// gain remote client status
	RemoteAddr() net.Addr
	// send data to remote client
	SendMsg(uint32, []byte) error
}
