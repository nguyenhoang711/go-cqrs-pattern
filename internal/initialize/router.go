package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhoang711/go-cqrs-pattern/global"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	return r
}
