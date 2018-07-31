package seelog

import (
	"golang.org/x/net/websocket"
	"qiniupkg.com/x/log.v7"
)

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/echo"
var conn *websocket.Conn

// 进行页面的socket的连接
func genConn(ws *websocket.Conn)  {
	conn = ws
	go func() {
		msg := <- buffer
		view(msg)
	}()
}

// 向页面写数据
func view(msg string)  {
	_,err := conn.Write([]byte(msg))
	if err != nil {
		log.Println("[seelog]",err.Error())
	}
}