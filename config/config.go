package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

type LogConfig struct {
	LogPath  string `json:"log_path"`
	LogLevel string `json:"log_level"`
}

type DBConfig struct {
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
}

type Config struct {
	LogConfig LogConfig `json:"log_config"`
	DBConfig  DBConfig  `json:"db_config"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		err = errors.Wrap(err, "Read config file failed.")
		return err
	}
	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal config file failed.")
		return err
	}
	return nil
}

func GetConfig() Config {
	return conf
}
