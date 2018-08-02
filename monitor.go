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
	offset := fileInfo.Size()
	for {
		fileInfo,_ = os.Stat(filePath)
		newOffset := fileInfo.Size()
		if offset < newOffset {
			msg := make([]byte,newOffset - offset)
			file,err := os.Open(filePath)
			if err != nil {
				log.Printf("[seelog] error:%v",err.Error())
			}
			_,err = file.Seek(offset,0)
			if err != nil {
				log.Printf("[seelog] error:%v",err.Error())
			}

			_,err = file.Read(msg)
			if err != nil {
				log.Printf("[seelog] error:%v",err.Error())
			}
			manager.broadcast <- msg
			offset = newOffset
		}
		time.Sleep(200 * time.Millisecond)
	}

}