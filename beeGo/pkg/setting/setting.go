package setting

import "github.com/astaxie/beego"

type ConfigParams struct {
	HttpAddr   string
	HttpPort   string
	MysqlUser  string
	MysqlPass  string
	MysqlUrls  string
	MysqlDb    string
	DriverName string
}

func LoadParams() *ConfigParams {
	config := &ConfigParams{}
	config.HttpAddr = beego.AppConfig.String("httpAddr")
	config.HttpPort = beego.AppConfig.String("httpport")
	config.DriverName = beego.AppConfig.String("driverName")
	config.MysqlUser = beego.AppConfig.String("mysqluser")
	config.MysqlPass = beego.AppConfig.String("mysqlpass")
	config.MysqlUrls = beego.AppConfig.String("mysqlurls")
	config.MysqlDb = beego.AppConfig.String("mysqldb")
	return config
}
