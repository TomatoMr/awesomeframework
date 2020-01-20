package db

import (
	"github.com/TomatoMr/awesomeframework/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

var engine *xorm.Engine

func InitEngine() error {
	var err error
	conf := config.GetConfig()
	engine, err = xorm.NewEngine("mysql", conf.DBConfig.DbUser+
		":"+conf.DBConfig.DbPassword+"@tcp("+conf.DBConfig.DbHost+":"+conf.DBConfig.DbPort+")/"+conf.DBConfig.DbName+"?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "InitEngine1")
		return err
	}
	err = engine.Ping()
	if err != nil {
		err = errors.Wrap(err, "InitEngine2")
		return err
	}
	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
