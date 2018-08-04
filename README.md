# seelog 
> 有了seelog,妈妈再也不用担心我登录服务器查看日志le...   
项目地址：https://github.com/xmge/seelog    
演示地址：http://seelog.xmge.top    
欢迎各位gopher使用指正:smiley: 

### 项目背景
> A：我去:confused:,,程序又出问题了，你去看看日志到底是哪里错啦  
  B：好好好，我马上登陆服务器看看    
  B：对啦，服务器密码是啥来着....    
  A：:cold_sweat: 密码是 &#￥&*@*~    
  B：额，日志文件在哪呢    
  A：...................    
  C：要不项目集成seelog吧，集成特别简单，集成后直接打开浏览器就能看日志了

### 项目介绍
* 与golang项目集成、提供浏览器实时查看日志的功能，类似 [tail -f xxx.log](https://www.cnblogs.com/fps2tao/p/7698224.html)
* 支持多浏览器同时访问
* 支持浏览器websocket断线重连
* 支持暂停、清屏、截图功能

### 集成方式
* 在项目中引入seelog, **go get github.com/xmge/seelog**
* 在代码中 执行 **seelog.See(logpath,port)**
* 在浏览器中访问 *http://host:port*

### 项目展示
![image](https://github.com/xmge/file/blob/master/images/seelog1.gif)

### 项目TODO
* 程序中直接输出日志
* 日志过滤
* os.stat()延迟
* 页面布局
* 分布式

