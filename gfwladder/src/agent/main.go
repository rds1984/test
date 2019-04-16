package main

import (
	"common"
	"net"
)

func main() {

	laddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:10000")
	common.HandleError("获取监听地址出错", err)
	net.ListenTCP("tcp4")

}
