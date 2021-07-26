package main

import (
	"awesomeProject/gin_demo/app/httpserver"
	"awesomeProject/gin_demo/utils"
)

func main() {
	//解析ini文件，调出监听端口
	port := utils.RandIni()
	//启动服务
	httpserver.InitRun(port)
}



