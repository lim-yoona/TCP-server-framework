package seNet

import "TCP-server-framework/seInterface"

type Request struct {
	conn seInterface.IConnection
	data []byte
}

func (r *Request) GetConnection() seInterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
