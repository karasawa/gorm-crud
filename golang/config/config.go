package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Db Db
}

type Db struct {
	Host string
	User string
	Password string
	DbName string
	Port string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}

	return &cfg, nil
}