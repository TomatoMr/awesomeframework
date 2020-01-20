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

type Config struct {
	LogConfig LogConfig `json:"log_config"`
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
