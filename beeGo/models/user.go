package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id          string `orm:"pk;column(id)";form:"id"`
	UserAccount string `orm:"column(user_account)";form:"userAccount"`
	Password string `orm:"column(password)";form:"password"`
	UserName    string `orm:"column(user_name)";form:"userName"`
	Sex string `orm:"column(sex)";form:"sex"`
}

func init()  {
	orm.RegisterModel(new(User))
}

//用户注册
func UserRegister(user User) bool{
	_,err := orm.NewOrm().Insert(&user)
	if err != nil{
		return false
	}
	return true
}

//用户登陆
func UserLogin(user User)(bool,User) {
	err := orm.NewOrm().Raw("select * from user where user_account= ? and password = ?",user.UserAccount,user.Password).QueryRow(&user)
	if err == nil{
		return  true,user
	}
	return false,user
}
