package filter

import (
	"beeGo/models"
	"beeGo/util/Md5"
	"beeGo/util/jwtlib"
	"fmt"
	"github.com/astaxie/beego/context"
)

var DumpHttpFilter = func(Ctx *context.Context) {
	method := Ctx.Request.Method
	header := Ctx.Request.Header
	body := Ctx.Request.Body
	fmt.Println("[dump http] method: ", method)
	fmt.Println("header: ", header)
	fmt.Println("body: ", body)
}

//获取Header验证  token
var ValidUserToken = func(Ctx *context.Context) bool{
	header := Ctx.Request.Header
	token:= header.Get("token")
	bl ,err,claims := jwtlib.EasyToken{}.ValidateToken(token)
	if claims != nil {
		var logs models.Logs
		logs.Id = Md5.MakeMD5("")
		logs.UserId = claims["jti"].(string)
		logs.UserName = claims["iss"].(string)
		logs.Path = fmt.Sprint(Ctx.Request.URL)
		logs.Method = fmt.Sprint(Ctx.Request.Method)
		models.WriteUserActionInfo(logs)
	}
	if !bl{
		Ctx.Output.Body([]byte(fmt.Sprint(err)))
		return false
	}
	return true
}