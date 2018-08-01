package seelog

import (
	"os"
	"time"
)

var buffer = make(chan string,1024)

func monitor(filePath string) error {
	fileInfo,err := os.Stat(filePath)
	if err != nil {
		return err
	}
	buff := make([]byte,1024)
	offset := fileInfo.Size()
	for {
		fileInfo,_ = os.Stat(filePath)
		newOffset := fileInfo.Size()
		if offset < newOffset {
			file,_ := os.Open(filePath)
			file.Seek(offset,0)
			file.Read(buff)
			buffer <- string(buff)
			offset = newOffset
		}
		time.Sleep(200 * time.Millisecond)
	}
}