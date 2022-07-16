package util

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
	SSLmode  string `mapstructure:"sslmode"`
	TimeZone string `mapstructure:"timezone"`
}

func LoadDBConfig(path string) (config DBConfig, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.UnmarshalKey("database", &config)
	return
}

func LoadServerConfig(path string) (config ServerConfig, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.UnmarshalKey("server", &config)
	return
}
