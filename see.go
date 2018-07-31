package seelog

import (
	"net/http"
	"log"
	"golang.org/x/net/websocket"
)


// 入口函数
func See(filePath,addr string)  {
	if filePath == "" {
		log.Println("filePath 不可为空")
	}
	if addr == "" {
		log.Println("addr 不可为空")
	}
	go monitor(filePath)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(genConn))
	log.Fatal(http.ListenAndServe(addr,nil))
}





