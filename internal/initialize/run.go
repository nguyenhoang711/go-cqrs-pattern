package initialize

import (
	"fmt"

	"github.com/nguyenhoang711/go-cqrs-pattern/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	StartTracing()
	InitEventStoreDB()
	_, err := InitElasticSearch()
	if err != nil {
		global.Logger.Error("errot to init elastic search", zap.Error(err))
	}

	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
