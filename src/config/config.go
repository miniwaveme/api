package config

import (
	"github.com/spf13/viper"
)

var (
	gConfig *viper.Viper
)

func GetConfig() *viper.Viper {
	return gConfig
}

func LoadConfig() {

	gConfig = viper.New()
	gConfig.SetEnvPrefix("miniwaveme")
	gConfig.AutomaticEnv()

	gConfig.SetDefault("api_config_path", "../../config/api/")
	gConfig.SetDefault("log.enabled", true)

	gConfig.SetConfigName("config")
	gConfig.AddConfigPath(gConfig.GetString("api_config_path"))
	gConfig.ReadInConfig()
}
