package utils

import (
	"TCP-server-framework/seInterface"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// 存储一切关于框架的全局参数，供其他模块使用
// 一些参数是可以通过zinx,json由用户进行配置
type GlobalObj struct {
	// server
	// 当前服务器全局server对象
	TcpServer seInterface.IServer
	// 服务器监听的IP
	Host    string
	TcpPort int
	Name    string
	// TCP server
	Version string
	// 最大连接数
	MaxConn int
	// 最大包大小
	MaxPackageSize uint32
}

// 定义一个全局的对象
var GlobalObject *GlobalObj

func (g *GlobalObj) LoadConf() {
	confFile, err := os.Open("conf/TCPServer.json")
	if err != nil {
		log.Println("Read config file false", err)
		panic(err)
	}
	defer confFile.Close()
	data, errr := ioutil.ReadAll(confFile)
	if errr != nil {
		log.Println("Read config file false", errr)
		panic(errr)
	}
	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		log.Println("Read config file false", err)
		panic(err)
	}
}

// 提供一个init方法，初始化当前的对象
func init() {
	// 如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:           "TCPServerApp",
		TcpPort:        8999,
		Version:        "V0.4",
		MaxConn:        10000,
		MaxPackageSize: 4096,
		Host:           "0.0.0.0",
	}
	GlobalObject.LoadConf()
}
