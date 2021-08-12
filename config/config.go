package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("./conf")
}

func GetMySQLConfig() MysqlConf {
	viper.SetConfigName("db_conf")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper readInitConfig error")
	}
	var mysqlConf MysqlConf
	if err := viper.Unmarshal(&mysqlConf); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper Unmarshal error")
	}
	return mysqlConf
}
