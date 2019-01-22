package database

import (
	"beeGo/pkg/setting"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterMysql() {
	config := setting.LoadParams()
	orm.RegisterDriver(config.DriverName, orm.DRMySQL)
	err:=orm.RegisterDataBase(
		"default",
		config.DriverName,
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8",
			config.MysqlUser, config.MysqlPass, config.MysqlUrls, config.MysqlDb,
		),
		30)
	//orm.RunSyncdb("default", false, true)
	if err!=nil{
		fmt.Println("数据库连接失败!")
	} else {
		fmt.Println("数据库连接成功!")
	}
	orm.Debug = true
}
