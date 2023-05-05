package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {
		log.Println("Dial err", err)
		return
	}
	for {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			log.Println("Write err", err)
		}

		buf := make([]byte, 512)
		num, errr := conn.Read(buf)
		if errr != nil {
			log.Println("client Read err", err)
		}
		fmt.Println("recieve server message", string(buf[:num]))
		time.Sleep(time.Second)
	}
}
