package conf

import (
	"github.com/spf13/viper"
)

var conf = LoadConf()

func C() *viper.Viper {
	return conf
}

func LoadConf() *viper.Viper {

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
