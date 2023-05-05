package seNet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		return
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						return
					}
					msg, err := dp.UnPack(headData)
					if err != nil {
						return
					}
					if msg.GetMesLen() > 0 {
						msge := msg.(*Message)
						msge.Data = make([]byte, msge.GetMesLen())
						_, err := io.ReadFull(conn, msge.Data)
						if err != nil {
							return
						}
					}
					fmt.Println("Receive message MsgId:", msg.GetMesId(), "MsgLen:", msg.GetMesLen(),
						"MsgData:", string(msg.GetMesData()))
				}
			}(conn)
		}

	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		return
	}
	dp := NewDataPack()
	// 模拟粘包过程，封装两个msg一起发送
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		return
	}
	msg2 := &Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte("worldee"),
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		return
	}
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	select {}
}
