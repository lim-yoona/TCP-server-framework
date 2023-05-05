package seNet

import "TCP-server-framework/seInterface"

type Request struct {
	conn seInterface.IConnection
	msg  seInterface.IMessage
}

func (r *Request) GetConnection() seInterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMesData()
}
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMesId()
}
