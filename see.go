package seelog

import (
	"net/http"
	"log"
	"golang.org/x/net/websocket"
)

var buffer = make(chan string,10)

// 入口函数
func See(filePath,addr string)  {
	if filePath == "" {
		log.Println("filePath 不可为空")
	}
	if addr == "" {
		log.Println("addr 不可为空")
	}
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/conn", websocket.Handler(genConn))
	log.Fatal(http.ListenAndServe(addr,nil))
}





