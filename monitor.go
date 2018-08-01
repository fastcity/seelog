package seelog

import (
	"os"
	"time"
	"log"
)


// 监控日志文件
func monitor(filePath string){
	defer func() {
		if err := recover();err != nil{
			log.Printf("[seelog] error:%+v",err)
		}
	}()

	fileInfo,err := os.Stat(filePath)
	if err != nil {
		log.Printf("[seelog] error:%v",err.Error())
	}
	msg := make([]byte,1024)
	offset := fileInfo.Size()
	for {
		fileInfo,_ = os.Stat(filePath)
		newOffset := fileInfo.Size()
		if offset < newOffset {
			file,_ := os.Open(filePath)
			file.Seek(offset,0)
			file.Read(msg)
			manager.broadcast <- msg
			offset = newOffset
		}
		time.Sleep(200 * time.Millisecond)
	}

}