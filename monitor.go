package seelog

import (
	"log"
	"os"
	"time"
)

// 监控日志文件
func monitor(filePath string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[seelog] error:%+v", err)
		}
	}()

	var fileInfo os.FileInfo
	var err error
	for i :=1; i <= 10; i++{
		fileInfo, err = os.Stat(filePath)
		if err != nil {
			log.Printf("[seelog] error:%v", err.Error())
			continue
		}
		break
	}

	offset := fileInfo.Size()
	for {
		fileInfo, err = os.Stat(filePath)
		if err != nil {
			log.Printf("[seelog] error:%v", err.Error())
			continue
		}
		newOffset := fileInfo.Size()
		if offset < newOffset {
			msg := make([]byte, newOffset-offset)
			file, err := os.Open(filePath)
			if err != nil {
				log.Printf("[seelog] error:%v", err.Error())
				continue
			}
			_, err = file.Seek(offset, 0)
			if err != nil {
				log.Printf("[seelog] error:%v", err.Error())
			}

			_, err = file.Read(msg)
			if err != nil {
				log.Printf("[seelog] error:%v", err.Error())
			}
			manager.broadcast <- msg
			offset = newOffset
			file.Close()
		}
		offset = newOffset
		time.Sleep(200 * time.Millisecond)
	}

}
