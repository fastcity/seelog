package seelog_test

import (
	"testing"
	"os"
	"fmt"
	"time"
)

func TestSee(t *testing.T) {

/*	go func() {
		os.Remove("test.log")
		f,err := os.Create("test.log")
		if err != nil {
			t.Fatal(err)
		}
		for i := 1; i<= 100; i++ {
			time.Sleep(1 * time.Second)
			testLog := fmt.Sprintf("第[%d]行日志\n",i)
			fmt.Print(testLog)
			f.WriteString(testLog)
		}
	}()*/
	//seelog.See("http://localhost:9999","test.log")

	f,err := os.Open("test.log")
	if err != nil {
	}
	fmt.Println(f.Stat())
}

func TestSee2(t *testing.T) {

	fileInfo,err := os.Stat("test.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	buff := make([]byte,1024)
	offset := fileInfo.Size()
	for {
		fmt.Println(offset)
		fileInfo,_ = os.Stat("test.log")
		fmt.Println(fileInfo.Size())
		if newOffset := fileInfo.Size();offset < newOffset {
			file,_ := os.Open("test.log")
			file.Seek(offset,0)
			file.Read(buff)
			fmt.Println(string(buff))
			offset = newOffset
		}
		time.Sleep(1 *time.Second)
		fmt.Println("正在监听")
	}
}



