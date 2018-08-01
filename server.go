package seelog

import (
	"net/http"
	"fmt"
	"log"
	"golang.org/x/net/websocket"
	"time"
)

// 开启 httpServer
func server(port int)  {

	defer func() {
		if err := recover();err != nil{
			log.Printf("[seelog] error:%+v",err)
		}
	}()

	// 返回页面
	http.HandleFunc("/",page)
	// socket链接
	http.Handle("/ws", websocket.Handler(genConn))
	// 测试
	http.Handle("/test",http.FileServer(http.Dir("index.html")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port),nil))
}

// 创建client对象
func genConn(ws *websocket.Conn)  {
	client := &Client{time.Now().String(),ws,make(chan []byte,1024)}
	manager.register <- client
	go client.read()
	client.write()
}
