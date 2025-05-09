package eventstoredb

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"go.uber.org/zap"

	"github.com/nguyenhoang711/go-cqrs-pattern/global"
	"github.com/nguyenhoang711/go-cqrs-pattern/pkg/config"
)

func NewEventStoreDB(cfg *config.EventStoreConfig) (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(cfg.ConnectionString)
	if err != nil {
		global.Logger.Error("erorr to parse connect string", zap.Error(err))
		return nil, err
	}

	return esdb.NewClient(settings)
}
