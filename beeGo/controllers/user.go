package controllers

import (
	"beeGo/models"
	"beeGo/util/Md5"
	"beeGo/util/jwtlib"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

// @Title 用户注册
// @Description 用户注册
// @Param	userAccount		FormData 	string	true		"账号"
// @Param	passWord		FormData 	string	true		"加密后的密码"
// @Param	userName		FormData 	string	true		"姓名"
// @Param	sex		FormData 	string	true		"性别"
// @Success 200 {code,msg} interface{}
// @Failure 403 :valid fail
// @router /Register [post]
func (u *UserController)Register()  {
	var user models.User
	user.Id = Md5.MakeMD5(u.GetString("userAccount"))
	user.UserName = u.GetString("userName")
	user.UserAccount = u.GetString("userAccount")
	user.Password = u.GetString("password")
	user.Sex = u.GetString("sex")
	bl :=models.UserRegister(user)
	if bl{
		u.Data["json"]=Response{0,"注册成功",""}
	}else {
		u.Data["json"]=ErrResponse{-1,"系统错误"}
	}
	u.ServeJSON()
}

// @Title 用户登陆
// @Description 用户登陆
// @Param	userAccount		FormData 	string	true		"账号"
// @Param	passWord		FormData 	string	true		"加密后的密码"
// @Success 200 {code,msg} interface{}
// @Failure 403 :valid fail
// @router /Login [post]
func (u *UserController) Login() {
	var user models.User
	user.UserAccount = u.GetString("userAccount")
	user.Password = u.GetString("password")
	bl,userInfo := models.UserLogin(user)
	if bl{
		token ,err:= jwtlib.EasyToken{UserName:userInfo.UserName,Expires:time.Now().Unix()+3600,Uid:userInfo.Id}.GetToken()
		if err != nil{
			u.Data["json"]=ErrResponse{-1,err}
		}else {
			if token != ""{
				u.Data["json"]=Response{0,"登陆成功", map[string]interface{}{
					"userInfo":userInfo,"token":token,
				}}
			}
		}
	}else {
		u.Data["json"]=ErrResponse{-1,"系统错误"}
	}
	u.ServeJSON()
}

//token验证失败返回信息
func (u *UserController) ValidFailToken(err error) {
	u.Data["json"]=ErrResponse{-1,err}
}