package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type PublicController struct {
	beego.Controller
}

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

const (
	SUCCESS = iota
	FAIL
)

//公共返回部分
func (p *PublicController) SendResponse(code int, msg string, data interface{}) {
	result := &Result{code,msg,data}
	fmt.Println(result)
	p.Data["json"] = map[string]int{"data": code}
	p.ServeJSON()
}
