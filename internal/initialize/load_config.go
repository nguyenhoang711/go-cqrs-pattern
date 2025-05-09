package initialize

import (
	"fmt"

	"github.com/nguyenhoang711/go-cqrs-pattern/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viperLoad := viper.New()
	viperLoad.AddConfigPath("./configs")
	viperLoad.SetConfigName("config")
	viperLoad.SetConfigType("yaml")

	err := viperLoad.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	if err := viperLoad.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}

}
