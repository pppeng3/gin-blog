package mysql

import (
	"Blog/config"
	"Blog/db/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

const base string = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true"

func Init() {
	mysqlConf := config.GetMySQLConfig()
	dsn := fmt.Sprintf(base, mysqlConf.Mysql.Username, mysqlConf.Mysql.Password, mysqlConf.Mysql.Host, mysqlConf.Mysql.Port, mysqlConf.Mysql.Dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "new_blog_",
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Blog{}, &model.Comment{})
}

func GetDB() *gorm.DB {
	return db.Debug()
}
