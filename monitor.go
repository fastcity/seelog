package seelog

import (
	"os"
	"fmt"
)

var msgBuffer = make(chan string,10)

func monitor(filePath string) error {
	file,err := os.Open(filePath)
	if err != nil {
		return err
	}
	fileInfo,err := file.Stat()
	if err != nil {
		return err
	}

	buff := make([]byte,1024)
	offset := fileInfo.Size()
	for {
		if offset <= fileInfo.Size() {
			file.Seek(offset,0)
			file.Read(buff)
			fmt.Println(string(buff))
		}
	}

	return nil
}