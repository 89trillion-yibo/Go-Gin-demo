package httpserver

import (
	"awesomeProject/gin_demo/internal/router"
	"github.com/gin-gonic/gin"
)

//启动服务
func InitRun(port string) {
	r:= gin.Default()
	router.Routers(r)
	r.Run(port)
}
