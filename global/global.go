package global

import (
	database "github.com/nguyenhoang711/go-cqrs-pattern/database/sqlc"
	"github.com/nguyenhoang711/go-cqrs-pattern/pkg/config"
	"github.com/nguyenhoang711/go-cqrs-pattern/pkg/logger"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	Db     *database.Store
)
