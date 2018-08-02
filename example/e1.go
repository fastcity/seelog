package main

import (
	"github.com/xmge/seelog"
	"os"
	"time"
	"fmt"
	"log"
)

const (
	TestLog  = "test.log"	// yourproject 日志位置
	TestPort = 9999			// 查看日志时端口
)

// example
func main()  {

	// 在程序开始时 开启seelog 即可
	seelog.See(TestLog, TestPort)

	// 模拟your项目
	yourProject()
}

func yourProject()  {
	f, err := os.OpenFile(TestLog, os.O_RDWR|os.O_CREATE, 0766);	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 100; i++ {
		time.Sleep(500 * time.Millisecond)
		testLog := fmt.Sprintf("「模拟日志」[%s] 第[%d]行日志\n", time.Now().String(),i)
		_, err := f.WriteString(testLog)
		if err != nil {
			log.Println(err.Error())
		}
	}
	os.Remove(TestLog)
}
