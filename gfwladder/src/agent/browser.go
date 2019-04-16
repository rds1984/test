package main

import (
	"common"
	"net"
)

func StartBrowser() {
	laddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:10000")
	common.HandleError(err, "获取监听地址出错")
	listen, err := net.ListenTCP("tcp4", laddr)
	common.HandleError(err, "监听地址出错", laddr)
	for {
		con, err := listen.Accept()
		common.HandleInfo(err, con)
	}
}
