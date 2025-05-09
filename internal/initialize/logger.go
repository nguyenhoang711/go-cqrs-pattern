package initialize

import (
	"github.com/nguyenhoang711/go-cqrs-pattern/global"
	"github.com/nguyenhoang711/go-cqrs-pattern/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
