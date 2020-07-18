package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/TomatoMr/awesomeframework/admin"
	"github.com/TomatoMr/awesomeframework/config"
	"github.com/TomatoMr/awesomeframework/db"
	"github.com/TomatoMr/awesomeframework/logger"
	"github.com/TomatoMr/awesomeframework/mq"
	httpsrv "github.com/TomatoMr/awesomeframework/process/http"
	"github.com/TomatoMr/awesomeframework/process/rpc"
	"github.com/TomatoMr/awesomeframework/redis"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
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
	go httpsrv.StartHttpServer(config.GetConfig().HttpConfig.Addr)

	//启动rpc服务
	go rpc.StartRpcServer(config.GetConfig().RpcConfig.Addr)

	//启动nsq消费者
	go mq.StartMqServer()

	logger.GetLogger().Info("Init success.")

	pCh := admin.Get()
	for {
		select {
		case <-pCh.OnOff:
			timeout := time.Second * time.Duration(pCh.Timeout)
			go func() {
				mux := http.NewServeMux()
				mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
				mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
				mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
				mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
				mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
				s := &http.Server{Addr: pCh.Host, Handler: mux}
				logger.GetLogger().Info("start serve pprof")
				go s.ListenAndServe()
				time.Sleep(timeout)
				s.Shutdown(context.Background())
				<-pCh.Used
				logger.GetLogger().Info("stop serve pprof")
			}()
			pCh.Used <- struct{}{}
			logger.GetLogger().Info("fuck")
		}
	}
}
