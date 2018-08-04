package seelog

import (
	"net/http"
	"fmt"
	"log"
	"golang.org/x/net/websocket"
	"time"
	"html/template"
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
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("index.html")
		if (err != nil) {
			log.Println(err)
		}
		t.Execute(writer, nil)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port),nil))
}

// 创建client对象
func genConn(ws *websocket.Conn)  {
	client := &client{time.Now().String(),ws,make(chan []byte,1024)}
	manager.register <- client
	go client.read()
	client.write()
}
