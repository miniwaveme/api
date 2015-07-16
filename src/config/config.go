package config

import (
	"github.com/spf13/viper"
	"os"
)

var (
	gConfig *viper.Viper
)

func GetConfig() *viper.Viper {
	return gConfig
}

func LoadConfig() {

	configPath := os.Getenv("MINIWAVEME_API_CONFIG_PATH")

	if &configPath == nil {
		configPath = "../../config/api/"
	}

	gConfig = viper.New()

	gConfig.SetConfigName("config")
	gConfig.AddConfigPath(configPath)
	err := gConfig.ReadInConfig()

	if err != nil {
		//log.Warn("No configuration file loaded - using defaults")
		gConfig.SetDefault("miniwaveme.log.enabled", "default")
	}
}
