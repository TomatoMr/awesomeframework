package mq

import (
	"github.com/TomatoMr/awesomeframework/config"
	"github.com/TomatoMr/awesomeframework/logger"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"os"
)

func StartMqServer() {
	conf := nsq.NewConfig()
	q, _ := nsq.NewConsumer(config.GetConfig().NsqConfig.Topic, config.GetConfig().NsqConfig.Channel, conf)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//do something when you receive a message
		logger.GetLogger().Info("receive", zap.String(config.GetConfig().NsqConfig.Topic, string(message.Body)))
		return nil
	}))
	err := q.ConnectToNSQLookupd(config.GetConfig().NsqConfig.NsqLookupdAddr)
	if err != nil {
		logger.GetLogger().Error("connect to nsqlookupd failed.", zap.Error(err))
		os.Exit(-1)
	}
}
