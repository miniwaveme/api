package config

import (
	"github.com/spf13/viper"
)

var (
	gConfig = LoadConfig()
)

func GetConfig() *viper.Viper {
	return gConfig
}

func LoadConfig() *viper.Viper {

	vip := viper.New()
	vip.SetEnvPrefix("miniwaveme")
	vip.AutomaticEnv()

	vip.SetDefault("api_config_path", "../../config/api/")
	vip.SetDefault("log.enabled", true)

	vip.SetConfigName("config")
	vip.AddConfigPath(vip.GetString("api_config_path"))
	vip.ReadInConfig()

	return vip
}
