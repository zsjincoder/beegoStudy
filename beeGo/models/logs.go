package models

import "github.com/astaxie/beego/orm"

type Logs struct {
	Id       string `orm:"pk;column(id)";json:"_"`
	UserId   string `orm:"column(user_id)";json:"jti"`
	UserName string `orm:"column(user_name)";json:"iss"`
	Path     string `orm:"column(path)";json:"_"`
	Method   string `orm:"column(method)"`
}

func init() {
	orm.RegisterModel(new(Logs))
}

//写入用户操作信息
func WriteUserActionInfo(logs Logs) bool {
	_, err := orm.NewOrm().Insert(&logs)
	if err == nil {
		return true
	}
	return false
}
