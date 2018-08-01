package seelog_test

import (
	"testing"
	"os"
	"fmt"
	"time"
	"log"
	"github.com/xmge/seelog"
)

func TestSee(t *testing.T) {

	go func() {
		err := os.Remove("test.log")
		if err != nil {
			log.Println(err.Error())
		}
		//f,err := os.OpenFile("test.log",os.O_APPEND|os.O_WRONLY,0666)
		f,err := os.Create("test.log")
		if err != nil {
			t.Log(err.Error())
			return
		}
		for i := 1; i<= 100; i++ {
			time.Sleep(1 * time.Second)
			testLog := fmt.Sprintf("第[%d]行日志\n",i)
			_,err := f.WriteString(testLog)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()
	seelog.See("test.log",9999)
}

func TestSee2(t *testing.T)  {
	for {
		fileInfo,_ := os.Stat("test.log")
		fmt.Println(fileInfo.Size())
		fileInfo,_ = os.Stat("test.log")
		fmt.Println(fileInfo.Size())
		time.Sleep(1 *time.Second)
	}

}