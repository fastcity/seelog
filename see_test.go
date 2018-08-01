package seelog_test

import (
	"fmt"
	"github.com/xmge/seelog"
	"log"
	"os"
	"testing"
	"time"
)

const (
	TestLog  = "test.log"
	TestPort = 9999
)

func TestSee(t *testing.T) {

	// 测试
	seelog.See(TestLog, TestPort)

	// 模拟日志输出
	err := os.Remove(TestLog)
	if err != nil {
		log.Println(err.Error())
	}
	f, err := os.Create(TestLog)
	if err != nil {
		t.Log(err.Error())
		return
	}
	for i := 1; i <= 10000; i++ {
		time.Sleep(1 * time.Second)
		testLog := fmt.Sprintf("「模拟日志」[%s] 第[%d]行日志\n", time.Now().String(),i)
		_, err := f.WriteString(testLog)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
