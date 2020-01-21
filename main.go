package main

import (
	"flag"
	"fmt"
	"github.com/TomatoMr/awesomeframework/config"
	"github.com/TomatoMr/awesomeframework/db"
	"github.com/TomatoMr/awesomeframework/logger"
	"github.com/TomatoMr/awesomeframework/process/http"
	"github.com/TomatoMr/awesomeframework/process/rpc"
	"github.com/TomatoMr/awesomeframework/redis"
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

	err = db.InitEngine()
	if err != nil {
		fmt.Printf("Init DB failed. Error is %v", err)
		os.Exit(1)
	}

	err = redis.InitRedis()
	if err != nil {
		fmt.Printf("Init Redis failed. Error is %v", err)
		os.Exit(1)
	}

	//启动http服务
	go http.StartHttpServer(config.GetConfig().HttpConfig.Addr)

	//启动rpc服务
	go rpc.StartRpcServer(config.GetConfig().RpcConfig.Addr)

	logger.GetLogger().Info("Init success.")

	select {}
}
