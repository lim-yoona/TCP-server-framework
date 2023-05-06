package main

import (
	"TCP-server-framework/seNet"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("client1 start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {
		log.Println("Dial err", err)
		return
	}
	for {
		dp := seNet.NewDataPack()
		binaryMsg, err := dp.Pack(seNet.NewMessage(1, []byte("I am 1 message!")))
		if err != nil {
			log.Println("Pack err", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			log.Println("Write err", err)
			return
		}

		// 读
		headData := make([]byte, dp.GetHeadLen())
		_, errr := io.ReadFull(conn, headData)
		if errr != nil {
			return
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			return
		}
		if msg.GetMesLen() > 0 {
			msge := msg.(*seNet.Message)
			msge.Data = make([]byte, msge.GetMesLen())
			_, err := io.ReadFull(conn, msge.Data)
			if err != nil {
				return
			}
		}
		fmt.Println("Receive message MsgId:", msg.GetMesId(), "MsgLen:", msg.GetMesLen(),
			"MsgData:", string(msg.GetMesData()))
		time.Sleep(time.Second * 5)
	}
}
