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
	Send([]byte) error
}

// define a function to deal with the conn work
type HandleFunc func(*net.TCPConn, []byte, int) error
