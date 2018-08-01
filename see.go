package seelog

import (
	"log"
)

//  启动seelog
func See(filePath string,port int) {

		// 检查参数
		if !checkParam(filePath,port){
			return
		}

		// 开启socket管理器
		go manager.start()
		// 监控文件
		go monitor(filePath)
		// 开启httpServer
		go server(port)
}


// 参数验证
func checkParam(filePath string,port int) bool {
	if filePath == "" {
		log.Println("filePath 不可为空")
		return false
	}
	if port == 0 {
		log.Println("port 不可为空")
		return false
	}
	return true
}



