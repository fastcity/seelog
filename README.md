# seelog
> 有了seelog,妈妈再也不用担心我登录服务器查看日志le...

### 项目介绍
* 与golang项目集成、提供浏览器实时查看日志的功能，类似 [tail -f xxx.log](https://www.cnblogs.com/fps2tao/p/7698224.html)

### 集成方式
* 在项目中引入seelog, **go get github.com/xmge/seelog**
* 在代码中 执行 **seelog.See("logpath",port)**
* 在浏览器中访问 **http://host:port**

### 项目展示

