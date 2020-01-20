package main

import (
	"flag"
	"fmt"
	"github.com/TomatoMr/awesomeframework/config"
	"github.com/TomatoMr/awesomeframework/logger"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "配置文件路径")
	flag.Parse()

	if configPath == "" {
		fmt.Printf("Config Path must be assigned.")
		os.Exit(1)
	}

	var err error
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("Init config failed. Error is %v", err)
		os.Exit(1)
	}

	logConfig := config.GetConfig().LogConfig

	err = logger.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		fmt.Printf("Init logger failed. Error is %v", err)
		os.Exit(1)
	}

	logger.GetLogger().Info("Init success.")
}
