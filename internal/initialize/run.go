package initialize

import (
	"fmt"

	"github.com/nguyenhoang711/go-cqrs-pattern/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	StartTracing()

	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
