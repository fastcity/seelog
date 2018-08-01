package seelog

import (
	"net/http"
	"log"
	"golang.org/x/net/websocket"
	"fmt"
)

var addr string

// 入口函数
func See(filePath string,port int)  {
	go func() {

		defer func() {
			if err := recover();err != nil{
				log.Println("seelog 发生错误")
			}
		}()

		if filePath == "" {
			log.Println("filePath 不可为空")
		}
		if port == 0 {
			log.Println("port 不可为空")
		}
		addr = fmt.Sprintf(":%d",port)
		go monitor(filePath)
		http.HandleFunc("/",page)
		http.Handle("/ws", websocket.Handler(genConn))
		log.Fatal(http.ListenAndServe(addr,nil))
	}()
}





