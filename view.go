package seelog

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

var conn *websocket.Conn

// 进行页面的socket的连接
func genConn(ws *websocket.Conn)  {
	conn = ws
	for {
		view(<- buffer)
	}
}

// 向页面写数据
func view(msg string)  {
	_,err := conn.Write([]byte(msg))
	if err != nil {
		log.Println("[seelog]",err.Error())
	}
}

func page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	w.WriteHeader(200)
	content := `<!DOCTYPE html>
				<html>
				<head>
					<meta charset="UTF-8"/>
					<title>seelog</title>
					<link rel="shortcut icon" href="http://www.eiguo.cn/favicon.ico" type="image/x-icon" />
					<script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
				
					<script>
					 function connect (){
								var ws = new WebSocket("ws://"+ window.location.host +"/ws");
								ws.onmessage = function(e) {
									$('#log').append("<p style='color: white'>"+ e.data +"</p>").scrollTop($('#log')[0].scrollHeight)
								};
								ws.onclose = function () {
									$('#status').css("background-color","red").text("链接断开")
									reConnect()
								}
								ws.onopen = function () {
									$('#status').css("background-color","chartreuse").text("连接成功")
								}
								ws.onerror = function (e) {
									$('#status').css("background-color","red").text("链接断开")
								}
							}

                      function reConnect(){
								setTimeout(function(){
									connect();
								},1000);
							}
						connect()
					</script>
				
				</head>
				<body>
				
				<header>
					<h2 id="title">实时查看日志信息 &nbsp;<button id="status" style="background-color: darkorange">正在连接...</button></h2>
				</header>
				<div id="log"></div>
				</body>
				
				<style>
					body {
						margin-left: 2%
					}
					#title {
				
					}
					#log {
						width:96%;
						height: 800px;
						background-color:#181818;
						border: 1px #ccc solid;
						overflow-y: scroll;
						float: left;
					}
				</style>
				</html>`
	w.Write([]byte(content))
	w.Header()
}