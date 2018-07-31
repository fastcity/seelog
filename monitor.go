package seelog

import (
	"os"
	"fmt"
	"time"
)

var buffer = make(chan string,10000)

func monitor(filePath string) error {
	fileInfo,err := os.Stat(filePath)
	if err != nil {
		return err
	}
	buff := make([]byte,1024)
	offset := fileInfo.Size()
	for {
		fmt.Printf("offset=%d\n",offset)
		fileInfo,_ = os.Stat(filePath)
		newOffset := fileInfo.Size()
		fmt.Printf("newOffset=%d\n",newOffset)
		if offset < newOffset {
			file,_ := os.Open(filePath)
			file.Seek(offset,0)
			file.Read(buff)
			fmt.Println(string(buff))
			buffer <- string(buff)
			offset = newOffset
		}
		time.Sleep(1 *time.Second)
	}
}